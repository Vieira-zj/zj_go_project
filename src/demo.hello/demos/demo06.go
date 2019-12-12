package demos

import (
	"fmt"
	"reflect"
	"runtime"

	myutils "tools.app/utils"
)

// demo, bits operation
func testBitsOperation() {
	fmt.Println("左移 1<<2:", 1<<2)
	fmt.Println("右移 10>>2:", 10>>2)
	// 异或：两个二进位，不同则该位为1, 否则该位为0
	fmt.Println("异或 10^2:", 10^2)
	fmt.Println("异或 10^10:", 10^10)
	// 或：两个二进位中只要有一个为1, 该位的结果值为1
	fmt.Println("或 10|2:", 10|2)
	// 与：两个二进位都为1, 该位的结果值才为1, 否则为0
	fmt.Println("与 10&2:", 10&2)
}

// demo, base64 encode for bytes
func testBase64Encode() {
	size := 16
	b := make([]byte, size, size)
	for i := 0; i < size; i++ {
		b[i] = uint8(60 + i)
	}
	fmt.Println("\nbytes:", string(b))
	fmt.Println("base64 encode string:", myutils.GetBase64Text(b))
}

// demo, init slice by make, and recover func
func testInitSliceAndRecovery() {
	defer func() {
		if p := recover(); p != nil {
			if err, ok := p.(runtime.Error); ok {
				fmt.Println("\n[runtime error]", err)
			} else {
				fmt.Println("\n[error]", p.(error))
			}
		}
	}()

	b := make([]byte, 0, 10)
	// b := make([]byte, 10, 10)
	for i := 0; i < 10; i++ {
		b[i] = uint8(60 + i)
	}
	fmt.Println("\nbytes:", string(b))
}

// demo, interface and type assert
type mockError struct {
	message string
}

func (e mockError) Error() string {
	return "mock " + e.message
}

func printError01(err interface{}) {
	if e, ok := err.(mockError); ok { // type assert
		fmt.Println("Mock Error:", e.Error())
		return
	}
	fmt.Println("not an error!")
}

func printError02(err interface{}) {
	if e, ok := err.(interface{ Error() string }); ok { // interface assert
		fmt.Println("Error:", e.Error())
		return
	}
	fmt.Println("not an error!")
}

func testInterfaceTypeAssert() {
	mockErr := mockError{
		message: "file write error!",
	}
	fmt.Println("\ntype assert:")
	printError01(mockErr)
	printError01("type")

	fmt.Println("\ninterface assert:")
	printError02(mockErr)
	printError02("interface")
}

// demo, point type assert
func isPointer(object interface{}) bool {
	t := reflect.TypeOf(object)
	fmt.Println("\nobject kind:", t.Kind())
	return t.Kind() == reflect.Ptr
}

func testPointTypeAssert() {
	mockErr := mockError{
		message: "file write error!",
	}
	fmt.Println("is point:", isPointer(mockErr))
	fmt.Println("is point:", isPointer(&mockErr))
}

// MainDemo06 main for golang demo06.
func MainDemo06() {
	testBitsOperation()
	// testBase64Encode()
	// testInitSliceAndRecovery()
	// testInterfaceTypeAssert()
	// testPointTypeAssert()

	fmt.Println("golang demo06 DONE.")
}
