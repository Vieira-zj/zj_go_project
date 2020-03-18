package sort

import (
	"strconv"
	"strings"
	"testing"
)

func TestIntToString(t *testing.T) {
	s := make([]string, 0)
	for i := 0; i <= 10; i++ {
		s = append(s, strconv.Itoa(i))
	}
	t.Log(strings.Join(s, ""))

	str := ""
	for i := 0; i <= 10; i++ {
		str += strconv.Itoa(i)
	}
	t.Log(str)
}

func TestCompressString(t *testing.T) {
	input := []string{"aabcccccaaa", "abbccd"}
	expect := []string{"a2b1c5a3", "abbccd"}
	for i := 0; i < len(input); i++ {
		ret := compressString(input[i])
		if ret != expect[i] {
			t.Fatalf("input %s, exepct %s, and actual %s", input[i], expect[i], ret)
		}
	}
}
