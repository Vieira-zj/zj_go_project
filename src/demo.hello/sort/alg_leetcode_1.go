package sort

import (
	"fmt"
	"strings"
)

// ------------------------------
// #1. 验证回文串：给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写
// ------------------------------

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

// ------------------------------
// #2. 两数之和：给定一个整数数组 nums 和一个目标值 target, 请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标
// ------------------------------

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

// ------------------------------
// #3. 删除链表中的节点
// ------------------------------

type listNode struct {
	Val  int
	Next *listNode
}

// 输入参数只给定要求被删除的节点
func deleteNode(node *listNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func createListNodes(values []int) *listNode {
	header := &listNode{
		Val:  values[0],
		Next: nil,
	}
	curNode := header
	for _, val := range values[1:] {
		curNode.Next = &listNode{
			Val:  val,
			Next: nil,
		}
		curNode = curNode.Next
	}
	return header
}

func getListNodeByValue(header *listNode, value int) *listNode {
	for header != nil {
		if header.Val == value {
			return header
		}
		header = header.Next
	}
	return nil
}

func printListNodes(header *listNode) {
	for header.Next != nil {
		fmt.Printf("%d, ", header.Val)
		header = header.Next
	}
	fmt.Println(header.Val)
}

// ------------------------------
// #4. 转换成小写字母
// ------------------------------

func toLowerCase(str string) string {
	retStr := ""
	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			c = c + 'a' - 'A'
		}
		retStr += string(c)
	}
	return retStr
}

// ------------------------------
// #5. 分割平衡字符串
// ------------------------------

type stack struct {
	slice []int
	top   int
}

func (s *stack) size() int {
	return s.top
}

func (s *stack) getTop() (int, error) {
	if s.size() == 0 {
		return -1, fmt.Errorf("stack is empty")
	}
	return s.slice[s.top-1], nil
}

func (s *stack) push(val int) {
	s.top++
	s.slice[s.top-1] = val
}

func (s *stack) pop() (int, error) {
	if s.size() == 0 {
		return -1, fmt.Errorf("stack is empty")
	}
	retVal := s.slice[s.top-1]
	s.top--
	return retVal, nil
}

func balancedStringSplit(s string) int {
	count := 0
	st := &stack{
		slice: make([]int, len(s)),
		top:   0,
	}

	for _, c := range s {
		if val, err := st.getTop(); err != nil { // stack empty
			st.push(int(c))
		} else {
			if val == int(c) {
				st.push(int(c))
			} else {
				st.pop()
				if st.size() == 0 {
					count++
				}
			}
		}
	}
	return count
}

// LeetCodeMain contains leetcode algorithms.
func LeetCodeMain() {
	if false {
		fmt.Println("\n#1. 验证回文串")
		fmt.Println("excpect true, actual:", isPalindrome("A man, a plan, a canal: Panama"))
		fmt.Println("excpect false, actual:", isPalindrome("race a car"))

		fmt.Println("\n#2.1 两数之和")
		fmt.Println("expect [0,1], actual:", twoSum01([]int{2, 7, 11, 15}, 9))
		fmt.Println("expect [0,2], actual:", twoSum01([]int{-3, 4, 3, 90}, 0))

		fmt.Println("\n#2.2 两数之和")
		fmt.Println("expect [1,2], actual:", twoSum02([]int{3, 2, 4}, 6))
		fmt.Println("expect [0,2], actual:", twoSum02([]int{-3, 4, 3, 90}, 0))

		fmt.Println("\n#3. 删除链表中的节点")
		listNodes := createListNodes([]int{4, 5, 1, 9})
		deleteNode(getListNodeByValue(listNodes, 1))
		fmt.Print("expect [4,5,9], actual: ")
		printListNodes(listNodes)

		fmt.Println("\n#4. 转换成小写字母")
		fmt.Println("expect 'hello', actual:", toLowerCase("Hello"))

		fmt.Println("\n#5. 分割平衡字符串")
		fmt.Println("expect 4, actual:", balancedStringSplit("RLRRLLRLRL"))
		fmt.Println("expect 3, actual:", balancedStringSplit("RLLLLRRRLR"))
		fmt.Println("expect 1, actual:", balancedStringSplit("LLLLRRRR"))
	}

	fmt.Println("leetcode sample done.")
}
