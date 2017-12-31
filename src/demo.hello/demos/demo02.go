package demos

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"sync"
	"time"
)

// demo 01-01, interface
type abser interface {
	abs() float64
	string() string
}

type myFloat float64

func (f myFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f *myFloat) string() string {
	return fmt.Sprintf("%.2f", f.abs())
}

func testMyFloatInterface() {
	var a abser
	f := myFloat(-math.Sqrt2)
	a = &f // pointer here
	fmt.Println(a.abs())
	fmt.Println(a.string())
}

// demo 01-02, interface
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
	zjData := zjGetterAndSetter{
		name: "zhengjin",
		age:  30,
	}

	var getAndSet myGetterAndSetter
	getAndSet = &zjData
	fmt.Println(getAndSet.myGet())

	zjInitAndPrintInfoByStruct(zjData)     // object
	zjInitAndPrintInfoByPointer(&zjData)   // pointer
	zjInitAndPrintInfoByInterface(&zjData) // pointer
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

// demo 05, mutex
func testRwMutex() {
	mutex := new(sync.RWMutex)
	fmt.Println("ready in main")
	mutex.Lock()
	fmt.Println("mutex locked in main")

	chs := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		chs[i] = make(chan int)
		go func(i int, ch chan<- int) {
			fmt.Println("ready in routine:", i)
			mutex.RLock()
			fmt.Println("mutex read locked in routine:", i)
			time.Sleep(2 * time.Second)
			fmt.Println("mutex read unlocked in routine:", i)
			mutex.RUnlock()
			ch <- i
		}(i, chs[i])
	}

	time.Sleep(2 * time.Second)
	fmt.Println("mutex unlocked in main")
	mutex.Unlock()

	for _, ch := range chs {
		<-ch
	}
}

// MainDemo02 : main
func MainDemo02() {
	// testMyFloatInterface()
	// testInterface()

	// testPanicAndRecover()
	// testErrorType()
	// testBase64()

	testRwMutex()

	fmt.Println("demo 02 done.")
}
