package sort

import (
	"fmt"
	"math"
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

// ------------------------------
// #3. 多数元素
// 给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
// ------------------------------

func majorityElement(nums []int) int {
	var group, count int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			group = nums[i]
			count++
			continue
		}
		if group == nums[i] {
			count++
		} else {
			count--
		}
	}
	return group
}

// ------------------------------
// #4. 罗马数字转整数
// 罗马数字 2 写做 II, 即为两个并列的 1. 12 写做 XII, 即为 X + II. 27 写做  XXVII, 即为 XX + V + II.
// 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII, 而是 IV.
// ------------------------------

func romanToInt(s string) int {
	retVal := 0
	prev := getMapperInt(s[0])
	for i := 1; i < len(s); i++ {
		cur := getMapperInt(s[i])
		if prev < cur {
			retVal -= prev
		} else {
			retVal += prev
		}
		prev = cur
	}
	return retVal + prev
}

func getMapperInt(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

// ------------------------------
// #5. 合并两个有序链表
// ------------------------------

func mergeTwoLists(l1 *listNode, l2 *listNode) *listNode {
	// 考虑 l1 或 l2 为空的情况
	head := &listNode{}
	cur := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head.Next
}

// ------------------------------
// #6. 移动零
// ------------------------------

func moveZeroes01(nums []int) {
	// 冒泡交换
	n := 0
	for i := 0; i < len(nums)-n; {
		if nums[i] != 0 {
			i++
			continue
		}
		for j := i; j < len(nums)-n-1; j++ {
			nums[j], nums[j+1] = nums[j+1], nums[j]
		}
		n++
	}
}

func moveZeroes02(nums []int) {
	// 非零数字前移
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[cur] = nums[i]
			cur++
		}
	}

	for ; cur < len(nums); cur++ {
		nums[cur] = 0
	}
}

// ------------------------------
// #7. 快乐数
// 输入: 19
// 输出: true
// 解释:
// 1^2 + 9^2 = 82
// 8^2 + 2^2 = 68
// 6^2 + 8^2 = 100
// 1^2 + 0^2 + 0^2 = 1
// ------------------------------

func isHappy(n int) bool {
	seen := []int{1}
	for {
		n = happySum(n)
		if contains(seen, n) {
			break
		} else {
			seen = append(seen, n)
		}
	}
	if n == 1 {
		return true
	}
	return false
}

func happySum(n int) int {
	sum := 0
	for n >= 10 {
		sum += int(math.Pow(float64(n%10), 2))
		n /= 10
	}
	return sum + int(math.Pow(float64(n), 2))
}

func contains(s []int, num int) bool {
	for _, n := range s {
		if n == num {
			return true
		}
	}
	return false
}

// LeetCodeMain02 contains leetcode algorithms.
func LeetCodeMain02() {
	if false {
		fmt.Println("\n#1. 位1的个数")
		fmt.Println("expect 3, actual:", hammingWeight(11))

		fmt.Println("\n#2. Fizz Buzz")
		fmt.Println("fizz buzz results:", fizzBuzz(15))

		fmt.Println("\n#3. 多数元素")
		fmt.Println("expect 3, actual:", majorityElement([]int{3, 3, 4}))
		fmt.Println("expect 2, actual:", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))

		fmt.Println("\n#4. 罗马数字转整数")
		fmt.Println("expect 58, actual:", romanToInt("LVIII"))
		fmt.Println("expect 1994, actual:", romanToInt("MCMXCIV"))

		fmt.Println("\n#5. 合并两个有序链表")
		listNodes1 := createListNodes([]int{1, 2, 4})
		listNodes2 := createListNodes([]int{1, 3, 4})
		fmt.Print("expect [1->1->2->3->4->4], actual: ")
		printListNodes(mergeTwoLists(listNodes1, listNodes2))

		fmt.Println("\n#6. 移动零")
		nums := []int{0, 1, 0, 3, 12}
		// moveZeroes01(nums)
		moveZeroes02(nums)
		fmt.Println("expect [1,3,12,0,0], actual:", nums)

		fmt.Println("\n#7. 快乐数")
		fmt.Println("expect true, actual:", isHappy(19))
	}

	fmt.Println("leetcode sample2 done.")
}
