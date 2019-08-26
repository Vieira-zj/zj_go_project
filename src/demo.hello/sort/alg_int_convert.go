package sort

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 十进制转二进制
func intOctToBinary(num int) string {
	stack := list.New()
	for num > 0 {
		stack.PushBack(strconv.Itoa(num % 2))
		num /= 2
	}

	var (
		ret []string
		len = stack.Len()
	)
	for i := 0; i < len; i++ {
		ele := stack.Back()
		stack.Remove(ele)
		if num, ok := ele.Value.(string); ok {
			ret = append(ret, num)
		}
	}
	return strings.Join(ret, "")
}

// 十进制转二进制
func intOctToBinary2(num int) string {
	var ret string
	for num > 0 {
		ret = strconv.Itoa(num%2) + ret
		num /= 2
	}
	return ret
}

// 二进制转十进制
func intBinaryToOct(bin string) int {
	bin = myReverse(bin)

	var ret float64
	if string(bin[0]) == "1" {
		ret = 1
	}
	for i := 1; i < len(bin); i++ {
		if string(bin[i]) == "1" {
			ret += math.Pow(2, float64(i))
		}
	}
	return int(ret)
}

func myReverse(s string) string {
	r := []rune(s)
	start := 0
	end := len(s) - 1
	for start < end {
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
	return string(r)
}

// TestIntOctAndBinary test for intOctToBinary and intBinaryToOct.
func TestIntOctAndBinary() {
	num := 10 // 1010
	fmt.Println("\nint oct to binary:")
	fmt.Printf("%d=%s\n", num, intOctToBinary(num))
	num = 110 // 1101110
	fmt.Printf("%d=%s\n", num, intOctToBinary2(num))

	fmt.Println("\nint binary to oct:")
	b := "1010"
	fmt.Printf("%s=%d\n", b, intBinaryToOct(b))
	b = "1101110"
	fmt.Printf("%s=%d\n", b, intBinaryToOct(b))
}
