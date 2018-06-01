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
	"time"

	"github.com/jmcvetta/randutil"
	"github.com/larspensjo/config"
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
	tmpl, err := template.New("test").Parse("hello, {{.}}\n")
	if err != nil {
		panic(err)
	}

	name := "zhengjin"
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
	patternEn := "{{.Count}} items are made of {{.Material}}\n"
	tmpl, err := template.New(tagEn).Parse(patternEn)
	if err != nil {
		panic(err)
	}
	tagCn := "Chinese"
	patternCn := "{{.Count}}个物料的材料是{{.Material}}\n"
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

// demo 06-04, parse single file
func testGoTemplate04() {
	// filePath := os.Getenv("ZJGOPRJ") + "src/demo.hello/demos/tmpl_cn.txt"
	filePath := "src/demo.hello/demos/tmpl_cn.txt"
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

// demo 06-05, parse file with nest template
func testGoTemplate05() {
	fileSubTmpl := "src/demo.hello/demos/sub.tmpl"
	fileTmpl := "src/demo.hello/demos/tmpl_en.txt"
	tmpl, err := template.ParseFiles(fileSubTmpl, fileTmpl)
	if err != nil {
		panic(err)
	}

	// sub.tmpl + Inventory => tmpl_en.txt => stdout
	err = tmpl.ExecuteTemplate(os.Stdout, "main", Inventory{"wool", 21})
	if err != nil {
		panic(err)
	}
}

// demo 06-06, parse files with nest template
func testGoTemplate06() {
	pattern := "src/demo.hello/demos/tmpl_*.txt"
	tmpl, err := template.ParseGlob(pattern)
	if err != nil {
		panic(err)
	}
	fileSubTmpl := "src/demo.hello/demos/sub.tmpl"
	tmpl, err = tmpl.New("sub").ParseFiles(fileSubTmpl)
	if err != nil {
		panic(err)
	}

	fmt.Println("#1. template en output:")
	err = tmpl.ExecuteTemplate(os.Stdout, "main", Inventory{"wool", 21})
	if err != nil {
		panic(err)
	}
	fmt.Println("#2. template cn output:")
	err = tmpl.ExecuteTemplate(os.Stdout, "tmpl_cn.txt", Inventory{"wool", 27})
	if err != nil {
		panic(err)
	}
}

// demo 07, read configs and set template
func readValuesFromConfigs(path, section string) map[string]string {
	// $ go get github.com/larspensjo/config
	cfg, err := config.ReadDefault(path)
	if err != nil {
		fmt.Errorf("Fail to find %s, error: %s", path, err)
	}

	retMap := make(map[string]string)
	if cfg.HasSection(section) {
		options, err := cfg.SectionOptions(section)
		if err == nil {
			for _, option := range options {
				value, err := cfg.String(section, option)
				if err == nil {
					retMap[option] = value
				}
			}
		}
	}
	if len(retMap) == 0 {
		panic(fmt.Sprintf("no options in section [%s]", section))
	}
	return retMap
}

const confFile = "src/demo.hello/demos/test.conf"

func testReadConfigs() {
	fmt.Println("read configs and set template string")
	goInfos := readValuesFromConfigs(confFile, "default")
	fmt.Println("go infos:")
	for k, v := range goInfos {
		fmt.Printf("%s=%s\n", k, v)
	}

	pattern := "Go version {{.version}}, and bin path {{.path}}\n"
	tmpl, err := template.New("goInfos").Parse(pattern)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, goInfos) // map instead of struct
	if err != nil {
		panic(err)
	}
}

func testBuildTemplate() {
	fmt.Println("sub.tmpl + test.conf => test_tmpl.conf => output.txt")

	// for sub tmpl, it supports diff data types, like array
	fileSubTmpl := "src/demo.hello/demos/sub.tmpl"
	fileTmpl := "src/demo.hello/demos/test_tmpl.conf"
	tmpl, err := template.ParseFiles(fileTmpl, fileSubTmpl)
	if err != nil {
		panic(err)
	}

	// for conf, it supports only key and value
	testInfos := readValuesFromConfigs(confFile, "test")

	pathOutput := "src/demo.hello/demos/output.txt"
	fOutput, err := os.OpenFile(pathOutput, os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(fOutput, "testInfos", testInfos)
	if err != nil {
		panic(err)
	}
	fOutput.Close()
	fmt.Println("write configs done.")
}

// demo 08, if or map
var fnGetMsgByID = func(id string) {
	fmt.Println("message id:", id)
}

var fnGetMsgByName = func(name string) {
	fmt.Println("message name:", name)
}

func getMsgByIf(tag, input string) {
	if tag == "id" {
		fnGetMsgByID(input)
	} else if tag == "name" {
		fnGetMsgByName(input)
	} else {
		fmt.Println("invalid argument!")
	}
}

func getMsgByMap(tag, input string) {
	fns := make(map[string]func(string))
	fns["id"] = fnGetMsgByID
	fns["name"] = fnGetMsgByName
	fns[tag](input)
}

func testGetMsgByIfAndMap() {
	tag := "name"
	name := "test"
	getMsgByIf(tag, name)
	getMsgByMap(tag, name)
}

// demo 09, time calculation
func testTimeSub() {
	start := time.Now()
	time.Sleep(2 * time.Second)
	duration := time.Now().Sub(start)
	fmt.Printf("time duration: %.2f\n", duration.Seconds())

	for int(time.Now().Sub(start).Seconds()) < 5 {
		fmt.Println("wait 1 second ...")
		time.Sleep(time.Second)
	}
}

// demo 10, test get random strings
func testRandomValues() {
	if num, err := randutil.IntRange(1, 10); err == nil {
		fmt.Println("get a random number:", num)
	}

	if str, err := randutil.String(10, randutil.Numerals); err == nil {
		fmt.Println("get string of random numbers:", str)
	}
	if str, err := randutil.String(10, randutil.Alphabet); err == nil {
		fmt.Println("get string of random string:", str)
	}
	if str, err := randutil.String(10, randutil.Alphanumeric); err == nil {
		fmt.Println("get string of random string:", str)
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
	// testGoTemplate06()

	// testReadConfigs()
	// testBuildTemplate()

	// testGetMsgByIfAndMap()
	// testTimeSub()
	// testRandomValues()

	fmt.Println("demo 04 done.")
}
