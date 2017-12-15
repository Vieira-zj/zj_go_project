package examples

import (
	"fmt"
	"os"
	"sort"
)

func hello(firstName, lastName string) {
	fmt.Printf("hello, %s %s\n", firstName, lastName)
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

	i := 1
	j := 3
	i, j = j, i
	fmt.Println("vars exchange:")
	fmt.Println(i)
	fmt.Println(j)

	fmt.Println("enum:")
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
	fmt.Printf("the length of %s is %d.\n", str, len(str))
	fmt.Printf("the first character of %s is %c.\n", str, ch)

	fmt.Println("iterator on string:")
	for i := 0; i < len(str); i++ {
		fmt.Println(i, str[i])
	}
	for i, ch := range str {
		fmt.Println(i, ch)
	}

	fmt.Println("create value by new")
	p := new(int)
	fmt.Println("value by new():", *p)
	*p = 2
	fmt.Print("p value:", *p)
}

func getName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}

func modifyArray(array [5]int) {
	array[0] = 10
	fmt.Println("In modifyArray(), array values:", array)
}

func arrayExamples() {
	q := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q)

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
		fmt.Print(v, " ")
	}

	fmt.Println("\nElements of mySlice: ")
	for _, v := range mySlice {
		fmt.Print(v, " ")
	}

	fmt.Println("\ninit slice:")
	mySlice2 := make([]int, 5, 10)
	fmt.Println("slice length:", len(mySlice2))
	fmt.Println("slice capbility:", cap(mySlice2))

	fmt.Println("slice pass as value:")
	mySlice2 = append(mySlice2, 1, 2, 3)
	myUpdateSlice(mySlice2)
	fmt.Println("slice after append:", mySlice2)

	mySlice3 := []int{11, 12, 13}
	copy(mySlice2, mySlice3)
	fmt.Println("slice after copied:", mySlice2)
}

func myUpdateSlice(s []int) {
	s = append(s, 4, 5)
	fmt.Println("slice in update:", s)
}

type personInfo struct {
	ID      string
	Name    string
	Address string
}

func mapExamples() {
	fmt.Println("map examples:")
	fmt.Println("map init and iterator:")
	tmpMap1 := make(map[string]int)
	tmpMap1["one"] = 1
	tmpMap1["two"] = 2
	for k, v := range tmpMap1 {
		fmt.Printf("%d - %s\n", v, k)
	}

	tmpMap2 := map[string]int{"three": 3, "four": 4}
	for k := range tmpMap2 {
		fmt.Printf("key: %s\n", k)
	}

	fmt.Println("map pass as reference:")
	var personDB map[string]personInfo
	personDB = make(map[string]personInfo)

	personDB["test1"] = personInfo{"test1", "Tom", "Room 203,..."}
	personDB["test2"] = personInfo{"test2", "Jack", "Room 101,..."}

	myUpdateMap(personDB)
	fmt.Println(personDB)

	fmt.Println("print map values as sorted:")
	// var ids []string
	ids := make([]string, 0, len(personDB))
	for id := range personDB {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	for _, id := range ids {
		fmt.Printf("%s\t%s\n", id, personDB[id].Name)
	}

	fmt.Println("delete map entry by key:")
	delete(personDB, "test1")
	person, ok := personDB["test1"]
	if ok {
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
	fmt.Println("arguments:")
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func myPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case nil:
			fmt.Println("null")
		case int:
			fmt.Println(arg, "is an int value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		case string:
			fmt.Println(arg, "is a string value.")
		default:
			panic(fmt.Sprintf("unexpected type %T: %v", arg, arg))
		}
	}
}

func fnTest() {
	fn := func(x, y int) int {
		return x + y
	}
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
	fmt.Println((*p).String())
	fmt.Println(p.String())

	fmt.Println("method value:")
	p1 := point{1, 3}
	fnScale1 := p1.scaleBy
	fnScale1(2)
	fmt.Println(p1.String())
	fmt.Printf("%T\n", fnScale1)

	p2 := point{1, 4}
	fnScale2 := (*point).scaleBy
	fnScale2(&p2, 2)
	fmt.Println(p2.String())
	fmt.Printf("%T\n", fnScale2)
}

// PrintGoEnvValues : print go root and path env values
func PrintGoEnvValues() {
	fmt.Printf("$GOROOT: %s\n", os.Getenv("GOROOT"))
	fmt.Printf("$GOPATH: %s\n", os.Getenv("GOPATH"))
}

// MainHello : main function for general examples.
func MainHello() {
	hello("zheng", "jin")
	// varsExamples()

	// _, _, nickName := getName()
	// fmt.Println("nick name: " + nickName)

	// arrayExamples()
	// mapExamples()

	// switchTest(7)
	// controlTest()
	// argsTest(1, 2, 3)

	// var v1 = 1
	// var v2 int64 = 1234
	// var v3 = "test"
	// var v4 float32 = 1.234
	// myPrintf(v1, v2, v3, v4)

	// fnTest()
	// fnClosureTest()

	// methodTest()

	fmt.Println("hello demo.")
}
