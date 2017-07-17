package gcache

import (
	"strconv"
)

/*
 *set命令
 */
func setEvent(db *gdb, key []byte, valueType int, value []byte) (string, bool) {
	var realValue interface{}
	var err error
	switch valueType {
	case TYPE_INT8:
		realValue, err = strconv.ParseInt(string(value), 10, 8)
	case TYPE_INT32:
		realValue, err = strconv.ParseInt(string(value), 10, 32)
	case TYPE_INT64:
		realValue, err = strconv.ParseInt(string(value), 10, 64)
	case TYPE_FLOAT64:
		realValue, err = strconv.ParseFloat(string(value), 64)
	case TYPE_STRING:
		realValue = string(value)
	default:
		return STR_FAIL, false
	}

	if err != nil {
		return STR_FAIL, false
	}

	if db.setNode(string(key), uint8(valueType), realValue) {
		return STR_SUCC, true
	} else {
		return STR_FAIL, false
	}
}

/*
 *get事件
 */
func getEvent(db *gdb, key []byte) (string, bool) {
	return db.getNode(string(key))
}

/*
 *delete事件
 */
func deleteEvent(db *gdb, key []byte) (string, bool) {
	if db.deleteNode(string(key)) {
		return STR_SUCC, true
	} else {
		return STR_FAIL, false
	}
}
