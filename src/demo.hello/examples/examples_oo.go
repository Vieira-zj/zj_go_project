package examples

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"time"
)

func myUpper(s *string) {
	*s = strings.ToUpper(*s)
}

func valueAndRefExamples() {
	fmt.Println("test value and ref for string:")
	str := "hello"
	myUpper(&str)
	fmt.Printf("string upper: %s\n", str)

	fmt.Println("\ntest value and ref for array:")
	srcArr := [3]int{1, 2, 3}
	fmt.Printf("src array: %v\n", srcArr)

	var vCopiedArr = srcArr
	var pCopiedArr *[3]int // pointer to array
	pCopiedArr = &srcArr
	fmt.Printf("p type: %T\n", pCopiedArr)

	fmt.Printf("length by value: %v\n", len(vCopiedArr))
	fmt.Printf("length by ref: %v\n", len(pCopiedArr))

	srcArr[1]++
	fmt.Printf("copied array by value: %v\n", vCopiedArr)
	fmt.Printf("copied array by ref: %v\n", *pCopiedArr)

	pCopiedArr[1]++
	fmt.Printf("value at 1 by value: %v\n", vCopiedArr[1])
	fmt.Printf("value at 1 by ref: %v\n", pCopiedArr[1])
	fmt.Printf("src array: %v\n", srcArr)
}

// struct
type person struct {
	Name    string
	Age     int
	Address string
}

func updateStructTest() {
	// struct pass as value but not reference
	fmt.Println("pass by value:")
	p1 := person{"Tom", 30, "ShangHai China"}
	fmt.Println("before update:", p1)
	myUpdateStruct(p1)
	fmt.Println("after update:", p1)

	fmt.Println("\nupdate by reference:")
	p2 := person{"Henry", 35, "BeiJing China"}
	fmt.Println("before update:", p2)
	myUpdateStructByRef(&p2)
	fmt.Println("after update:", p2)
}

func myUpdateStruct(p person) {
	p.Age++
	fmt.Println("in update by value, struct:", p)
}

func myUpdateStructByRef(p *person) {
	p.Age++
	fmt.Println("in update by ref, struct:", *p)
}

// method, int
type myInteger int

func (a myInteger) less(b myInteger) bool {
	return a < b
}

func (a *myInteger) selfAdd(b myInteger) {
	*a += b
}

func myIntMethodTest() {
	var a myInteger = 1
	a.selfAdd(3)
	if a.less(2) {
		fmt.Printf("%d less 2\n", a)
	} else {
		fmt.Printf("%d greater 2\n", a)
	}
}

// method, string
type myString string

func (s *myString) selfUpper() {
	tmpStr := strings.ToUpper(string(*s))
	*s = myString(tmpStr)
}

func myStringMethodTest() {
	s := myString("hello")
	s.selfUpper()
	fmt.Printf("my string upper: %v\n", s)
}

// check types
func checkTypeTest() {
	var w io.Writer
	fmt.Printf("%T\n", w) // nil

	fmt.Println("\nexample: w init as os.Stdout")
	w = os.Stdout
	fmt.Printf("%T\n", w)
	myCheckType(w)
	w.Write([]byte("hello\n"))

	fmt.Println("\nexample: w init as bytes.Buffer")
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)
	myCheckType(w)
	w.Write([]byte("world"))
}

func myCheckType(w interface{}) {
	if v, ok := w.(*os.File); ok { // check by struct
		fmt.Printf("w support File interface: %T\n", v)
	} else {
		fmt.Printf("w not support File interface: %T\n", v)
	}

	if v, ok := w.(io.ReadWriter); ok { // check by interface
		fmt.Printf("w support RW interface: %T\n", v)
	} else {
		fmt.Printf("w not support RW interface: %T\n", v)
	}
}

// print types
type stringer interface {
	String() string
}

type myStringer struct {
}

func (*myStringer) String() string {
	return "this is a method implement from Stringer"
}

func myPrintTest() {
	myInt := 1
	myStr1 := "hello world"
	myStr2 := new(myStringer)
	fmt.Printf("new string type: %T\n", myStr2)
	myPrintTypes(myInt, myStr1, myStr2, 1.0)
}

func myPrintTypes(args ...interface{}) {
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			fmt.Printf("%v => type is int\n", v)
		case string:
			fmt.Printf("%v => type is string\n", v)
		default:
			// if v, ok := v.(*myStringer); ok {
			if v, ok := v.(stringer); ok {
				fmt.Println("default stringer:", v.String())
			} else {
				fmt.Println("other types")
			}
		}
	}
}

// Ex01, sort
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

func arraySortTest01() {
	names := []string{"candy", "brow", "dan", "emm", "anny"}
	myNames := myStringSlice(names)
	fmt.Println("src names value:", names)
	sort.Sort(myNames)
	fmt.Println("names after sorted:", names)
}

// Ex02, sort
type track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

// sort by artist
type byArtist []*track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// custom sort
type customSort struct {
	t    []*track
	less func(x, y *track) bool // custom fn
}

func (x customSort) Len() int {
	return len(x.t)
}

func (x customSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x customSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func arraySortTest02() {
	var tracks = []*track{
		{"Go", "Delilah", "Root Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Ahead", "Alicia", "As I Am", 2007, length("4m36s")},
		{"Ready", "Martin", "Smash", 2011, length("4m24s")},
	}

	fmt.Println("before sort:")
	printTracks(tracks)

	fmt.Println("\nafter sort by artist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("\nafter reverse sort by artist:")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("\nafter sort by custom:")
	sort.Sort(customSort{tracks, func(x, y *track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
}

func printTracks(tracks []*track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	fmt.Printf(format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Printf(format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Printf(format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// MainOO : main function for struct and interface examples.
func MainOO() {
	// valueAndRefTest()

	// updateStructTest()
	// myIntMethodTest()
	// myStringMethodTest()

	// checkTypeTest()
	// myPrintTest()

	// arraySortTest01()
	// arraySortTest02()

	fmt.Println("oo demo.")
}
