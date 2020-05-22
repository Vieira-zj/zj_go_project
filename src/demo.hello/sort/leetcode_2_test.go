package sort

import (
	"fmt"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	result := hammingWeight01(11)
	if result != 3 {
		t.Fatal("expect 3, and actual:", result)
	}

	result = hammingWeight02(11)
	if result != 3 {
		t.Fatal("expect 3, and actual:", result)
	}
}

func TestFizzBuzz(t *testing.T) {
	result := fizzBuzz(15)
	t.Log("fizz buzz results:", result)
}

func TestMajorityElement(t *testing.T) {
	input := []int{3, 3, 4}
	result := majorityElement(input)
	if result != 3 {
		t.Fatal("expect 3, and actual:", result)
	}

	input = []int{2, 2, 1, 1, 4, 2, 2}
	result = majorityElement(input)
	if result != 2 {
		t.Fatal("expect 2, and actual:", result)
	}
}

func TestRomanToInt(t *testing.T) {
	result := romanToInt("LVIII")
	if result != 58 {
		t.Fatal("expect 58, and actual:", result)
	}

	result = romanToInt("MCMXCIV")
	if result != 1994 {
		t.Fatal("expect 1994, and actual:", result)
	}
}

func TestMergeTwoLists(t *testing.T) {
	listNodes1 := createListNodes([]int{1, 2, 4})
	listNodes2 := createListNodes([]int{1, 3, 4})
	result := mergeTwoLists(listNodes1, listNodes2)
	fmt.Print("expect [1->1->2->3->4->4], and actual: ")
	printListNodes(result)
}

func TestMoveZeroes(t *testing.T) {
	nums1 := []int{0, 1, 0, 3, 12}
	moveZeroes01(nums1)
	t.Log("expect [1,3,12,0,0], and actual:", nums1)

	nums2 := []int{0, 1, 0, 3, 12}
	moveZeroes02(nums2)
	t.Log("expect [1,3,12,0,0], and actual:", nums2)
}

func TestIsHappy(t *testing.T) {
	result := isHappy(19)
	if !result {
		t.Fatal("expect true, and actual:", result)
	}
}

func TestMaxProfit(t *testing.T) {
	result := maxProfit01([]int{7, 1, 5, 3, 6, 4})
	if result != 7 {
		t.Fatal("expect 7, and actual:", result)
	}
	result = maxProfit01([]int{1, 2, 3, 4, 5})
	if result != 4 {
		t.Fatal("expect 4, and actual:", result)
	}
	result = maxProfit01([]int{7, 6, 4, 3, 1})
	if result != 0 {
		t.Fatal("expect 0, and actual:", result)
	}

	result = maxProfit02([]int{7, 1, 5, 3, 6, 4})
	if result != 7 {
		t.Fatal("expect 7, and actual:", result)
	}
	result = maxProfit02([]int{1, 2, 3, 4, 5})
	if result != 4 {
		t.Fatal("expect 4, and actual:", result)
	}
	result = maxProfit02([]int{7, 6, 4, 3, 1})
	if result != 0 {
		t.Fatal("expect 0, and actual:", result)
	}
}

func TestInvertTree(t *testing.T) {
	tree := createBinTree([]int{4, 2, 7, 1, 3, 6, 9})
	fmt.Print("expect [4 7 9 6 2 3 1], and actual: ")
	printTree(invertTree(tree))
	fmt.Print()
}

func TestGetDecimalValue(t *testing.T) {
	list1 := createListNodes([]int{1, 0, 1})
	result := getDecimalValue01(list1)
	if result != 5 {
		t.Fatal("expect 5, and actual:", result)
	}
	list1 = createListNodes([]int{0})
	result = getDecimalValue01(list1)
	if result != 0 {
		t.Fatal("expect 0, and actual:", result)
	}

	list2 := createListNodes([]int{1, 0, 1})
	result = getDecimalValue02(list2)
	if result != 5 {
		t.Fatal("expect 5, and actual:", result)
	}
	list2 = createListNodes([]int{0})
	result = getDecimalValue02(list2)
	if result != 0 {
		t.Fatal("expect 0, and actual:", result)
	}
}

func TestMaxSubArray(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(input)
	if result != 6 {
		t.Fatal("expect 6, and actual:", result)
	}
}

func TestMinStack(t *testing.T) {
	stack := Constructor()
	for _, val := range []int{-2, 0, -3} {
		stack.Push(val)
	}
	stack.debugPrint()

	result := stack.GetMin()
	if result != -3 {
		t.Fatal("expect -3, and actual:", result)
	}

	stack.Pop()
	result = stack.Top()
	if result != 0 {
		t.Fatal("expect 0, and actual:", result)
	}

	result = stack.GetMin()
	if result != -2 {
		t.Fatal("expect -2, and actual:", result)
	}
}
