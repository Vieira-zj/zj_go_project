package sort

import (
	"fmt"
)

// 快速排序
func quickSort(s []int, left, right int) {
	if left > right {
		return
	}

	base := s[left]
	i := left
	j := right
	for left != right {
		for s[right] >= base && left < right {
			right--
		}
		for s[left] <= base && left < right {
			left++
		}

		if left < right {
			s[left], s[right] = s[right], s[left]
		}
	}
	s[i] = s[left]
	s[left] = base

	quickSort(s, i, left-1)
	quickSort(s, left+1, j)
}

// TestQuickSort test for quickSort.
func TestQuickSort() {
	s := []int{1, 16, 15, 7, 99, 50, 0, 99, 11}
	quickSort(s, 0, len(s)-1)
	fmt.Println("\nquick sort results:", s)
}
