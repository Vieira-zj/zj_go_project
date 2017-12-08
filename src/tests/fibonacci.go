package tests

// Fibonacci : for test benchmark in ginkgo
func Fibonacci(n int) int {
	if n < 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
