package sort

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// ------------------------------
// tree struct
// ------------------------------

type myBinTreeNode struct {
	value int
	left  *myBinTreeNode
	right *myBinTreeNode
}

// stack for treenodes
type treeNodesStack struct {
	nodes []*myBinTreeNode
	top   int
}

func (stack *treeNodesStack) init(cap int) {
	stack.nodes = make([]*myBinTreeNode, cap)
}

func (stack *treeNodesStack) add(node *myBinTreeNode) {
	// expect (by index): add(0) => [0]; pop() => [nil]; add(1) => [1]
	// error (by append): add(0) => [0]; pop() => [nil]; add(1) => [nil,1]
	// stack.nodes = append(stack.nodes, node)
	stack.nodes[stack.top] = node
	stack.top++
}

func (stack *treeNodesStack) pop() *myBinTreeNode {
	if stack.top <= 0 {
		panic("stack is empty!")
	}

	stack.top--
	node := stack.nodes[stack.top]
	stack.nodes[stack.top] = nil
	return node
}

func (stack *treeNodesStack) size() int {
	return stack.top
}

func (stack *treeNodesStack) toString() string {
	if stack.top <= 0 {
		return "[]"
	}

	ret := make([]string, 0, stack.top)
	for i := stack.top - 1; i >= 0; i-- {
		ret = append(ret, strconv.Itoa(stack.nodes[i].value))
	}
	return fmt.Sprintf("[%s]", strings.Join(ret, ","))
}

// ------------------------------
// #1. 从数组创建二叉树
// ------------------------------

func createBinaryTree(arr []int) *myBinTreeNode {
	tree := make([]*myBinTreeNode, 0, len(arr))
	for _, val := range arr {
		node := &myBinTreeNode{
			value: val,
		}
		tree = append(tree, node)
	}

	for i := 0; i < len(tree)/2; i++ {
		tree[i].left = tree[i*2+1]
		tmp := i*2 + 2
		if tmp < len(tree) {
			tree[i].right = tree[tmp]
		}
	}
	return tree[0]
}

// ------------------------------
// #2. 按层打印二叉树 从上往下 从左往右（先序遍历-递归）
// ------------------------------

func preOdderBinTree1(root *myBinTreeNode) {
	if root == nil {
		return
	}
	fmt.Print(strconv.Itoa(root.value) + ",")
	preOdderBinTree1(root.left)
	preOdderBinTree1(root.right)
}

// ------------------------------
// #3. 按层打印二叉树 从上往下 从左往右（先序遍历-非递归）
// ------------------------------

func preOdderBinTree2(root *myBinTreeNode) {
	stack := treeNodesStack{}
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

// ------------------------------
// #4. 中序遍历
// ------------------------------

func inOdderBinTree(root *myBinTreeNode) {
	if root == nil {
		return
	}
	inOdderBinTree(root.left)
	fmt.Print(strconv.Itoa(root.value) + ",")
	inOdderBinTree(root.right)
}

// ------------------------------
// #5. 二叉树的深度（递归）
// ------------------------------

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

// ------------------------------
// #6. 二叉树的深度（非递归）
// ------------------------------

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
	if val, ok := item.Value.(*myBinTreeNode); ok {
		return val
	}
	panic("invalid type, expect BinTreeNode!")
}

// TestTreeAlgorithms test for tree algorithms.
func TestTreeAlgorithms() {
	if false {
		fmt.Println("\n#1. 从数组创建二叉树")
		arr := make([]int, 0, 12)
		for i := 0; i < cap(arr); i++ {
			arr = append(arr, i)
		}
		root := createBinaryTree(arr)

		fmt.Println("\n#2. 按层打印二叉树 从上往下 从左往右（先序遍历-递归）")
		fmt.Println("tree nodes with pre-order:")
		preOdderBinTree1(root)
		fmt.Println()

		fmt.Println("\n#3. 按层打印二叉树 从上往下 从左往右（先序遍历-非递归）")
		fmt.Println("tree nodes with pre-order (by stack):")
		preOdderBinTree2(root)
		fmt.Println()

		fmt.Println("\n#4. 中序遍历")
		inOdderBinTree(root)
		fmt.Println()

		fmt.Println("\n#5. 二叉树的深度（递归）")
		fmt.Println("tree depth:", getBinTreeDepth1(root))

		fmt.Println("\n#6. 二叉树的深度（非递归）")
		fmt.Println("tree depth:", getBinTreeDepth2(root))
	}

	fmt.Println("tree algorithms done.")
}
