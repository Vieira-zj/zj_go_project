package main

import (
	"fmt"
	"sort"
)

// sort.Interface
// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool // i, j are indices of sequence elements
// 	Swap(i, j int)
// }

type myStringSlice []string

func (p myStringSlice) Len() int {
	return len(p)
}

func (p myStringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p myStringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func sortTest() {
	names := []string{"candy", "brow", "dan", "emm", "anny"}
	fmt.Println("names before sorted:", names)
	sort.Sort(myStringSlice(names))
	fmt.Println("names after sorted:", names)
}

func main() {
	sortTest()

	fmt.Println("interface demo.")
}
