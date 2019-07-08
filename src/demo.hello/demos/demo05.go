package demos

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

// demo, type assert
func testVarTypeAssert() {
	// #1, type assert failed
	var w io.Writer = os.Stdout
	if b, ok := w.(*bytes.Buffer); ok {
		data := make([]byte, 10)
		if n, err := b.Read(data); err != nil {
			panic(err)
		} else {
			fmt.Println("read bytes count:", n)
		}
	} else {
		fmt.Println("\nnot support Reader!")
	}

	// #2, type assert pass
	var r io.Reader = bytes.NewBufferString("hello ")
	if b, ok := r.(io.Writer); ok {
		if _, err := b.Write([]byte("world!")); err != nil {
			panic(err)
		} else {
			fmt.Fprintln(os.Stdout, b)
		}
	} else {
		fmt.Println("not support Writer!")
	}
}

// demo, test pointer variable
func testPointerVar01() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("\narray a: addr=%p, addr_val=%p, val=%v\n", &a, a, a)

	p1 := &a
	fmt.Printf("pointer p1: addr=%p, addr_val=%p, val=%d\n", &p1, p1, *p1)
	p1[0] = 10

	var p2 = new([5]int)
	*p2 = a // value copy
	fmt.Printf("pointer p2: addr=%p, addr_val=%p, val=%d\n", &p2, p2, *p2)
	p2[0] = 100

	fmt.Printf("src a: %v, p1: %v, p2: %v\n", a, *p1, *p2)
}

// demo, test pointer variable
func testPointerVar02() {
	// note: here, var "s" like a pointer to slice, but it's used as a slice, and not a pointer.
	s := make([]int, 0, 1)
	s = append(s, 1)
	fmt.Printf("\nslice s: type=%T, addr=%p, addr_val=%p, val=%v, len=%d, cap=%d\n",
		s, &s, s, s, len(s), cap(s))

	// #1
	p := &s
	fmt.Printf("\npointer p: type=%T, addr=%p, addr_val=%p, val=%v\n", p, &p, p, *p)
	// error: not support index
	// fmt.Println(*p[10])

	*p = append(*p, 2) // change s
	fmt.Printf("pointer p: addr=%p, addr_val=%p, val=%v\n", &p, p, *p)
	fmt.Printf("slice s: addr=%p, addr_val=%p, val=%v, len=%d, cap=%d\n", &s, s, s, len(s), cap(s))

	// #2
	s1 := s
	fmt.Printf("\nslice s1: addr=%p, addr_val=%p, val=%v, len=%d, cap=%d\n", &s1, s1, s1, len(s1), cap(s1))
	s1[0] = 10               // change s
	s1 = append(s1, 3, 4, 5) // not change s
	fmt.Printf("slice s1: addr=%p, addr_val=%p, val=%v, len=%d, cap=%d\n", &s1, s1, s1, len(s1), cap(s1))
	fmt.Printf("slice s: addr=%p, addr_val=%p, val=%v, len=%d, cap=%d\n", &s, s, s, len(s), cap(s))

	fmt.Printf("\nsrc s: %v, p: %v, s1: %v\n", s, *p, s1)
}

// demo, panic in routine
func testPanicInRoutine() {
	fmt.Println("start go routine and panic.")
	ch := make(chan int)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				fmt.Println("routine internal error:", p)
				close(ch)
			}
		}()

		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			ch <- i
			if i == 3 {
				panic(fmt.Errorf("mock panic in routine"))
			}
		}
	}()

	for i := range ch {
		fmt.Println("receive:", i)
	}
	fmt.Println("main routine done.")
}

// demo, WaitGroup
func testSyncWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Printf("go routine 1_%d is running...\n", idx)
			time.Sleep(time.Duration(2) * time.Second)
		}(i, &wg)
	}

	// use WaitGroup as global var
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fmt.Printf("go routine 2_%d is running...\n", idx)
			time.Sleep(time.Duration(3) * time.Second)
		}(j)
	}

	time.Sleep(time.Duration(500) * time.Millisecond)
	fmt.Println("go routines count:", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("all go routines are done")
}

// demo, invoke method which receiver as value or reference
type myCalculation struct {
	val int
}

func (c *myCalculation) Greater(i int) bool {
	return i > c.val
}

func (c myCalculation) Less(i int) bool {
	return i < c.val
}

func testMethodsWithDiffRev() {
	const base = 3
	calRef := &myCalculation{val: base}
	fmt.Println("\n4 > base:", calRef.Greater(4))
	fmt.Println("2 < base:", calRef.Less(2))

	calVal := myCalculation{val: base}
	fmt.Println("2 > base:", calVal.Greater(2))
	fmt.Println("5 < base:", calVal.Less(5))
}

// MainDemo05 main for golang demo05.
func MainDemo05() {
	// testVarTypeAssert()
	// testPointerVar01()
	// testPointerVar02()

	testPanicInRoutine()
	// testSyncWaitGroup()

	// testMethodsWithDiffRev()

	fmt.Println("golang demo05 DONE.")
}
