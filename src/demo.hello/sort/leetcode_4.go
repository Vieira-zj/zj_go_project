package sort

import (
	"container/list"
	"fmt"
	"sort"
	"strconv"
)

// ------------------------------
// 1. 字符串压缩
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

// ------------------------------
// 双端队列
// ------------------------------

type dequeNode struct {
	val  int
	pre  *dequeNode
	next *dequeNode
}

// MyDeque 双端队列
type MyDeque struct {
	length int
	head   *dequeNode
	tail   *dequeNode
}

// NewDeque returns an instance of deque.
func NewDeque() *MyDeque {
	return &MyDeque{
		length: 0,
	}
}

// Size returns length of queue.
func (q *MyDeque) Size() int {
	return q.length
}

// GetHead returns value of queue head.
func (q *MyDeque) GetHead() int {
	if q.Size() == 0 {
		return -1
	}
	return q.head.val
}

// GetTail returns value of queue tail.
func (q *MyDeque) GetTail() int {
	if q.Size() == 0 {
		return -1
	}
	return q.tail.val
}

// PushHead push an element at head of queue.
func (q *MyDeque) PushHead(value int) {
	node := &dequeNode{
		val: value,
	}
	if q.Size() == 0 {
		q.head = node
		q.tail = node
	} else {
		node.next = q.head
		q.head.pre = node
		q.head = node
	}
	q.length++
}

// PushTail push an element at tail of queue.
func (q *MyDeque) PushTail(value int) {
	node := &dequeNode{
		val: value,
	}
	if q.Size() == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		node.pre = q.tail
		q.tail = node
	}
	q.length++
}

// PopHead pop an element from head of queue.
func (q *MyDeque) PopHead() int {
	ret := -1
	if q.Size() == 0 {
		return ret
	}

	if q.Size() == 1 {
		ret = q.head.val
	} else {
		ret = q.head.val
		q.head = q.head.next
		q.head.pre = nil
	}
	q.length--
	return ret
}

// PopTail pop an element from tail of queue.
func (q *MyDeque) PopTail() int {
	ret := -1
	if q.Size() == 0 {
		return ret
	}

	if q.Size() == 1 {
		ret = q.tail.val
	} else {
		ret = q.tail.val
		q.tail = q.tail.pre
		q.tail.next = nil
	}
	q.length--
	return ret
}

// Print show queue elements from head to tail.
func (q *MyDeque) Print() {
	cur := q.head
	for cur != nil {
		fmt.Printf("%d,", cur.val)
		cur = cur.next
	}
	fmt.Println()
}

// PrintReverse show queue elements from tail to head.
func (q *MyDeque) PrintReverse() {
	cur := q.tail
	for cur != nil {
		fmt.Printf("%d,", cur.val)
		cur = cur.pre
	}
	fmt.Println()
}

// ------------------------------
// 2. 队列的最大值
// 定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value, push_back 和 pop_front 的均摊时间复杂度都是O(1)
// 若队列为空，pop_front 和 max_value 需要返回 -1
// ------------------------------

type queueNode struct {
	val  int
	next *queueNode
}

// MaxQueue returns max value in queue.
type MaxQueue struct {
	head   *queueNode
	tail   *queueNode
	deque  *MyDeque
	length int
}

// NewMaxQueue returns an instance of MaxQueue.
func NewMaxQueue() MaxQueue {
	return MaxQueue{
		length: 0,
		deque:  NewDeque(),
	}
}

// MaxValue returns max value of queue.
func (q *MaxQueue) MaxValue() int {
	if q.length == 0 {
		return -1
	}
	return q.deque.GetHead()
}

// PushBack push element back of queue.
func (q *MaxQueue) PushBack(value int) {
	node := &queueNode{
		val: value,
	}

	if q.length == 0 {
		q.head = node
		q.tail = node
		q.deque.PushTail(value)
	} else {
		q.tail.next = node
		q.tail = node
		for q.deque.Size() > 0 && q.deque.GetTail() < value {
			q.deque.PopTail()
		}
		q.deque.PushTail(value)
	}
	q.length++
}

// PopFront pop element from front of queue.
func (q *MaxQueue) PopFront() int {
	if q.length == 0 {
		return -1
	}

	node := q.head
	q.head = q.head.next
	q.length--
	if node.val == q.deque.GetHead() {
		q.deque.PopHead()
	}
	return node.val
}

// ------------------------------
// 3. 最长回文串
// 给定一个包含大写和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
// 输入: "abbccccddd"
// 输出: 9
// ------------------------------

func longestPalindrome(s string) int {
	m := make(map[rune]int, 52)
	for _, r := range s {
		m[r]++
	}

	ret := 0
	for _, v := range m {
		ret += v - v%2
	}

	if ret < len(s) {
		return ret + 1
	}
	return ret
}

// ------------------------------
// 4. 链表的中间结点
// 给定一个带有头结点 head 的非空单链表，返回链表的中间结点。
// 如果有两个中间结点，则返回第二个中间结点。
// ------------------------------

func middleNode(head *listNode) *listNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// ------------------------------
// 5. 二叉树的直径
// 一棵二叉树的直径长度是任意两个结点路径长度中的最大值。
// 这条路径可能穿过也可能不穿过根结点。
// ------------------------------

var ans int

