package sort

import (
	"fmt"
	"strconv"
)

// ------------------------------
// #1. 位1的个数
// 编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 '1' 的个数。
// ------------------------------

func hammingWeight(num uint32) int {
	var count int
	for num > 0 {
		count += int(num & 1)
		num >>= 1
	}
	return count
}

// ------------------------------
// #2. Fizz Buzz
// 1. 如果 n 是3的倍数，输出 Fizz
// 2. 如果 n 是5的倍数，输出 Buzz
// 3. 如果 n 同时是3和5的倍数，输出 FizzBuzz
// ------------------------------

func fizzBuzz(n int) []string {
	strs := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			strs[i-1] = "FizzBuzz"
		} else if i%5 == 0 {
			strs[i-1] = "Buzz"
		} else if i%3 == 0 {
			strs[i-1] = "Fizz"
		} else {
			strs[i-1] = strconv.Itoa(i)
		}
	}
	return strs
}

// LeetCodeMain02 contains leetcode algorithms.
func LeetCodeMain02() {
	if true {
		fmt.Println("\n#1. 位1的个数")
		fmt.Println("expect 3, actual:", hammingWeight(11))

		fmt.Println("\n#2. Fizz Buzz")
		fmt.Println("fizz buzz results:", fizzBuzz(15))
	}

	fmt.Println("leetcode sample2 done.")
}
