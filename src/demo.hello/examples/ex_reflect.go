package examples

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"time"
)

func testReflect() {
	fmt.Println("\n#1: example TypeOf:")
	t := reflect.TypeOf(3)
	fmt.Println("int typeof:", t)
	fmt.Printf("int type: %T\n", 3)
	fmt.Println("int typeof kind:", t.Kind())

	var w io.Writer = os.Stdout
	fmt.Println("\nio.writer typeof:", reflect.TypeOf(w))
	fmt.Println("io.writer typeof kind:", reflect.TypeOf(w).Kind()) // ptr

	fmt.Println("\n#2: example ValueOf:")
	v := reflect.ValueOf(3)
	fmt.Println("int valueof:", v)
	fmt.Printf("int valueof: %v\n", v)
	fmt.Println("int valueof string:", v.String())

	t = v.Type()
	fmt.Println("int type:", t.String())
	fmt.Println("int type kind:", v.Kind())

	fmt.Println("\n#3: type transfer:")
	x := v.Interface()
	i := x.(int)
	fmt.Printf("int value: %d\n", i)
}

// example: print by type
type iStringer2 interface {
	String() string
}

type myStringer2 struct {
	name string
}

func (s myStringer2) String() string {
	return fmt.Sprintf("hello %s!", s.name)
}

func testMySprint() {
	fmt.Printf("\nmy print: %s\n", mySprint(9))
	fmt.Printf("my print: %s\n", mySprint("test"))
	fmt.Printf("my print: %s\n", mySprint(true))

	myStr := myStringer2{
		name: "Henry",
	}
	fmt.Printf("\nmy print: %s\n", mySprint(myStr))
}

func mySprint(x interface{}) string {
	switch x := x.(type) {
	case iStringer:
		return x.String()
	case int:
		return strconv.Itoa(x)
	case string:
		return x
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "unknown"
	}
}

// example: any
func typeFormatTest() {
	var x int64 = 1
	var d = 1 * time.Nanosecond
	fmt.Println(any(x))
	fmt.Println(any(d))
	fmt.Println(any([]int64{x}))
	fmt.Println(any([]time.Duration{d}))
}

// any formats any value as a string.
func any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}

// example: display
func typeDisplayTest() {
	fmt.Println("\ndisplay array:")
	testArr := [5]int{1, 2, 3, 4, 5}
	display("testArr", reflect.ValueOf(testArr))

	fmt.Println("\ndisplay map:")
	testMap := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}
	display("testMap", reflect.ValueOf(testMap))

	fmt.Println("\ndisplay struct:")
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
	}
	display("strangelove", reflect.ValueOf(strangelove))

	fmt.Println("\ndisplay pointer:")
	var testPtr = &strangelove
	display("testPtr", reflect.ValueOf(testPtr))

	fmt.Println("\ndisplay demo:")
	var i interface{} = 3
	display("i", reflect.ValueOf(i))
	display("&i", reflect.ValueOf(&i))
	// TODO:
}

func display(name string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", name)
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", name, i), v.Index(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", name, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", name, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", name)
		} else {
			display(fmt.Sprintf("(*%s)", name), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", name)
		} else {
			fmt.Printf("%s.type = %s\n", name, v.Elem().Type())
			display(name+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", name, formatAtom(v))
	}
}

// MainReflect : main for reflect examples.
func MainReflect() {
	// testReflect()
	// testMySprint()

	// typeFormatTest()
	// typeDisplayTest()

	fmt.Println("golang reflect examples DONE.")
}
