package sort

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	expect := containsDuplicate([]int{1, 2, 3, 1})
	if !expect {
		t.Fatal("expect true, and actual:", expect)
	}

	expect = containsDuplicate([]int{1, 2, 3, 4})
	if expect {
		t.Fatal("expect false, and actual:", expect)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 1, 2}
	expect := removeDuplicates01(input)
	if expect != 2 {
		t.Fatal("expect 2, and actual:", expect)
	}
	input = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expect = removeDuplicates01(input)
	if expect != 5 {
		t.Fatal("expect 5, and actual:", expect)
	}

	input = []int{1, 1, 2}
	expect = removeDuplicates02(input)
	if expect != 2 {
		t.Fatal("expect 2, and actual:", expect)
	}
	input = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expect = removeDuplicates02(input)
	if expect != 5 {
		t.Fatal("expect 5, and actual:", expect)
	}
}

func TestMissingNumber(t *testing.T) {
	input := []int{3, 0, 1}
	expect := missingNumber01(input)
	if expect != 2 {
		t.Fatal("expect 2, and actual:", expect)
	}
	input = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	expect = missingNumber01(input)
	if expect != 8 {
		t.Fatal("expect 8, and actual:", expect)
	}

	input = []int{3, 0, 1}
	expect = missingNumber02(input)
	if expect != 2 {
		t.Fatal("expect 2, and actual:", expect)
	}
	input = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	expect = missingNumber02(input)
	if expect != 8 {
		t.Fatal("expect 8, and actual:", expect)
	}
}

func TestIsPalindromeLinkedList(t *testing.T) {
	list1 := createListNodes([]int{1, 2})
	expect := isPalindromeLinkedList(list1)
	if expect {
		t.Fatal("expect false, and actual:", expect)
	}

	list2 := createListNodes([]int{1, 2, 2, 1})
	expect = isPalindromeLinkedList(list2)
	if !expect {
		t.Fatal("expect true, and actual:", expect)
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
	expect := reverseLeftWords01("abcdefg", 2)
	if expect != "cdefgab" {
		t.Fatal("expect 'cdefgab', and actual:", expect)
	}
	expect = reverseLeftWords01("lrloseumgh", 6)
	if expect != "umghlrlose" {
		t.Fatal("expect 'umghlrlose', and actual:", expect)
	}

	expect = reverseLeftWords02("abcdefg", 2)
	if expect != "cdefgab" {
		t.Fatal("expect 'cdefgab', and actual:", expect)
	}
	expect = reverseLeftWords02("lrloseumgh", 6)
	if expect != "umghlrlose" {
		t.Fatal("expect 'umghlrlose', and actual:", expect)
	}
}

func TestReverseWords(t *testing.T) {
	input := "Let's take LeetCode contest"
	expect := reverseWords(input)
	t.Log("expect (s'teL ekat edoCteeL tsetnoc), and actual:")
	t.Log(expect)
}

func TestIsPalindromeNumber(t *testing.T) {
	expect := isPalindromeNumber(1001)
	if !expect {
		t.Fatal("expect true, and actual:", expect)
	}
	expect = isPalindromeNumber(12321)
	if !expect {
		t.Fatal("expect true, and actual:", expect)
	}

	expect = isPalindromeNumber(10)
	if expect {
		t.Fatal("expect false, and actual:", expect)
	}
	expect = isPalindromeNumber(1000021)
	if expect {
		t.Fatal("expect false, and actual:", expect)
	}
}

func TestReverseNumber(t *testing.T) {
	expect := reverseNumber(123)
	if expect != 321 {
		t.Fatal("expect 321, and actual:", expect)
	}

	expect = reverseNumber(-123)
	if expect != -321 {
		t.Fatal("expect -321, and actual:", expect)
	}
}

func TestIsValidBrackets(t *testing.T) {
	expect := isValidBrackets("()[]{}")
	if !expect {
		t.Fatal("expect true, and actual:", expect)
	}

	expect = isValidBrackets("([)]")
	if expect {
		t.Fatal("expect false, and actual:", expect)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	expect := longestCommonPrefix([]string{"aa", "ab"})
	if expect != "a" {
		t.Fatal("expect 'a', and actual:", expect)
	}

	expect = longestCommonPrefix([]string{"flower", "flow", "flight"})
	if expect != "fl" {
		t.Fatal("expect 'fl', and actual:", expect)
	}
}

func TestFirstUniqChar(t *testing.T) {
	expect := firstUniqChar("leetcode")
	if expect != 0 {
		t.Fatal("expect 0, and actual:", expect)
	}

	expect = firstUniqChar("loveleetcode")
	if expect != 2 {
		t.Fatal("expect 2, and actual:", expect)
	}
}

func TestFindNumbers(t *testing.T) {
	input := []int{12, 345, 2, 6, 7896}
	expect := findNumbers(input)
	if expect != 2 {
		t.Fatal("expect 2, and actual:", expect)
	}
}

func TestCheckPermutation(t *testing.T) {
	s1 := "abc"
	s2 := "bca"
	expect := checkPermutation(s1, s2)
	if !expect {
		t.Fatal("expect true, and actual:", expect)
	}

	s1 = "abc"
	s2 = "bad"
	expect = checkPermutation(s1, s2)
	if expect {
		t.Fatal("expect false, and actual:", expect)
	}
}
