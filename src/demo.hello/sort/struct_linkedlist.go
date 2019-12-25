package sort

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// ------------------------------
// #1. linkedlist struct
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
		return "null"
	}

	var ret []string
	for tmp := list.head; tmp != nil; tmp = tmp.next {
		ret = append(ret, strconv.Itoa(tmp.value))
	}
	return fmt.Sprintf("[%s]", strings.Join(ret, ","))
}

// ------------------------------
// linkedlist alg
// ------------------------------

type myNode struct {
	value int
	next  *myNode
}

func createLinkedList(nums []int) *myNode {
	head := &myNode{
		value: nums[0],
	}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.next = &myNode{
			value: nums[i],
		}
		cur = cur.next
	}

	return head
}

func createRandLinkedList(size, seed int) *myNode {
	head := &myNode{
		value: 0,
	}
	rand.Seed(int64(seed))
	cur := head
	for i := 1; i <= size; i++ {
		cur.next = &myNode{
			value: rand.Intn(100),
		}
		cur = cur.next
	}

	return head
}

func createCycleLinkedList(size, seed int) *myNode {
	head := &myNode{
		value: 0,
	}
	rand.Seed(123)
	cur := head
	for i := 1; i <= size; i++ {
		cur.next = &myNode{
			value: rand.Intn(100),
		}
		cur = cur.next
	}
	cur.next = head

	return head
}

func printLinkedList(head *myNode) {
	s := make([]string, 0, 20)
	for cur := head; cur != nil; cur = cur.next {
		s = append(s, strconv.Itoa(cur.value))
	}
	fmt.Printf("linkedlist: [%s]\n", strings.Join(s, ","))
}

// ------------------------------
// #2. 链表去重
// ------------------------------

func distinctLinkedList(head *myNode) {
	for cur := head; cur.next != nil; cur = cur.next {
		for node := cur; node.next != nil; {
			if node.next.value == cur.value {
				node.next = node.next.next
			} else {
				node = node.next
			}
		}
	}
}

// ------------------------------
// #3. 单链表排序 选择排序（交换值）
// ------------------------------

func sortLinkedList01(head *myNode) {
	for cNode := head; cNode.next != nil; cNode = cNode.next {
		for nNode := cNode.next; nNode != nil; nNode = nNode.next {
			if cNode.value > nNode.value { // 找到最小值
				cNode.value, nNode.value = nNode.value, cNode.value
			}
		}
	}
}

// ------------------------------
// #4. 单链表排序 冒泡排序（交换值）
// ------------------------------

func sortLinkedList02(head *myNode) {
	for node := head; node.next != nil; node = node.next {
		cNode := head
		for cNode.next != nil {
			nNode := cNode.next
			if cNode.value > nNode.value { // 比较相邻元素
				cNode.value, nNode.value = nNode.value, cNode.value
			}
			cNode = cNode.next
		}
	}
}

// ------------------------------
// #5. 单链表反转（数组）
// https://www.cnblogs.com/mafeng/p/7149980.html
// ------------------------------

