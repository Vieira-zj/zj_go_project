package examples

import (
	"fmt"
	"os"
	"sort"
)

func hello(firstName, lastName string) {
	fmt.Printf("hello, %s %s\n", firstName, lastName)
}

func stringFormatTest() {
	myRegExp := `[0-9]\{1,\}`
	filePath := "/home/qboxserver/service_config/_package/iorate.conf"
	cmd := fmt.Sprintf("sed -i '13s/%s/100/' %s", myRegExp, filePath)
	fmt.Println(cmd)
	cmd = fmt.Sprintf(`sed -i '13s/%s/100/' %s`, myRegExp, filePath)
	fmt.Println(cmd)
}

func getNames() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}

func varsExamples() {
	// var v1 int = 10
	var v2 = 10
	v3 := 10

	fmt.Println("vars examples:")
	fmt.Println("vars define and init:")
	// fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)

	fmt.Println("vars exchange:")
	i := 1
	j := 3
	i, j = j, i
	fmt.Println(i)
	fmt.Println(j)

	fmt.Println("enum type:")
	type weekday int
	const (
		Sunday weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	fmt.Println(Monday)

	fmt.Println("string examples:")
	str := "hello world"
	ch := str[0]
	fmt.Printf("length of %s is %d.\n", str, len(str))
	fmt.Printf("first char of %s is %c.\n", str, ch)
	fmt.Printf("char at 0 to 5 is %v.\n", str[0:5])

	fmt.Println("iterator on string:")
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c,", i, str[i]) // char
	}
	fmt.Println()
	for i, ch := range str {
		fmt.Printf("str[%d] = %v,", i, ch) // int
	}
	fmt.Println()

	fmt.Println("create value by new:")
	p := new(int)
	fmt.Printf("p type: %T\n", p) // pointer
	fmt.Printf("p value: %d\n", *p)
	*p = 2
	fmt.Printf("update p value: %d\n", *p)
}

func arrayExamples() {
	q := [...]int{1, 2, 3}
	fmt.Printf("array q: %T\n", q)

	fmt.Println("array examples:")
	fmt.Println("array pass as value not reference:")
	array := [5]int{1, 2, 3, 4, 5}
	modifyArray(array)
	fmt.Println("In main(), array values:", array)

	fmt.Println("create slice from array:")
	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var mySlice = myArray[:5]

	fmt.Println("Elements of myArray: ")
	for _, v := range myArray {
		fmt.Printf("%d ", v)
	}

	fmt.Println("\nElements of mySlice: ")
	for _, v := range mySlice {
		fmt.Printf("%d ", v)
	}

	fmt.Println("\ninit slice:")
	// init 5 elements with default value 0, capbility is 10
	mySlice2 := make([]int, 5, 10)
	fmt.Println("slice length:", len(mySlice2))
	fmt.Println("slice capbility:", cap(mySlice2))

	fmt.Println("slice pass as value:")
	mySlice2 = append(mySlice2, 1, 2, 3)
	fmt.Println("slice after append:", mySlice2)
	myUpdateSlice(mySlice2)
	fmt.Println("in main(), slice:", mySlice2)

	mySlice3 := []int{11, 12, 13}
	copy(mySlice2, mySlice3)
	fmt.Printf("slices2 after copied: %v\n", mySlice2)
	fmt.Printf("slices3 after copied: %v\n", mySlice3)
}

func modifyArray(array [5]int) {
	array[0] = 10
	fmt.Println("In modifyArray(), array values:", array)
}

func myUpdateSlice(s []int) {
	s = append(s, 4, 5)
	fmt.Println("in update(), slice:", s)
}

type personInfo struct {
	ID      string
	Name    string
	Address string
}

