package main

import (
	"fmt"

	"demo.hello/demos"
	"demo.hello/examples"
)

func init() {
	fmt.Println("run init")
}

// cmd: go install src/demo.hello/main/main.go
func main() {
	// examples.MainHello()
	// examples.MainIO()
	// examples.MainOO()
	examples.MainGoRoutine()

	demos.MainDemo01()

	fmt.Println("main done.")
}
