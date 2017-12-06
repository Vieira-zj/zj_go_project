package main

import "fmt"

func varsExamples() {
	// var v1 int = 10
	var v2 = 10
	v3 := 10

	fmt.Println("VARS EXAMPLES:")
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
	const (
		Sunday = iota
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
}

func getName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}

func modifyArray(array [5]int) {
	array[0] = 10
	fmt.Println("In modifyArray(), array values:", array)
}

func arrayExamples() {
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

	mySlice2 = append(mySlice2, 1, 2, 3)
	fmt.Println("slice after append:", mySlice2)

	mySlice3 := []int{11, 12, 13}
	copy(mySlice2, mySlice3)
	fmt.Println("slice after copied:", mySlice2)
}

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func mapTest() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	personDB["test1"] = PersonInfo{"test1", "Tom", "Room 203,..."}
	personDB["test2"] = PersonInfo{"test2", "Jack", "Room 101,..."}

	fmt.Println("map examples:")
	delete(personDB, "test1")
	person, ok := personDB["test1"]
	if ok {
		fmt.Println("Person found:", person.Name)
	} else {
		fmt.Println("Did not find.")
	}
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
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
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

func mainHello() {
	// varsExamples()

	// _, _, nickName := getName()
	// fmt.Println("nick name: " + nickName)

	// arrayExamples()
	// mapTest()

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

	fmt.Printf("\nhello, world\n")
}