func diameterOfBinaryTree(root *treeNode) int {
	maxTreeDepth(root)
	ret := ans
	ans = 0
	return ret
}

// 最大深度
func maxTreeDepth(root *treeNode) int {
	if root == nil {
		return 0
	}

	left := maxTreeDepth(root.Left) + 1
	right := maxTreeDepth(root.Right) + 1
	ans = maxInt(left+right-2, ans)
	return maxInt(left, right)
}

// ------------------------------
// 6. 找到所有数组中消失的数字
// 给定一个范围在  1 <= a[i] <= n ( n = 数组大小 ) 的整型数组，数组中的元素一些出现了2次，另一些只出现1次。
// 找到所有在 [1, n] 范围之间没有出现在数组中的数字。
// 输入: [4,3,2,7,8,2,3,1]
// 输出: [5,6]
// ------------------------------

func findDisappearedNumbers(nums []int) []int {
	idx := 0
	for _, num := range nums {
		if num < 0 {
			idx = (-num) - 1
		} else {
			idx = num - 1
		}
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	ret := make([]int, 0)
	for idx, num := range nums {
		if num > 0 {
			ret = append(ret, idx+1)
		}
	}
	return ret
}

// ------------------------------
// 7. 最短无序连续子数组
// 输入: [2, 6, 4, 8, 10, 9, 15]
// 输出: 5
// 解释: 你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。
// ------------------------------

func findUnsortedSubarray(nums []int) int {
	copied := make([]int, len(nums))
	copy(copied, nums)
	sort.Ints(copied)
	fmt.Println("src:", nums)
	fmt.Println("sorted:", copied)

	start := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != copied[i] {
			start = i
			break
		}
	}
	if start == -1 {
		return 0
	}

	end := -1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != copied[i] {
			end = i
			break
		}
	}
	return end - start + 1
}

// ------------------------------
// 8. 二叉搜索树的第k大节点
// 二叉搜索树：
// 1）若任意节点的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
// 2）若任意节点的右子树不空，则右子树上所有节点的值均大于它的根节点的值；
// ------------------------------

func kthLargest(root *myBinTreeNode, k int) int {
	search := &kthSearch{
		k: k,
	}
	search.recursion(root)
	return search.kthMax
}

type kthSearch struct {
	k      int
	kthMax int
}

func (search *kthSearch) recursion(root *myBinTreeNode) {
	if root == nil {
		return
	}

	search.recursion(root.right)
	if search.k == 0 {
		return
	}
	search.k--
	if search.k == 0 {
		search.kthMax = root.value
	}
	search.recursion(root.left)
}

// ------------------------------
// 9. 判定字符是否唯一
// ------------------------------

func isUnique(astr string) bool {
	// ASCII码字符个数为128个
	var arr [128]int
	for _, rune := range astr {
		arr[rune]++
		if arr[rune] >= 2 {
			return false
		}
	}
	return true
}

// ------------------------------
// 10. 从尾到头打印链表
// 输入：head = [1,3,2]
// 输出：[2,3,1]
// ------------------------------

func reversePrint01(head *listNode) []int {
	if head == nil {
		return nil
	}

	stack := list.New()
	for ; head != nil; head = head.Next {
		stack.PushFront(head.Val)
	}

	ret := []int{}
	for e := stack.Front(); e != nil; e = e.Next() {
		ret = append(ret, e.Value.(int))
	}
	return ret
}

func reversePrint02(head *listNode) []int {
	if head == nil {
		return nil
	}

	arr := make([]int, 0)
	for ; head != nil; head = head.Next {
		arr = append(arr, head.Val)
	}

	for start, end := 0, len(arr)-1; start < end; {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}

func reversePrint03(head *listNode) []int {
	if head == nil {
		return nil
	}

	count := 0
	cur := head
	for ; head != nil; head = head.Next {
		count++
	}

	ret := make([]int, count)
	fmt.Println(count)
	for i := count - 1; i >= 0; i-- {
		ret[i] = cur.Val
		cur = cur.Next
	}
	return ret
}

// ------------------------------
// 11. 按既定顺序创建目标数组
// 目标数组 target 最初为空。
// 按从左到右的顺序依次读取 nums[i] 和 index[i], 在 target 数组中的下标 index[i] 处插入值 nums[i].
//
// 0 <= nums[i] <= 100
// 0 <= index[i] <= i
// ------------------------------

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		target[i] = -1
	}

	for i := 0; i < len(nums); i++ {
		key := index[i]
		val := nums[i]
		if i == key {
			target[key] = val
			continue
		}
		if target[key] == -1 {
			target[key] = val
			continue
		}

		for j := len(nums) - 1; j >= key+1; j-- {
			target[j] = target[j-1]
		}
		target[key] = val
	}
	return target
}

// ------------------------------
// 12. 访问所有点的最小时间
// 1）每一秒沿水平或者竖直方向移动一个单位长度，或者跨过对角线；
// 2）必须按照数组中出现的顺序来访问这些点。
// ------------------------------

func minTimeToVisitAllPoints(points [][]int) int {
	var ans int
	for i := 0; i < len(points)-1; i++ {
		p1x := points[i][0]
		p1y := points[i][1]
		p2x := points[i+1][0]
		p2y := points[i+1][1]
		ans += maxInt(absInt(p1x-p2x), absInt(p1y-p2y))
	}
	return ans
}

func absInt(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
