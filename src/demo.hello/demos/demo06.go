package demos

import (
	"fmt"

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

// MainDemo06 main for golang demo06.
func MainDemo06() {
	// testBase64Encode()

	fmt.Println("golang demo06 DONE.")
}
