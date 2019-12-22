package sort

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// ------------------------------
// linkedlist struct
// ------------------------------

type lnode struct {
	value int
	next  *lnode
}

type linkedList struct {
	head *lnode
	last *lnode
}

func (list *linkedList) append(val int) {
	n := &lnode{
		value: val,
		next:  nil,
	}

	if list.head == nil {
		list.head = n
	} else {
		list.last.next = n
	}
	list.last = n
}

// 插入节点后保持有序
func (list *linkedList) insert(val int) {
	n := &lnode{
		value: val,
		next:  nil,
	}
	// list is empty
	if list.head == nil {
		list.head = n
		list.last = n
		return
	}

	// insert at head
	if list.head.value > val {
		n.next = list.head
		list.head = n
		return
	}

	// insert at mid
	for tmp := list.head; tmp.next != nil; tmp = tmp.next {
		if tmp.next.value > val {
			n.next = tmp.next
			tmp.next = n
			return
		}
	}

	// insert at tail
	list.last.next = n
	list.last = n
}

func (list *linkedList) toString() string {
	if list.head == nil {
		return "nil"
	}

	var ret []string
	for tmp := list.head; tmp != nil; tmp = tmp.next {
		ret = append(ret, strconv.Itoa(tmp.value))
	}
	return strings.Join(ret, ",")
}

// ------------------------------
// linkedlist alg
// ------------------------------

type myNode struct {
	value int
	next  *myNode
}

func printLinkedList(head *myNode) {
	s := make([]string, 0, 20)
	for cur := head; cur != nil; cur = cur.next {
		s = append(s, strconv.Itoa(cur.value))
	}
	fmt.Printf("linked list: [%s]\n", strings.Join(s, ","))
}

// ------------------------------
// 链表去重
// ------------------------------

func distinctLinkedList(head *myNode) {
	for cur := head; cur.next != nil; cur = cur.next {
		for node := cur; node.next != nil; node = node.next {
			if node.next.value == cur.value {
				node.next = node.next.next
				if node.next == nil {
					break
				}
			}
		}
	}
}

// ------------------------------
// 单链表排序 选择排序（交换值）
// ------------------------------

func linkedListSort01(head *myNode) {
	for cNode := head; cNode.next != nil; cNode = cNode.next {
		for nNode := cNode.next; nNode != nil; nNode = nNode.next {
			if cNode.value > nNode.value { // 找到最小值
				cNode.value, nNode.value = nNode.value, cNode.value
			}
		}
	}
}

// ------------------------------
// 单链表排序 冒泡排序（交换值）
// ------------------------------

func linkedListSort02(head *myNode) {
	for node := head; node.next != nil; node = node.next {
		cNode := head
		nNode := head.next
		for cNode.next != nil {
			if cNode.value > nNode.value { // 比较相邻元素
				cNode.value, nNode.value = nNode.value, cNode.value
			}
			cNode = nNode
			nNode = nNode.next
		}
	}
}

// ------------------------------
// 单链表反转（交换结点）
// https://www.cnblogs.com/mafeng/p/7149980.html
// ------------------------------

func linkedListReverse01(head *myNode) *myNode {
	nArr := make([]*myNode, 0, 10) // node array
	for cur := head; cur != nil; cur = cur.next {
		nArr = append(nArr, cur)
	}

	last := len(nArr) - 1
	for i := last; i > 0; i-- {
		nArr[i].next = nArr[i-1]
	}
	nArr[0].next = nil
	return nArr[last]
}

func linkedListReverse02(head *myNode) *myNode {
	var nNode *myNode  // next node
	pNode := head      // pre node
	cNode := head.next // current node
	head.next = nil

	for cNode != nil {
		nNode = cNode.next
		cNode.next = pNode
		// move next
		pNode = cNode
		cNode = nNode
	}
	return pNode
}

// ------------------------------
// 链表中是否有环（两个指针）
// ------------------------------

