package sort

import (
	"fmt"
	"math"
	"strings"
)

// ------------------------------
// #1. 验证回文串
// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
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
// #2. 两数之和
// 给定一个整数数组 nums 和一个目标值 target, 请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
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

func createCycleListNodes(values []int, end int) *listNode {
	header := &listNode{
		Val:  values[0],
		Next: nil,
	}
	curNode := header
	var endNode *listNode

	for idx, val := range values[1:] {
		curNode.Next = &listNode{
			Val:  val,
			Next: nil,
		}
		if end == idx {
			endNode = curNode.Next
		}
		curNode = curNode.Next
	}
	curNode.Next = endNode
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

// ------------------------------
// #6. 反转链表
// ------------------------------

func reverseList(head *listNode) *listNode {
	// head为nil或只有一个元素的情况
	var pre *listNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// ------------------------------
// #7. 环形链表
// ------------------------------

func hasCycle(head *listNode) bool {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// ------------------------------
// #8. 颠倒二进制位（考虑使用位运算）
// ------------------------------

func reverseBits(num uint32) uint32 {
	size := 32
	bits := make([]byte, size)
	for i := size - 1; i >= 0 && num > 0; i-- {
		bits[i] = byte(num % 2)
		num /= 2
	}

	var retVal uint32
	for i := 0; i < size; i++ {
		if bits[i] == 1 {
			retVal += uint32(math.Pow(float64(2), float64(i)))
		}
	}
	return retVal
}

// ------------------------------
// #9. 实现 strStr()
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。
// 如果不存在，则返回 -1.
// 注：当 needle 是空字符串时我们应当返回 0. 这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
// ------------------------------

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		found := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}

// ------------------------------
// #10. 二叉树的最大深度
// ------------------------------

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func createBinTree(values []int) *treeNode {
	treeNodes := make([]*treeNode, len(values))
	for i := 0; i < len(values); i++ {
		treeNodes[i] = &treeNode{
			Val: values[i],
		}
	}

	for i := 0; i < len(values)/2; i++ {
		treeNodes[i].Left = treeNodes[i*2+1]
		if i*2+2 < len(values) {
			treeNodes[i].Right = treeNodes[i*2+2]
		}
	}
	return treeNodes[0]
}

func maxDepth(root *treeNode) int {
	if root == nil {
		return 0
	}

	lDepth := maxDepth(root.Left) + 1
	rDepth := maxDepth(root.Right) + 1
	if lDepth > rDepth {
		return lDepth
	}
	return rDepth
}

// ------------------------------
// #11. Excel表列序号
// A -> 1, AA -> 27
// ------------------------------

func titleToNumber(s string) int {
	base := 1
	retNum := 0
	for i := len(s) - 1; i >= 0; i-- {
		retNum += base * (int(s[i]) - 'A' + 1)
		base *= 26
	}
	return retNum
}

// ------------------------------
// #12. 合并两个有序数组
// 给定两个有序整数数组 nums1 和 nums2, 将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
// 输入：
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3
// 输出：[1,2,2,3,5,6]
// ------------------------------

func mergeSortedNums(nums1 []int, m int, nums2 []int, n int) {
	cur := len(nums1) - 1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[cur] = nums1[m-1]
			m--
		} else {
			nums1[cur] = nums2[n-1]
			n--
		}
		cur--
	}

	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
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

		fmt.Println("\n#6. 反转链表")
		listNodes2 := createListNodes([]int{1, 2, 3, 4, 5})
		fmt.Print("expect [5,4,3,2,1], actual: ")
		printListNodes(reverseList(listNodes2))

		fmt.Println("\n#7. 环形链表")
		listNode3 := createCycleListNodes([]int{3, 2, 0, -4}, 1)
		fmt.Println("expect true, actual:", hasCycle(listNode3))
		listNode3 = createCycleListNodes([]int{3, 2, 0, -4}, -1)
		fmt.Println("expect false, actual:", hasCycle(listNode3))

		fmt.Println("\n#8. 颠倒二进制位")
		fmt.Println("expect 964176192, actual:", reverseBits(43261596))

		fmt.Println("\n#9. 实现strStr()")
		fmt.Println("expect 2, actual: ", strStr("hello", "ll"))
		fmt.Println("expect -1, actual: ", strStr("aaaaa", "bba"))

		fmt.Println("\n#10. 二叉树的最大深度")
		fmt.Println("expect 3, actual:", maxDepth(createBinTree([]int{3, 9, 20, -1, -1, 15, 7})))

		fmt.Println("\n#11. Excel表列序号")
		fmt.Println("expect 28, actual:", titleToNumber("AB"))
		fmt.Println("expect 701, actual:", titleToNumber("ZY"))

		fmt.Println("\n#12. 合并两个有序数组")
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2, 5, 6}
		mergeSortedNums(nums1, 3, nums2, len(nums2))
		fmt.Println("expect [1,2,2,3,5,6], actual:", nums1)
	}

	fmt.Println("leetcode sample done.")
}
