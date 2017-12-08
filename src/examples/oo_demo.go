package examples

import (
	"fmt"
	"sort"
)

type integer int

func (a integer) less(b integer) bool {
	return a < b
}

func (a *integer) add(b integer) {
	*a += b
}

func addMethodTest() {
	var a integer = 1
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

type stringer interface {
	String() string
}

type myStringer struct {
}

func (*myStringer) String() string {
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
			if v, ok := v.(stringer); ok {
				val := v.String()
				fmt.Println("default stringer:", val)
			} else {
				fmt.Println("other types")
			}
		}
	}
}

// sort.Interface
// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool // i, j are indices of sequence elements
// 	Swap(i, j int)
// }

type myStringSlice []string

func (p myStringSlice) Len() int {
	return len(p)
}

func (p myStringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p myStringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortTest() {
	names := []string{"candy", "brow", "dan", "emm", "anny"}
	fmt.Println("names before sorted:", names)
	sort.Sort(myStringSlice(names))
	fmt.Println("names after sorted:", names)
}

// MainOO : main function for struct and interface examples.
func MainOO() {
	// addMethodTest()
	// valueAndReferenceTest()

	// var myStr stringer = new(myStringer)
	// myPrintTest(1, "test", myStr, 1.0)

	sortTest()

	fmt.Println("oo demo.")
}
