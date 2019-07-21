package demos

import (
	"bufio"
	"bytes"
	"container/heap"
	"container/list"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
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
			// if not recover(), panic in routine will cause main routine exit with 2.
			// after panic(), defer func() will be invoked and we can check error by recover().
			if p := recover(); p != nil {
				fmt.Println("routine internal error:", p)
			}
			fmt.Println("routine exit, and close channel.")
			close(ch)
		}()

		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			ch <- i
			if i == 3 {
				panic(fmt.Errorf("mock panic"))
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
	fmt.Println("all go routines are done.")
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

// demo, inherit in struct
type t1 struct {
	s1 string
}

func (*t1) f1() {
	fmt.Println("t1.f1")
}

type t2 struct {
	s2 string
}

func (*t2) f2() {
	fmt.Println("t2.f2")
}

type ts struct {
	*t1 // pointer
	t2  // value
}

func testStructInherit() {
	t := ts{
		t1: &t1{"foo"},
		t2: t2{"bar"},
	}
	t.f1()
	t.f2()
	fmt.Println(t.s1)
	fmt.Println(t.s2)
}

// demo, bufio Writer
type myWriter struct{}

func (myWriter) Write(b []byte) (n int, err error) {
	fmt.Println(len(b))
	return len(b), nil
}

func testBufioWriter() {
	fmt.Println("\n|", strings.Repeat("-", 30))
	fmt.Println("Unbuffered I/O")
	w := myWriter{}
	w.Write([]byte{'a'})
	w.Write([]byte{'b'})
	w.Write([]byte{'c'})
	w.Write([]byte{'d'})

	fmt.Println("Buffered I/O")
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte{'a'})
	bw.Write([]byte{'b'})
	bw.Write([]byte{'c'})
	bw.Write([]byte{'d'})

	fmt.Println("buffered bytes:", bw.Buffered())
	fmt.Println("buffered available:", bw.Available())
	if err := bw.Flush(); err != nil {
		panic(err)
	}
}

// demo, bufio Reader (WriterTo)
type myReader struct {
	n int
}

func (r *myReader) Read(p []byte) (n int, err error) {
	fmt.Printf("read #%d\n", r.n)
	if r.n >= 10 {
		return 0, io.EOF
	}

	copy(p, "abcd")
	r.n++
	return 4, nil
}

func testBufioReader() {
	r := &myReader{}
	br := bufio.NewReaderSize(r, 16)
	n, err := br.WriteTo(ioutil.Discard)
	if err != nil {
		panic(err)
	}
	fmt.Printf("written bytes: %d\n", n)
}

// demo, bufio ReadWriter
func testBufioRW() {
	s := strings.NewReader("abcd")
	br := bufio.NewReaderSize(s, 4)
	w := new(bytes.Buffer)
	bw := bufio.NewWriterSize(w, 4)
	rw := bufio.NewReadWriter(br, bw)

	// buffer read
	b := make([]byte, 2)
	if _, err := rw.Read(b); err != nil {
		panic(err)
	}
	fmt.Printf("\nread bytes: %q\n", b)
	fmt.Println("reader buffer:", rw.Reader.Buffered())

	// buffer write
	if _, err := rw.Write([]byte("efgh")); err != nil {
		panic(err)
	}
	if err := rw.Flush(); err != nil {
		panic(err)
	}
	fmt.Println("write bytes:", w.String())
	fmt.Println("writer buffer:", rw.Writer.Buffered())
}

// demo, container/list 双向链表
func testContainerList() {
	fnPrint := func(l *list.List) {
		fmt.Println("\nlist elements:")
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}
	}

	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	fnPrint(l)

	l.PushFront(0)
	fnPrint(l)

	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == 1 {
			l.InsertAfter(1.1, e)
		}
		if e.Value == 2 {
			l.InsertBefore(1.2, e)
		}
	}
	fnPrint(l)

	fmt.Println("\nlist elements in reserve:")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}

// demo, container/heap 最小二叉树
type student struct {
	name  string
	score int
}

type studentHeap []student

func (h studentHeap) Len() int {
	return len(h)
}

