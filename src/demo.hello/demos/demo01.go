package demos

import (
	"bytes"
	"context"
	"fmt"
	"image/color"
	"math"
	"strconv"
	"strings"
	"time"
)

// demo, inner function
func testPrintFormatName() {
	printFormatName("zheng", "jin")
}

func printFormatName(firstName, lastName string) {
	fnGetShortName := func(firstName, lastName string) string {
		tmp := fmt.Sprintf("%c%c", firstName[0], lastName[0])
		return strings.ToUpper(tmp)
	}
	fmt.Printf("fname=%s, lname=%s, sname=%s\n",
		firstName, lastName, fnGetShortName(firstName, lastName))
}

// demo, defer and recover()
func testRecoverFromPanic() {
	if ret, err := myDivision(4, 0); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("4/0 results:", ret)
	}

	if ret, err := myDivision(4, 2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("4/2 results:", ret)
	}
}

func myDivision(x, y int) (ret int, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()

	if y == 0 {
		panic("y value is zero!")
	}
	ret = x / y
	return
}

// demo, struct init
type fullName struct {
	fName    string
	lName    string
	nickName string
}

func testPrintStructValue() {
	fullName := fullName{
		fName: "fname",
		lName: "lname",
	}
	fmt.Printf("\nfull name: %v\n", fullName)

	fullName.nickName = "nick"
	fmt.Println("full name with nick name:", fullName)
}

// demo, struct and method
type point struct {
	x, y float64
}

func (p point) distance(q point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p *point) scaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func testStructMethod() {
	fmt.Println("\nstruct method:")
	p := point{1, 2}
	q := point{4, 6}
	fmt.Printf("distance: %.1f\n", p.distance(q))

	fmt.Println("\nmethod by reference:")
	p.scaleBy(3)
	fmt.Printf("scaled p: %v\n", p)
	(&q).scaleBy(2)
	fmt.Printf("scaled q: %v\n", q)

	fmt.Println("\nmethod as variable:")
	distanceFromP := p.distance
	fmt.Printf("distanceFromP: %T\n", distanceFromP)
	fmt.Printf("distance: %.1f\n", distanceFromP(q))

	scaleP := p.scaleBy
	fmt.Printf("scaleP: %T\n", scaleP)
	scaleP(2)
	fmt.Printf("p: %v\n", p)

	fmt.Println("\nmethod as variable:")
	p = point{1, 2}
	q = point{4, 6}
	distance := point.distance
	fmt.Printf("distance() type: %T\n", distance)
	fmt.Printf("distance: %.1f\n", distance(p, q))

	scale := (*point).scaleBy
	scale(&p, 2)
	fmt.Printf("scale() type: %T\n", scale)
	fmt.Printf("scaled p: %v\n", p)
}

// demo, innner struct
type coloredPoint struct {
	point
	Color color.RGBA
}

func testInnerStruct() {
	fmt.Println("\nstruct fields:")
	var cp coloredPoint
	cp.x = 1
	fmt.Printf("point x: %v\n", cp.point.x)
	cp.point.y = 2
	fmt.Printf("point y: %v\n", cp.y)

	fmt.Println("\nstruct method:")
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = coloredPoint{point{1, 1}, red}
	var q = coloredPoint{point{5, 4}, blue}
	fmt.Printf("distance: %.1f\n", p.distance(q.point))
	p.scaleBy(2)
	q.scaleBy(2)
	fmt.Printf("scaled distance: %.1f\n", p.distance(q.point))
}

// demo, method var
func (p point) add(q point) point {
	return point{p.x + q.x, p.y + q.y}
}

func (p point) sub(q point) point {
	return point{p.x - q.x, p.y - q.y}
}

type path []point

func (pa path) translateBy(offset point, add bool) {
	var op func(p, q point) point // method var
	if add {
		op = point.add
	} else {
		op = point.sub
	}
	fmt.Printf("operation type: %T\n", op)

	for i := range pa {
		pa[i] = op(pa[i], offset)
	}
}

func testTranslateBy() {
	var pa path = []point{{1, 2}, {4, 6}}
	pa.translateBy(point{1, 1}, true)

	for _, p := range pa {
		fmt.Printf("point: %v\n", p)
	}
}

// MyObject container public and private fields.
// demo, test access control demo.
type MyObject struct {
	VarPublic  string
	varPrivate string
}

