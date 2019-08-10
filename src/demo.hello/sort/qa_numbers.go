package sort

import (
	"fmt"
	"math/rand"
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

// 查找最小的k个元素（topK）
func topKMinNumbers(nums []int, k int) []int {
	// 部分排序 维护一个大小为K的数组 由大到小排序 保持有序
	list := mySortFixedIntList{}
	list.init(nums[:k])
	for _, num := range nums[k:] {
		list.add(num)
	}
	return list.getNumbers()
}

type mySortFixedIntList struct {
	size    int
	numbers []int
}

func (l *mySortFixedIntList) getNumbers() []int {
	return l.numbers
}

func (l *mySortFixedIntList) init(nums []int) {
	l.size = len(nums)
	l.numbers = make([]int, l.size, l.size)
	copy(l.numbers, nums)

	// bubble sort
	var isExchange bool
	for i := 0; i < l.size-1; i++ {
		isExchange = false
		for j := 0; j < l.size-1-i; j++ {
			if l.numbers[j] < l.numbers[j+1] {
				l.numbers[j], l.numbers[j+1] = l.numbers[j+1], l.numbers[j]
				isExchange = true
			}
		}
		if !isExchange {
			break
		}
	}
}

func (l *mySortFixedIntList) add(num int) {
	if num < l.numbers[0] {
		l.numbers[0] = num
	}
	for i := 0; i < l.size-1; i++ {
		if l.numbers[i] > l.numbers[i+1] {
			break
		}
		l.numbers[i], l.numbers[i+1] = l.numbers[i+1], l.numbers[i]
	}
}

// 抽样, 从n个中抽m个
func numberSampling(nums []int, m int) []int {
	selected := make([]int, 0, m)
	remaining := len(nums)

	// 轮流判断n个数组成的列表中每个数的概率(m/n), 每次判断后n=n-1, 若当前被判断的数被选择, 则m=m-1, 否则m不变
	for _, num := range nums {
		// rand.Float32() 返回 0 ~ 1的随机数
		if float32(remaining)*rand.Float32() < float32(m) {
			selected = append(selected, num)
			m--
		}
		remaining--
	}
	return selected
}

// TestNumbersAlgorithms test for numbers algorithms.
func TestNumbersAlgorithms() {
	// #1
	fmt.Println("\nprime numbers in 100:", getPrimeNumbersWithN(100))

	// #2
	numbers := []int{3, 4, 6, 2, 1, 6, 7, 10, 13}
	fmt.Println("\nsrc numbers:", numbers)
	sortOddNumbersFront(numbers)
	fmt.Println("numebers with odd in front:", numbers)

	// #3
	numbers = []int{3, 11, 6, 2, 13, 1, 6, 7, 10}
	fmt.Printf("\n(%v) top 4 min numbers: %v\n", numbers, topKMinNumbers(numbers, 4))

	// #4
	numbers = make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
	}
	fmt.Printf("\n(%v) sampling 3 numbers: %v\n", numbers, numberSampling(numbers, 3))
}
