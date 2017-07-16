package gcache

type gdata struct {
	key       string
	value     interface{}
	next      *gdata
	valueType uint8
}

func (g *gdata) setValue(key string, valueType uint8, value interface{}) {
	g.key = key
	g.value = value
	g.valueType = valueType
}

func (g *gdata) getValue() (uint8, interface{}) {
	return g.valueType, g.value
}
