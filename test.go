package main

import (
	"fmt"
	_ "reflect"
	_ "strconv"
	_ "unsafe"
)

type test struct {
	value int
}

func main() {
	x := make([]*test, 100)
	x[0] = new(test)
	x[0].value = 1
	y := x[0]
	y.value = 2
	fmt.Println(*x[0], *y)
}
