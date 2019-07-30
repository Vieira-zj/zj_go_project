package sort

import (
	"fmt"
)

// 快速排序 O(N*logN)
func quickSort(s []int, left, right int) {
	if left >= right {
		return
	}

	base := s[left] // 基准数
	i := left
	j := right
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
	s[i] = s[left]
	s[left] = base

	quickSort(s, i, left-1)
	quickSort(s, left+1, j)
}

// TestQuickSort test for quickSort.
func TestQuickSort() {
	s := []int{1, 16, 15, 7, 99, 50, 0, 99, 11}
	quickSort(s, 0, len(s)-1) // 引用传递
	fmt.Println("\nquick sort results:", s)
}
