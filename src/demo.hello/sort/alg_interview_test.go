package sort

import (
	"fmt"
	"testing"
)

func TestOctToBinary(t *testing.T) {
	result := octToBinary(0, 2)
	fmt.Println("expect '00', and actual:", result)
	result = octToBinary(1, 2)
	fmt.Println("expect '01', and actual:", result)
	result = octToBinary(2, 2)
	fmt.Println("expect '10', and actual:", result)
}

func TestGetCombinationCnt(t *testing.T) {
	result := getCombinationCnt(2)
	fmt.Println("expect 4, and actual:", result)
	result = getCombinationCnt(3)
	fmt.Println("expect 6, and actual:", result)
	result = getCombinationCnt(4)
	fmt.Println("expect 10, and actual:", result)
}

func TestFormatByRange(t *testing.T) {
	arr := []int{1, 2, 3, 5, 6, 7, 10, 13}
	fmt.Println("expect [1-3,5-7,10,13], and actual:", formatByRange(arr))
	arr = []int{1, 5, 6, 7, 8, 10, 12, 13, 14}
	fmt.Println("expect [1,5-8,10,12-14], and actual:", formatByRange(arr))
}

func TestDepthIterator(t *testing.T) {
	tree := createBinaryTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println("expect [1,2,4,8,9,5,10,3,6,7], and actual:")
	depthIterator(tree)
	fmt.Println()
	depthIterator02(tree)
	fmt.Println()
}

func TestWidthIterator(t *testing.T) {
	tree := createBinTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println("expect [1,2,3,4,5,6,7,8,9,10], and actual:")
	widthIterator(tree)
	fmt.Println()
}
