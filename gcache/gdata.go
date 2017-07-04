package gcache

type gdata interface {
	Get()
	Set()
	GetType()
}

type gInt struct {
	value int
}

func (v *gInt) Get() int {
	return v.value
}

func (v *gInt) Set(value int) {
	v.value = value
}

func (v *gInt) GetType() string {
	return "int"
}
