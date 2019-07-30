package sort

import (
	"fmt"
)

// 桶排序 O(2*(M+N)) 浪费空间 不兼容小数
func bucketSort(inputS []int) []int {
	var buckets [100]int
	fmt.Println("\ninit buckets:", buckets[:10])
	for _, num := range inputS {
		buckets[num]++
	}

	retS := make([]int, 0, len(inputS))
	for idx, num := range buckets {
		if num == 0 {
			continue
		}
		for i := 0; i < num; i++ {
			retS = append(retS, idx)
		}
	}
	return retS
}

// TestBucketSort test for bucketSort.
func TestBucketSort() {
	arr := bucketSort([]int{1, 16, 15, 99, 50, 0, 99})
	fmt.Println("bucket sort results:", arr)
}
