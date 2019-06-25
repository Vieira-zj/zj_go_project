package examples

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func hello(firstName, lastName string) {
	fmt.Printf("hello, %s %s\n", firstName, lastName)
	fmt.Println("col1\tcol2")
	fmt.Println(`col1\tcol2`)
}

func intFormatTest() {
	// 10进制转36进制
	var num int64 = 1380469261
	fmt.Println("\nformat int64 / 36:", strconv.FormatInt(num, 36))
}

func returnTest() {
	firstName, _, _ := getNames()
	_, lastName, _ := getNames()
	fmt.Printf("name: %s %s\n", firstName, lastName)
}

func getNames() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}

func intVarTest() {
	fmt.Println("\n#1. vars declare and init:")
	// var v1 int = 10
	// fmt.Println(v1)
	var v2 = 10
	// fmt.Println("v2=" + strconv.Itoa(v2))
	fmt.Printf("v2=%d\n", v2)
	v3 := 10
	// fmt.Println("v3=" + strconv.Itoa(v3))
	fmt.Printf("v3=%d\n", v3)

	fmt.Println("\n#2. vars exchange:")
	i := 1
	j := 3
	i, j = j, i
	fmt.Println(i)
	fmt.Println(j)

	fmt.Println("\n#3. enum type:")
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

	today := Saturday
	fmt.Println("today:", today)
}

func stringVarTest() {
	fmt.Println("\n#1. string basic:")
	str := "hello world"
	fmt.Printf("length of %s is %d\n", str, len(str))
	fmt.Printf("first char of %s is %c\n", str, str[0])
	fmt.Printf("char at 0 to 5 is %v\n", str[0:5])

	fmt.Println("\n#2. iterator on string:")
	str = "hello"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c,", i, str[i]) // char
	}
	fmt.Println()
	for i, ch := range str {
		fmt.Printf("str[%d] = %v,", i, ch) // int
	}
	fmt.Println()

	fmt.Println("\n#3. create value by new:")
	p := new(int)
	fmt.Printf("p type: %T\n", p) // type pointer
	fmt.Printf("p value: %d\n", *p)
	*p = 2
	fmt.Printf("update p value: %d\n", *p)
}

