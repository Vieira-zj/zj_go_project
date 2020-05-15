package sort

import (
	"container/list"
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// ------------------------------
// #1. 十进制转二进制
// ------------------------------

func intOctToBinary01(num int) string {
	l := list.New()
	for num > 0 {
		l.PushBack(strconv.Itoa(num % 2))
		num /= 2
	}

	ret := ""
	len := l.Len()
	for i := 0; i < len; i++ {
		ele := l.Back()
		l.Remove(ele)
		if b, ok := ele.Value.(string); ok {
			ret += b
		}
	}
	return ret
}

func intOctToBinary02(num int) string {
	var ret string
	for num > 0 {
		ret = strconv.Itoa(num%2) + ret
		num /= 2
	}
	return ret
}

// ------------------------------
// #2. 二进制转十进制
// ------------------------------

func intBinaryToOct01(bits string) int {
	var ret float64
	if bits[len(bits)-1] == '1' {
		ret = 1
	}
	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == '1' {
			ret += math.Pow(2, float64(len(bits)-1-i))
		}
	}
	return int(ret)
}

func intBinaryToOct02(bits string) int {
	var ret int
	for i := 0; i < len(bits); i++ {
		ret *= 2
		if bits[i] == '1' {
			ret++
		}
	}
	return ret
}

func intBinaryToOct03(bits string) int {
	ret := 0
	for _, b := range bits {
		ret <<= 1
		if b == '1' {
			ret++
		}
	}
	return ret
}

func myReverse(s string) string {
	r := []rune(s)
	start := 0
	end := len(s) - 1
	for start < end {
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
	return string(r)
}

// ------------------------------
// #3. 求 n 内的质数
// ------------------------------

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

// ------------------------------
// Arrays
// #4. 数组中的奇数排在前面
// ------------------------------

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

// ------------------------------
// #5. 将有序数组a[]和b[]合并到c[]中
// ------------------------------

func mergeSortedSlice(a, b []int) []int {
	var (
		idxA = 0
		idxB = 0
		lenA = len(a)
		lenB = len(b)
		ret  = make([]int, 0, len(a)+len(b))
	)

	for idxA < lenA && idxB < lenB {
		if a[idxA] < b[idxB] {
			ret = append(ret, a[idxA])
			idxA++
		} else {
			ret = append(ret, b[idxB])
			idxB++
		}
	}
	if idxA < lenA {
		ret = append(ret, a[idxA:]...)
	}
	if idxB < lenB {
		ret = append(ret, b[idxB:]...)
	}
	return ret
}

// ------------------------------
// #6. 查找最小的k个元素（topK）
// ------------------------------

func topKMinNumbers(nums []int, k int) []int {
	ints := mySortedFixedInts{}
	ints.init(nums[:k])
	for _, num := range nums[k:] {
		ints.add(num)
	}
	return ints.getNumbers()
}

// 部分排序 维护一个大小为K的数组 由大到小排序 保持有序
type mySortedFixedInts struct {
	size    int
	numbers []int
}

func (l *mySortedFixedInts) getNumbers() []int {
	return l.numbers
}

func (l *mySortedFixedInts) init(nums []int) {
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

func (l *mySortedFixedInts) add(num int) {
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

// ------------------------------
// #7. 数组2*n个元素，n个奇数、n个偶数，使得数组奇数下标位置放置的都是奇数，偶数下标位置放置的都是偶数
// ------------------------------

func numbersSelect(nums []int) {
	var (
		evenIdx = 0
		oddIdx  = 1
		size    = len(nums)
	)

	for oddIdx < size && evenIdx < size {
		for evenIdx < size && nums[evenIdx]%2 == 0 {
			evenIdx += 2
		}
		for oddIdx < size && nums[oddIdx]%2 == 1 {
			oddIdx += 2
		}
		if oddIdx < size && evenIdx < size {
			nums[evenIdx], nums[oddIdx] = nums[oddIdx], nums[evenIdx]
		}
	}
}

// ------------------------------
// #8. 抽样，从n个中抽m个
// ------------------------------

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

// ------------------------------
// fibonacci
// 0,1,1,2,3,5,8,13
// n >= 0, and n is index of number.
// ------------------------------

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacci2(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	before := 0
	after := 1
	for i := 2; i <= n; i++ {
		tmp := after
		after += before
		before = tmp
	}
	return after
}

// TestNumbersAlgorithms test for numbers algorithms.
func TestNumbersAlgorithms() {
	if false {
		fmt.Println("\n#1. 十进制转二进制")
		num := 10 // 1010
		fmt.Printf("%d=%s\n", num, intOctToBinary01(num))
		num = 110 // 1101110
		fmt.Printf("%d=%s\n", num, intOctToBinary02(num))

		fmt.Println("\n#2. 二进制转十进制")
		bits := "1010"
		fmt.Printf("%s=%d\n", bits, intBinaryToOct01(bits))
		bits = "1101110"
		fmt.Printf("%s=%d\n", bits, intBinaryToOct02(bits))
		fmt.Printf("%s=%d\n", bits, intBinaryToOct03(bits))

		fmt.Println("\n#3. 求 n 内的质数")
		fmt.Println("prime numbers in 100:", getPrimeNumbersWithN(100))

		fmt.Println("\n#4. 数组中的奇数排在前面")
		numbers := []int{3, 4, 6, 2, 1, 6, 7, 10, 13}
		fmt.Println("src numbers:", numbers)
		sortOddNumbersFront(numbers)
		fmt.Println("numebers with odd in front:", numbers)

		fmt.Println("\n#5. 将有序数组a[]和b[]合并到c[]中")
		a := []int{2, 4, 6, 13, 15}
		b := []int{1, 7, 9, 11, 14}
		fmt.Println("merge for sorted slices:", mergeSortedSlice(a, b))

		fmt.Println("\n#6. 查找最小的k个元素（topK）")
		numbers1 := []int{3, 11, 6, 2, 13, 1, 6, 7, 10}
		fmt.Printf("(%v) top 4 min numbers: %v\n", numbers1, topKMinNumbers(numbers1, 4))

		fmt.Println("\n#7. 数组2*n个元素，n个奇数、n个偶数，使得数组奇数下标位置放置的都是奇数，偶数下标位置放置的都是偶数")
		numbers = []int{3, 11, 17, 6, 2, 13, 6, 7, 10, 20}
		numbersSelect(numbers)
		fmt.Println("odd/even numbers in odd/even index:", numbers)

		fmt.Println("\n#8. 抽样，从n个中抽m个")
		numbers = make([]int, 0, 10)
		for i := 0; i < 10; i++ {
			numbers = append(numbers, i)
		}
		fmt.Printf("(%v) sampling 3 numbers: %v\n", numbers, numberSampling(numbers, 3))

		fmt.Println("\n#9. fibonacci")
		fmt.Println("fibonacci(7):", fibonacci(7))
		fmt.Println("fibonacci2(7):", fibonacci2(7))
	}

	fmt.Println("numbers algorithms done.")
}
