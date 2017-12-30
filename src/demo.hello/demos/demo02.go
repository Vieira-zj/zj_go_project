package demos

import (
	"encoding/base64"
	"errors"
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

// demo 02, panic and recover
func myWork(isOccur bool) {
	myLog("myWork start")
	if isOccur {
		panic("mock error")
	}
	myLog("myWork done")
}

func myLog(args ...interface{}) {
	fmt.Println(args...)
}

func testPanicAndRecover() {
	defer func() {
		fmt.Println("recover.")
		if r := recover(); r != nil {
			myLog(r)
		}
	}()
	myWork(true)
}

// demo 03, Error
type myError struct {
	infoa string
	infob string
	err   error
}

func (merr *myError) Error() string {
	errInfo := fmt.Sprintf(
		"infoa: %s, infob: %s, original error info: %s", merr.infoa, merr.infob, merr.err.Error())
	return errInfo
}

func testErrorType() {
	// #1
	var err = errors.New("mock original error")
	fmt.Println(err.Error())

	// #2
	err = fmt.Errorf("%s", "error from fmt.Errorf()")
	fmt.Println(err.Error())

	// #3
	myErr := &myError{
		infoa: "error info a",
		infob: "error info b",
		err:   errors.New("test custom error"),
	}
	fmt.Println(myErr.Error())
}

// demo 04, base64 encode and decode
func base64EncodeAndDecode(enc *base64.Encoding, input string) {
	encStr := enc.EncodeToString([]byte(input))
	fmt.Printf("base64 encode string: %s\n", encStr)

	decStr, err := enc.DecodeString(encStr)
	if err != nil {
		panic("base64 decode error")
	}
	fmt.Printf("base64 decode string: %s\n", decStr)

	if input != string(decStr) {
		panic(errors.New("not equal"))
	}
}

func testBase64() {
	const str = "Go 言语编程 "
	base64EncodeAndDecode(base64.StdEncoding, str)
	base64EncodeAndDecode(base64.URLEncoding, str)
	base64EncodeAndDecode(base64.RawStdEncoding, str)
	base64EncodeAndDecode(base64.RawURLEncoding, str)
}

// MainDemo02 : main
func MainDemo02() {
	// testInterface()

	// testPanicAndRecover()
	// testErrorType()
	// testBase64()

	fmt.Println("demo 02 done.")
}
