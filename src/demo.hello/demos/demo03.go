package demos

import "fmt"

func testMapGetEmpty() {
	m := map[int]string{
		1: "one",
		2: "two",
	}
	fmt.Println("item at 2 =>", m[2])
	fmt.Printf("first char: %c\n", m[2][0])
	fmt.Printf("item length: %d\n", len(m[2]))

	if len(m) > 0 && len(m[3]) > 0 {
		fmt.Println("item at 3 =>", m[3])
	}
}

// MainDemo03 : main
func MainDemo03() {
	testMapGetEmpty()

	fmt.Println("demo 03 done.")
}
