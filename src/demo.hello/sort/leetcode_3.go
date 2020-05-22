package sort

import (
	"fmt"
	"strconv"
	"strings"
)

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

// ------------------------------
// #5. 回文链表
// ------------------------------

func isPalindromeLinkedList(head *listNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	cur1 := head
	cur2 := reverseList(slow.Next)
	for cur2 != nil { // cur2指向的链表长度较短
		if cur1.Val != cur2.Val {
			return false
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	return true
}

// ------------------------------
// #6. 链表中倒数第k个节点
// ------------------------------

func getKthFromEnd(head *listNode, k int) *listNode {
	slow := head
	fast := head

	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// ------------------------------
// #7. 替换空格
// 时间复杂度：O(n) 空间复杂度：O(n)
// ------------------------------

func replaceSpace(s string) string {
	ret := make([]rune, 0, len(s)*3)
	for _, r := range s {
		if r == ' ' {
			ret = append(ret, '%', '2', '0')
		} else {
			ret = append(ret, r)
		}
	}
	return string(ret)
}

// ------------------------------
// #8. 左旋转字符串
// 把字符串前面的若干个字符转移到字符串的尾部。
// ------------------------------

func reverseLeftWords01(s string, n int) string {
	b := []rune(s)
	for i := 0; i < n; i++ {
		for j := 0; j < len(b)-1; j++ {
			b[j], b[j+1] = b[j+1], b[j]
		}
	}
	return string(b)
}

func reverseLeftWords02(s string, n int) string {
	return s[n:] + s[:n]
}

// ------------------------------
// #9. 反转字符串中的单词
// ------------------------------

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	for i := 0; i < len(strs); i++ {
		strs[i] = reverseString(strs[i])
	}
	return strings.Join(strs, " ")
}

func reverseString(s string) string {
	b := []rune(s)
	start := 0
	end := len(s) - 1
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
	return string(b)
}

// ------------------------------
// #10. 回文整数
// ------------------------------

func isPalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}

	base := 1
	tmp := x / base
	for tmp >= 10 {
		base *= 10
		tmp = x / base
	}

	left := 0
	right := 0
	for x > 0 {
		left = x / base
		right = x % 10
		if left != right {
			return false
		}
		x = (x % base) / 10
		base /= 100
	}
	return true
}

// ------------------------------
// #11. 整数反转
// ------------------------------

func reverseNumber(x int) int {
	const maxInt = int(^uint(0) >> 1)

	ret := 0
	pop := 0
	for x != 0 {
		pop = x % 10
		ret = ret*10 + pop
		x /= 10
		// 如果反转后整数溢出那么就返回 0
		if maxInt/10-pop < ret {
			return 0
		}
	}
	return ret
}

// ------------------------------
// #12. 有效的括号
// ------------------------------

func isValidBrackets(s string) bool {
	st := &stack{
		slice: make([]int, len(s)),
		top:   0,
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			st.push(int(char))
		}
		if char == ')' || char == ']' || char == '}' {
			c, err := st.pop()
			if err != nil {
				return false
			}
			if (char == ')' && c != '(') || (char == ']' && c != '[') || (char == '}' && c != '{') {
				return false
			}
		}
	}
	return st.size() == 0
}

// ------------------------------
// #13. 最长公共前缀
// ------------------------------

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(strs[i]) < len(prefix) || strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1] // str[闭区间:开区间]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}

// ------------------------------
// #14. 字符串中的第一个唯一字符
// ------------------------------

func firstUniqChar(s string) int {
	var chars [26]int16
	for _, c := range s {
		idx := c - 'a'
		chars[idx]++
	}

	for i, c := range s {
		idx := c - 'a'
		if chars[idx] == 1 {
			return i
		}
	}
	return -1
}

// ------------------------------
// #15. 统计位数为偶数的数字
// ------------------------------

func findNumbers(nums []int) int {
	count := 0
	for _, num := range nums {
		str := strconv.Itoa(num)
		if len(str)%2 == 0 {
			count++
		}
	}
	return count
}

// ------------------------------
// #16. 判定是否互为字符重排
// 给定两个字符串 s1 和 s2, 请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。
// ------------------------------

func checkPermutation(s1 string, s2 string) bool {
	arr := make([]int, 128)
	for _, ch := range s1 {
		arr[ch]++
	}
	for _, ch := range s2 {
		arr[ch]--
	}

	for _, val := range arr {
		if val != 0 {
			return false
		}
	}
	return true
}

// ------------------------------
// #17. 移除重复节点
// 移除未排序链表中的重复节点，保留最开始出现的节点。
// 输入：[1, 2, 3, 3, 2, 1]
// 输出：[1, 2, 3]
// ------------------------------

func removeDuplicateNodes(head *listNode) *listNode {
	m := make(map[int]struct{}, 0)
	var pre *listNode
	cur := head
	for cur != nil {
		if _, ok := m[cur.Val]; ok {
			pre.Next = cur.Next
		} else {
			m[cur.Val] = struct{}{}
			pre = cur
		}
		cur = cur.Next
	}
	return head
}