func (h studentHeap) Less(i, j int) bool {
	return h[i].score < h[j].score // 最小堆
}

func (h studentHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *studentHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length, not just its contents.
	*h = append(*h, x.(student))
}

func (h *studentHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func testContainerHeap() {
	fnPrint := func(h *studentHeap) {
		fmt.Println("\nstudent heap items:")
		for _, ele := range *h {
			fmt.Printf("student name %s, score %d\n", ele.name, ele.score)
		}
	}

	h := &studentHeap{
		{name: "xiaoming", score: 82},
		{name: "xiaozhang", score: 88},
		{name: "laowang", score: 85},
	}
	heap.Init(h)
	heap.Push(h, student{name: "xiaoli", score: 66})
	fnPrint(h)

	for i, ele := range *h {
		if ele.name == "xiaozhang" {
			(*h)[i].score = 60
			heap.Fix(h, i)
		}
	}
	fnPrint(h)

	fmt.Println("\nstudent heap pop items:")
	for h.Len() > 0 {
		item := heap.Pop(h).(student)
		fmt.Printf("student name %s,score %d\n", item.name, item.score)
	}
}

// demo, sync Cond
func testSyncCond() {
	var (
		count = 3
		num   = 3
	)
	ch := make(chan struct{}, num)

	var lock sync.Mutex
	cond := sync.NewCond(&lock)

	for i := 0; i < num; i++ {
		go func(i int) {
			cond.L.Lock()
			defer func() {
				cond.L.Unlock()
				ch <- struct{}{}
			}()

			for i < count {
				fmt.Printf("goroutine_%d start and wait\n", i)
				cond.Wait()
				fmt.Printf("goroutine_%d receive a notify\n", i)
			}
			fmt.Printf("goroutine_%d end\n", i)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("broadcast...")
	cond.L.Lock()
	count--
	cond.Broadcast()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("signal...")
	cond.L.Lock()
	count--
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(time.Second)
	fmt.Println("broadcast...")
	cond.L.Lock()
	count--
	cond.Broadcast()
	cond.L.Unlock()

	for i := 0; i < num; i++ {
		<-ch
	}
}

// demo, sync atomic add
func testSyncAtomicAdd() {
	var ops uint32

	for i := 0; i < 10; i++ {
		go func() {
			for {
				atomic.AddUint32(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	finalOps := atomic.LoadUint32(&ops)
	fmt.Println("\nfinal ops:", finalOps)
}

// demo, interface reflection
type user struct {
	ID   int
	Name string
	Age  int
}

func (u user) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func testStructReflect() {
	u := user{1, "Allen.Wu", 25}
	printFieldsAndMethod(u)
}

func printFieldsAndMethod(input interface{}) {
	rType := reflect.TypeOf(input)
	fmt.Println("interface type:", rType)
	rValue := reflect.ValueOf(input)
	fmt.Println("interface value:", rValue)

	for i := 0; i < rType.NumField(); i++ {
		field := rType.Field(i)
		value := rValue.Field(i).Interface()
		fmt.Printf("%s (%v) = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < rType.NumMethod(); i++ {
		m := rType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

// demo, reflection, update value by ref
func testUpdateValByReflect() {
	var num float32 = 1.2345
	fmt.Println("\nold value of number:", num)

	pointer := reflect.ValueOf(&num)
	elem := pointer.Elem()
	fmt.Println("type of point element:", elem.Type())
	fmt.Println("settability of element:", elem.CanSet())

	if elem.CanSet() {
		elem.SetFloat(6.66)
		fmt.Println("new value of number:", num)
	}
}

// MainDemo05 main for golang demo05.
func MainDemo05() {
	// testVarTypeAssert()
	// testPointerVar01()
	// testPointerVar02()

	// testPanicInRoutine()
	// testSyncWaitGroup()

	// testMethodsWithDiffRev()
	// testStructInherit()

	// testBufioWriter()
	// testBufioReader()
	// testBufioRW()

	// testContainerList()
	// testContainerHeap()

	// testSyncCond()
	// testSyncAtomicAdd()

	// testStructReflect()
	// testUpdateValByReflect()

	fmt.Println("golang demo05 DONE.")
}
