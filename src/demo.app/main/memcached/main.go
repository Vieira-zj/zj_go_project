package main

import (
	"fmt"

	"demo.app/memcached"
)

func main() {
	memcached.ConnectMemcacheAndTest()
	fmt.Println("memcached demo done.")
}
