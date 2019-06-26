package demos

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/jmcvetta/randutil"
	"gopkg.in/mgo.v2/bson"
	myutils "tools.test/utils"
)

// demo, var "HelloMessage" init before init() function
func init() {
	fmt.Println("[demo04.go] init") // #2
}

// HelloMsg : public var, invoked from main.go
var HelloMsg = sayHello()

func sayHello() string {
	fmt.Println("start run sayHello()") // #1
	return "hello world!"
}

// demo, get file base and full name
func testGetFileName() {
	srcPath := os.Getenv("HOME") + "/Downloads/tmp_files"
	f, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nfile full name:", f.Name())

	info, err := f.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("file base name:", info.Name())
}

// demo, verify go version
func testGoVersion() {
	curVersion := runtime.Version()
	fmt.Printf("\n%s >= go1.15: %v\n", curVersion, isGoVersionOK("1.15"))
	fmt.Printf("%s >= go1.10: %v\n", curVersion, isGoVersionOK("1.10"))
	fmt.Printf("%s >= go1.9.3: %v\n", curVersion, isGoVersionOK("1.9.3"))
}

func isGoVersionOK(baseVersion string) bool {
	curVersion := runtime.Version()[2:]
	curArr := strings.Split(curVersion, ".")
	baseArr := strings.Split(baseVersion, ".")

	for i := 0; i < 2; i++ { // check first 2 digits
		cur, _ := strconv.ParseInt(curArr[i], 10, 32)
		base, _ := strconv.ParseInt(baseArr[i], 10, 32)
		if cur == base {
			continue
		}
		return cur > base
	}
	return true // cur == base
}

// demo, time calculation
func testTimeOpSub() {
	start := time.Now()
	time.Sleep(time.Duration(2) * time.Second)
	duration := time.Now().Sub(start)
	fmt.Printf("time duration: %.2f\n", duration.Seconds())

	for int(time.Now().Sub(start).Seconds()) < 5 {
		fmt.Println("wait 1 second ...")
		time.Sleep(time.Second)
	}
}

// demo, test get random strings
func testRandomValues() {
	if num, err := randutil.IntRange(1, 10); err == nil {
		fmt.Println("get a random number 1-10:", num)
	}

	if str1, err := randutil.String(10, randutil.Numerals); err == nil {
		fmt.Println("get string of 10 chars (random number):", str1)
	}
	if str2, err := randutil.String(10, randutil.Alphabet); err == nil {
		fmt.Println("get string of 10 chars (random alphabet):", str2)
	}
	if str3, err := randutil.String(10, randutil.Alphanumeric); err == nil {
		fmt.Println("get string of 10 chars (random number and alphabet):", str3)
	}
}

// demo, init random bytes
func testInitBytes() {
	buf := initBytesBySize(32)
	fmt.Printf("init bytes print as numbers: %d\n", buf)
	fmt.Printf("init bytes print as chars: %c\n", buf)

	str := base64.StdEncoding.EncodeToString(buf)
	fmt.Printf("init bytes print as base64 string: %s\n", str)
}

func initBytesBySize(size int) []byte {
	// init []byte "buf" with size of zero
	buf := make([]byte, size)
	for i := 0; i < len(buf); i++ {
		buf[i] = uint8(i % 16)
	}
	return buf
}

// demo, struct reference
type mySuperStruct struct {
	id  uint
	val string
}

type mySubStruct struct {
	super mySuperStruct // by value
	desc  string
}

type mySubStructRef struct {
	super *mySuperStruct // by refrence
	desc  string
}

func testStructRefValue() {
	s := mySuperStruct{
		id:  10,
		val: "test10",
	}

	subVal := mySubStruct{
		super: s,
		desc:  "inherit from super by value",
	}
	fmt.Printf("before => sub struct: %+v\n", subVal)

	subRef := mySubStructRef{
		super: &s,
		desc:  "inherit from super by reference",
	}
	fmt.Printf("before => sub struct ref: %+v\n", subRef.super)

	s.val = strings.ToUpper(s.val)
	fmt.Printf("after => sub struct: %+v\n", subVal)
	fmt.Printf("after => sub struct Ref: %+v\n", subRef.super)
}

// demo, if and map
var fnPrintMsgID = func(id string) {
	fmt.Println("message id:", id)
}

var fnPrintMsgName = func(name string) {
	fmt.Println("message name:", name)
}

func printMsgByIf(tag, input string) {
	fmt.Println("\nprint message by if condition.")
	if tag == "id" {
		fnPrintMsgID(input)
	} else if tag == "name" {
		fnPrintMsgName(input)
	} else {
		fmt.Println("invalid argument!")
	}
}

