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
	expect := []string{"a2b1c5a3", "abbccd"}
	for i := 0; i < len(input); i++ {
		ret := compressString(input[i])
		if ret != expect[i] {
			t.Fatalf("input %s, exepct %s, and actual %s", input[i], expect[i], ret)
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
	expect := queue.MaxValue()
	if expect != -1 {
		t.Fatal("expect -1, and actual:", expect)
	}

	queue.PushBack(4)
	queue.PushBack(2)
	queue.PushBack(0)
	queue.PushBack(3)
	expect = queue.MaxValue()
	if expect != 4 {
		t.Fatal("expect 4, and actual:", expect)
	}
	expect = queue.PopFront()
	if expect != 4 {
		t.Fatal("expect 4, and actual:", expect)
	}
	expect = queue.MaxValue()
	if expect != 3 {
		t.Fatal("expect 3, and actual:", expect)
	}
}

func TestLongestPalindrome(t *testing.T) {
	expect := longestPalindrome("abbccccddd")
	if expect != 9 {
		t.Fatal("expect 9, and actual:", expect)
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
	expect := findDisappearedNumbers(input)
	t.Log("miss numbers:", expect)
	if len(expect) != 2 {
		t.Fatal("expect 2, and actual:", len(expect))
	}
}

func TestFindUnsortedSubarray(t *testing.T) {
	input := []int{2, 6, 4, 8, 10, 9, 15}
	expect := findUnsortedSubarray(input)
	if expect != 5 {
		t.Fatal("expect 5, and actual:", expect)
	}
}

func TestIsUnique(t *testing.T) {
	expect := isUnique("leetcode")
	if expect {
		t.Fatal("expect false, and actual:", expect)
	}

	expect = isUnique("abc")
	if !expect {
		t.Fatal("expect true, and actual:", expect)
	}
}
