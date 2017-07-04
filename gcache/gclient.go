package gcache

import (
	"bufio"
	"fmt"
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

func handleClientConn(gconn *gconnnect) {
	defer gconn.close()

	for {
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')

		if !checkWarn(err) {
			gconn.write(input)
			response, _ := gconn.read()
			fmt.Println(response)
		}
	}
}
