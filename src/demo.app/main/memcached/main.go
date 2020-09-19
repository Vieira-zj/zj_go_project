package main

import (
	"fmt"

	"src/demo.app/memcached"
)

func main() {
	memcached.ConnectMemcacheAndTest()
	fmt.Println("memcached demo done.")
}
