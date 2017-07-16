package main

import (
	"fmt"
	_ "reflect"
	_ "strconv"
	_ "unsafe"
)

func main() {
	x := []byte("12345")
	fmt.Println(x[:4])
}
