package sort

import (
	"fmt"
)

// 桶排序 O(2*(M+N)) 浪费空间 不兼容小数
func bucketSort(numbers []int) []int {
	var buckets [100]int
	fmt.Println("\ninit array of buckets:", buckets[:10])
	for _, num := range numbers {
		buckets[num]++
	}

	ret := make([]int, 0, len(numbers))
	for idx, num := range buckets {
		if num == 0 {
			continue
		}
		for i := 0; i < num; i++ {
			ret = append(ret, idx)
		}
	}
	return ret
}

// 冒泡排序（交换排序）O(N*N)
func bubbleSort(s []int) {
	n := len(s)
	var isExchange bool

	for i := 0; i < n-1; i++ { // n-1次
		isExchange = false
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] { // 比较相邻的两个数
				s[j], s[j+1] = s[j+1], s[j]
				isExchange = true
			}
		}
		if !isExchange {
			break
		}
	}
}

// 快速排序（交换排序）O(N*logN)
func quickSort(s []int, start, end int) {
	if start >= end {
		return
	}

	base := s[start] // 基准数
	left := start
	right := end
	for left != right {
		// 从右开始往左移动
		for s[right] >= base && left < right {
			right--
		}
		for s[left] <= base && left < right {
			left++
		}
		// 没有相遇时，交互两个数在数组中的位置
		if left < right {
			s[left], s[right] = s[right], s[left]
		}
	}
	s[start] = s[left]
	s[left] = base

	quickSort(s, start, left-1)
	quickSort(s, left+1, end)
}

// 归并排序 O(N*logN)
// https://my.oschina.net/mutoushirana/blog/1854644
func mergeSort(s []int) []int {
	if len(s) == 1 {
		return s
	}

	mid := len(s) / 2
	s1 := mergeSort(s[:mid])
	s2 := mergeSort(s[mid:])
	return merge(s1, s2)
}

// 二路归并：两个有序的子序列合并为一个新的有序序列
func merge(s1, s2 []int) []int {
	var (
		i   = 0
		j   = 0
		ret = make([]int, 0, len(s1)+len(s2))
	)

	for i < len(s1) && j < len(s2) {
		if s1[i] < s2[j] {
			ret = append(ret, s1[i])
			i++
		} else {
			ret = append(ret, s2[j])
			j++
		}
	}
	if i < len(s1) {
		ret = append(ret, s1[i:]...)
	}
	if j < len(s2) {
		ret = append(ret, s2[j:]...)
	}
	return ret
}

// TestSortAlgorithms test for sort algorithms.
func TestSortAlgorithms() {
	// bucket sort
	numbers := bucketSort([]int{1, 16, 15, 99, 50, 0, 99, 13})
	fmt.Println("bucket sort results:", numbers)

	// bubble sort
	s := []int{1, 16, 15, 7, 99, 50, 0, 99, 13, 7}
	fmt.Println("\ninit slice of numbers:", s)
	bubbleSort(s) // 引用传递
	fmt.Println("bubble sort results:", s)

	// quick sort
	s = []int{1, 16, 15, 7, 99, 50, 0, 99, 11, 32}
	quickSort(s, 0, len(s)-1) // 引用传递
	fmt.Println("\nquick sort results:", s)

	// merge sort
	s = []int{3, 16, 14, 8, 99, 53, 0, 99, 8, 32, 66}
	fmt.Println("\nmerge sort results:", mergeSort(s))
}
