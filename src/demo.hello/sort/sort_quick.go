package sort

import (
	"fmt"
)

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

// TestQuickSort test for quickSort.
func TestQuickSort() {
	s := []int{1, 16, 15, 7, 99, 50, 0, 99, 11, 32}
	quickSort(s, 0, len(s)-1) // 引用传递
	fmt.Println("\nquick sort results:", s)
}
