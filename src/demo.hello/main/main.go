package main

import (
	"flag"
	"fmt"
	"runtime"

	"demo.hello/demos"
	"demo.hello/examples"
)

func init() {
	fmt.Println("run init")
	fmt.Println("go version:", runtime.Version())
}

// testAccessControl : use struct from demos/demo01
func testAccessControl() {
	// obj := demos.MyObject{"pub", "pri"} // error
	obj := demos.GetMyObject()
	(&obj).Init("pub_test", "pri_test")

	fmt.Printf("public value: %s\n", obj.VarPublic)
	fmt.Printf("private value: %s\n", obj.MethodPublicGet())
}

// flag test
var (
	retCode = 200
	port    = 8080
	help    = false
)

func testFlagParser() {
	fmt.Println("flag test")
	flag.IntVar(&retCode, "c", 200, "return status code")
	flag.IntVar(&port, "p", 8080, "port number")
	flag.BoolVar(&help, "h", false, "help")

	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	fmt.Printf("url=>local:%d, status code=>%d\n", port, retCode)
}

func main() {
	// https://github.com/gopl-zh/gopl-zh.github.com.git
	examples.MainHello()
	// examples.MainIO()
	// examples.MainOO()
	// examples.MainGoRoutine()
	// examples.MainCrawl()
	// examples.MainReflect()
	// examples.MainHTTP()

	// testAccessControl()
	// testFlagParser()

	// demos.MainDemo01()
	// demos.MainDemo02()
	// demos.MainDemo03()
	demos.MainDemo04()

	// demos.MainCache()
	demos.MainUtils()

	fmt.Println("message:", demos.HelloMessage)
	fmt.Println("main done.")
}
