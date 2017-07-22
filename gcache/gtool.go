package gcache

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"time"
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
func hashValue(str *string, max *int) int {
	var hash uint32 = 5381
	length := len(*str)

	for i := 0; i < length; i++ {
		hash = uint32(hash<<5) + hash + uint32((*str)[i])
	}

	return int(hash % uint32(*max))
}

/*
 *检查类型
 */
func checkType(value []byte) (int, bool) {
	//字符串
	if len(value) > 2 && value[0] == '"' && value[len(value)-1] == '"' {
		return TYPE_STRING, true
	}

	str := string(value)
	//int
	if _, err := strconv.Atoi(str); err == nil {
		return TYPE_INT, true
	}
	//float
	if _, err := strconv.ParseFloat(str, 64); err == nil {
		return TYPE_FLOAT, true
	}

	return 0, false
}

/**
 * 转换interface到string
 */
func interfaceToString(valueType uint8, value interface{}) string {
	switch valueType {
	case TYPE_INT:
		return strconv.Itoa(value.(int))
	case TYPE_FLOAT:
		return strconv.FormatFloat(value.(float64), 'f', -1, 64)
	case TYPE_STRING:
		return value.(string)
	default:
		return ""
	}
}

/*
 *返回随机字符串
 */
func randStr() string {
	ret := md5.Sum([]byte(time.Now().String()))
	return string(ret[:5])
}
