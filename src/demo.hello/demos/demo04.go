package demos

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/jmcvetta/randutil"
	"gopkg.in/mgo.v2/bson"
	zjutils "tools.test/utils"
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

// demo 06, if or map
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

// demo 07, time calculation
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

// demo 08, test get random strings
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

// demo 09, init random bytes
func initBytesBySize(size int) []byte {
	buf := make([]byte, size)
	for i := 0; i < len(buf); i++ {
		buf[i] = uint8(i % 10)
	}
	return buf
}

func testInitBytes() {
	b := initBytesBySize(16)
	fmt.Printf("init bytes by number: %d\n", b)
	fmt.Printf("init bytes by char: %c\n", b)
	fmt.Printf("init bytes by base64 str: %s\n", base64.StdEncoding.EncodeToString(b))
}

// demo 10, gzip encode and decode
func testGzipCode() {
	// srcb := []byte("gzip encode and decode, test.")
	srcb, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("src length:", len(srcb))

	destb, err := zjutils.GzipEncode(srcb)
	if err != nil {
		panic(err)
	}
	fmt.Println("gzip encode length:", len(destb))

	b, err := zjutils.GzipDecode(destb)
	if err != nil {
		panic(err)
	}
	fmt.Println("gzip decode length:", len(b))
	if len(b) <= 128 {
		fmt.Println("encode and decode bytes:", string(b))
	}
}

// demo 11, compress and decompress
func testGetFileName() {
	src := os.Getenv("HOME") + "/Downloads/tmp_files/tmp_dir_backup"
	if f, err := os.Open(src); err == nil {
		fmt.Println("file full name:", f.Name())
		if info, err := f.Stat(); err == nil {
			fmt.Println("file base name:", info.Name())
		}
	}
}

func testTarCompressFile() {
	src := os.Getenv("HOME") + "/Downloads/tmp_files/pics/upload.jpg"
	dest := os.Getenv("HOME") + "/Downloads/tmp_files/pics/upload.tar.gz"

	f, err := os.Open(src)
	if err != nil {
		fmt.Println("read src file error:", err.Error())
	}
	err = zjutils.Compress([]*os.File{f}, dest)
	if err != nil {
		fmt.Println("comporess error:", err.Error())
	}
}

func testTarCompressDir() {
	src := os.Getenv("HOME") + "/Downloads/tmp_files/tmp_dir"
	dest := os.Getenv("HOME") + "/Downloads/tmp_files/tmp_dir.tar.gz"

	if f, err := os.Open(src); err == nil {
		fmt.Printf("compress file (%s) with tar.gz\n", src)
		err := zjutils.Compress([]*os.File{f}, dest)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func testTarDecompress() {
	src := os.Getenv("HOME") + "/Downloads/tmp_files/tmp_dir.tar.gz"
	dest := os.Getenv("HOME") + "/Downloads/tmp_files"

	err := zjutils.DeCompress(src, dest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("decompress to:", dest)
}

func testGetTotalGoroutines() {
	printTotalGoroutines := func() {
		fmt.Println("*** total goroutines:", runtime.NumGoroutine())
	}

	const waitTime = 5
	ch := make(chan int, 10)

	printTotalGoroutines()
	for i := 0; i < 10; i++ {
		go func(ch chan<- int, num int) {
			time.Sleep(time.Duration(rand.Intn(waitTime)) * time.Second)
			ch <- num
		}(ch, i)
	}

	go func(ch chan int) {
		time.Sleep(time.Duration(waitTime+1) * time.Second)
		fmt.Println("close channel")
		close(ch)
		printTotalGoroutines()
	}(ch)

	printTotalGoroutines()
	time.Sleep(2 * time.Second)
	printTotalGoroutines()

	for num := range ch {
		fmt.Println("iterator at:", num)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("testGetTotalGoroutines done.")
}

// MainDemo04 : main
func MainDemo04() {
	// testStructRefValue()
	// testGoVersion()

	// testJSONOmitEmpty()
	// testBSONCases()

	// testGetMsgByIfAndMap()
	// testTimeSub()

	// testRandomValues()
	// testInitBytes()

	// testGzipCode()
	// testGetFileName()
	// testTarCompressFile()
	// testTarCompressDir()
	// testTarDecompress()

	// testGetTotalGoroutines()

	fmt.Println("golang demo04 DONE.")
}
