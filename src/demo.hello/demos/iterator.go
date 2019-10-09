package demos

import (
	"fmt"
	"time"
)

/*
iterator by class
*/

type myIterator struct {
	data  []int
	index int
}

func (iter *myIterator) HasNext() bool {
	return iter.index < len(iter.data)
}

func (iter *myIterator) Next() int {
	ret := iter.data[iter.index]
	iter.index++
	return ret
}

type intsOne []int

func (i intsOne) Iterator() *myIterator {
	return &myIterator{
		data:  i,
		index: 0,
	}
}

/*
iterator by closure
*/

type intsTwo []int

func (i intsTwo) Iterator() func() (int, bool) {
	index := 0
	return func() (val int, ok bool) {
		if index >= len(i) {
			return
		}
		val, ok = i[index], true
		index++
		return
	}
}

/*
iterator by channel
*/

type intsThree []int

func (i intsThree) Iterator() <-chan int {
	c := make(chan int)
	go func() {
		for _, val := range i {
			c <- val
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
		close(c)
	}()
	return c
}

/*
iterator by foreach
*/

type intsFour []int

func (i intsFour) ForEach(fn func(ele int)) {
	for _, val := range i {
		fn(val)
	}
}

// MainIterator test for iterator examples.
func MainIterator() {
	intsone := intsOne{1, 2, 3, 4, 5}
	for iter := intsone.Iterator(); iter.HasNext(); {
		fmt.Println("element:", iter.Next())
	}
	fmt.Println()

	intstwo := intsTwo{2, 4, 6, 8, 10}
	iter := intstwo.Iterator()
	for {
		if val, ok := iter(); !ok {
			break
		} else {
			fmt.Println("element:", val)
		}
	}
	fmt.Println()

	intsthree := intsThree{1, 3, 5, 7, 9}
	for val := range intsthree.Iterator() {
		fmt.Println("element:", val)
	}
	fmt.Println()

	intsfour := intsFour{11, 12, 13, 14, 15}
	intsfour.ForEach(func(ele int) {
		fmt.Println("element:", ele)
	})
}
