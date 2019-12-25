package sort

import (
	"fmt"
	"strconv"
)

// ------------------------------
// #1. 回文字符串
// ------------------------------

func isRecycleStr(s string) bool {
	start := 0
	end := len(s) - 1
	for start < end { // error: start != end
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// ------------------------------
// #2. 在一个字符串中找到第一个只出现一次的字符
// ------------------------------

func firstCharAppearOnce(s string) string {
	keys := make([]byte, 0, len(s))
	dict := make(map[byte]int, len(s)) // map无序

	for _, b := range []byte(s) {
		keys = append(keys, b)
		if _, ok := dict[b]; ok {
			dict[b]++
		} else {
			dict[b] = 1
		}
	}

	for _, k := range keys {
		if dict[k] == 1 {
			return string(k)
		}
	}
	return ""
}

// ------------------------------
// #3. 小写字母排在大写字母的前面
// ------------------------------

func charLowerFrontOfUpper(s string) string {
	// 'A'=65,'a'=97
	b := []byte(s)
	start := 0
	end := len(b) - 1

	for start < end {
		for int(b[start]) >= 'a' && start < end {
			start++
		}
		for int(b[end]) < 'Z' && start < end {
			end--
		}
		if start < end {
			b[start], b[end] = b[end], b[start]
		}
	}
	return string(b)
}

// ------------------------------
// #4. 找出字符串中最长的连续数字
// ------------------------------

func longestContinuiousNums(s string) string {
	var tmp, max, start int
	for i, ch := range s {
		if _, err := strconv.Atoi(string(ch)); err != nil {
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

// ------------------------------
// #5. 字符串查找，返回该子字符串在文本中的位置
// ------------------------------

func searchSubString(s, sub string) int {
	size := len(s)
	subSize := len(sub)
	for i := 0; i <= size-subSize; i++ {
		isFound := true
		for j := 0; j < subSize; j++ {
			if s[i+j] != sub[j] {
				isFound = false
				break
			}
		}
		if isFound {
			return i
		}
	}
	return -1
}

// ------------------------------
// #6. 两个字符串的相同字符
// ------------------------------

func getSameCharsInStrs(s1, s2 string) string {
	same := make([]rune, 0, 10)
	s := myString{
		val: s2,
	}

	for _, val := range []rune(s1) {
		if s.containsChar(val) {
			same = append(same, val)
		}
	}
	if len(same) == 0 {
		return ""
	}
	return string(same)
}

type myString struct {
	val string
}

func (s *myString) containsChar(c rune) bool {
	for _, val := range []rune(s.val) {
		if val == c {
			return true
		}
	}
	return false
}

// ------------------------------
// #7. reverse sentence words divied by space
// ------------------------------

func reverseByWord(s []rune, start, end int) {
	strReverse(s, start, end)

	slow := start
	fast := start
	for fast <= end {
		if string(s[fast]) == " " {
			strReverse(s, slow, fast-1)
			fast++
			slow = fast
		} else {
			fast++
		}
	}
	strReverse(s, slow, fast-1) // last word
}

func strReverse(s []rune, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

// TestStringsAlgorithms test for strings algorithms.
func TestStringsAlgorithms() {
	if false {
		fmt.Println("\n#1. 回文字符串")
		vals := []string{"ahaha", "ahha", "haha"}
		for _, val := range vals {
			fmt.Printf("%s is recycle string: %v\n", val, isRecycleStr(val))
		}

		fmt.Println("\n#2. 在一个字符串中找到第一个只出现一次的字符")
		s := "ahacchdeff"
		fmt.Println("first char appear once:", firstCharAppearOnce(s))

		fmt.Println("\n#3. 小写字母排在大写字母的前面")
		str := "HaJKPnobAAdCPc"
		fmt.Println("lower letters front of upper letters:", charLowerFrontOfUpper(str))

		fmt.Println("\n#4. 找出字符串中最长的连续数字")
		str1 := "abcd13579ed124ss123456789z"
		fmt.Println("longest continuious numbers:", longestContinuiousNums(str1))

		fmt.Println("\n#5. 字符串查找，返回该子字符串在文本中的位置")
		str2 := "this is a string test, search sub string."
		sub := "ing"
		fmt.Printf("sub string (%s) index at: %d\n", sub, searchSubString(str2, sub))

		fmt.Println("\n#6. 两个字符串的相同字符")
		s1 := "abcde"
		s2 := "bwcxyz"
		fmt.Printf("(%s) and (%s) same chars are: %s\n", s1, s2, getSameCharsInStrs(s1, s2))

		fmt.Println("\n#7. reverse string")
		word := []rune("hello")
		fmt.Println("src string:", string(word))
		strReverse(word, 0, len(word)-1)
		fmt.Println("reverse string:", string(word))

		fmt.Println("\n#7. reverse sentence words divied by space")
		words := []rune("this is a test!")
		fmt.Println("src string:", string(words))
		reverseByWord(words, 0, len(words)-1)
		fmt.Println("string reverse by word:", string(words))
	}

	fmt.Println("strings algorithms done.")
}
