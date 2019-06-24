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

func testValueAndRef() {
	myUpper := func(s *string) {
		*s = strings.ToUpper(*s)
	}

	fmt.Println("\n#1: test value and ref for string:")
	str := "hello"
	myUpper(&str)
	fmt.Printf("upper string: %s\n", str)

	fmt.Println("\n#2: test value and ref for array:")
	srcArr := [3]int{1, 2, 3}
	fmt.Printf("src array: %v\n", srcArr)

	var vCopiedArr = srcArr
	var pCopiedArr *[3]int // type pointer
	pCopiedArr = &srcArr
	fmt.Printf("p type: %T\n", pCopiedArr)

	fmt.Printf("\nvCopiedArr length: %v\n", len(vCopiedArr))
	fmt.Printf("pCopiedArr length: %v\n", len(pCopiedArr))

	srcArr[1]++
	fmt.Printf("\nsrc changed, vCopiedArr: %v\n", vCopiedArr)
	fmt.Printf("src changed, pCopiedArr: %v\n", *pCopiedArr)

	pCopiedArr[1]++
	fmt.Printf("\nref changed, vCopiedArr[1]: %v\n", vCopiedArr[1])
	fmt.Printf("ref changed, pCopiedArr[1]: %v\n", pCopiedArr[1])

	fmt.Printf("\nafter updated, srcArr: %v\n", srcArr)
}

// struct
type person struct {
	Name    string
	Age     int
	Address string
}

func testStructValAndRef() {
	// default, struct pass as value but not reference
	fmt.Println("\npass by value:")
	p1 := person{"Tom", 30, "ShangHai China"}
	fmt.Println("before update:", p1)
	myUpdatePerson(p1)
	fmt.Println("after update:", p1)

	fmt.Println("\npass by reference:")
	p2 := person{"Henry", 35, "BeiJing China"}
	fmt.Println("before update:", p2)
	myUpdatePersonByRef(&p2)
	fmt.Println("after update:", p2)
}

func myUpdatePerson(p person) {
	p.Age++
	fmt.Println("[myUpdatePerson], struct:", p)
}

func myUpdatePersonByRef(p *person) {
	p.Age++
	fmt.Println("[myUpdatePersonByRef], struct:", *p)
}

// struct
type fullName struct {
	firstName string
	lastName  string
	nickName  string
}

func testStructArgs() {
	name1 := fullName{
		firstName: "fname1",
		lastName:  "lname1",
		nickName:  "nname1",
	}
	myPrintFullName(name1)
	myPrintFullName(fullName{firstName: "fname2", lastName: "lname2", nickName: "nname2"})
}

func myPrintFullName(name fullName) {
	fmt.Printf("full name: %s %s, nick name: %s\n",
		name.firstName, name.lastName, name.nickName)
}

// object myInteger
type myInteger int

func (a myInteger) less(b myInteger) bool {
	return a < b
}

func (a *myInteger) selfAdd(b myInteger) {
	*a += b
}

func testMyIntObject() {
	var i myInteger = 1
	i.selfAdd(3)
	if i.less(2) {
		fmt.Printf("\n%d less 2\n", i)
	} else {
		fmt.Printf("\n%d greater 2\n", i)
	}
}

// object myString
type myString string

func (s *myString) selfUpper() {
	tmpStr := strings.ToUpper(string(*s))
	*s = myString(tmpStr)
}

func testMyStringObject() {
	s := myString("hello world")
	s.selfUpper()
	fmt.Printf("\nmy upper string: %v\n", s)
}

// check types
func testCheckIOType() {
	var w io.Writer
	fmt.Printf("\nw type: %T\n", w) // nil

	fmt.Println("\n#1: w init as os.Stdout")
	w = os.Stdout
	fmt.Printf("w type: %T\n", w)
	myCheckIOType(w)
	w.Write([]byte("hello\n"))

	fmt.Println("\n#2: w init as bytes.Buffer")
	w = new(bytes.Buffer)
	fmt.Printf("w type: %T\n", w)
	myCheckIOType(w)
	w.Write([]byte("world"))
}

func myCheckIOType(w interface{}) {
	if v, ok := w.(*os.File); ok { // check by struct
		fmt.Printf("w support File interface: %T\n", v)
	} else {
		fmt.Printf("w NOT support File interface: %T\n", v)
	}

	if v, ok := w.(io.ReadWriter); ok { // check by interface
		fmt.Printf("w support rw interface: %T\n", v)
	} else {
		fmt.Printf("w NOT support rw interface: %T\n", v)
	}
}

// check types (interface and struct)
type iStringer interface {
	String() string
}

type myStringer struct {
}

func (*myStringer) String() string {
	return "this is a method implement from [iStringer]"
}

func testMyPrintType() {
	myInt := 1
	myStr := "hello world"
	myNewStr := new(myStringer)
	fmt.Printf("\nnew string type: %T\n", myNewStr)
	myPrintTypes(myInt, myStr, myNewStr, 1.0)
}

func myPrintTypes(args ...interface{}) {
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			fmt.Printf("type int: %v\n", v)
		case string:
			fmt.Printf("type string: %v\n", v)
		default:
			// if v, ok := v.(*myStringer); ok {
			if v, ok := v.(iStringer); ok {
				fmt.Println("default type iStringer:", v.String())
			} else {
				fmt.Println("unknow types")
			}
		}
	}
}

// sort examples
// sort.Interface
// type Interface interface {
// 	Len() int
// 	Less(i, j int) bool // i, j are indices of sequence elements
// 	Swap(i, j int)
// }

// Ex01, sort
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

func testSliceSort01() {
	names := []string{"candy", "brow", "dan", "emm", "anny"}
	myNames := myStringSlice(names)
	fmt.Println("\nsrc names value:", names)
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
	less func(x, y *track) bool // custom func
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

func testSliceSort02() {
	var tracks = []*track{
		{"Go", "Delilah", "Root Up", 2012, getDuration("3m38s")},
		{"Go", "Moby", "Moby", 1992, getDuration("3m37s")},
		{"Ahead", "Alicia", "As I Am", 2007, getDuration("4m36s")},
		{"Ready", "Martin", "Smash", 2011, getDuration("4m24s")},
	}

	fmt.Println("\nbefore sort:")
	printTracks(tracks)
	sort.Sort(byArtist(tracks))
	fmt.Println("after sort by artist:")
	printTracks(tracks)

	fmt.Println("after reverse sort by artist:")
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

func getDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

// MainOO : main function for struct and interface examples.
func MainOO() {
	// testValueAndRef()
	// testStructValAndRef()
	// testStructArgs()

	// testMyIntObject()
	// testMyStringObject()

	// testCheckIOType()
	// testMyPrintType()

	// testSliceSort01()
	// testSliceSort02()

	fmt.Println("golang oo demo DONE.")
}
