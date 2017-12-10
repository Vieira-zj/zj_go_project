package main

import (
	"fmt"

	"demo.hello/examples"
)

func init() {
	fmt.Println("run init")
}

func main() {
	// cmd: go install src/main/main.go
	examples.MainHello()
	// examples.MainIO()
	// examples.MainOO()
	// examples.MainGoRoutine()

	fmt.Println("main done.")
}
