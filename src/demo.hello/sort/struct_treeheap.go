package sort

import (
	"fmt"
	"strings"
)

/**
 * 堆 完全二叉树
 * 1. 使用数组来实现
 * 节点的左子节点是(2*index+1) 节点的右子节点是(2*index+2) 节点的父节点是((index-1)/2)
 * 2. 堆中的每一个节点的值都大于（或等于）这个节点的子节点的值
 * 3. 插入和删除的时间复杂度都为O(logN)
 */

type treeHeap struct {
	heapArray []int
	cap       int
	size      int
}

func (tree *treeHeap) init(cap int) {
	tree.cap = cap
	tree.heapArray = make([]int, cap, cap)
}

func (tree *treeHeap) isEmpty() bool {
	return tree.size == 0
}

func (tree *treeHeap) isFull() bool {
	return tree.size == tree.cap
}

func (tree *treeHeap) insert(value int) {
	if tree.isFull() {
		panic("tree heap is full!")
	}
	tree.heapArray[tree.size] = value
	tree.trickleUp(tree.size)
	tree.size++
}

func (tree *treeHeap) trickleUp(index int) {
	pIndex := (index - 1) / 2 // parent index
	for index > 0 && tree.heapArray[pIndex] < tree.heapArray[index] {
		tree.heapArray[pIndex], tree.heapArray[index] = tree.heapArray[index], tree.heapArray[pIndex]
		index = pIndex
		pIndex = (pIndex - 1) / 2
	}
}

func (tree *treeHeap) remove() int {
	root := tree.heapArray[0]
	tree.size--
	tree.heapArray[0] = tree.heapArray[tree.size]
	tree.trickleDown(0)
	return root
}

func (tree *treeHeap) trickleDown(index int) {
	var largeChildIdx int
	for index < tree.size/2 {
		lChildIdx := index*2 + 1
		rChildIdx := index*2 + 2

		if rChildIdx < tree.size && tree.heapArray[rChildIdx] > tree.heapArray[lChildIdx] {
			largeChildIdx = rChildIdx
		} else {
			largeChildIdx = lChildIdx
		}
		if tree.heapArray[index] >= tree.heapArray[largeChildIdx] {
			break
		}
		tree.heapArray[index], tree.heapArray[largeChildIdx] =
			tree.heapArray[largeChildIdx], tree.heapArray[index]
		index = largeChildIdx
	}
}

func (tree *treeHeap) change(index, newValue int) {
	oldValue := tree.heapArray[index]
	if newValue == oldValue {
		return
	}

	tree.heapArray[index] = newValue
	if newValue > oldValue {
		tree.trickleUp(index)
		return
	}
	tree.trickleDown(index)
}

func (tree *treeHeap) toString() string {
	var rootValue, lValue, rValue int
	ret := make([]string, 0, tree.size/2)

	for index := 0; index < tree.size/2; index++ {
		rootValue = tree.heapArray[index]
		lValue = tree.heapArray[index*2+1]
		rValue = -1
		if (index*2 + 2) < tree.size {
			rValue = tree.heapArray[index*2+2]
		}
		ret = append(ret, fmt.Sprintf("[root:%d,lchild:%d,rchild:%d]", rootValue, lValue, rValue))
	}
	return fmt.Sprint(strings.Join(ret, "\n"))
}

// TestTreeHeap test for TreeHeap struct.
func TestTreeHeap() {
	const size = 10
	tree := treeHeap{}
	tree.init(size * 2)
	for i := 0; i < size; i++ {
		tree.insert(i)
	}
	fmt.Println("\ntree heap:\n", tree.toString())

	tree.insert(1)
	tree.insert(7)
	fmt.Println("\ninsert values, and tree heap:\n", tree.toString())

	fmt.Printf("\nremove %d from tree heap.\n", tree.remove())
	fmt.Println("remove values, and tree heap:\n", tree.toString())

	tree.change(3, 11)
	fmt.Println("\nchange values, and tree heap:\n", tree.toString())
}
