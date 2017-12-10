package main

import (
	"fmt"
	"os"
)

func init() {
	fmt.Println("$GOROOT: " + os.Getenv("GOROOT"))
	fmt.Println("$GOPATH: " + os.Getenv("GOPATH"))
}

// cmd: go install src/demo.tests/main/test.go
func main() {
	fmt.Println("test main done.")
}
