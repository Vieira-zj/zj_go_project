package sort

import (
	"fmt"
	"strings"
)

// #1. 验证回文串：给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写
func isPalindrome(s string) bool {
	str := strings.ToLower(s)
	start := 0
	end := len(str) - 1
	for start < end {
		for !isLetter(str[start]) && !isNumber(str[start]) && start < end {
			start++
		}
		for !isLetter(str[end]) && !isNumber(str[end]) && start < end {
			end--
		}
		if str[start] != str[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func isNumber(b byte) bool {
	return b >= '0' && b <= '9'
}

func isLetter(b byte) bool {
	return (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z')
}

// 两数之和：给定一个整数数组 nums 和一个目标值 target, 请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标
// #2.1 时间复杂度：O(n^2) 空间复杂度：O(1)
func twoSum01(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

// #2.2 时间复杂度：O(n) 空间复杂度：O(n)
func twoSum02(nums []int, target int) []int {
	tmpMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		n := target - nums[i]
		if val, ok := tmpMap[n]; ok {
			return []int{val, i}
		}
		tmpMap[nums[i]] = i
	}
	return []int{-1, -1}
}

// LeetCodeMain contains leetcode algorithms.
func LeetCodeMain() {
	fmt.Println("\n#1. 验证回文串")
	fmt.Println("excpect true, actual:", isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println("excpect false, actual:", isPalindrome("race a car"))

	fmt.Println("\n#2.1 两数之和")
	fmt.Println("expect [0,1], acutal:", twoSum01([]int{2, 7, 11, 15}, 9))
	fmt.Println("expect [0,2], acutal:", twoSum01([]int{-3, 4, 3, 90}, 0))

	fmt.Println("\n#2.2 两数之和")
	fmt.Println("expect [1,2], acutal:", twoSum02([]int{3, 2, 4}, 6))
	fmt.Println("expect [0,2], acutal:", twoSum02([]int{-3, 4, 3, 90}, 0))
}
