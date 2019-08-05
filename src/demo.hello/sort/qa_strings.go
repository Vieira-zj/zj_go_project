package sort

import (
	"fmt"
	"strconv"
)

// 在一个字符串中找到第一个只出现一次的字符
func charFirstAppearOnce(s string) string {
	keys := make([]byte, 0, len(s))
	dict := make(map[byte]int, len(s))

	for _, b := range []byte(s) {
		keys = append(keys, b)
		if _, ok := dict[b]; ok {
			dict[b]++
		} else {
			dict[b] = 1
		}
	}

	for _, b := range keys {
		if dict[b] == 1 {
			return string(b)
		}
	}
	return "nil"
}

// 小写字母排在大写字母的前面
func charLowerFrontOfUpper(s string) string {
	// 'A'=65,'a'=97
	b := []byte(s)
	start := 0
	end := len(b) - 1

	for start < end {
		for int(b[start]) >= 97 && start < end {
			start++
		}
		for int(b[end]) < 97 && start < end {
			end--
		}
		if start < end {
			b[start], b[end] = b[end], b[start]
		}
	}
	return string(b)
}

// 字符串中找出连续最长的数字
func longestContinuiousNumbers(s string) string {
	var (
		tmp   = 0
		max   = 0
		start = 0
	)

	for i, b := range s {
		if _, err := strconv.Atoi(string(b)); err != nil {
			if tmp > max {
				max = tmp
			}
			tmp = 0
		} else { // 数字
			tmp++
			if tmp == 1 {
				start = i
			}
		}
	}
	if tmp > max {
		max = tmp
	}
	return s[start : start+max]
}

// 字符串查找，返回该字字符串在文本中的位置
func searchSubString(s, sub string) int {
	size := len(s)
	subSize := len(sub)
	for i := 0; i <= size-subSize; i++ {
		j := 0
		for ; j < subSize; j++ {
			if s[i+j] != sub[j] {
				break
			}
		}
		if j == subSize {
			return i
		}
	}
	return -1
}

// TestStringsAlgorithms test for strings algorithms.
func TestStringsAlgorithms() {
	s := "ahbaccdeff"
	fmt.Println("\nfirst appear once char:", charFirstAppearOnce(s))

	s = "HaJKPnobAAdCPc"
	fmt.Println("\nlower letters front of upper letters:", charLowerFrontOfUpper(s))

	s = "abcd13579ed124ss123456789"
	fmt.Println("\nlongest continuious numbers:", longestContinuiousNumbers(s))

	s = "this is a string test, find sub string in text."
	sub := "ing"
	fmt.Printf("\nsearch sub string (%s) at: %d\n", sub, searchSubString(s, sub))
}
