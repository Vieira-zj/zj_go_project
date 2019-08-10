package sort

import (
	"fmt"
)

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

// TestBubbleSort test for bubbleSort.
func TestBubbleSort() {
	s := []int{1, 16, 15, 7, 99, 50, 0, 99, 13}
	fmt.Println("\ninit slice:", s)
	bubbleSort(s) // 引用传递
	fmt.Println("bubble sort results:", s)
}
