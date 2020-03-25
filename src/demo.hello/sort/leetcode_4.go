package sort

import (
	"fmt"
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
// 给定一个包含大写和小写字母的字符串，找到通过这些字母构造成的最长的回文串
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
