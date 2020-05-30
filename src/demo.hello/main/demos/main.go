package main

import (
	"fmt"

	"demo.hello/demos"
)

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

func main() {
	// demos.MainDemo01()
	// demos.MainDemo02()
	// demos.MainDemo03()
	// demos.MainDemo04()
	// demos.MainDemo05()
	demos.MainDemo06()

	// demos.MainUtils()
	// demos.MainIterator()

	// testAccControl()
	// testInvokeOrder()

	fmt.Println("Go demos done.")
}