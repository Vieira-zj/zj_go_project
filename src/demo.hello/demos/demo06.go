package demos

import (
	"fmt"
	"runtime"

	myutils "tools.app/utils"
)

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

// MainDemo06 main for golang demo06.
func MainDemo06() {
	// testBase64Encode()
	// testInitSliceAndRecovery()

	fmt.Println("golang demo06 DONE.")
}
