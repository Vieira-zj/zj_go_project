package sort

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	next  *node
}

type linkedList struct {
	head *node
	last *node
}

func (list *linkedList) append(val int) {
	n := &node{
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
	n := &node{
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

// TestLinkedList test for linked list.
func TestLinkedList() {
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
