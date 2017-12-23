package gotests_test

import (
	"fmt"
	"testing"

	"demo.tests/gotests"
)

// cmd: go test -v src/demo.tests/gotests/word_test.go
func ExampleIsPalindrome() {
	fmt.Println(gotests.IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(gotests.IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}

func TestIsPalindrome(t *testing.T) {
	// data driver
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}

	for _, test := range tests {
		if got := gotests.IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
