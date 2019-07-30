package sort

import (
	"fmt"
)

func reverse(s []rune, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func revertByWord(s []rune, start, end int) {
	reverse(s, start, end)

	i := start
	j := start
	for j <= end {
		if string(s[j]) == " " {
			reverse(s, i, j-1)
			j++
			i = j
		} else {
			j++
		}
	}
	reverse(s, i, j-1)
}

// TestRevertByWord test for revertByWord.
func TestRevertByWord() {
	tmp := []rune("hello")
	fmt.Println("\nsrc string:", string(tmp))
	reverse(tmp, 0, len(tmp)-1)
	fmt.Println("reverse string:", string(tmp))

	tmp = []rune("this is a test.")
	fmt.Println("\nsrc string:", string(tmp))
	revertByWord(tmp, 0, len(tmp)-1)
	fmt.Println("revert by word string:", string(tmp))
}