func printMsgByMap(tag, input string) {
	fmt.Println("\nprint message by map.")
	fns := make(map[string]func(string))
	fns["id"] = fnPrintMsgID
	fns["name"] = fnPrintMsgName
	fns[tag](input)
}

func testPrintMsgByCond() {
	tag := "id"
	name := "message01"
	printMsgByIf(tag, name)
	printMsgByMap(tag, name)
}

// demo, json keyword "omitempty"
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
		fmt.Println("\nmarshal json string:", string(data))
	}

	p2 := project{
		Name: "CleverGo",
		URL:  "https://github.com/headwindfly/clevergo",
	}
	if data, err := json.MarshalIndent(p2, "", "  "); err == nil {
		fmt.Println("marshal json string with omitempty:", string(data))
	}
}

// demo, bson parser
func testBSONParser() {
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
		savePath := filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files/fh.test.bson")
		if err := ioutil.WriteFile(savePath, data, 0666); err != nil {
			panic(err)
		}
		fmt.Printf("save bson bin file: %s\nparse bson: 'bsondump fh.test.bson'\n", savePath)
	}
}

// demo, gzip encode and decode
func testGzipCode() {
	// srcb := []byte("gzip encode and decode, test.")
	srcb, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("src length:", len(srcb))

	destb, err := myutils.GzipEncode(srcb)
	if err != nil {
		panic(err)
	}
	fmt.Println("gzip encode length:", len(destb))

	b, err := myutils.GzipDecode(destb)
	if err != nil {
		panic(err)
	}
	fmt.Println("gzip decode length:", len(b))
	if len(b) <= 128 {
		fmt.Println("encode and decode bytes:", string(b))
	}
}

// demo, compress and decompress
func testTarCompressFile() {
	src := os.Getenv("HOME") + "/Downloads/tmp_files/pics/upload.jpg"
	dest := os.Getenv("HOME") + "/Downloads/tmp_files/pics/upload.tar.gz"

	f, err := os.Open(src)
	if err != nil {
		fmt.Println("read src file error:", err.Error())
	}
	err = myutils.Compress([]*os.File{f}, dest)
	if err != nil {
		fmt.Println("comporess error:", err.Error())
	}
}

func testTarCompressDir() {
	src := os.Getenv("HOME") + "/Downloads/tmp_files/tmp_dir"
	dest := os.Getenv("HOME") + "/Downloads/tmp_files/tmp_dir.tar.gz"

	if f, err := os.Open(src); err == nil {
		fmt.Printf("compress file (%s) with tar.gz\n", src)
		err := myutils.Compress([]*os.File{f}, dest)
		if err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func testTarDecompress() {
	src := os.Getenv("HOME") + "/Downloads/tmp_files/tmp_dir.tar.gz"
	dest := os.Getenv("HOME") + "/Downloads/tmp_files"

	err := myutils.DeCompress(src, dest)
	if err != nil {
		panic(err)
	}
	fmt.Println("decompress to:", dest)
}

// demo, get routines count
func testGetGoroutinesCount() {
	printRoutineCount := func() {
		fmt.Println("***** goroutines count:", runtime.NumGoroutine())
	}

	printRoutineCount() // 1
	const waitTime = 5
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go func(ch chan<- int, num int) {
			sleep, err := randutil.IntRange(2, waitTime)
			if err != nil {
				fmt.Println(err)
				sleep = waitTime
			}
			time.Sleep(time.Duration(sleep) * time.Second)
			ch <- num
		}(ch, i)
	}

	go func(ch chan int) {
		time.Sleep(time.Duration(waitTime+2) * time.Second)
		printRoutineCount() // 2
		fmt.Println("close channel")
		close(ch)
	}(ch)

	time.Sleep(time.Second)
	printRoutineCount() // 12 (10 + 1 + 1)

	for num := range ch {
		fmt.Println("iterator at:", num)
	}
	time.Sleep(time.Second)
	printRoutineCount() // 1
	fmt.Println("testGetGoroutinesCount DONE.")
}

// MainDemo04 : main
func MainDemo04() {
	// testGetFileName()
	// testGoVersion()
	// testTimeOpSub()
	// testRandomValues()
	// testInitBytes()

	// testStructRefValue()
	// testPrintMsgByCond()

	// testJSONOmitEmpty()
	// testBSONParser()

	// testGzipCode()
	// testTarCompressFile()
	// testTarCompressDir()
	// testTarDecompress()

	// testGetGoroutinesCount()

	fmt.Println("golang demo04 DONE.")
}