// Init for init MyObject struct.
func (o *MyObject) Init(pub, pri string) {
	o.VarPublic = pub
	o.varPrivate = pri
}

// MethodPublicGet returns public value.
func (o MyObject) MethodPublicGet() string {
	return o.varPrivate
}

func (o MyObject) methodPrivateGet() string {
	return o.VarPublic
}

func testAccControl() {
	fmt.Println("\naccess private fields/methods internal:")
	obj := MyObject{"public_in", "private_in"}
	fmt.Printf("public var: %s\n", obj.VarPublic)
	fmt.Printf("private var: %s\n", obj.varPrivate)
	fmt.Printf("\npublic method get: %s\n", obj.MethodPublicGet())
	fmt.Printf("private method get: %s\n", obj.methodPrivateGet())
}

// GetMyObject invoked from external and returns an empty object.
func GetMyObject() MyObject {
	return MyObject{}
}

// demo, args package and unpackage
func testArgsPkgAndUnpkg() {
	fnSum := func(args ...int32) {
		fmt.Printf("\nargs type: %T\n", args)
		var sum int32
		for _, arg := range args {
			sum = sum + arg
		}
		fmt.Println("Sum:", sum)
	}
	fnSum(1, 2, 3, 4, 5)
	fnSum([]int32{5, 6, 7, 8}...)

	s := []int{1, 2}
	s = append(s, []int{3, 4, 5}...)
	fmt.Println("\nslice:", s)
}

// demo, time format
func testTimeFormat() {
	t := time.Now()
	fmt.Printf("\nweek day: %d, time: %d:%d\n", t.Weekday(), t.Hour(), t.Minute())

	s := strconv.FormatInt(t.Unix(), 10)
	fmt.Println("unix time (seconds from 1970):", s)

	t = time.Unix(t.Unix()+60, 0)
	baseDate := "2006-01-02 15:04:05"
	fmt.Println("cur date (+60s):", t.Format(baseDate))
}

// demo, code block
func testCodeBlock() {
	testSuper := "super"
	fmt.Println("\ntestSupser=" + testSuper)
	{
		testSub1 := "sub1"
		{
			testSub2 := "sub2"
			fmt.Printf("testSupser=%s, testSub1=%s, testSub2=%s\n", testSuper, testSub1, testSub2)
		}
		fmt.Printf("testSupser=%s, testSub1=%s\n", testSuper, testSub1)
		testSuper += ", change in sub"
	}
	fmt.Println("testSupser=" + testSuper)
}

// demo, array and slice
func testArrayAndSlice01() {
	// 数组中元素是值传递, 切片中元素是引用传递
	fnUpdateArray := func(arr [5]int32) {
		fmt.Printf("\n[fnUpdateArray] array: addr=%p, val_addr=%p\n", &arr, arr)
		arr[1] = 123
		for i := 0; i < len(arr); i++ {
			fmt.Printf("[fnUpdateArray] array item: val=%d, addr=%p\n", arr[i], &arr[i])
		}
	}

	fnUpdateSlice := func(s []int32) {
		fmt.Printf("\n[fnUpdateSlice] slice: addr=%p, val_addr=%p\n", &s, s)
		s[1] = 456
		for i := 0; i < len(s); i++ {
			fmt.Printf("[fnUpdateSlice] slice item: val=%d, addr=%p\n", s[i], &s[i])
		}
	}

	// #1
	var array1 = [...]int32{1, 2, 3, 4, 5}
	fmt.Printf("\narray1: addr=%p, val_addr=%p\n", &array1, array1)
	for i := 0; i < len(array1); i++ {
		fmt.Printf("array1 item: val=%d, addr=%p\n", array1[i], &array1[i])
	}

	array2 := array1
	array2[0] = 100
	fmt.Printf("\narray2, addr=%p, val_addr=%p\n", &array2, array2)
	for i := 0; i < len(array2); i++ {
		fmt.Printf("array2 item: val=%d, addr=%p\n", array2[i], &array2[i])
	}

	fnUpdateArray(array1)
	fmt.Printf("\narray1: %v\ncopied array2: %v\n", array1, array2)

	// #2
	var slice1 = []int32{1, 2, 3, 4, 5}
	fmt.Printf("\nslice1: addr=%p, val_addr=%p\n", &slice1, slice1)
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1 item: val=%d, addr=%p\n", slice1[i], &slice1[i])
	}

	slice2 := slice1
	slice2[0] = 200
	fmt.Printf("\nslice2, addr=%p, val_addr=%p\n", &slice2, slice2)
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("slice2 item: val=%d, addr=%p\n", slice2[i], &slice2[i])
	}

	fnUpdateSlice(slice1)
	fmt.Printf("\nslice: %v, copied slice2: %v\n", slice1, slice2)
}

