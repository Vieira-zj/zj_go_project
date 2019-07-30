package sort

import (
	"fmt"
)

func isRecycleString(s string) bool {
	var (
		mid, next int
	)
	if len(s)%2 == 1 {
		mid = (len(s) - 1) / 2
		next = mid + 1
	} else {
		mid = len(s) / 2
		next = mid
	}

	stack := make([]uint8, mid, mid)
	top := -1
	for i := 0; i < mid; i++ {
		top++
		stack[top] = s[i]
	}

	for i := next; i < len(s); i++ {
		if stack[top] != s[i] {
			break
		}
		top--
	}

	if top == -1 {
		return true
	}
	return false
}

// TestRecycleString test for isRecycleString.
func TestRecycleString() {
	tmp := []string{"ahaha", "ahha", "haha"}
	fmt.Println()
	for _, s := range tmp {
		fmt.Printf("%s is recycle string: %v\n", s, isRecycleString(s))
	}
}
