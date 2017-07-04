package main

import (
	"gcache/gcache"
	_ "strings"
	_ "time"
)

func main() {
	gcache.ServerInit()
}

/**
 * 处理命令
 */
// func handleCommand(command *string) (string, []string) {
// 	strArr := strings.Split(*command, " ")
// 	commandLength := len(strArr)

// 	switch(strArr[0]) {
// 	case "use":
// 	}
// }
