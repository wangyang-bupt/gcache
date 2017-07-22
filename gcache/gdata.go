package gcache

type gdata struct {
	key       string
	value     interface{}
	next      *gdata
	valueType uint8
}

/*
 *设置Key-value
 */
func (g *gdata) setValue(key string, valueType uint8, value interface{}) {
	g.key = key
	g.value = value
	g.valueType = valueType
}

/*
 *获取value
 */
func (g *gdata) getValue() (uint8, interface{}) {
	return g.valueType, g.value
}

/*
 *获取类型名
 */
func (g *gdata) getTypeString() string {
	switch g.valueType {
	case TYPE_INT:
		return "int"
	case TYPE_FLOAT:
		return "float"
	case TYPE_STRING:
		return "string"
	}
	return "unknow"
}

func (g *gdata) getKey() string {
	return g.key
}
