package main

import (
	"flag"
	"fmt"
	"runtime"
	"strconv"

	"demo.hello/apps"
	"demo.hello/demos"
	"demo.hello/examples"
	"demo.hello/sort"
)

func init() {
	fmt.Println("\n[main.go] init")
	fmt.Println("go version:", runtime.Version())
	fmt.Println("system arch:", runtime.GOARCH)
	fmt.Println("default int size:", strconv.IntSize)

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Printf("runtime allocated memory: %d Kb\n", mem.Alloc/1024)
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
	// demos.MainDemo01()
	// demos.MainDemo02()
	// demos.MainDemo03()
	// demos.MainDemo04()
	// demos.MainDemo05()
	demos.MainDemo06()

	// demos.MainUtils()
	// demos.MainIterator()
}

func mainAppDemo() {
	apps.TestCache()
}

func sortDemo() {
	// sort.TestIntOctAndBinary()
	// sort.TestNumbersAlgorithms()

	// sort.TestRevertByWord()
	// sort.TestStringsAlgorithms()

	// sort.TestLinkedList()
	// sort.TestLinkedListAlgorithms()

	// sort.TestTreeHeap()
	// sort.TestTreeAlgorithms()

	// sort.TestSkipList()

	// sort.TestSortAlgorithms()
	// sort.TestSearchAlgorithms()
	sort.LeetCodeMain()
}

func main() {
	// testFlagParser()
	// testAccControl()
	// testInvokeOrder()

	// mainDemo()
	// mainExample()
	// mainAppDemo()

	sortDemo()

	fmt.Println("GO demo main done.")
}
