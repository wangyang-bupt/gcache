package gcache

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func ClientInit() {
	conn, err := net.DialTimeout("tcp", SEVER_ADDRESS, CLIENT_CONNCET_WAIT_TIME*time.Second)
	checkExitError(err)
	fmt.Println("welcome!")
	gconn := gconnnect{conn}
	handleClientConn(&gconn)
}

/*
 *处理连接
 */
func handleClientConn(gconn *gconnnect) {
	defer gconn.close()

	for {
		/**命令输入**/
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadBytes(IO_BUFF_END)
		var success bool
		var transData []byte

		if !checkWarn(err) {
			input = input[:len(input)-1]
			inputArray := bytes.Split(input, []byte{' '})
			/**匹配命令**/
			switch string(inputArray[0]) {
			case "set":
				transData, success = setCommand(inputArray)
			case "get":
				transData, success = getCommand(inputArray)
			case "delete":
				transData, success = deleteCommand(inputArray)
			default:
				success = false
			}
		}
		if success {
			gconn.write(transData)
		} else {
			fmt.Println("command error!")
			continue
		}

		reponse, err := gconn.read()
		if err == nil {
			fmt.Println(string(reponse))
		} else if err == io.EOF {
			break
		} else if checkWarn(err) {
			fmt.Println("network error!")
		}
	}
}

/*
 *set命令
 */
func setCommand(inputArray [][]byte) ([]byte, bool) {
	command := make([]byte, 0, 100)

	if len(inputArray) != 3 {
		return nil, false
	}

	//检查参数类型
	valueType, ok := checkType(inputArray[2])
	if !ok {
		return nil, false
	}

	//命令参数 key 类型参数 value
	command = append(command, SET)
	command = append(command, DELIMITER)

	command = append(command, inputArray[1]...)
	command = append(command, DELIMITER)

	command = append(command, byte(valueType))
	command = append(command, DELIMITER)

	if valueType != TYPE_STRING {
		command = append(command, inputArray[2]...)
	} else {
		command = append(command, inputArray[2][1:len(inputArray[2])-1]...)
	}

	return command, true
}

/*
 *get命令
 */
func getCommand(inputArray [][]byte) ([]byte, bool) {
	command := make([]byte, 0, 100)

	if len(inputArray) != 2 {
		return nil, false
	}

	//命令参数 key 类型参数 value
	command = append(command, GET)
	command = append(command, DELIMITER)

	command = append(command, inputArray[1]...)
	command = append(command, DELIMITER)
	return command, true
}

/*
 *delete命令
 */
func deleteCommand(inputArray [][]byte) ([]byte, bool) {
	command := make([]byte, 0, 100)

	if len(inputArray) != 2 {
		return nil, false
	}

	//命令参数 key 类型参数 value
	command = append(command, DELETE)
	command = append(command, DELIMITER)

	command = append(command, inputArray[1]...)
	command = append(command, DELIMITER)
	return command, true
}
