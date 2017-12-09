package examples

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

type person struct {
	Name    string
	Age     int
	Address string
}

func myUpdateStruct(p person) {
	p.Age++
	fmt.Println("in struct update:", p)
}

func myUpdateStructByRef(p *person) {
	p.Age++
	fmt.Println("in struct update by reference:", *p)
}

func testUpdateStruct() {
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

type integer int

func (a integer) less(b integer) bool {
	return a < b
}

func (a *integer) add(b integer) {
	*a += b
}

func addMethodTest() {
	var a integer = 1
	a.add(3)
	if a.less(2) {
		fmt.Println(a, "less 2")
	} else {
		fmt.Println(a, "greater 2")
	}
}

func valueAndReferenceTest() {
	srcArr := [3]int{1, 2, 3}
	var copiedArr = srcArr
	srcArr[1]++
	fmt.Println("src array:", srcArr)
	fmt.Println("copied array by value:", copiedArr)

	var pCopiedArr = &srcArr // *[3]int
	fmt.Println("copied array by reference:", *pCopiedArr)
}

type stringer interface {
	String() string
}

type myStringer struct {
}

func (*myStringer) String() string {
	return "this is a method implement from Stringer"
}

func myPrintTest(args ...interface{}) {
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			fmt.Println(v, "=> type is int")
		case string:
			fmt.Println(v, "=> type is string")
		default:
			if v, ok := v.(stringer); ok {
				val := v.String()
				fmt.Println("default stringer:", val)
			} else {
				fmt.Println("other types")
			}
		}
	}
}

func interfaceVarTest() {
	var w io.Writer
	fmt.Printf("%T\n", w)

	fmt.Println("example: w init as os.Stdout")
	w = os.Stdout
	fmt.Printf("%T\n", w)
	myCheckType(w)
	w.Write([]byte("hello\n"))

	fmt.Println("\nexample: w init as bytes.Buffer")
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)
	myCheckType(w)
	w.Write([]byte("hello"))
}

func myCheckType(w interface{}) {
	if f, ok := w.(*os.File); ok { // check by struct
		fmt.Printf("w support File interface: %T\n", f)
	} else {
		fmt.Printf("w not support File interface: %T\n", f)
	}

	if f, ok := w.(io.ReadWriter); ok { // check by interface
		fmt.Printf("w support RW interface: %T\n", f)
	} else {
		fmt.Printf("w not support RW interface: %T\n", f)
	}
}

// sort.Interface
// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool // i, j are indices of sequence elements
// 	Swap(i, j int)
// }

// sort example 01
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

func sortTest() {
	names := []string{"candy", "brow", "dan", "emm", "anny"}
	fmt.Println("names before sorted:", names)
	sort.Sort(myStringSlice(names))
	fmt.Println("names after sorted:", names)
}

// sort example 02
type track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*track{
	{"Go", "Delilah", "Root Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Ahead", "Alicia", "As I Am", 2007, length("4m36s")},
	{"Ready", "Martin", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	fmt.Printf(format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Printf(format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Printf(format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
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
	less func(x, y *track) bool
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

func sortTest2() {
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

// MainOO : main function for struct and interface examples.
func MainOO() {
	// testUpdateStruct()

	// addMethodTest()
	// valueAndReferenceTest()

	// var myStr stringer = new(myStringer)
	// myPrintTest(1, "test", myStr, 1.0)

	// sortTest()
	// sortTest2()

	interfaceVarTest()

	fmt.Println("oo demo.")
}
