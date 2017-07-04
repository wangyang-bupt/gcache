package gcache

import (
	"fmt"
	"io"
	"net"
	_ "strings"
	_ "time"
)

func ServerInit() {
	listener, err := net.Listen("tcp", SEVER_ADDRESS)
	checkExitError(err)
	fmt.Println("welcome!")

	for {
		conn, _ := listener.Accept()
		gconn := gconnnect{conn}
		go handleServerConn(&gconn)
	}
}

func handleServerConn(gconn *gconnnect) {
	defer gconn.close()

	for {
		command, err := gconn.read()

		if err == io.EOF {
			break
		}
		if checkWarn(err) {
			gconn.write("wrong command!\n")
			continue
		}

		gconn.write("I get your massage\n")
		fmt.Println(command)
	}

	fmt.Println("one connect end")
}
