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

// demo 01, inner function
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

// demo 02, defer and recover()
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

// demo 03, struct init
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

// demo 04-01, struct and method
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

// demo 04-02, innner struct
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

// demo 04-03, method var
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

// MyObject : demo05, test access control
type MyObject struct {
	VarPublic  string
	varPrivate string
}

// Init : init MyObject struct
func (o *MyObject) Init(pub, pri string) {
	o.VarPublic = pub
	o.varPrivate = pri
}

// MethodPublicGet : return public value
func (o MyObject) MethodPublicGet() string {
	return o.varPrivate
}

func (o MyObject) methodPrivateGet() string {
	return o.VarPublic
}

// GetMyObject : return an empty object
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

// demo 06, context
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

// demo 07, update bytes
func testUpdateBytes() {
	s := "it's a test"
	b := []byte(s)

	fmt.Println("\n#1: by value:")
	fmt.Printf("before update: %s\n", string(b))
	myUpdateBytesByValue(b)
	fmt.Printf("after update: %s\n", string(b))

	fmt.Println("\n#2: by pointer:")
	fmt.Printf("before update: %s\n", string(b))
	myUpdateBytesByPointer(&b)
	fmt.Printf("after update: %s\n", string(b))
}

func myUpdateBytesByValue(b []byte) {
	b = bytes.ToUpper(b)
	b = append(b, '!', '!', '!')
	fmt.Printf("[UpdateBytes] by value: %s\n", string(b))
}

func myUpdateBytesByPointer(b *[]byte) {
	*b = bytes.ToUpper(*b)
	*b = append(*b, []byte{'!', '!', '!'}...)
	fmt.Printf("[UpdateBytes] by pointer: %s\n", string(*b))
}

// demo 08, time format
func testTimeFormat() {
	t := time.Now()
	fmt.Printf("\nweek day: %d, time: %d:%d\n", t.Weekday(), t.Hour(), t.Minute())

	s := strconv.FormatInt(t.Unix(), 10)
	fmt.Println("unix time (seconds from 1970):", s)

	t = time.Unix(t.Unix()+60, 0)
	baseDate := "2006-01-02 15:04:05"
	fmt.Println("cur date (+60s):", t.Format(baseDate))
}

// demo 09, code block
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

// MainDemo01 : main
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

	// testUpdateBytes()

	// testTimeFormat()
	// testCodeBlock()

	fmt.Println("golang demo01 DONE.")
}
