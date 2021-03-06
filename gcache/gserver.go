package gcache

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

var (
	db          gdb
	commandChan chan []byte
	gconnArray  map[string]*gconnnect
	connectNum  int
)

func init() {
	db.size = INIT_GDATA_NUM
	db.gdatas = make([]*gdata, INIT_GDATA_NUM)
	commandChan = make(chan []byte, 1000)
	gconnArray = make(map[string]*gconnnect, MAX_CONNECT)
	connectNum = 0
	os.Mkdir(BACKUP_FOLDER, 0644)
}

func ServerInit() {
	listener, err := net.Listen("tcp", SEVER_ADDRESS)
	checkExitError(err)
	fmt.Println("welcome!")

	//单个goruntine处理命令
	go commandEvent()

	for {
		conn, _ := listener.Accept()
		//最大连接数限制
		if connectNum == MAX_CONNECT {
			conn.Close()
		} else {
			connectNum++
			gconn := gconnnect{conn}
			for {
				gkey := randStr()
				if _, ok := gconnArray[gkey]; !ok {
					gconnArray[gkey] = &gconn
					go handleServerConn(&gconn, gkey)
					break
				}
			}
		}
	}
}

/*
 *处理到来的命令行事件
 */
func commandEvent() {
	var command []byte
	var response string

	for {
		command = <-commandChan
		/**获取conn在全局数组中的key**/
		gKey := string(command[len(command)-5:])
		/**处理参数**/
		commandArray := bytes.Split(command[:len(command)-5], []byte{DELIMITER})
		switch commandArray[0][0] {
		case SET:
			/**db key typeValue value*/
			response = setEvent(&db, string(commandArray[1]), int(commandArray[2][0]), string(commandArray[3]))
		case GET:
			/**db key*/
			response = getEvent(&db, string(commandArray[1]))
		case DELETE:
			response = deleteEvent(&db, string(commandArray[1]))
		case TYPE:
			response = typeEvent(&db, string(commandArray[1]))
		case INCR:
			response = incrDecrEvent(&db, string(commandArray[1]), INCR)
		case DECR:
			response = incrDecrEvent(&db, string(commandArray[1]), DECR)
		case BACKUP:
			response = backupEvent(&db, string(commandArray[1]))
		case RECOVERY:
			response = recoveryEvent(&db, string(commandArray[1]))
		}

		gconnArray[gKey].write([]byte(response))
	}
}

/*
 *处理一个连接
 */
func handleServerConn(gconn *gconnnect, gkey string) {
	defer gconn.close()
	defer delete(gconnArray, gkey)

	for {
		command, err := gconn.read()
		command = append(command, []byte(gkey)...)
		if err == nil {
			commandChan <- command
		} else if err == io.EOF {
			break
		} else if checkWarn(err) {
			gconn.write([]byte("wrong command!"))
			continue
		}
	}

	fmt.Println("one connect end")
}
