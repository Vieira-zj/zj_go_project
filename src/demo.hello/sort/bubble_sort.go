package sort

import (
	"fmt"
)

// 冒泡排序 O(N*N)
func bubbleSort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ { // n-1次
		for j := 0; j < n-1-i; j++ { // n-i次
			if s[j] > s[j+1] { // 比较相邻的两个数
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

// TestBubbleSort test for bubbleSort.
func TestBubbleSort() {
	s := []int{1, 16, 15, 7, 99, 50, 0, 99}
	fmt.Println("\ninit slice:", s)
	bubbleSort(s) // 引用传递
	fmt.Println("bubble sort results:", s)
}
