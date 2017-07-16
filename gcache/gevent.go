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

	/**是第一个元素**/
	realKey := string(key)
	hash := hashValue(&realKey, &db.size)
	node := &db.gdatas[hash]
	if node.key == "" {
		node.setValue(realKey, uint8(valueType), realValue)
		return STR_SUCC, true
	}

	/**不是第一个元素**/
	for {
		if node.key == realKey {
			node.setValue(realKey, uint8(valueType), realValue)
			break
		}
		if node.next == nil {
			break
		}
		node = node.next
	}
	newNode := new(gdata)
	newNode.setValue(realKey, uint8(valueType), realValue)
	node.next = newNode
	return STR_SUCC, true
}

/*
 *get事件
 */
func getEvent(db *gdb, key []byte) (string, bool) {
	realKey := string(key)
	hash := hashValue(&realKey, &db.size)

	node := &db.gdatas[hash]
	for {
		if node == nil {
			return "", false
		}
		if node.key == realKey {
			return interfaceToString(node.getValue()), true
		}
		node = node.next
	}
	return "", false
}
