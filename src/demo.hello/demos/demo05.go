package demos

import (
	"bytes"
	"fmt"
	"io"
	"os"
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

// MainDemo05 main for golang demo05.
func MainDemo05() {
	testVarTypeAssert()
	// testPointerVar01()
	// testPointerVar02()

	fmt.Println("golang demo05 DONE.")
}
