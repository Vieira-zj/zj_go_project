package gotests_test

import (
	"testing"

	"src/demo.tests/gotests"
)

// cmd: go test -v -bench=. src/demo.tests/gotests/word_perf_test.go
func BenchmarkIsPalindrome01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotests.IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func BenchmarkIsPalindrome02(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotests.IsPalindrome2("A man, a plan, a canal: Panama")
	}
}

func BenchmarkIsPalindrome03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotests.IsPalindrome3("A man, a plan, a canal: Panama")
	}
}
