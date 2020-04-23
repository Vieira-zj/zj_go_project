package sort

import (
	"fmt"
	"testing"
)

func TestHammingWeight(t *testing.T) {
	expect := hammingWeight01(11)
	if expect != 3 {
		t.Fatal("expect 3, and actual:", expect)
	}

	expect = hammingWeight02(11)
	if expect != 3 {
		t.Fatal("expect 3, and actual:", expect)
	}
}

func TestFizzBuzz(t *testing.T) {
	result := fizzBuzz(15)
	t.Log("fizz buzz results:", result)
}

func TestMajorityElement(t *testing.T) {
	input := []int{3, 3, 4}
	expect := majorityElement(input)
	if expect != 3 {
		t.Fatal("expect 3, and actual:", expect)
	}

	input = []int{2, 2, 1, 1, 4, 2, 2}
	expect = majorityElement(input)
	if expect != 2 {
		t.Fatal("expect 2, and actual:", expect)
	}
}

func TestRomanToInt(t *testing.T) {
	expect := romanToInt("LVIII")
	if expect != 58 {
		t.Fatal("expect 58, and actual:", expect)
	}

	expect = romanToInt("MCMXCIV")
	if expect != 1994 {
		t.Fatal("expect 1994, and actual:", expect)
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
	expect := isHappy(19)
	if !expect {
		t.Fatal("expect true, and actual:", expect)
	}
}

func TestMaxProfit(t *testing.T) {
	expect := maxProfit01([]int{7, 1, 5, 3, 6, 4})
	if expect != 7 {
		t.Fatal("expect 7, and actual:", expect)
	}
	expect = maxProfit01([]int{1, 2, 3, 4, 5})
	if expect != 4 {
		t.Fatal("expect 4, and actual:", expect)
	}
	expect = maxProfit01([]int{7, 6, 4, 3, 1})
	if expect != 0 {
		t.Fatal("expect 0, and actual:", expect)
	}

	expect = maxProfit02([]int{7, 1, 5, 3, 6, 4})
	if expect != 7 {
		t.Fatal("expect 7, and actual:", expect)
	}
	expect = maxProfit02([]int{1, 2, 3, 4, 5})
	if expect != 4 {
		t.Fatal("expect 4, and actual:", expect)
	}
	expect = maxProfit02([]int{7, 6, 4, 3, 1})
	if expect != 0 {
		t.Fatal("expect 0, and actual:", expect)
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
	expect := getDecimalValue01(list1)
	if expect != 5 {
		t.Fatal("expect 5, and actual:", expect)
	}
	list1 = createListNodes([]int{0})
	expect = getDecimalValue01(list1)
	if expect != 0 {
		t.Fatal("expect 0, and actual:", expect)
	}

	list2 := createListNodes([]int{1, 0, 1})
	expect = getDecimalValue02(list2)
	if expect != 5 {
		t.Fatal("expect 5, and actual:", expect)
	}
	list2 = createListNodes([]int{0})
	expect = getDecimalValue02(list2)
	if expect != 0 {
		t.Fatal("expect 0, and actual:", expect)
	}
}

func TestMaxSubArray(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	expect := maxSubArray(input)
	if expect != 6 {
		t.Fatal("expect 6, and actual:", expect)
	}
}

func TestMinStack(t *testing.T) {
	stack := Constructor()
	for _, val := range []int{-2, 0, -3} {
		stack.Push(val)
	}
	stack.debugPrint()

	expect := stack.GetMin()
	if expect != -3 {
		t.Fatal("expect -3, and actual:", expect)
	}

	stack.Pop()
	expect = stack.Top()
	if expect != 0 {
		t.Fatal("expect 0, and actual:", expect)
	}

	expect = stack.GetMin()
	if expect != -2 {
		t.Fatal("expect -2, and actual:", expect)
	}
}
