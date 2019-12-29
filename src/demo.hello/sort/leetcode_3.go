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

// ------------------------------
// #3. 缺失数字
// 给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。
// ------------------------------

func missingNumber01(nums []int) int {
	m := make(map[int]struct{})
	for _, num := range nums {
		m[num] = struct{}{}
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return -1
}

func missingNumber02(nums []int) int {
	var sumAll int
	for i := 1; i <= len(nums); i++ {
		sumAll += i
	}

	var sum int
	for _, num := range nums {
		sum += num
	}
	return sumAll - sum
}

// ------------------------------
// #4. 相交链表
// 找到两个单链表相交的起始节点。
// ------------------------------

func getIntersectionNode(headA, headB *listNode) *listNode {
	// 两个链表长度不等，找到长度相同时的起点，向后迭代
	if headA == nil || headB == nil {
		return nil
	}

	curNodeA := headA
	curNodeB := headB
	for curNodeA != curNodeB {
		if curNodeA == nil {
			curNodeA = headB
		} else {
			curNodeA = curNodeA.Next
		}
		if curNodeB == nil {
			curNodeB = headA
		} else {
			curNodeB = curNodeB.Next
		}
	}
	return curNodeA
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

		fmt.Println("\n#3.1 缺失数字")
		fmt.Println("expect 2, and actual:", missingNumber01([]int{3, 0, 1}))
		fmt.Println("expect 8, and actual:", missingNumber01([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
		fmt.Println("\n#3.2 缺失数字")
		fmt.Println("expect 2, and actual:", missingNumber02([]int{3, 0, 1}))
		fmt.Println("expect 8, and actual:", missingNumber02([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
	}

	fmt.Println("leetcode sample3 done.")
}
