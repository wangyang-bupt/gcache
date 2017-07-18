package gcache

import (
	"strconv"
)

/*
 *set命令
 */
func setEvent(db *gdb, key []byte, valueType int, value []byte) string {
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
		return STR_FAIL
	}

	if err != nil {
		return STR_FAIL
	}

	if db.setNode(string(key), uint8(valueType), realValue) {
		return STR_SUCC
	} else {
		return STR_FAIL
	}
}

/*
 *get事件
 */
func getEvent(db *gdb, key []byte) string {
	node, _ := db.getNode(string(key))
    if node == nil {
        return ""
    }
    return interfaceToString(node.valueType, node.value)
}

/*
 *delete事件
 */
func deleteEvent(db *gdb, key []byte) string {
	if db.deleteNode(string(key)) {
		return STR_SUCC
	} else {
		return STR_FAIL
	}
}

/*
 *type事件
 */
func typeEvent(db *gdb, key[]byte) string {
    node, _ := db.getNode(string(key))
    if node == nil {
        return ""
    }
    return node.getTypeString()
}

/*
 *incr/decr事件
 */
func incrDecrEvent(db *gdb, key []byte, cr int) string {
    node, _ := db.getNode(string(key))
    
    if node == nil {
        var value int8
        if cr == INCR {
            value = 1 
        } else {
            value = -1
        }

        if db.setNode(string(key), uint8(TYPE_INT8), value) {
            ret := strconv.Itoa(int(value))
            return ret
        }
        return STR_FAIL
    }

    valueType := node.getTypeString()
    
    if valueType != "int" {
        return valueType
    }
    
    if cr == INCR {
        node.value = strconv.Atoi(interfaceToString(node.valueType, node.value)) + 1
    } else {
        node.value = strconv.Atoi(interfaceToString(node.valueType, node.value)) - 1
    }

    return strconv.itoa(node.value)
}
