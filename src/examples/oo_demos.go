package main

import "fmt"

type Integer int

func (a Integer) less(b Integer) bool {
	return a < b
}

func (a *Integer) add(b Integer) {
	*a += b
}

func addMethodTest() {
	var a Integer = 1
	a.add(3)
	if a.less(2) {
		fmt.Println(a, "less 2")
	} else {
		fmt.Println(a, "greater 2")
	}
}

func valueAndReferenceTest() {
	srcArr := [3]int{1, 2, 3}
	var copiedArr = srcArr
	srcArr[1]++
	fmt.Println("src array:", srcArr)
	fmt.Println("copied array by value:", copiedArr)

	var pCopiedArr = &srcArr // *[3]int
	fmt.Println("copied array by reference:", *pCopiedArr)
}

type Stringer interface {
	String() string
}

type MyStringer struct {
}

func (*MyStringer) String() string {
	return "this is a method implement from Stringer"
}

func myPrintTest(args ...interface{}) {
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			fmt.Println(v, "=> type is int")
		case string:
			fmt.Println(v, "=> type is string")
		default:
			if v, ok := v.(Stringer); ok {
				val := v.String()
				fmt.Println("default stringer:", val)
			} else {
				fmt.Println("other types")
			}
		}
	}
}

func main() {
	// addMethodTest()
	// valueAndReferenceTest()

	var myStr Stringer = new(MyStringer)
	myPrintTest(1, "test", myStr, 1.0)

	fmt.Println("oo demo.")
}
