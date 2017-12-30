package demos

import (
	"fmt"
)

// demo 01, interface
type myGetter interface {
	myGet() string
}

type mySetter interface {
	mySet(string)
}

type myGetterAndSetter interface {
	myGetter
	mySetter
}

type zjGetterAndSetter struct {
	name string
	age  int
	desc string
}

func (zj *zjGetterAndSetter) mySet(val string) {
	zj.desc = val
}

func (zj *zjGetterAndSetter) myGet() string {
	if len(zj.desc) == 0 {
		zj.desc = "empty"
	}
	return fmt.Sprintf("name: %s, age: %d, desc: %s", zj.name, zj.age, zj.desc)
}

func zjInitAndPrintInfoByStruct(input zjGetterAndSetter) {
	input.mySet("this is a struct.")
	fmt.Println(input.myGet())
}

func zjInitAndPrintInfoByPointer(input *zjGetterAndSetter) {
	input.mySet("this is a pointer.")
	fmt.Println(input.myGet())
}

func zjInitAndPrintInfoByInterface(input myGetterAndSetter) {
	input.mySet("this is an interface.")
	fmt.Println(input.myGet())
}

func testInterface() {
	data := zjGetterAndSetter{
		name: "zhengjin",
		age:  30,
	}
	zjInitAndPrintInfoByStruct(data)     // object
	zjInitAndPrintInfoByPointer(&data)   // pointer
	zjInitAndPrintInfoByInterface(&data) // pointer
}

// MainDemo02 : main
func MainDemo02() {
	testInterface()

	fmt.Println("demo 02 done.")
}
