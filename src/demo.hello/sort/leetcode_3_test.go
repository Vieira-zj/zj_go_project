package sort

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	result := containsDuplicate([]int{1, 2, 3, 1})
	if !result {
		t.Fatal("expect true, and actual:", result)
	}

	result = containsDuplicate([]int{1, 2, 3, 4})
	if result {
		t.Fatal("expect false, and actual:", result)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 1, 2}
	result := removeDuplicates01(input)
	if result != 2 {
		t.Fatal("expect 2, and actual:", result)
	}
	input = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result = removeDuplicates01(input)
	if result != 5 {
		t.Fatal("expect 5, and actual:", result)
	}

	input = []int{1, 1, 2}
	result = removeDuplicates02(input)
	if result != 2 {
		t.Fatal("expect 2, and actual:", result)
	}
	input = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result = removeDuplicates02(input)
	if result != 5 {
		t.Fatal("expect 5, and actual:", result)
	}
}

func TestMissingNumber(t *testing.T) {
	input := []int{3, 0, 1}
	result := missingNumber01(input)
	if result != 2 {
		t.Fatal("expect 2, and actual:", result)
	}
	input = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	result = missingNumber01(input)
	if result != 8 {
		t.Fatal("expect 8, and actual:", result)
	}

	input = []int{3, 0, 1}
	result = missingNumber02(input)
	if result != 2 {
		t.Fatal("expect 2, and actual:", result)
	}
	input = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	result = missingNumber02(input)
	if result != 8 {
		t.Fatal("expect 8, and actual:", result)
	}
}

func TestIsPalindromeLinkedList(t *testing.T) {
	list1 := createListNodes([]int{1, 2})
	result := isPalindromeLinkedList(list1)
	if result {
		t.Fatal("expect false, and actual:", result)
	}

	list2 := createListNodes([]int{1, 2, 2, 1})
	result = isPalindromeLinkedList(list2)
	if !result {
		t.Fatal("expect true, and actual:", result)
	}
}

func TestGetKthFromEnd(t *testing.T) {
	list := createListNodes([]int{1, 2, 3, 4, 5})
	fmt.Print("expect [4,5], and actual: ")
	printListNodes(getKthFromEnd(list, 2))
}

func TestReplaceSpace(t *testing.T) {
	result := replaceSpace("We are happy.")
	t.Log("expect 'We%20are%20happy.', and actual:", result)
}

func TestReverseLeftWords(t *testing.T) {
	result := reverseLeftWords01("abcdefg", 2)
	if result != "cdefgab" {
		t.Fatal("expect 'cdefgab', and actual:", result)
	}
	result = reverseLeftWords01("lrloseumgh", 6)
	if result != "umghlrlose" {
		t.Fatal("expect 'umghlrlose', and actual:", result)
	}

	result = reverseLeftWords02("abcdefg", 2)
	if result != "cdefgab" {
		t.Fatal("expect 'cdefgab', and actual:", result)
	}
	result = reverseLeftWords02("lrloseumgh", 6)
	if result != "umghlrlose" {
		t.Fatal("expect 'umghlrlose', and actual:", result)
	}
}

func TestReverseWords(t *testing.T) {
	input := "Let's take LeetCode contest"
	result := reverseWords(input)
	t.Log("expect (s'teL ekat edoCteeL tsetnoc), and actual:")
	t.Log(result)
}

func TestIsPalindromeNumber(t *testing.T) {
	result := isPalindromeNumber(1001)
	if !result {
		t.Fatal("expect true, and actual:", result)
	}
	result = isPalindromeNumber(12321)
	if !result {
		t.Fatal("expect true, and actual:", result)
	}

	result = isPalindromeNumber(10)
	if result {
		t.Fatal("expect false, and actual:", result)
	}
	result = isPalindromeNumber(1000021)
	if result {
		t.Fatal("expect false, and actual:", result)
	}
}

func TestReverseNumber(t *testing.T) {
	result := reverseNumber(123)
	if result != 321 {
		t.Fatal("expect 321, and actual:", result)
	}

	result = reverseNumber(-123)
	if result != -321 {
		t.Fatal("expect -321, and actual:", result)
	}
}

func TestIsValidBrackets(t *testing.T) {
	result := isValidBrackets("()[]{}")
	if !result {
		t.Fatal("expect true, and actual:", result)
	}

	result = isValidBrackets("([)]")
	if result {
		t.Fatal("expect false, and actual:", result)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	result := longestCommonPrefix([]string{"aa", "ab"})
	if result != "a" {
		t.Fatal("expect 'a', and actual:", result)
	}

	result = longestCommonPrefix([]string{"flower", "flow", "flight"})
	if result != "fl" {
		t.Fatal("expect 'fl', and actual:", result)
	}
}

func TestFirstUniqChar(t *testing.T) {
	result := firstUniqChar("leetcode")
	if result != 0 {
		t.Fatal("expect 0, and actual:", result)
	}

	result = firstUniqChar("loveleetcode")
	if result != 2 {
		t.Fatal("expect 2, and actual:", result)
	}
}

func TestFindNumbers(t *testing.T) {
	input := []int{12, 345, 2, 6, 7896}
	result := findNumbers(input)
	if result != 2 {
		t.Fatal("expect 2, and actual:", result)
	}
}

func TestCheckPermutation(t *testing.T) {
	s1 := "abc"
	s2 := "bca"
	result := checkPermutation(s1, s2)
	if !result {
		t.Fatal("expect true, and actual:", result)
	}

	s1 = "abc"
	s2 = "bad"
	result = checkPermutation(s1, s2)
	if result {
		t.Fatal("expect false, and actual:", result)
	}
}

func TestReference(t *testing.T) {
	// 引用传递
	node1 := &listNode{
		Val: 1,
	}
	node2 := &listNode{
		Val: 2,
	}
	node1.Next = node2
	printListNodes(node1)
	node2 = nil
	printListNodes(node1)
}

func TestRemoveDuplicateNodes(t *testing.T) {
	nodes := createListNodes([]int{1, 2, 3, 3, 2, 1})
	result := removeDuplicateNodes(nodes)
	t.Log("expect [1, 2, 3], and actual:")
	printListNodes(result)

	nodes = createListNodes([]int{1, 1, 1, 1, 2})
	result = removeDuplicateNodes(nodes)
	t.Log("expect [1, 2], and actual:")
	printListNodes(result)
}