// demo, array and slice
func testArrayAndSlice02() {
	// #1
	var array1 = [...]int32{1, 2, 3, 4, 5}
	fmt.Printf("\narray type: %T\n", array1)

	s1 := array1[2:4]
	fmt.Printf("\nbefore: array=%v, slice=%v\n", array1, s1)
	s1[0] = 100
	fmt.Printf("after: array=%v, slice=%v\n", array1, s1)

	// #2
	array2 := [...]int32{1, 2, 3, 4, 5}
	fmt.Println("\nexceed cap, and slice re-allocate")
	s2 := array2[2:4]
	fmt.Printf("s2 addr: %p, items:\n", &s2)
	for i := 0; i < len(s2); i++ {
		fmt.Printf("s2 item: addr=%p, val=%d\n", &s2[i], s2[i])
	}

	s2 = append(s2, 6, 7)
	fmt.Printf("\nnew s2 addr: %p, items:\n", &s2)
	for i := 0; i < len(s2); i++ {
		fmt.Printf("s2 item: addr=%p, val=%d\n", &s2[i], s2[i])
	}

	fmt.Printf("\nbefore: array=%v, slice=%v\n", array2, s2)
	s2[0] = 200
	fmt.Printf("after: array=%v, slice=%v\n", array2, s2)
}

// demo, map var reference
func testMapReference() {
	// map中val元素是引用传递, 但map[key]则返回一个新的元素
	fnUpdateMap := func(m map[int]string) {
		fmt.Printf("[fnUpdateMap] map val_addr=%p\n", m)
		m[1] = "One"
	}

	m := make(map[int]string)
	m[1] = "one"
	m[2] = "two"
	m[3] = "three"
	fmt.Printf("\nsrc map val_addr=%p\n", m)

	mCopied := m
	mCopied[3] = "THREE"
	mCopied[4] = "four"
	fnUpdateMap(m)
	m[5] = "five"
	fmt.Printf("cpoied map val_addr=%p\n", mCopied)

	key := 1
	fmt.Println("\n[main]")
	if item, ok := m[key]; ok {
		fmt.Printf("src map[%d]=%s, addr=%p\n", key, item, &item)
	}
	if item, ok := m[key]; ok {
		fmt.Printf("src map[%d]=%s, addr=%p\n", key, item, &item)
	}

	if item, ok := mCopied[key]; ok {
		fmt.Printf("copied map[%d]=%s, addr=%p\n", key, item, &item)
	}
}

// demo, struct var reference
func testStructReference() {
	// struct中元素是值传递
	type people struct {
		id  string
		age int
		job string
	}

	fnUpdateStruct := func(p people) {
		fmt.Printf("\n[fnUpdateStruct] p.job addr: %p\n", &p.job)
		p.id = "002"
		p.age = 35
		fmt.Printf("[fnUpdateStruct] p: %+v\n", p)
	}

	p1 := people{
		id:  "001",
		age: 30,
		job: "tester",
	}
	p2 := p1
	p2.job = "developer"

	fnUpdateStruct(p1)
	fmt.Println("[main]")
	fmt.Printf("p1.job addr: %p\n", &p1.job)
	fmt.Printf("p2.job addr: %p\n", &p2.job)
	fmt.Printf("p1: %+v, p2: %+v\n", p1, p2)
}

// demo, update bytes
func testUpdateBytes() {
	b := []byte("hello world")
	fmt.Printf("\n[main] b: addr=%p, val_addr=%p, len=%d, cap=%d\n", &b, b, len(b), cap(b))

	fmt.Println("\n#1: by value:")
	fmt.Printf("before update: %s\n", string(b))
	myUpdateBytesByVal(b)
	fmt.Printf("after update: %s\n", string(b))

	fmt.Println("\n#2: by reference:")
	fmt.Printf("before update: %s\n", string(b))
	myUpdateBytesByRef(&b)
	fmt.Printf("after update: %s\n", string(b))
}