func reverseLinkedList01(head *myNode) *myNode {
	nArr := make([]*myNode, 0, 10)
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

// ------------------------------
// #6. 单链表反转（交换结点）
// ------------------------------

func reverseLinkedList02(head *myNode) *myNode {
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
// #7. 链表中是否有环（两个指针）
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

// ------------------------------
// #8. 链表中是否有环（hash表）
// ------------------------------

// 通过hash表来检查一个结点此前是否被访问过来判断链表是否为环形链表。
// 过程：遍历所有结点并在hash表中存储每个结点引用（内存地址）。
// TODO:

// ------------------------------
// #9. 单链表排序 归并排序
// https://www.cnblogs.com/zhanghaiba/p/3534521.html
// ------------------------------

func linkedListMergeSort(head *myNode) *myNode {
	if head == nil || head.next == nil {
		return head
	}

	sub1, sub2 := divideLinkedList(head)
	sorted1 := linkedListMergeSort(sub1)
	sorted2 := linkedListMergeSort(sub2)
	return mergeLinkedList(sorted1, sorted2)
}

// 合并两个有序链表
func mergeLinkedList(head1, head2 *myNode) *myNode {
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

// 拆分链表
func divideLinkedList(head *myNode) (left, right *myNode) {
	slow := head
	fast := head.next
	for fast != nil {
		fast = fast.next
		if fast != nil {
			fast = fast.next
			slow = slow.next
		}
	}

	right = slow.next
	slow.next = nil
	left = head
	return
}

// TestLinkedListAlgorithms test for linkedlist algorithms.
func TestLinkedListAlgorithms() {
	if false {
		fmt.Println("\n#1. linkedlist struct")
		values := []int{1, 16, 15, 7, 99, 7, 50, 99, 0}
		l1 := &linkedList{}
		for _, val := range values {
			l1.append(val)
		}
		fmt.Println("append linked list values:", l1.toString())

		l2 := &linkedList{}
		for _, num := range values {
			l2.insert(num)
		}
		fmt.Println("insert linked list values:", l2.toString())

		fmt.Println("\n#2. 链表去重")
		head1 := createLinkedList([]int{1, 3, 7, 5, 6, 3, 3, 2, 5})
		fmt.Print("src ")
		printLinkedList(head1)
		fmt.Println("distinct ")
		distinctLinkedList(head1)
		printLinkedList(head1)

		fmt.Println("\n#3. 单链表排序 选择排序（交换值）")
		head2 := createRandLinkedList(10, 666)
		fmt.Print("src ")
		printLinkedList(head2)
		fmt.Printf("sorted ")
		sortLinkedList01(head2)
		printLinkedList(head2)

		fmt.Println("\n#4. 单链表排序 冒泡排序（交换值）")
		head3 := createRandLinkedList(10, 123)
		fmt.Print("src ")
		printLinkedList(head3)
		fmt.Printf("sorted ")
		sortLinkedList02(head3)
		printLinkedList(head3)

		fmt.Println("\n#5. 单链表反转（数组）")
		head4 := createRandLinkedList(10, 234)
		fmt.Print("src ")
		printLinkedList(head4)
		fmt.Print("reverse ")
		printLinkedList(reverseLinkedList01(head4))

		fmt.Println("\n#6. 单链表反转（交换结点）")
		head5 := createRandLinkedList(10, 234)
		fmt.Print("src ")
		printLinkedList(head5)
		fmt.Print("reverse ")
		printLinkedList(reverseLinkedList02(head5))

		fmt.Println("\n#7. 链表中是否有环（两个指针）")
		head6 := createRandLinkedList(10, 456)
		fmt.Println("recycle linkedlist:", isRecycleLinkedlist(head6))
		cycle := createCycleLinkedList(20, 456)
		fmt.Println("recycle linkedlist:", isRecycleLinkedlist(cycle))

		fmt.Println("\n#9.1 合并两个有序链表")
		head7 := createLinkedList([]int{3, 5, 7, 11, 15})
		head8 := createLinkedList([]int{4, 6, 8, 10, 14, 20})
		fmt.Println("src linkedlists:")
		printLinkedList(head7)
		printLinkedList(head8)
		fmt.Print("merged ")
		printLinkedList(mergeLinkedList(head7, head8))

		fmt.Println("\n#9.2 拆分链表")
		head9 := createRandLinkedList(10, 567)
		fmt.Print("src ")
		printLinkedList(head9)
		fmt.Println("divide linkedlists:")
		sub1, sub2 := divideLinkedList(head9)
		printLinkedList(sub1)
		printLinkedList(sub2)

		fmt.Println("\n#9. 单链表排序 归并排序")
		head10 := createLinkedList([]int{7, 3, 15, 5, 11, 6, 20, 10, 4, 14, 8})
		fmt.Print("src ")
		printLinkedList(head10)
		fmt.Print("sorted ")
		printLinkedList(linkedListMergeSort(head10))
	}

	fmt.Println("LinkedList algorithms done.")
}
