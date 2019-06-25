package demos

import (
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"strings"
	"time"
)

// demo 01, map
func testCheckMapEntry() {
	m := map[int]string{
		1: "one",
		2: "two",
	}
	fmt.Println("\nentry[2] value:", m[2])
	fmt.Printf("entry[2] 1st char: %c\n", m[2][0])
	fmt.Printf("entry[2] length: %d\n", len(m[2]))

	if entry, ok := m[3]; ok {
		fmt.Println("entry[3] value:", entry)
	}
}

// demo 02-01, custom reader
type alphaReader1 struct {
	src string
	cur int
}

// Read : read bytes from current position, and copy to p.
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

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func newAlphaReader1(src string) *alphaReader1 {
	return &alphaReader1{src: src}
}

func testAlphaReader1() {
	reader := newAlphaReader1("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	var b []byte

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		// fmt.Print(string(p[:n]))
		b = append(b, p[:n]...)
	}
	fmt.Println("\noutput:", string(b))
}

// demo 02-02, custom reader
type alphaReader2 struct {
	reader io.Reader
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

func newAlphaReader2(reader io.Reader) *alphaReader2 {
	return &alphaReader2{reader: reader}
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

func newChanWriter() *chanWriter {
	// return &chanWriter{ch: make(chan byte, 256)}
	return &chanWriter{make(chan byte, 256)}
}

func testChanWriter() {
	writer := newChanWriter()

	for i := 0; i < 10; i++ {
		go func(idx int) {
			writer.Write([]byte(fmt.Sprintf("Stream%d:", idx)))
			writer.Write([]byte("data\n"))
		}(i)
	}

	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("close chan writer")
		writer.Close()
	}()

	for c := range writer.Chan() {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

// demo 03-01, time ticker in select block
func testSelectTimeTicker01() {
	ticker := time.NewTicker(time.Duration(3) * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case time := <-ticker.C:
			fmt.Printf("ticker time: %v\n", time)
		default: // not block
			fmt.Println("wait 1 sec...")
			time.Sleep(time.Second)
		}
	}
	ticker.Stop()
}

func testSelectTimeTicker02() {
	tick := time.Tick(time.Duration(3) * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case time := <-tick:
			fmt.Printf("tick time: %d:%d\n", time.Hour(), time.Minute())
		default: // not block
			fmt.Println("wait 1 sec...")
			time.Sleep(time.Second)
		}
	}
}

// demo 03-02, time after in select block
func testSelectTimeAfter() {
	ch := make(chan string)
	go func() {
		wait := 10
		fmt.Printf("wait %d second in go routine...\n", wait)
		time.Sleep(time.Duration(wait) * time.Second)
		ch <- "done"
	}()

	select {
	case ret := <-ch:
		fmt.Println("return from routine:", ret)
	case <-time.After(3 * time.Second):
		fmt.Printf("3 seconds timeout!\n")
	}
}

// demo 04, channel queue
func testChanQueue() {
	const cap = 5
	queue := make(chan int, cap)
	for i := 0; i < cap; i++ {
		queue <- rand.Intn(10)
		time.Sleep(time.Duration(300) * time.Millisecond)
	}

	go func() {
		for i := 0; i < 10; i++ {
			queue <- rand.Intn(20)
			time.Sleep(time.Duration(300) * time.Millisecond)
		}
		close(queue)
	}()

	fmt.Println("queue value:")
	for v := range queue {
		fmt.Println(v)
	}
}

// demo 05, bufferred channel
func testBufferedChan() {
	queue := make(chan int, 10)
	go func() {
		producers(queue)
	}()
	go func() {
		consumer(queue)
	}()

	for i := 0; i < 15; i++ {
		fmt.Println("queue size:", len(queue))
		time.Sleep(time.Second)
	}
	fmt.Println("close queue")
	close(queue)
}

func producers(queue chan<- int) {
	for {
		select {
		case queue <- rand.Intn(10):
			fmt.Println("true => enqueued without blocking")
		default:
			fmt.Println("false => not enqueued, would have blocked because of queue is full")
		}
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
}

func consumer(queue <-chan int) {
	// OUTER:
	for {
		select {
		case item, valid := <-queue:
			if valid {
				fmt.Println("ok && valid => item is good, use it")
				fmt.Printf("pop off item: %d\n", item)
			} else {
				fmt.Println("ok && !valid => channel closed, quit polling")
			}
			// break OUTER
		default:
			fmt.Println("!ok => channel open, but empty, try later")
		}
		time.Sleep(time.Second)
	}
}

// demo 06, iterator for chars
func testIteratorChars() {
	s := "hello"
	for _, c := range s {
		fmt.Printf("%c", c)
	}
	fmt.Println()

	b := []byte("world")
	fmt.Printf("b type: %T\n", b) // type: []uint8
	for _, c := range b {
		fmt.Printf("%c", c)
	}
	fmt.Println()
}

// demo 07, function as variable
func testFuncVariable() {
	fmt.Printf("\nadd results: %d\n", myCalculation(2, 2, funcMyAdd))
	fmt.Printf("min results: %d\n", myCalculation(2, 8, funcMyMin))
}

func funcMyAdd(num1, num2 int) int {
	return num1 + num2
}

func funcMyMin(num1, num2 int) int {
	ret := num1 - num2
	return int(math.Abs(float64(ret)))
}

func myCalculation(num1, num2 int, fnCal func(n1, n2 int) int) int {
	return fnCal(num1, num2)
}

// demo 08, function decoration
type apiResponse struct {
	RetCode uint16
	Body    string
	Err     error
}

type apiArgsUser struct {
	UID      uint32
	UserName string
}

func mockAPIPass(args interface{}) *apiResponse {
	info := args.(apiArgsUser)
	content := fmt.Sprintf("user info: Uid=%d, name=%s", info.UID, info.UserName)
	return &apiResponse{
		RetCode: 200,
		Body:    content,
		Err:     nil,
	}
}

type apiArgsGroup struct {
	GID       uint32
	GroupName string
}

func mockAPIFailed(args interface{}) *apiResponse {
	info := args.(apiArgsGroup)
	content := fmt.Sprintf("group not found: Gid=%d, name=%s", info.GID, info.GroupName)
	return &apiResponse{
		RetCode: 204,
		Body:    content,
		Err:     errors.New("EOF"),
	}
}

func testDecorateAPIs() {
	fmt.Println("\n#1. decoration sample: pass")
	{
		args := apiArgsUser{
			UID:      101,
			UserName: "Henry",
		}
		resp := assertAPIs(args, mockAPIPass)
		fmt.Println("pass with resp body:", resp.Body)
	}

	fmt.Println("\n#2. decoration sample: failed")
	{
		args := apiArgsGroup{
			GID:       8,
			GroupName: "QA",
		}
		resp := assertAPIs(args, mockAPIFailed)
		fmt.Println("failed with resp body:", resp.Body)
	}
}

// decoration 装饰器
func assertAPIs(args interface{}, fn func(args interface{}) *apiResponse) *apiResponse {
	resp := fn(args)
	fmt.Printf("response: %+v\n", *resp)
	if resp.RetCode != 200 {
		fmt.Println("failed with ret code:", resp.RetCode)
	}
	if resp.Err != nil {
		fmt.Println("failed with error:", resp.Err.Error())
	}

	return resp
}

// MainDemo03 : main
func MainDemo03() {
	// testCheckMapEntry()

	// testAlphaReader1()
	// testAlphaReader2()
	// testChanWriter()

	// testSelectTimeTicker01()
	// testSelectTimeTicker02()
	// testSelectTimeAfter()

	// testChanQueue()
	// testBufferedChan()

	// testIteratorChars()
	// testFuncVariable()
	// testDecorateAPIs()

	fmt.Println("golang demo03 DONE.")
}
