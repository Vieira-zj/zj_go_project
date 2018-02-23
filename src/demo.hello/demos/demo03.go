package demos

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

// demo 01, map
func testMapGetEmpty() {
	m := map[int]string{
		1: "one",
		2: "two",
	}
	fmt.Println("item at 2 =>", m[2])
	fmt.Printf("first char: %c\n", m[2][0])
	fmt.Printf("item length: %d\n", len(m[2]))

	if len(m) > 0 && len(m[3]) > 0 {
		fmt.Println("item at 3 =>", m[3])
	}
}

// demo 02-01, custom reader
type alphaReader1 struct {
	src string
	cur int
}

func newAlphaReader1(src string) *alphaReader1 {
	return &alphaReader1{src: src}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader1) Read(p []byte) (int, error) {
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	x := len(a.src) - a.cur
	bound := 0
	if x >= len(p) {
		bound = len(p)
	} else {
		bound = x
	}

	buf := make([]byte, bound)
	for n := 0; n < bound; n++ {
		if char := alpha(a.src[a.cur]); char != 0 {
			buf[n] = char
		}
		a.cur++
	}
	copy(p, buf)
	return bound, nil
}

func testAlphaReader1() {
	reader := newAlphaReader1("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		// fmt.Printf("%d\n", n)
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}

// demo 02-02, custom reader
type alphaReader2 struct {
	reader io.Reader
}

func newAlphaReader2(reader io.Reader) *alphaReader2 {
	return &alphaReader2{reader: reader}
}

func (a *alphaReader2) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}

	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if char := alpha(p[i]); char != 0 {
			buf[i] = char
		}
	}
	copy(p, buf)
	return n, nil
}

func testAlphaReader2() {
	reader := newAlphaReader2(strings.NewReader("Hello! It's 9am, where is the sun?"))
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Print(string(p[:n]))
				break
			}
			panic(err.Error())
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}

// demo 02-03, custom writer
type chanWriter struct {
	ch chan byte
}

func newChanWriter() *chanWriter {
	return &chanWriter{make(chan byte, 1024)}
}

func (w *chanWriter) Chan() <-chan byte {
	return w.ch
}

func (w *chanWriter) Write(p []byte) (int, error) {
	n := 0
	for _, b := range p {
		w.ch <- b
		n++
	}
	return n, nil
}

func (w *chanWriter) Close() error {
	close(w.ch)
	return nil
}

func testChanWriter() {
	writer := newChanWriter()
	go func() {
		defer writer.Close()
		writer.Write([]byte("Stream "))
		writer.Write([]byte("me"))
	}()
	for c := range writer.Chan() {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

// demo 03-01, time ticker in select block
func testSelectForTimeTicker() {
	ticker := time.NewTicker(3 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case time := <-ticker.C:
			fmt.Printf("%v\n", time)
		default: // not block
			fmt.Println("wait...")
			time.Sleep(time.Second)
		}
	}
	ticker.Stop()
}

// demo 03-02, time after in select block
func testSelectForTimeAfter() {
	ch := make(chan string)

	go func() {
		wait := 20
		fmt.Printf("wait %d second in go routine...\n", wait)
		time.Sleep(time.Duration(wait) * time.Second)
		ch <- "done"
	}()

	select {
	case ret := <-ch:
		fmt.Println("return from go routine:", ret)
	case <-time.After(5 * time.Second):
		fmt.Println("timeout for return from go routine")
	}
}

// demo 04, channel queue
func testChanQueue() {
	const total = 5
	queue := make(chan int, total)
	for i := 0; i < total; i++ {
		queue <- rand.Intn(10)
		time.Sleep(300 * time.Millisecond)
	}

	go func() {
		for i := 0; i < 10; i++ {
			queue <- rand.Intn(20)
			time.Sleep(300 * time.Millisecond)
		}
		close(queue)
	}()

	for v := range queue {
		fmt.Printf("queue value: %d\n", v)
	}
}

// demo 05, buffered channel
func producers(queue chan int) {
	item := rand.Intn(10)
OUTER:
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		select {
		case queue <- item:
			fmt.Println("true => enqueued without blocking")
			break OUTER
		default:
			fmt.Println("false => not enqueued, would have blocked because of queue full")
		}
	}
}

func consumer(queue chan int) {
OUTER:
	for i := 0; i < 3; i++ {
		select {
		case item, valid := <-queue:
			if valid {
				fmt.Println("ok && valid => item is good, use it")
				fmt.Printf("pop off item: %d\n", item)
			} else {
				fmt.Println("ok && !valid => channel closed, quit polling")
			}
			break OUTER
		default:
			fmt.Println("!ok => channel open, but empty, try later")
		}
		time.Sleep(time.Second)
	}
}

func testBufferedChan() {
	queue := make(chan int, 3)
	count := 6

	go func() {
		for i := 0; i < count; i++ {
			producers(queue)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < count; i++ {
			time.Sleep(2 * time.Second)
			consumer(queue)
		}
	}()

	time.Sleep(15 * time.Second)
	close(queue)
}

// demo 06, iterator for chars
func testIteratorChars() {
	s := "hello"
	for _, c := range s {
		fmt.Printf("%c", c)
	}
	fmt.Println()

	b := []byte("world")
	for _, c := range b {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

// demo 07, function as argument
func fnMyAdd(num1 int, num2 int) int {
	return num1 + num2
}

func fnMyMin(num1 int, num2 int) int {
	if num1 > num2 {
		return num1 - num2
	}
	return num2 - num1
}

func myCalculation(testNum1 int, testNum2 int, fnCal func(num1 int, num2 int) int) int {
	return fnCal(testNum1, testNum2)
}

func testFuncPointer() {
	ret := myCalculation(2, 2, fnMyAdd)
	fmt.Printf("results for add: %d\n", ret)

	ret = myCalculation(8, 2, fnMyMin)
	fmt.Printf("results for min: %d\n", ret)
}

// MainDemo03 : main
func MainDemo03() {
	// testMapGetEmpty()

	// testAlphaReader1()
	// testAlphaReader2()
	// testChanWriter()

	// testSelectForTimeTicker()
	// testSelectForTimeAfter()

	// testChanQueue()
	// testBufferedChan()

	// testIteratorChars()
	// testFuncPointer()

	fmt.Println("demo 03 done.")
}
