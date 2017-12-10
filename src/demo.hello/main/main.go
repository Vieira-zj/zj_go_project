package main

import (
	"fmt"

	"demo.hello/examples"
)

func init() {
	fmt.Println("run init")
}

// cmd: go install src/demo.hello/main/main.go
func main() {
	examples.MainHello()
	// examples.MainIO()
	// examples.MainOO()
	// examples.MainGoRoutine()

	fmt.Println("main done.")
}
