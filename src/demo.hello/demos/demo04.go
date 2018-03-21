package demos

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// demo 01, init value
func init() {
	fmt.Println("start run demo04") // #2
}

func sayHello() string {
	fmt.Println("start run sayHello()") // #1
	return "hello world!"
}

// HelloMessage : test init value
var HelloMessage = sayHello()

// demo 02, struct reference
type mySubStruct struct {
	id  uint
	val string
}

type mySuperStruct struct {
	sub mySubStruct // by value
	ex  string
}

type mySuperStructRef struct {
	sub *mySubStruct // by refrence
	ex  string
}

func testStructRefValue() {
	sub := mySubStruct{
		id:  10,
		val: "ten",
	}

	super := mySuperStruct{
		sub: sub,
		ex:  "number 10",
	}
	fmt.Printf("before => sub struct: %+v\n", super)

	superRef := mySuperStructRef{
		sub: &sub,
		ex:  "number 10",
	}
	fmt.Printf("before => sub struct ref: %+v\n", superRef.sub)

	sub.val = "TEN"
	fmt.Printf("after => sub struct: %+v\n", super)
	fmt.Printf("after => sub struct Ref: %+v\n", superRef.sub)
}

// demo 03, verify go version
func isGoVersionOK(baseVersion string) bool {
	currVersion := runtime.Version()[2:]
	currArr := strings.Split(currVersion, ".")
	baseArr := strings.Split(baseVersion, ".")

	for i := 0; i < 2; i++ { // check first 2 digits
		curr, _ := strconv.ParseInt(currArr[i], 10, 32)
		base, _ := strconv.ParseInt(baseArr[i], 10, 32)
		if curr == base {
			continue
		}
		return curr > base
	}
	return true // curr == base
}

func testGoVersion() {
	currVersion := runtime.Version()
	fmt.Printf("%s >= go1.15 is ok: %v\n", currVersion, isGoVersionOK("1.15"))
	fmt.Printf("%s >= go1.10 is ok: %v\n", currVersion, isGoVersionOK("1.10"))
	fmt.Printf("%s >= go1.9.3 is ok: %v\n", currVersion, isGoVersionOK("1.9.3"))
}

// demo 04, json keyword "omitempty"
func testJSONOmitEmpty() {
	type project struct {
		Name string `json:"name"`
		URL  string `json:"url"`
		Desc string `json:"desc"`
		Docs string `json:"docs,omitempty"`
	}

	p1 := project{
		Name: "CleverGo",
		URL:  "https://github.com/headwindfly/clevergo",
		Desc: "CleverGo Perf Framework",
		Docs: "https://github.com/headwindfly/clevergo/tree/master/docs",
	}
	if data, err := json.MarshalIndent(p1, "", "  "); err == nil {
		fmt.Println("json string:", string(data))
	}

	p2 := project{
		Name: "CleverGo",
		URL:  "https://github.com/headwindfly/clevergo",
	}
	if data, err := json.MarshalIndent(p2, "", "  "); err == nil {
		fmt.Println("json string:", string(data))
	}
}

// MainDemo04 : main
func MainDemo04() {
	// testStructRefValue()
	// testGoVersion()
	testJSONOmitEmpty()

	fmt.Println("demo 04 done.")
}
