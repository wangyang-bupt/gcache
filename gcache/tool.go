package gcache

import (
	"fmt"
	"os"
)

/**
 * 检查使程序退出的错误
 */
func checkExitError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(100001)
	}
}

/**
 * 检查程序一般的错误
 */
func checkWarn(err error) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		return true
	}
	return false
}

/**
 * DJBX33A算法实现的hash函数
 */
func Hash(str *string, max *int) int {
	var hash uint32 = 5381
	length := len(*str)

	for i := 0; i < length; i++ {
		hash = uint32(hash<<5) + hash + uint32((*str)[i])
	}

	return int(hash % uint32(*max))
}
