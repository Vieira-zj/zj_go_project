package sort

import (
	"fmt"
)

// 桶排序 O(2*(M+N)) 浪费空间 不兼容小数
func bucketSort(numbers []int) []int {
	var buckets [100]int
	fmt.Println("\ninit buckets:", buckets[:10])
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

// TestBucketSort test for bucketSort.
func TestBucketSort() {
	numbers := bucketSort([]int{1, 16, 15, 99, 50, 0, 99, 13})
	fmt.Println("bucket sort results:", numbers)
}
