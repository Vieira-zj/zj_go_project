package gotests_test

import (
	"testing"

	"demo.tests/gotests"
)

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotests.IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotests.IsPalindrome2("A man, a plan, a canal: Panama")
	}
}

func BenchmarkIsPalindrome3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotests.IsPalindrome3("A man, a plan, a canal: Panama")
	}
}
