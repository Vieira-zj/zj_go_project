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
	fmt.Println("example Type:")
	t := reflect.TypeOf(3)
	fmt.Println("type:", t)
	fmt.Printf("type: %T\n", 3)
	fmt.Println("kind:", t.Kind())

	var w io.Writer = os.Stdout
	fmt.Println("type:", reflect.TypeOf(w))
	fmt.Println("kind:", reflect.TypeOf(w).Kind()) // ptr

	fmt.Println("example Value:")
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())

	t = v.Type()
	fmt.Println("type:", t.String())
	fmt.Println("kind:", v.Kind())

	x := v.Interface()
	i := x.(int)
	fmt.Printf("value: %d\n", i)
}

// example: print
func mySprint(x interface{}) string {
	type stringer interface {
		String() string
	}

	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		return "???"
	}
}

// example: any
func anyTest() {
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
func displayTest() {
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

	// TODO:
	fmt.Println("\ndisplay demo:")
	var i interface{} = 3
	display("i", reflect.ValueOf(i))
	display("&i", reflect.ValueOf(&i))
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

// MainReflect : main function for reflect examples.
func MainReflect() {
	// testReflect()

	// fmt.Println(mySprint(true))
	// anyTest()
	displayTest()

	fmt.Println("reflect demo.")
}
