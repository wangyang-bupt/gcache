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
