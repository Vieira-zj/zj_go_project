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
func testRecover() {
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
	var pa path = []point{point{1, 2}, point{4, 6}}
	pa.translateBy(point{1, 1}, true)

	for _, p := range pa {
		fmt.Printf("point: %v\n", p)
	}
}

// MyObject test access control demo.
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

// GetMyObject returns an empty object.
func GetMyObject() MyObject {
	return MyObject{}
}

func testAccControl() {
	fmt.Println("\naccess private fields/methods internal:")
	obj := MyObject{"public_in", "private_in"}
	fmt.Printf("public var: %s\n", obj.VarPublic)
	fmt.Printf("private var: %s\n", obj.varPrivate)
	fmt.Printf("\npublic method get: %s\n", obj.MethodPublicGet())
	fmt.Printf("private method get: %s\n", obj.methodPrivateGet())
}

// demo, context
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
	ret := a + 1
	time.Sleep(time.Duration(1) * time.Second)
	return ret
}

func testContext01() {
	a := 1
	b := 2
	timeout := time.Duration(2) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	res := myAdd(ctx, 1, 2)
	go func() {
		time.Sleep(time.Duration(5) * time.Second)
		cancel()
	}()
	fmt.Printf("\nCompute: %d+%d, result: %d\n", a, b, res)
}

func testContext02() {
	a := 1
	b := 2
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		cancel()
	}()
	res := myAdd(ctx, 1, 2)
	fmt.Printf("\nCompute: %d+%d, result: %d\n", a, b, res)
}

// demo, args package and unpackage
func testArgsPkgAndUnpkg() {
	// package
	fnSum := func(args ...int32) {
		fmt.Printf("\nargs type: %T\n", args)
		var sum int32
		for _, arg := range args {
			sum = sum + arg
		}
		fmt.Println("Sum:", sum)
	}
	fnSum(1, 2, 3, 4, 5)

	// unpackage
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
		fmt.Printf("[fnUpdateArray] array address: %p\n", &arr)
		arr[1] = 123
	}

	fnUpdateSlice := func(s []int32) {
		fmt.Printf("[fnUpdateSlice] slice address: %p\n", &s)
		s[1] = 456
	}

	var array1 = [...]int32{1, 2, 3, 4, 5}
	fmt.Printf("\narray1 address: %p\n", &array1)
	array2 := array1
	fmt.Printf("array2 address: %p\n", &array2)
	array2[0] = 100
	fnUpdateArray(array1)
	fmt.Printf("array: %v, copied array: %v\n", array1, array2)

	var slice1 = []int32{1, 2, 3, 4, 5}
	fmt.Printf("\nslice1 address: %p\n", &slice1)
	slice2 := slice1
	fmt.Printf("slice2 address: %p\n", &slice2)
	slice2[0] = 200
	fnUpdateSlice(slice1)
	fmt.Printf("slice: %v, copied slice: %v\n", slice1, slice2)
}

// demo, array and slice
func testArrayAndSlice02() {
	var array = [...]int32{1, 2, 3, 4, 5}
	var slice = []int32{1, 2, 3, 4, 5}

	fmt.Printf("\narray type: %T\n", array)
	fmt.Printf("slice type: %T\n", slice)

	s1 := array[2:4]
	fmt.Printf("\nbefore: array=%v, slice=%v\n", array, s1)
	s1[0] = 100
	fmt.Printf("after: array=%v, slice=%v\n", array, s1)

	s2 := array[2:4]
	s2 = append(s2, 6, 7)
	fmt.Println("\nre-allocate slice")
	fmt.Printf("before: array=%v, slice=%v\n", array, s2)
	s2[0] = 200
	fmt.Printf("after: array=%v, slice=%v\n", array, s2)
}

// demo, update bytes
func testUpdateBytes() {
	s := "hello world"
	b := []byte(s)
	fmt.Printf("\n[main] b address: %p\n", &b)

	fmt.Println("\n#1: by value:")
	fmt.Printf("before update: %s\n", string(b))
	myUpdateBytesByVal(b)
	fmt.Printf("after update: %s\n", string(b))

	fmt.Println("\n#2: by pointer:")
	fmt.Printf("before update: %s\n", string(b))
	myUpdateBytesByRef(&b)
	fmt.Printf("after update: %s\n", string(b))
}

func myUpdateBytesByVal(b []byte) {
	fmt.Printf("[myUpdateBytesByVal] b address: %p\n", &b)
	b[0] = 'H'
	b = bytes.ToUpper(b)
	b = append(b, '!', '!', '!')
	fmt.Printf("[UpdateBytes] by value: %s\n", string(b))
}

func myUpdateBytesByRef(b *[]byte) {
	fmt.Printf("[myUpdateBytesByRef] b address: %p\n", b)
	*b = bytes.ToUpper(*b)
	*b = append(*b, []byte{'!', '!', '!'}...)
	fmt.Printf("[UpdateBytes] by reference: %s\n", string(*b))
}

// demo, copy bytes
func testCopyBytes() {
	copyBytesByVal := func(b []byte) {
		fmt.Printf("[testCopyBytes] b address: %p\n", &b)
		copy(b, []byte("hello"))
		b = append(b, []byte(" world")...)
	}

	copyBytesByRef := func(b *[]byte) {
		fmt.Printf("[copyBytesByRef] b address: %p\n", b)
		copy(*b, []byte("hello"))
		*b = append(*b, []byte(" world")...)
	}

	b1 := make([]byte, 10)
	fmt.Printf("\nb1 address: %p\n", &b1)
	fmt.Printf("b1 type: %T, cap: %d\n", b1, cap(b1))
	copyBytesByVal(b1)
	fmt.Println("#1: make bytes and copied by val:", string(b1))
	copyBytesByRef(&b1)
	fmt.Println("#2: make bytes and copied by ref:", string(b1))

	var b2 []byte
	fmt.Printf("\nb2 address: %p\n", &b2)
	fmt.Printf("b2 type: %T, cap: %d\n", b2, cap(b2))
	copyBytesByVal(b2)
	fmt.Println("#3: init bytes and copied by val:", string(b2))
	copyBytesByRef(&b2)
	fmt.Println("#4: init bytes and copied by ref:", string(b2))
}

func testMapReference() {
	// TODO:
}

func testStructReference() {
	// TODO:
}

// MainDemo01 main for golang demo01.
func MainDemo01() {
	// testPrintFormatName()
	// testPrintStructValue()
	// testRecover()

	// testStructMethod()
	// testInnerStruct()
	// testAccControl()

	// testTranslateBy()

	// testContext01()
	// testContext02()

	// testArgsPkgAndUnpkg()
	// testTimeFormat()
	// testCodeBlock()

	// testArrayAndSlice01()
	// testArrayAndSlice02()
	// testUpdateBytes()
	// testCopyBytes()

	fmt.Println("golang demo01 DONE.")
}
