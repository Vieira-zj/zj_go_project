package sort

import (
	"strconv"
)

// ------------------------------
// #1. 字符串压缩
// 输入："aabcccccaaa"
// 输出："a2b1c5a3"
// 若“压缩”后的字符串没有变短，则返回原先的字符串。
// ------------------------------

func compressString(S string) string {
	if len(S) == 0 {
		return S
	}

	var comStr string
	ch := S[0]
	cnt := 1
	for idx := 1; idx < len(S); idx++ {
		if S[idx] != ch {
			comStr += string(ch) + strconv.Itoa(cnt)
			ch = S[idx]
			cnt = 1
		} else {
			cnt++
		}
	}
	comStr += string(ch) + strconv.Itoa(cnt)

	if len(comStr) >= len(S) {
		return S
	}
	return comStr
}
