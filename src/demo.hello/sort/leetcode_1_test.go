package sort

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	expect := isPalindrome(input)
	if !expect {
		t.Fatal("excpect true, and actual:", expect)
	}

	input = "race a car"
	expect = isPalindrome(input)
	if expect {
		t.Fatal("excpect false, and actual:", expect)
	}
}

func TestTwoSum01(t *testing.T) {
	expect := twoSum01([]int{2, 7, 11, 15}, 9)
	t.Log("expect [0,1], and actual:", expect)

	expect = twoSum01([]int{-3, 4, 3, 90}, 0)
	t.Log("expect [0,2], and actual:", expect)
}

func TestTwoSum02(t *testing.T) {
	expect := twoSum02([]int{3, 2, 4}, 6)
	t.Log("expect [1,2], and actual:", expect)

	expect = twoSum02([]int{-3, 4, 3, 90}, 0)
	t.Log("expect [0,2], and actual:", expect)
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
	expect := toLowerCase(input)
	if expect != "hello" {
		t.Fatal("expect 'hello', and actual:", expect)
	}
}

func TestBalancedStringSplit(t *testing.T) {
	input := "RLRRLLRLRL"
	expect := balancedStringSplit(input)
	if expect != 4 {
		t.Fatal("expect 4, and actual:", expect)
	}

	input = "RLLLLRRRLR"
	expect = balancedStringSplit(input)
	if expect != 3 {
		t.Fatal("expect 3, and actual:", expect)
	}

	input = "LLLLRRRR"
	expect = balancedStringSplit(input)
	if expect != 1 {
		t.Fatal("expect 1, and actual:", expect)
	}
}

func TestReverseList(t *testing.T) {
	listNodes := createListNodes([]int{1, 2, 3, 4, 5})
	expect := reverseList(listNodes)
	fmt.Print("expect [5,4,3,2,1], and actual: ")
	printListNodes(expect)
}

func TestHasCycle(t *testing.T) {
	listNode := createCycleListNodes([]int{3, 2, 0, -4}, 1)
	expect := hasCycle(listNode)
	if !expect {
		t.Fatal("expect true, and actual:", expect)
	}

	listNode = createCycleListNodes([]int{3, 2, 0, -4}, -1)
	expect = hasCycle(listNode)
	if expect {
		t.Fatal("expect false, and actual:", expect)
	}
}

func TestReverseBits(t *testing.T) {
	input := 43261596
	expect := reverseBits(uint32(input))
	if expect != 964176192 {
		t.Fatal("expect 964176192, and actual:", expect)
	}
}

func TestStrStr(t *testing.T) {
	expect := strStr("hello", "ll")
	if expect != 2 {
		t.Fatal("expect 2, and actual:", expect)
	}

	expect = strStr("aaaaa", "bba")
	if expect != -1 {
		t.Fatal("expect -1, and actual:", expect)
	}
}

func TestMaxDepth(t *testing.T) {
	tree := createBinTree([]int{3, 9, 20, -1, -1, 15, 7})
	expect := maxDepth(tree)
	if expect != 3 {
		t.Fatal("expect 3, and actual:", expect)
	}
}

func TestTitleToNumber(t *testing.T) {
	expect := titleToNumber("AB")
	if expect != 28 {
		t.Fatal("expect 28, and actual:", expect)
	}

	expect = titleToNumber("ZY")
	if expect != 701 {
		t.Fatal("expect 701, and actual:", expect)
	}
}

func TestMergeSortedNums(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	mergeSortedNums(nums1, 3, nums2, len(nums2))
	t.Log("expect [1,2,2,3,5,6], and actual:", nums1)
}

func TestGeneTriangle(t *testing.T) {
	expect := geneTriangle(5)
	t.Log("expect [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]], and actual:", expect)
}

func TestIsSymmetric(t *testing.T) {
	tree := createBinTree([]int{1, 2, 2, 3, 4, 4, 3})
	expect := isSymmetric(tree)
	if !expect {
		t.Fatal("expect true, and actual:", expect)
	}
}

func TestSingleNumber(t *testing.T) {
	input := []int{2, 2, 1}
	expect := singleNumber(input)
	if expect != 1 {
		t.Fatal("expect 1, and actual:", expect)
	}

	input = []int{4, 1, 2, 1, 2}
	expect = singleNumber(input)
	if expect != 4 {
		t.Fatal("expect 4, and actual:", expect)
	}
}