func mapExamples() {
	fmt.Println("map examples:")
	fmt.Println("map init and iterator:")
	tmpMap1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	tmpMap1["three"] = 3
	fmt.Printf("map1 type: %T\n", tmpMap1)
	for k, v := range tmpMap1 {
		fmt.Printf("%d=%s\n", v, k)
	}

	tmpMap2 := make(map[string]int)
	tmpMap2["four"] = 4
	tmpMap2["five"] = 5
	fmt.Printf("map2 type: %T\n", tmpMap2)
	for k := range tmpMap2 {
		fmt.Printf("key: %s\n", k)
	}

	fmt.Println("map pass as reference:")
	var persons map[string]personInfo
	persons = make(map[string]personInfo)
	persons["test1"] = personInfo{"test1", "Tom", "Room 203,..."}
	persons["test2"] = personInfo{"test2", "Jack", "Room 101,..."}
	myUpdateMap(persons)
	fmt.Printf("in main(), person map: %v\n", persons)

	fmt.Println("print map values as sorted:")
	// var ids []string
	ids := make([]string, 0, len(persons))
	for id := range persons {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		fmt.Printf("%s\t%s\n", id, persons[id].Name)
	}

	fmt.Println("delete map entry by key:")
	delete(persons, "test1")
	if person, ok := persons["test1"]; ok {
		fmt.Println("Person found:", person.Name)
	} else {
		fmt.Println("Did not find.")
	}
}

func myUpdateMap(persons map[string]personInfo) {
	// persons["test2"].Address = "Room 101,...(update)"
	persons["test3"] = personInfo{ID: "test3", Name: "henry", Address: "Room 606..."}
}

func switchTest(number int) {
	switch {
	case 0 <= number && number <= 3:
		fmt.Printf("0-3")
	case 4 <= number && number <= 6:
		fmt.Printf("4-6")
	case 7 <= number && number <= 9:
		fmt.Printf("7-9")
	default:
		fmt.Printf("invalid number")
	}
}

func controlTest() {
JLoop:
	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			if i > 5 {
				break JLoop
			}
			fmt.Println(i)
		}
	}
}

func argsTest(args ...int) {
	fmt.Println("input args:")
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func fnTest() {
	fn := func(x, y int) int {
		return x + y
	}
	fmt.Printf("fn type: %T\n", fnTest)
	results := fn(1, 2)
	fmt.Println(results)

	func() {
		fmt.Println("fnTest done.")
	}()
}

func fnClosureTest() {
	j := 5
	fn := func() func() { // return function
		i := 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()

	fn()
	j *= 2
	fn()
}

func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case nil:
			fmt.Println("nil")
		case int:
			fmt.Println(arg, "int")
		case int64:
			fmt.Println(arg, "int64")
		case string:
			fmt.Println(arg, "string")
		default:
			panic(fmt.Sprintf("unexpected type %T: %v", arg, arg))
		}
	}
}

type point struct {
	x, y int
}

// receiver as reference
func (p *point) scaleBy(factor int) {
	p.x *= factor
	p.y *= factor
}

// receiver as value
func (p point) String() string {
	return fmt.Sprintf("x=%d, y=%d", p.x, p.y)
}

func methodTest() {
	fmt.Println("invoke methods of struct:")
	p := &point{1, 2}
	p.scaleBy(2)
	fmt.Printf("invoke1: %s\n", (*p).String())
	fmt.Printf("invoke2: %s\n", p.String())

	fmt.Println("method value:")
	p1 := point{1, 3}
	fnScale1 := p1.scaleBy
	fnScale1(2)
	fmt.Printf("fnScale1 type: %T\n", fnScale1)
	fmt.Println("after scale:", p1.String())

	p2 := point{2, 4}
	fnScale2 := (*point).scaleBy
	fnScale2(&p2, 2)
	fmt.Printf("fnScale2 type: %T\n", fnScale2)
	fmt.Println(p2.String())
}

// PrintGoEnvValues : print go root and path env values
func PrintGoEnvValues() {
	fmt.Printf("$GOROOT: %s\n", os.Getenv("GOROOT"))
	fmt.Printf("$GOPATH: %s\n", os.Getenv("GOPATH"))
}

// MainHello : main function for general examples.
func MainHello() {
	// hello("zheng", "jin")
	// stringFormatTest()

	// firstName, _, _ := getNames()
	// _, lastName, _ := getNames()
	// fmt.Printf("name: %s %s\n", firstName, lastName)

	// varsExamples()
	// arrayExamples()
	// mapExamples()

	// switchTest(7)
	// controlTest()
	// argsTest(1, 2, 3)

	// fnTest()
	// fnClosureTest()

	// var v1 = 1
	// var v2 int64 = 1234
	// var v3 = "test"
	// var v4 float32 = 1.234
	// myPrintf(v1, v2, v3, nil, v4)

	// methodTest()

	fmt.Println("hello demo.")
}
