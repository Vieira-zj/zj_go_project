package sort

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type myBinTreeNode struct {
	value int
	left  *myBinTreeNode
	right *myBinTreeNode
}

// 从数组创建二叉树
func createBinaryTree(arr []int) *myBinTreeNode {
	tree := make([]myBinTreeNode, 0, len(arr))
	for _, val := range arr {
		node := myBinTreeNode{
			value: val,
		}
		tree = append(tree, node)
	}

	for i := 0; i < len(tree)/2; i++ {
		tree[i].left = &tree[i*2+1]
		tmp := i*2 + 2
		if tmp < len(tree) {
			tree[i].right = &tree[tmp]
		}
	}
	return &tree[0]
}

// 按层打印二叉树 从上往下 从左往右（先序遍历-递归）
func preOdderBinTree1(root *myBinTreeNode) {
	if root == nil {
		return
	}
	fmt.Print(strconv.Itoa(root.value) + ",")
	preOdderBinTree1(root.left)
	preOdderBinTree1(root.right)
}

// 按层打印二叉树 从上往下 从左往右（先序遍历-非递归）
func preOdderBinTree2(root *myBinTreeNode) {
	stack := treeNodeStack{}
	stack.init(30)
	stack.add(root)

	for stack.size() > 0 {
		node := stack.pop()
		fmt.Print(strconv.Itoa(node.value) + ",")
		if node.right != nil {
			stack.add(node.right)
		}
		if node.left != nil {
			stack.add(node.left)
		}
	}
}

type treeNodeStack struct {
	nodes []*myBinTreeNode
	top   int
}

func (stack *treeNodeStack) init(cap int) {
	stack.nodes = make([]*myBinTreeNode, cap, cap)
}

func (stack *treeNodeStack) size() int {
	return stack.top
}

func (stack *treeNodeStack) toString() string {
	if stack.top <= 0 {
		return "[]"
	}

	ret := make([]string, 0, stack.top)
	for i := stack.top - 1; i >= 0; i-- {
		ret = append(ret, strconv.Itoa(stack.nodes[i].value))
	}
	return fmt.Sprintf("[%s]", strings.Join(ret, ","))
}

func (stack *treeNodeStack) add(node *myBinTreeNode) {
	stack.nodes[stack.top] = node
	stack.top++
}

func (stack *treeNodeStack) pop() *myBinTreeNode {
	if stack.top <= 0 {
		panic("stack is empty!")
	}

	node := stack.nodes[stack.top-1]
	stack.nodes[stack.top-1] = nil
	stack.top--
	return node
}

// 中序遍历
func inOdderBinTree(root *myBinTreeNode) {
	if root == nil {
		return
	}
	inOdderBinTree(root.left)
	fmt.Print(strconv.Itoa(root.value) + ",")
	inOdderBinTree(root.right)
}

// 二叉树的深度（递归）
func getBinTreeDepth1(root *myBinTreeNode) int {
	if root == nil {
		return 0
	}
	depth1 := getBinTreeDepth1(root.left)
	depth2 := getBinTreeDepth1(root.right)
	if depth1 > depth2 {
		return depth1 + 1
	}
	return depth2 + 1
}

// 二叉树的深度（非递归）
func getBinTreeDepth2(root *myBinTreeNode) int {
	var depth int
	queue := list.New()
	queue.PushBack(root) // queue保存每一层的所有节点

	for queue.Len() != 0 {
		depth++
		count := queue.Len()
		for i := 0; i < count; i++ {
			node := popFront(queue)
			if node.left != nil {
				queue.PushBack(node.left)
			}
			if node.right != nil {
				queue.PushBack(node.right)
			}
		}
	}
	return depth
}

func popFront(queue *list.List) *myBinTreeNode {
	item := queue.Front()
	queue.Remove(item)
	return item.Value.(*myBinTreeNode)
}

// TestTreeAlgorithms test for tree algorithms.
func TestTreeAlgorithms() {
	arr := make([]int, 0, 12)
	for i := 0; i < cap(arr); i++ {
		arr = append(arr, i)
	}
	root := createBinaryTree(arr)

	// #1
	fmt.Println("\n#1: tree nodes with pre-order:")
	preOdderBinTree1(root)
	fmt.Println()

	fmt.Println("#2: tree nodes with pre-order:")
	preOdderBinTree2(root)
	fmt.Println()

	// #2
	fmt.Println("\ntree nodes with in-order:")
	inOdderBinTree(root)
	fmt.Println()

	// #3
	fmt.Println("\n#1: tree depth:", getBinTreeDepth1(root))
	fmt.Println("#2: tree depth:", getBinTreeDepth2(root))
}
