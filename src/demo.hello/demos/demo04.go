package demos

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"text/template"

	"gopkg.in/mgo.v2/bson"
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

// demo 05, bson
func testBSONCases() {
	type testStruct struct {
		FH  []byte `bson:"fh"`
		NFH []byte `bson:"nfh"`
	}

	srcFh := "Bpb_fwEAAAB3eK148Y4dFSvzt1ILAAAAMUMVAAAAAAAKqnHPAAAAAAny-rvibYqoFP-lPkI53JfmoIx5"
	srcNfh := "CJYxQxUAAAAAAAny-rvibYqoFP-lPkI53JfmoIx5a29kby10ZXN0LwUAAHJjUUyDsxizWg=="
	fh, err := base64.URLEncoding.DecodeString(srcFh)
	if err != nil {
		panic(err)
	}
	nfh, err := base64.URLEncoding.DecodeString(srcNfh)
	if err != nil {
		panic(err)
	}

	s := testStruct{
		FH:  fh,
		NFH: nfh,
	}
	if data, err := bson.Marshal(&s); err == nil {
		// parse bson bin file => $ bsondump fh.test1.bson
		ioutil.WriteFile("/Users/zhengjin/Downloads/tmp_files/fh.test.bson", data, 0666)
	}
}

// demo 06-01, go template, parse string
func testGoTemplate01() {
	name := "zhengjin"
	tmpl, err := template.New("test").Parse("hello, {{.}}\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, name)
	if err != nil {
		panic(err)
	}
}

// Inventory : struct used for template test
type Inventory struct {
	Material string
	Count    uint
}

// demo 06-02, go template, parse struct
func testGoTemplate02() {
	pattern := "{{.Count}} items are made of {{.Material}}\n"
	tmpl, err := template.New("test").Parse(pattern)
	if err != nil {
		panic(err)
	}

	sweaters := Inventory{"wool", 17}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

// demo 06-03, go template, multiple tmpls
func testGoTemplate03() {
	tagEn := "English"
	tagCn := "Chinese"
	patternEn := "{{.Count}} items are made of {{.Material}}\n"
	patternCn := "{{.Count}}个物料的材料是{{.Material}}\n"
	tmpl, err := template.New(tagEn).Parse(patternEn)
	if err != nil {
		panic(err)
	}
	tmpl, err = tmpl.New(tagCn).Parse(patternCn)
	if err != nil {
		panic(err)
	}

	sweaters := Inventory{"wool", 17}
	tmpl = tmpl.Lookup(tagEn)
	fmt.Println("Current template:", tmpl.Name())
	err = tmpl.ExecuteTemplate(os.Stdout, tagEn, sweaters)
	if err != nil {
		panic(err)
	}
	tmpl = tmpl.Lookup(tagCn)
	fmt.Println("Current template:", tmpl.Name())
	err = tmpl.ExecuteTemplate(os.Stdout, tagCn, sweaters)
	if err != nil {
		panic(err)
	}
}

const projectPath = "/Users/zhengjin/Workspaces/zj_projects/ZjGoProject/"

// demo 06-04, go template, parse single file
func testGoTemplate04() {
	filePath := projectPath + "src/demo.hello/demos/tmpl_en.txt"
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Current template:", tmpl.Name())
	err = tmpl.Execute(os.Stdout, Inventory{"wool", 21})
	if err != nil {
		panic(err)
	}
}

// demo 06-05, go template, parse files
func testGoTemplate05() {
	filePath := projectPath + "src/demo.hello/demos/tmpl_*.txt"
	tmpl, err := template.ParseGlob(filePath)
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(os.Stdout, "tmpl_en.txt", Inventory{"wool", 21})
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "tmpl_cn.txt", Inventory{"wool", 27})
	if err != nil {
		panic(err)
	}
}

// MainDemo04 : main
func MainDemo04() {
	// testStructRefValue()
	// testGoVersion()

	// testJSONOmitEmpty()
	// testBSONCases()

	// testGoTemplate01()
	// testGoTemplate02()
	// testGoTemplate03()
	// testGoTemplate04()
	// testGoTemplate05()

	fmt.Println("demo 04 done.")
}