func isRecycleLinkedlist(head *myNode) bool {
	// 快慢指针, 定义p,q两个指针, p指针每次向前走1步, q每次向前走2步, 若在某个时刻出现p==q, 则存在环
	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

// 链表中是否有环（hash表）
// 通过hash表来检查一个结点此前是否被访问过来判断链表是否为环形链表。
// 过程：遍历所有结点并在hash表中存储每个结点引用（内存地址）。

// ------------------------------
// 单链表排序 归并排序
// https://www.cnblogs.com/zhanghaiba/p/3534521.html
// ------------------------------

func linkedListMergeSort(head *myNode) *myNode {
	if head == nil || head.next == nil {
		return head
	}

	sub1, sub2 := linkedListDivide(head)
	sorted1 := linkedListMergeSort(sub1)
	sorted2 := linkedListMergeSort(sub2)
	return linkedListMerge(sorted1, sorted2)
}

// 合并两个有序链表
func linkedListMerge(head1, head2 *myNode) *myNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	retHead := &myNode{value: -1}
	curNode := retHead
	// note: input src linkedList "head1" and "head2" are changed
	for ; head1 != nil && head2 != nil; curNode = curNode.next {
		if head1.value < head2.value {
			curNode.next = head1
			head1 = head1.next
		} else {
			curNode.next = head2
			head2 = head2.next
		}
	}

	if head1 != nil {
		curNode.next = head1
	}
	if head2 != nil {
		curNode.next = head2
	}
	return retHead.next
}

func linkedListDivide(head *myNode) (left, right *myNode) {
	slow := head
	fast := head.next
	for fast != nil {
		fast = fast.next
		if fast != nil {
			fast = fast.next
			slow = slow.next
		}
	}

	left = head
	right = slow.next
	slow.next = nil
	return
}

// TestLinkedListAlgorithms test for linkedlist algorithms.
func TestLinkedListAlgorithms() {
	if false {
		fmt.Println("\n# linkedlist struct")
		values := []int{1, 16, 15, 7, 99, 7, 50, 99, 0}
		l := &linkedList{}
		for _, val := range values {
			l.append(val)
		}
		fmt.Println("\nappend linked list values:", l.toString())

		l = &linkedList{}
		for _, num := range values {
			l.insert(num)
		}
		fmt.Println("insert linked list values:", l.toString())
	}

	// #1
	head := &myNode{
		value: 1,
	}
	cur := head
	for _, i := range []int{1, 3, 7, 5, 6, 3, 3, 2, 5} {
		cur.next = &myNode{
			value: i,
		}
		cur = cur.next
	}
	fmt.Println("\nsrc linkedlist:")
	printLinkedList(head)
	fmt.Println("linkedlist distinct:")
	distinctLinkedList(head)
	printLinkedList(head)
	fmt.Println()

	// #2-1
	head = &myNode{
		value: 0,
	}
	rand.Seed(666)
	cur = head
	for i := 1; i <= 10; i++ {
		cur.next = &myNode{
			value: rand.Intn(100),
		}
		cur = cur.next
	}

	printLinkedList(head)
	fmt.Println("linkedlist sort:")
	// linkedListSort01(head)
	linkedListSort02(head)
	printLinkedList(head)
	fmt.Println()

	// #2-2
	fmt.Println("\nlinkedlist reverse:")
	// printLinkedList(linkedListReverse01(head))
	printLinkedList(linkedListReverse02(head))

	// #3
	cycle := &myNode{
		value: 0,
	}
	rand.Seed(123)
	cur = cycle
	for i := 1; i <= 20; i++ {
		cur.next = &myNode{
			value: rand.Intn(100),
		}
		cur = cur.next
	}
	cur.next = cycle
	fmt.Println("\nrecycle linkedlist:", isRecycleLinkedlist(head))
	fmt.Println("recycle linkedlist:", isRecycleLinkedlist(cycle))

	// #4-1
	head1 := &myNode{
		value: 1,
	}
	cur = head1
	for _, i := range []int{3, 5, 7, 11, 15} {
		cur.next = &myNode{
			value: i,
		}
		cur = cur.next
	}
	head2 := &myNode{
		value: 2,
	}
	cur = head2
	for _, i := range []int{4, 6, 8, 10, 14, 20} {
		cur.next = &myNode{
			value: i,
		}
		cur = cur.next
	}
	fmt.Println("\nsorted linkedlist merge:")
	printLinkedList(linkedListMerge(head1, head2))

	// #4-2
	fmt.Println("\nsrc linkedlist:")
	printLinkedList(head1)
	fmt.Println("linkedlist divide:")
	sub1, sub2 := linkedListDivide(head1)
	printLinkedList(sub1)
	printLinkedList(sub2)

	// #4-3 merge-sort linkedlist
	head = &myNode{
		value: 1,
	}
	cur = head
	for _, i := range []int{7, 3, 15, 5, 11} {
		cur.next = &myNode{
			value: i,
		}
		cur = cur.next
	}
	for _, i := range []int{6, 20, 10, 4, 14, 8} {
		cur.next = &myNode{
			value: i,
		}
		cur = cur.next
	}
	fmt.Println("\nmerge-sort linkedlist:")
	printLinkedList(linkedListMergeSort(head))
}