func arrayVarTest() {
	fmt.Println("\n#1. define array:")
	q := [...]int{1, 2, 3}
	fmt.Printf("array q: %T\n", q)

	fmt.Println("\n#2. array pass as value (not reference):")
	array := [5]int{1, 2, 3, 4, 5}
	modifyArray(array)
	fmt.Println("in main(), array values:", array)

	fmt.Println("\n#3. create slice from array:")
	myArray := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("myArray elements:")
	for _, v := range myArray {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	var mySlice = myArray[:5]
	myArray[0] = 10
	fmt.Println("mySlice elements:")
	for _, v := range mySlice {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func modifyArray(array [5]int) {
	array[0] = 10
	fmt.Println("in modifyArray(), array values:", array)
}

func sliceVarTest() {
	fmt.Println("\n#1. init slice:")
	mySlice1 := make([]int, 0)
	fmt.Println("mySlice1 init length:", len(mySlice1))
	mySlice1 = append(mySlice1, 1, 2)
	fmt.Printf("mySlice1 length: %d values: %v\n", len(mySlice1), mySlice1)

	fmt.Println("\n#2. slice pass as value:")
	// init 5 elements with default value 0, capbility is 10
	mySlice2 := make([]int, 5, 10)
	fmt.Println("mySlice2 length:", len(mySlice2))
	fmt.Println("mySlice2 capbility:", cap(mySlice2))

	mySlice2 = append(mySlice2, 1, 2, 3)
	tmpSlice := []int{11, 12}
	mySlice2 = append(mySlice2, tmpSlice...)
	fmt.Println("mySlice2 values:", mySlice2)
	myUpdateSlice(mySlice2)
	fmt.Println("in main(), slice:", mySlice2)

	fmt.Println("\n#3. copy slice:")
	mySlice3 := []int{21, 22, 23}
	copy(mySlice2, mySlice3)
	fmt.Printf("mySlice2 after copied: %v\n", mySlice2)
	fmt.Printf("mySlice3 after copied: %v\n", mySlice3)
}

func myUpdateSlice(s []int) {
	s = append(s, 99)
	fmt.Println("in updateSlice(), slice:", s)
}

func mapVarTest01() {
	fmt.Println("\n#1. map init and iterator:")
	tmpMap1 := map[string]int{
		"one": 1,
		"two": 2,
	}
	tmpMap1["three"] = 3
	fmt.Printf("map1 type: %T\n", tmpMap1)
	fmt.Printf("map1 length: %d\n", len(tmpMap1))

	fmt.Println("\n#2. iterator on map1:")
	for k, v := range tmpMap1 {
		fmt.Printf("%d=%s\n", v, k)
	}

	fmt.Println("\n#3. map2 keys:")
	tmpMap2 := make(map[string]int)
	tmpMap2["four"] = 4
	tmpMap2["five"] = 5
	fmt.Printf("map2 type: %T\n", tmpMap2)
	for k := range tmpMap2 {
		fmt.Printf("key: %s\n", k)
	}
}

type personInfo struct {
	ID      string
	Name    string
	Address string
}

func mapVarTest02() {
	fmt.Println("\n#1. map pass as reference:")
	var persons map[string]personInfo
	persons = make(map[string]personInfo)
	persons["01"] = personInfo{"id01", "Tom", "Room 203,..."}
	persons["02"] = personInfo{"id02", "Jack", "Room 101,..."}
	myUpdateMap(persons)
	fmt.Printf("in main(), person map: %v\n", persons)

	fmt.Println("\n#2. print map as sorted:")
	// var ids []string
	ids := make([]string, 0, len(persons))
	for id := range persons {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		fmt.Printf("%s\t%s\n", id, persons[id].Name)
	}

	fmt.Println("\n#3. delete map entry by key:")
	delete(persons, "01")
	if person, ok := persons["01"]; ok {
		fmt.Println("person found:", person.Name)
	} else {
		fmt.Println("person not found.")
	}
}

func myUpdateMap(persons map[string]personInfo) {
	// persons["test2"].Address = "Room 101,...(update)"
	persons["03"] = personInfo{ID: "id03", Name: "henry", Address: "Room 606..."}
}

func switchTest() {
	number := 7
	switch {
	case 0 <= number && number <= 3:
		fmt.Printf("0-3")
	case 4 <= number && number <= 6:
		fmt.Printf("4-6")
	case 7 <= number && number <= 9:
		fmt.Printf("7-9")
	default:
		fmt.Printf("invalid number!")
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

func argsTest() {
	printNumbers := func(args ...int) {
		fmt.Printf("input args type: %T\n", args)
		fmt.Println("input args:")
		for _, arg := range args {
			fmt.Println(arg)
		}
	}

	printNumbers(1, 2, 3)
}

func innerFuncTest() {
	fnSum := func(x, y int) int {
		return x + y
	}
	fmt.Printf("\nfn type: %T\n", fnSum)
	fmt.Println("fnSum(1, 2):", fnSum(1, 2))

	func(name string) {
		fmt.Printf("hello %s, inner funcTest done\n", name)
	}("Henry")
}

func funcVarsTest() {
	add := func(x, y int) int {
		return x + y
	}
	subtract := func(x, y int) int {
		return x - y
	}
	myClaculation := func(x, y int, fn func(int, int) int) int {
		fmt.Printf("\nfunction type: %T\n", fn)
		return fn(x, y)
	}

	x := 10
	y := 7
	fmt.Printf("x + y = %d\n", myClaculation(x, y, add))
	fmt.Printf("x - y = %d\n", myClaculation(x, y, subtract))
}

func funcClosureTest() {
	j := 10
	myFunc := func() func() { // return function
		i := 0
		return func() {
			i++
			j++
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()

	myFunc()
	j *= 2
	myFunc()
}

func varTypeCheckTest() {
	var v1 = 1
	var v2 int64 = 1234
	var v3 = "test"
	var v4 float32 = 1.234
	myPrintf(v1, v2, v3, nil, v4)
}

func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case nil:
			fmt.Println("nil")
		case int:
			fmt.Println(arg, "=> type int")
		case int64:
			fmt.Println(arg, "=> type int64")
		case string:
			fmt.Println(arg, "=> type string")
		default:
			panic(fmt.Sprintf("unexpected type %T: %v", arg, arg))
		}
	}
}

// object point
type point struct {
	x, y int
}

// object, receiver as reference
func (p *point) scaleBy(factor int) {
	p.x *= factor
	p.y *= factor
}

// object, receiver as value
func (p point) String() string {
	return fmt.Sprintf("x=%d, y=%d", p.x, p.y)
}

func objectTest() {
	fmt.Println("\n#1")
	p := &point{1, 2}
	p.scaleBy(2)
	fmt.Printf("(*p).String(): %s\n", (*p).String())
	fmt.Printf("p.String(): %s\n", p.String())

	fmt.Println("\n#2")
	p1 := point{1, 3}
	funcScale1 := p1.scaleBy
	funcScale1(2)
	fmt.Printf("funcScale1 type: %T\n", funcScale1)
	fmt.Println("after scale:", p1.String())

	fmt.Println("\n#3")
	p2 := point{2, 4}
	funcScale2 := (*point).scaleBy
	funcScale2(&p2, 2)
	fmt.Printf("funcScale2 type: %T\n", funcScale2)
	fmt.Println("after scale:", p2.String())
}

// PrintGoEnvValues : print go root and path env values.
func PrintGoEnvValues() {
	fmt.Printf("$GOROOT: %s\n", os.Getenv("GOROOT"))
	fmt.Printf("$GOPATH: %s\n", os.Getenv("GOPATH"))
}

// MainHello : main function for general golang examples.
func MainHello() {
	// hello("zheng", "jin")

	// intFormatTest()
	// returnTest()

	// intVarTest()
	// stringVarTest()
	// arrayVarTest()
	// sliceVarTest()
	// mapVarTest01()
	// mapVarTest02()

	// switchTest()
	// controlTest()
	// argsTest()

	// innerFuncTest()
	// funcVarsTest()
	// funcClosureTest()

	// varTypeCheckTest()

	// objectTest()

	fmt.Println("golang hello example DONE.")
}
