package sort

import "fmt"

// ------------------------------
// #1. 存在重复元素
// 给定一个整数数组，判断是否存在重复元素。
// ------------------------------

// 时间复杂度 O(n) 空间复杂度 O(n)
func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := hashMap[nums[i]]; ok {
			return true
		}
		hashMap[nums[i]] = 1
	}
	return false
}

// ------------------------------
// #2. 删除排序数组中的重复项 (*)
// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间。
// ------------------------------

func removeDuplicates01(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	slow := 0
	fast := 0
	for fast < len(nums) {
		for nums[slow] == nums[fast] {
			fast++
			if fast >= len(nums) {
				return slow + 1
			}
		}
		if nums[slow+1] != nums[fast] {
			nums[slow+1], nums[fast] = nums[fast], nums[slow+1]
		}
		slow++
		fast++
	}
	fmt.Println(nums)
	return slow + 1
}

func removeDuplicates02(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	slow := 0
	fast := 1
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			if slow+1 != fast {
				nums[slow+1] = nums[fast]
			}
			slow++
		}
		fast++
	}
	fmt.Println(nums)
	return slow + 1
}

// LeetCodeMain03 contains leetcode algorithms.
func LeetCodeMain03() {
	if false {
		fmt.Println("\n#1. 存在重复元素")
		fmt.Println("expect true, and actual:", containsDuplicate([]int{1, 2, 3, 1}))
		fmt.Println("expect false, and actual:", containsDuplicate([]int{1, 2, 3, 4}))

		fmt.Println("\n#2.1 删除排序数组中的重复项")
		fmt.Println("expect 2, and actual:", removeDuplicates01([]int{1, 1, 2}))
		fmt.Println("expect 5, and actual:", removeDuplicates01([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

		fmt.Println("\n#2.2 删除排序数组中的重复项")
		fmt.Println("expect 2, and actual:", removeDuplicates02([]int{1, 1, 2}))
		fmt.Println("expect 5, and actual:", removeDuplicates02([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	}

	fmt.Println("leetcode sample2 done.")
}
