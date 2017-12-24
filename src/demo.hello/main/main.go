package main

import (
	"fmt"

	"demo.hello/demos"
	"demo.hello/examples"
)

func init() {
	fmt.Println("run init")
}

// testAccessControl : use struct from demos/demo01
func testAccessControl() {
	// obj := demos.MyObject{"pub", "pri"} // error
	obj := demos.GetMyObject()
	(&obj).Init("pub_test", "pri_test")

	fmt.Printf("public value: %s\n", obj.VarPublic)
	fmt.Printf("private value: %s\n", obj.MethodPublicGet())
}

// cmd: go install src/demo.hello/main/main.go
func main() {
	// https://github.com/gopl-zh/gopl-zh.github.com.git
	// examples.MainHello()
	// examples.MainIO()
	// examples.MainOO()
	// examples.MainGoRoutine()
	examples.MainCrawl()

	demos.MainDemo01()
	testAccessControl()

	fmt.Println("main done.")
}
