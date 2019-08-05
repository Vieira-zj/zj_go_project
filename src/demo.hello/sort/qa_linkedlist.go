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
	for cur := head; cur.next != nil; cur = cur.next {
		s = append(s, strconv.Itoa(cur.value))
	}
	fmt.Printf("linked list: [%s]\n", strings.Join(s, ","))
}

// 单链表排序 选择排序
func linkedListSort01(head *myNode) {
	for curNode := head; curNode.next != nil; curNode = curNode.next {
		for nextNode := curNode.next; nextNode.next != nil; nextNode = nextNode.next {
			if curNode.value > nextNode.value { // 找到最小值
				curNode.value, nextNode.value = nextNode.value, curNode.value
			}
		}
	}
}

// 单链表排序 冒泡排序
func linkedListSort02(head *myNode) {
	for node := head; node.next != nil; node = node.next {
		curNode := head
		for nextNode := head.next; nextNode.next != nil; nextNode = nextNode.next {
			if curNode.value > nextNode.value {
				curNode.value, nextNode.value = nextNode.value, curNode.value
			}
			curNode = nextNode
		}
	}
}

// 单链表反转
func linkedListReverse(head *myNode) {
	// TODO:
}

// TestLinkedListAlgorithms test for linkedlist algorithms.
func TestLinkedListAlgorithms() {
	head := &myNode{
		value: 0,
	}

	rand.Seed(7891)
	cur := head
	for i := 1; i <= 10; i++ {
		new := &myNode{
			value: rand.Intn(100),
		}
		cur.next = new
		cur = new
	}

	printLinkedList(head)
	fmt.Println("sorted:")
	// linkedListSort01(head)
	linkedListSort02(head)
	printLinkedList(head)
}
