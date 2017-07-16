package main

import (
	"gcache/gcache"
	_ "strings"
	_ "time"
)

func main() {
	gcache.ServerInit()
}
