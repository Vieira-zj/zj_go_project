package gotests_test

import (
	"testing"

	"demo.tests/gotests"
)

// cmd: go test -v src/demo.tests/gotests/word_test.go
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
			t.Errorf("IsPalindrome(%q) = %v, want: %v", test.input, got, test.want)
		}
	}
}
