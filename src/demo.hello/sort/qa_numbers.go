package sort

import (
	"fmt"
)

// 求n内的质数
func getPrimeNumbersWithN(n int) []int {
	primes := make([]int, 0, 10)
	var isPrime bool

	for i := 2; i <= n; i++ {
		isPrime = true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}

// 数组中的奇数排在前面
func sortOddNumbersFront(arr []int) {
	start := 0
	end := len(arr) - 1

	for start != end {
		if arr[start]%2 == 1 && start < end {
			start++
		}
		if arr[end]%2 == 0 && start < end {
			end--
		}
		if start < end {
			arr[start], arr[end] = arr[end], arr[start]
		}
	}
}

// TestNumbersAlgorithms test for numbers algorithms.
func TestNumbersAlgorithms() {
	// #1
	fmt.Println("\nprime numbers in 100:", getPrimeNumbersWithN(100))

	// #2
	n := []int{3, 4, 6, 2, 1, 6, 7, 13, 10}
	fmt.Println("\nsrc numbers:", n)
	sortOddNumbersFront(n)
	fmt.Println("numebers with odd in front:", n)
}
