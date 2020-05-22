package sort

import (
	"strconv"
	"strings"
	"testing"
)

func TestIntToString(t *testing.T) {
	s := make([]string, 0)
	for i := 0; i <= 10; i++ {
		s = append(s, strconv.Itoa(i))
	}
	t.Log(strings.Join(s, ""))

	str := ""
	for i := 0; i <= 10; i++ {
		str += strconv.Itoa(i)
	}
	t.Log(str)
}

func TestCompressString(t *testing.T) {
	input := []string{"aabcccccaaa", "abbccd"}
	result := []string{"a2b1c5a3", "abbccd"}
	for i := 0; i < len(input); i++ {
		ret := compressString(input[i])
		if ret != result[i] {
			t.Fatalf("input %s, exepct %s, and actual %s", input[i], result[i], ret)
		}
	}
}

func TestMyDeque(t *testing.T) {
	deque := NewDeque()
	for i := 0; i < 10; i++ {
		deque.PushHead(i)
	}
	deque.Print()

	for i := 0; i < 3; i++ {
		t.Log("pop:", deque.PopHead())
	}
	deque.Print()

	for i := 0; i < 2; i++ {
		t.Log("pop:", deque.PopTail())
	}
	deque.Print()

	deque.PushTail(10)
	deque.PushTail(11)
	deque.Print()
	deque.PrintReverse()
}

func TestMaxQueue(t *testing.T) {
	queue := NewMaxQueue()
	result := queue.MaxValue()
	if result != -1 {
		t.Fatal("expect -1, and actual:", result)
	}

	queue.PushBack(4)
	queue.PushBack(2)
	queue.PushBack(0)
	queue.PushBack(3)
	result = queue.MaxValue()
	if result != 4 {
		t.Fatal("expect 4, and actual:", result)
	}
	result = queue.PopFront()
	if result != 4 {
		t.Fatal("expect 4, and actual:", result)
	}
	result = queue.MaxValue()
	if result != 3 {
		t.Fatal("expect 3, and actual:", result)
	}
}

func TestLongestPalindrome(t *testing.T) {
	result := longestPalindrome("abbccccddd")
	if result != 9 {
		t.Fatal("expect 9, and actual:", result)
	}
}

func TestMiddleNode(t *testing.T) {
	header := createListNodes([]int{1, 2, 3, 4, 5})
	mid := middleNode(header)
	printListNodes(mid)
	if mid.Val != 3 {
		t.Fatal("expect 3, and actual:", mid.Val)
	}

	header = createListNodes([]int{1, 2, 3, 4, 5, 6})
	mid = middleNode(header)
	printListNodes(mid)
	if mid.Val != 4 {
		t.Fatal("expect 4, and actual:", mid.Val)
	}
}

func TestDiameterOfBinaryTree(t *testing.T) {
	header := createBinTree([]int{1, 2, 3, 4, 5})
	diameter := diameterOfBinaryTree(header)
	if diameter != 3 {
		t.Fatal("expect 3, and actual:", diameter)
	}

	diameter = diameterOfBinaryTree(nil)
	if diameter != 0 {
		t.Fatal("expect 0, and actual:", diameter)
	}
}

func TestFindDisappearedNumbers(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := findDisappearedNumbers(input)
	t.Log("miss numbers:", result)
	if len(result) != 2 {
		t.Fatal("expect 2, and actual:", len(result))
	}
}

func TestFindUnsortedSubarray(t *testing.T) {
	input := []int{2, 6, 4, 8, 10, 9, 15}
	result := findUnsortedSubarray(input)
	if result != 5 {
		t.Fatal("expect 5, and actual:", result)
	}
}

func TestIsUnique(t *testing.T) {
	result := isUnique("leetcode")
	if result {
		t.Fatal("expect false, and actual:", result)
	}

	result = isUnique("abc")
	if !result {
		t.Fatal("expect true, and actual:", result)
	}
}

func TestReversePrint(t *testing.T) {
	list := createListNodes([]int{1, 3, 2, 4})
	result := reversePrint01(list)
	t.Log("expect [4,2,3,1], and actual:", result)

	result = reversePrint02(list)
	t.Log("expect [4,2,3,1], and actual:", result)

	result = reversePrint03(list)
	t.Log("expect [4,2,3,1], and actual:", result)
}

func TestCreateTargetArray(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	result := createTargetArray(nums, index)
	t.Log("expect [0,4,1,3,2], and actual:", result)

	nums = []int{4, 2, 4, 3, 2}
	index = []int{0, 0, 1, 3, 1}
	result = createTargetArray(nums, index)
	t.Log("expect [2,2,4,4,3], and actual:", result)
}

func TestMinTimeToVisitAllPoints(t *testing.T) {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	result := minTimeToVisitAllPoints(points)
	if result != 7 {
		t.Fatal("expect 7, and actual:", result)
	}

	points = [][]int{{3, 2}, {-2, 2}}
	result = minTimeToVisitAllPoints(points)
	if result != 5 {
		t.Fatal("expect 5, and actual:", result)
	}
}
