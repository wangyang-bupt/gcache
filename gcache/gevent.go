package gcache

import (
	"strconv"
)

/*
 *set命令
 */
func setEvent(db *gdb, key string, valueType int, value string) string {
	var realValue interface{}
	var err error
	switch valueType {
	case TYPE_INT:
		realValue, err = strconv.Atoi(value)
	case TYPE_FLOAT:
		realValue, err = strconv.ParseFloat(value, 64)
	case TYPE_STRING:
		realValue = value
	default:
		return STR_FAIL
	}

	if err != nil {
		return STR_FAIL
	}

	if db.setNode(key, uint8(valueType), realValue) {
		return STR_SUCC
	} else {
		return STR_FAIL
	}
}

/*
 *get事件
 */
func getEvent(db *gdb, key string) string {
	node, _ := db.getNode(string(key))
	if node == nil {
		return ""
	}
	return interfaceToString(node.valueType, node.value)
}

/*
 *delete事件
 */
func deleteEvent(db *gdb, key string) string {
	if db.deleteNode(key) {
		return STR_SUCC
	} else {
		return STR_FAIL
	}
}

/*
 *type事件
 */
func typeEvent(db *gdb, key string) string {
	node, _ := db.getNode(key)
	if node == nil {
		return ""
	}
	return node.getTypeString()
}

/*
 *incr/decr事件
 */
func incrDecrEvent(db *gdb, key string, cr int) string {
	node, _ := db.getNode(key)

	if node == nil {
		var value int
		if cr == INCR {
			value = 1
		} else {
			value = -1
		}

		if db.setNode(key, uint8(TYPE_INT), value) {
			ret := strconv.Itoa(value)
			return ret
		}
		return STR_FAIL
	}

	valueType := node.getTypeString()

	if valueType != "int" {
		return valueType
	}

	newValue, _ := strconv.Atoi(interfaceToString(node.valueType, node.value))
	if cr == INCR {
		node.value = newValue + 1
		return strconv.Itoa(newValue + 1)
	} else {
		node.value = newValue - 1
		return strconv.Itoa(newValue - 1)
	}

}