func myUpdateBytesByVal(b []byte) {
	fmt.Printf("[myUpdateBytesByVal] b: val_addr=%p, len=%d, cap=%d\n", b, len(b), cap(b))
	b[0] = 'H'
	b = bytes.ToUpper(b) // re-allocate
	fmt.Printf("[myUpdateBytesByVal] b: val_addr=%p, len=%d, cap=%d\n", b, len(b), cap(b))
	b = append(b, '!', '!') // re-size
	fmt.Printf("[myUpdateBytesByVal] b: val_addr=%p, len=%d, cap=%d\n", b, len(b), cap(b))
	fmt.Printf("[myUpdateBytesByVal] by value: %s\n", string(b))
}

func myUpdateBytesByRef(b *[]byte) {
	fmt.Printf("[myUpdateBytesByRef] b: val_addr=%p, len=%d, cap=%d\n", b, len(*b), cap(*b))
	*b = bytes.ToUpper(*b)
	fmt.Printf("[myUpdateBytesByRef] b: val_addr=%p, len=%d, cap=%d\n", b, len(*b), cap(*b))
	*b = append(*b, []byte{'!', '!', '!'}...)
	fmt.Printf("[myUpdateBytesByRef] b: val_addr=%p, len=%d, cap=%d\n", b, len(*b), cap(*b))
	// for pointer, cannot get slice item by index
	// fmt.Printf("[myUpdateBytesByRef] b[0]: addr=%v", *b[0])
	fmt.Printf("[myUpdateBytesByRef] by reference: %s\n", string(*b))
}

// demo, copy bytes
func testCopyBytes() {
	updateBytesByCopy := func(b []byte) {
		fmt.Printf("[updateBytesByCopy] b: addr=%p, val_addr=%p\n", &b, b)
		copy(b, []byte("hello")) // 内存中, 覆盖原字符数组中的字符
		fmt.Printf("[updateBytesByCopy] copied b: addr=%p, val_addr=%p\n", &b, b)
	}

	b := []byte("test_go")
	fmt.Printf("\nb: val_addr=%p, type=%T, len=%d, cap=%d\n", b, b, len(b), cap(b))
	updateBytesByCopy(b)
	fmt.Println("copied b val:", string(b))

	b = make([]byte, 3)
	fmt.Printf("\nmake b: len=%d, cap=%d, val=%v\n", len(b), cap(b), b)
}

// demo, context
func testContext01() {
	var (
		a = 1
		b = 2
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Second) // timeout
	res := myAdd(ctx, a, b)
	go func() {
		time.Sleep(time.Duration(5) * time.Second)
		cancel()
	}()
	fmt.Printf("\nCompute: %d+%d, result: %d\n", a, b, res)
}

func testContext02() {
	var (
		a = 1
		b = 2
	)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		cancel()
	}()
	res := myAdd(ctx, a, b)
	fmt.Printf("\nCompute: %d+%d, result: %d\n", a, b, res)
}

func myAdd(ctx context.Context, a, b int) int {
	res := 0
	for i := 0; i < a; i++ {
		res = incr(res)
		select {
		case <-ctx.Done():
			fmt.Println("a: cancel incr()")
			return -1
		default:
		}
	}

	for i := 0; i < b; i++ {
		res = incr(res)
		select {
		case <-ctx.Done():
			fmt.Println("b: cancel incr()")
			return -1
		default:
		}
	}
	return res
}

func incr(a int) int {
	time.Sleep(time.Second)
	return a + 1
}

// MainDemo01 main for golang demo01.
func MainDemo01() {
	// testPrintFormatName()
	// testPrintStructValue()
	// testRecoverFromPanic()

	// testStructMethod()
	// testInnerStruct()
	// testAccControl()

	// testTranslateBy()
	// testArgsPkgAndUnpkg()
	// testTimeFormat()
	// testCodeBlock()

	// testArrayAndSlice01()
	// testArrayAndSlice02()
	// testMapReference()
	// testStructReference()

	// testUpdateBytes()
	// testCopyBytes()

	// testContext01()
	// testContext02()

	fmt.Println("golang demo01 DONE.")
}
