package sort

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	result := isPalindrome(input)
	if !result {
		t.Fatal("expect true, and actual:", result)
	}

	input = "race a car"
	result = isPalindrome(input)
	if result {
		t.Fatal("expect false, and actual:", result)
	}
}

func TestTwoSum01(t *testing.T) {
	result := twoSum01([]int{2, 7, 11, 15}, 9)
	t.Log("expect [0,1], and actual:", result)

	result = twoSum01([]int{-3, 4, 3, 90}, 0)
	t.Log("expect [0,2], and actual:", result)
}

func TestTwoSum02(t *testing.T) {
	result := twoSum02([]int{3, 2, 4}, 6)
	t.Log("expect [1,2], and actual:", result)

	result = twoSum02([]int{-3, 4, 3, 90}, 0)
	t.Log("expect [0,2], and actual:", result)
}

func TestDeleteNode(t *testing.T) {
	listNodes := createListNodes([]int{4, 5, 1, 9})
	target := getListNodeByValue(listNodes, 1)
	deleteNode(target)
	fmt.Print("expect [4,5,9], and actual: ")
	printListNodes(listNodes)
}

func TestToLowerCase(t *testing.T) {
	input := "Hello"
	result := toLowerCase(input)
	if result != "hello" {
		t.Fatal("expect 'hello', and actual:", result)
	}
}

func TestBalancedStringSplit(t *testing.T) {
	input := "RLRRLLRLRL"
	result := balancedStringSplit(input)
	if result != 4 {
		t.Fatal("expect 4, and actual:", result)
	}

	input = "RLLLLRRRLR"
	result = balancedStringSplit(input)
	if result != 3 {
		t.Fatal("expect 3, and actual:", result)
	}

	input = "LLLLRRRR"
	result = balancedStringSplit(input)
	if result != 1 {
		t.Fatal("expect 1, and actual:", result)
	}
}

func TestReverseList(t *testing.T) {
	listNodes := createListNodes([]int{1, 2, 3, 4, 5})
	result := reverseList(listNodes)
	fmt.Print("expect [5,4,3,2,1], and actual: ")
	printListNodes(result)
}

func TestHasCycle(t *testing.T) {
	listNode := createCycleListNodes([]int{3, 2, 0, -4}, 1)
	result := hasCycle(listNode)
	if !result {
		t.Fatal("expect true, and actual:", result)
	}

	listNode = createCycleListNodes([]int{3, 2, 0, -4}, -1)
	result = hasCycle(listNode)
	if result {
		t.Fatal("expect false, and actual:", result)
	}
}

func TestReverseBits(t *testing.T) {
	input := 43261596
	result := reverseBits(uint32(input))
	if result != 964176192 {
		t.Fatal("expect 964176192, and actual:", result)
	}
}

func TestStrStr(t *testing.T) {
	result := strStr("hello", "ll")
	if result != 2 {
		t.Fatal("expect 2, and actual:", result)
	}

	result = strStr("aaaaa", "bba")
	if result != -1 {
		t.Fatal("expect -1, and actual:", result)
	}
}

func TestMaxDepth(t *testing.T) {
	tree := createBinTree([]int{3, 9, 20, -1, -1, 15, 7})
	result := maxDepth(tree)
	if result != 3 {
		t.Fatal("expect 3, and actual:", result)
	}
}

func TestTitleToNumber(t *testing.T) {
	result := titleToNumber("AB")
	if result != 28 {
		t.Fatal("expect 28, and actual:", result)
	}

	result = titleToNumber("ZY")
	if result != 701 {
		t.Fatal("expect 701, and actual:", result)
	}
}

func TestMergeSortedNums(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	mergeSortedNums(nums1, 3, nums2, len(nums2))
	t.Log("expect [1,2,2,3,5,6], and actual:", nums1)
}

func TestGeneTriangle(t *testing.T) {
	result := geneTriangle(5)
	t.Log("expect [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]], and actual:", result)
}

func TestIsSymmetric(t *testing.T) {
	tree := createBinTree([]int{1, 2, 2, 3, 4, 4, 3})
	result := isSymmetric(tree)
	if !result {
		t.Fatal("expect true, and actual:", result)
	}
}

func TestSingleNumber(t *testing.T) {
	input := []int{2, 2, 1}
	result := singleNumber(input)
	if result != 1 {
		t.Fatal("expect 1, and actual:", result)
	}

	input = []int{4, 1, 2, 1, 2}
	result = singleNumber(input)
	if result != 4 {
		t.Fatal("expect 4, and actual:", result)
	}
}
