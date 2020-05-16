package sort

import (
	"container/list"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ------------------------------
// #1. n个球，可以着色为白色或黑色。连接3个以上（包含3个）球相同着色视为非法。有多少种着色可能？
// 使用二叉树？
// ------------------------------

func getCombinationCnt(n int) int {
	ret := 0
	all := math.Pow(float64(2), float64(n))
	if n < 3 {
		return int(all)
	}

	for i := 0; i < int(all); i++ {
		bits := octToBinary(i, n)
		fmt.Println(bits)
		pass := true
		for x := 0; x < len(bits)-2; x++ {
			if bits[x] == bits[x+1] && bits[x+1] == bits[x+2] {
				pass = false
				break
			}
		}
		if pass {
			ret++
		}
	}
	return ret
}

func octToBinary(n, size int) string {
	bits := make([]int, size)
	for i := size - 1; n > 0; i-- {
		bits[i] = n % 2
		n /= 2
	}

	var ret string
	for _, val := range bits {
		ret += strconv.Itoa(val)
	}
	return ret
}

// ------------------------------
// #2. 输入一个有序数组：[1,2,3,5,6,7,10,13]
// 输出：1-3,5-7,10,13
// ------------------------------

func formatByRange(arr []int) string {
	var nums []string
	start := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != 1 {
			if start == arr[i] {
				nums = append(nums, strconv.Itoa(start))
			} else {
				nums = append(nums, strconv.Itoa(start)+"-"+strconv.Itoa(arr[i]))
			}
			start = arr[i+1]
		}
	}

	if start == arr[len(arr)-1] {
		nums = append(nums, strconv.Itoa(start))
	} else {
		nums = append(nums, strconv.Itoa(start)+"-"+strconv.Itoa(arr[len(arr)-1]))
	}
	return strings.Join(nums, ",")
}

// ------------------------------
// #3. 二叉树的深度优先遍历（前序遍历）
// ------------------------------

func depthIterator(root *myBinTreeNode) {
	if root == nil {
		return
	}
	fmt.Print(strconv.Itoa(root.value) + ",")
	depthIterator(root.left)
	depthIterator(root.right)
}

func depthIterator02(root *myBinTreeNode) {
	s := &treeNodesStack{}
	s.init(30)
	s.add(root)

	for s.size() != 0 {
		node := s.pop()
		fmt.Printf(strconv.Itoa(node.value) + ",")
		if node.right != nil {
			s.add(node.right)
		}
		if node.left != nil {
			s.add(node.left)
		}
	}
}

// ------------------------------
// #4. 二叉树的广度优先遍历
// ------------------------------

func widthIterator(root *treeNode) {
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		ele := queue.Front()
		queue.Remove(ele)
		if node, ok := ele.Value.(*treeNode); ok {
			fmt.Print(strconv.Itoa(node.Val) + ",")
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
}
