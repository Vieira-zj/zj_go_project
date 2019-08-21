package sort

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type myNode struct {
	value int
	next  *myNode
}

func printLinkedList(head *myNode) {
	s := make([]string, 0, 10)
	for cur := head; cur != nil; cur = cur.next {
		s = append(s, strconv.Itoa(cur.value))
	}
	fmt.Printf("linked list: [%s]\n", strings.Join(s, ","))
}

// 链表去重
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

// 单链表排序 选择排序（交换值）
func linkedListSort01(head *myNode) {
	for cNode := head; cNode.next != nil; cNode = cNode.next {
		for nNode := cNode.next; nNode != nil; nNode = nNode.next {
			if cNode.value > nNode.value { // 找到最小值
				cNode.value, nNode.value = nNode.value, cNode.value
			}
		}
	}
}

// 单链表排序 冒泡排序（交换值）
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

// 单链表排序 归并排序
func linkedListSort03(head *myNode) {
	// TODO:
}

// 单链表反转（交换结点）
// https://www.cnblogs.com/mafeng/p/7149980.html
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

// 单链表反转（交换结点）
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

// 链表中是否有环（两个指针）
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

// TestLinkedListAlgorithms test for linkedlist algorithms.
func TestLinkedListAlgorithms() {
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
	printLinkedList(head)
	fmt.Println("distinct:")
	distinctLinkedList(head)
	printLinkedList(head)
	fmt.Println()

	// #2 sort
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
	fmt.Println("sorted:")
	// linkedListSort01(head)
	linkedListSort02(head)
	printLinkedList(head)
	fmt.Println()

	fmt.Println("reverse:")
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
	fmt.Println("\nis recycle linkedlist:", isRecycleLinkedlist(head))
	fmt.Println("is recycle linkedlist:", isRecycleLinkedlist(cycle))
}
