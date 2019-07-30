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

func (list *linkedList) insert(val int) {
	n := &node{
		value: val,
		next:  nil,
	}
	if list.head == nil {
		// list is empty
		list.head = n
		list.last = n
		return
	}
	if list.head.value > val {
		// insert head
		n.next = list.head
		list.head = n
		return
	}

	var base *node
	for tmp := list.head; tmp != nil; tmp = tmp.next {
		if tmp.next != nil && tmp.next.value > val {
			base = tmp
			break
		}
	}
	if base == nil {
		// insert tail
		list.last.next = n
		list.last = n
	} else {
		// insert mid
		n.next = base.next
		base.next = n
	}
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
	l := &linkedList{}
	for _, num := range []int{1, 16, 15, 7, 99, 50, 0, 99} {
		l.append(num)
	}
	fmt.Println("\nappend linked list values:", l.toString())

	l = &linkedList{}
	for _, num := range []int{1, 16, 15, 7, 99, 50, 0, 99} {
		l.insert(num)
	}
	fmt.Println("insert linked list values:", l.toString())
}
