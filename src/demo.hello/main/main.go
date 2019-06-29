package main

import (
	"flag"
	"fmt"
	"runtime"

	"demo.hello/apps"
	"demo.hello/demos"
	"demo.hello/examples"
)

func init() {
	fmt.Println("\n[main.go] init")
	fmt.Println("go version:", runtime.Version())
}

// flag test
var (
	retCode = 200
	port    = 8080
	help    = false
)

func testFlagParser() {
	flag.IntVar(&retCode, "c", 200, "status code")
	flag.IntVar(&port, "p", 8080, "port number")
	flag.BoolVar(&help, "h", false, "help")

	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	fmt.Printf("url=>localhost:%d, status code=>%d\n", port, retCode)
}

func testAccControl() {
	// struct in demos/demo01
	fmt.Println("\nfrom main, and only public fileds/methods can access:")
	// obj := demos.MyObject{"pub", "pri"} // error
	obj := demos.GetMyObject()
	obj.Init("pub_test", "pri_test")
	fmt.Printf("\npublic value: %s\n", obj.VarPublic)
	fmt.Printf("private value: %s\n", obj.MethodPublicGet())
}

func testInvokeOrder() {
	fmt.Println("\nmessage:", demos.HelloMsg)
}

func mainExample() {
	examples.MainHello()
	// examples.MainIO()
	// examples.MainOO()
	// examples.MainGoRoutine()
	// examples.MainHTTP()
	// examples.MainCrawl()
	// examples.MainReflect()
}

func mainDemo() {
	demos.MainDemo01()
	// demos.MainDemo02()
	// demos.MainDemo03()
	// demos.MainDemo04()
	// demos.MainUtils()
}

func mainAppDemo() {
	apps.MainCache()
}

func main() {
	// testFlagParser()
	// testAccControl()
	// testInvokeOrder()

	// mainExample()
	// mainDemo()
	mainAppDemo()

	fmt.Println("GO main done.")
}
