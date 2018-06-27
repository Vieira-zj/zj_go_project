package mocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"strings"
	"time"

	httpv1 "qbox.us/httputil.v1"
	"qbox.us/rpc"
	zjutils "utils.zhengjin/utils"
)

const testFilePath = "./test.file"

var total int

// MockDefault : default page
func MockDefault(rw http.ResponseWriter, req *http.Request) {
	total++
	log.Printf("\n===> MockDefault, access at %d time\n", total)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	req.ParseForm()
	retContent := "PAGE NOT FOUND!\n"
	retContent += fmt.Sprintf("request uri: %s\n", req.RequestURI)
	if len(req.Form) > 0 {
		retContent += fmt.Sprintf("raw query: %+v\n", req.Form)
	}
	if len(retContent) == 0 {
		retContent = "null\n"
	}

	rw.Header().Set("Content-Md5", strconv.Itoa(len([]byte(retContent))))
	rw.WriteHeader(http.StatusNotFound)
	log.Println("return 404")

	io.Copy(rw, strings.NewReader(retContent))
	log.Print("===> MockDefault, send data done\n\n")
}

var total01 int

// Mock01 : mock data stream and file download
func Mock01(rw http.ResponseWriter, req *http.Request) {
	total01++
	log.Printf("\n===> Mock01, access at %d time\n", total01)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	req.ParseForm()
	var b []byte
	isFile := GetBoolInReqForm(req, "isFile")
	if isFile {
		b = ReadBytesFromFile(testFilePath)
	} else {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "null"
		}
		b = []byte(fmt.Sprintf("from Host: %s\n", hostname) + "from Mock01, mock returned text")
		// b = InitBytesBySize(1024)
	}

	wait := GetNumberInReqForm(req, "wait")
	// wait before send header
	if wait > 0 {
		fmt.Printf("wait %d seconds\n", wait)
		time.Sleep(time.Duration(wait) * time.Second)
	}

	// mockMD5 := "f900b997e6f8a772994876dff023801e"
	// rw.Header().Set("Content-Md5", mockMD5) // mock md5
	rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
	rw.WriteHeader(http.StatusOK)
	log.Println("return code => 200")

	// wait after send header
	// if wait > 0 {
	// 	fmt.Printf("wait %d seconds\n", wait)
	// 	time.Sleep(wait * time.Second)
	// }
	io.Copy(rw, bytes.NewReader(b))
	log.Print("===> Mock01, send data done\n\n")
}

var total02 int

// Mock02 : mock bytes stream by flush
func Mock02(rw http.ResponseWriter, req *http.Request) {
	total02++
	log.Printf("\n===> Mock02, access at %d time\n", total02)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	rw.WriteHeader(http.StatusOK)
	log.Println("return code => 200")

	const contentLen = 2048
	for i := 0; i < 50; i++ {
		log.Println("mock body")
		time.Sleep(time.Duration(200) * time.Millisecond)
		_, err := io.Copy(rw, bytes.NewReader(InitBytesBySize(contentLen)))
		if err != nil {
			log.Fatalf("error: %v\n", err)
			break
		}
		rw.(http.Flusher).Flush()
	}

	b := []byte("from Mock02, mock text end")
	io.Copy(rw, bytes.NewReader(b))
	log.Print("Mock02, send data done\n\n")
}

var total03 int

// Mock03 : mock ret error code, ex 404, 503
func Mock03(rw http.ResponseWriter, req *http.Request) {
	total03++
	log.Printf("\n===> Mock03, access at %d time\n", total03)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	req.ParseForm()
	retCode := GetNumberInReqForm(req, "retCode")
	if retCode == 0 {
		retCode = 200
	}

	b := []byte("from Mock03, mock returned text")
	rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
	rw.WriteHeader(retCode)
	log.Printf("return code => %d", retCode)

	io.Copy(rw, bytes.NewReader(b))
	log.Print("===> Mock03, send data done\n\n")
}

var total04 int

// Mock04 : mock file download by range
func Mock04(rw http.ResponseWriter, req *http.Request) {
	total04++
	log.Printf("\n===> Mock04, access at %d time\n", total04)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))
	// data block is set from request header => [Range]:[bytes=0-4095]
	// for qiniuproxy, default range is 4M

	req.ParseForm()

	retCode := GetNumberInReqForm(req, "retCode")
	// for 4xx, no connection retry; 5xx, connection retry
	if retCode != 0 && retCode != 200 {
		retText := "Mock04, mock error string"
		rw.Header().Set("Content-Length", strconv.Itoa(len(retText)))
		rw.WriteHeader(retCode)
		log.Printf("return code => %d\n", retCode)
		io.Copy(rw, strings.NewReader(retText))
		log.Print("===> Mock04, send data done\n\n")
		return
	}

	isFile := GetBoolInReqForm(req, "isFile")
	isTest := GetBoolInReqForm(req, "isTest")
	var buf []byte
	if isFile {
		fmt.Println("read bytes from file")
		buf = ReadBytesFromFile(testFilePath)
	} else if isTest {
		fmt.Println("mock string body")
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "null"
		}
		buf = []byte(fmt.Sprintf("from Host: %s\n", hostname) + "from Mock04, mock returned text")
	} else {
		fmt.Println("mock bytes body")
		buf = InitBytesBySize(1024 * 1024 * 10)
	}

	// md5 check for non range request
	_, ok := req.Header["Range"]
	if !ok {
		isMD5 := GetBoolInReqForm(req, "isMd5")
		if isMD5 {
			md5 := GetBoolInReqForm(req, "md5")
			if md5 {
				rw.Header().Set("Content-MD5", zjutils.GetMd5ForText(string(buf)))
				// rw.Header().Set("Content-MD5", encode.GetURLBasedMd5ForText(string(buf)))
			} else {
				errMD5 := "0980a9e10670ccc4895432d4b4ae9fff"
				rw.Header().Set("Content-MD5", errMD5)
			}
		}
	}

	// Content-Length and Return Code handle by ReplyRange()
	// rw.WriteHeader(retCode)
	// log.Println("return code => 200")

	// send data by range
	waitForEachRead := 0
	rw = rpc.ResponseWriter{rw}
	// rr := rpc.ReadSeeker2RangeReader{bytes.NewReader(buf)}
	rr := createMockReader(buf, waitForEachRead)
	rw.(rpc.ResponseWriter).ReplyRange(rr, int64(len(buf)), &rpc.Metas{}, req)
	log.Print("===> Mock04, send blocked data done\n\n")
}

func createMockReader(buf []byte, waitForReader int) rpc.ReadSeeker2RangeReader {
	return rpc.ReadSeeker2RangeReader{&mockReader{wait: waitForReader, r: bytes.NewReader(buf)}}
}

type mockReader struct {
	wait int
	r    *bytes.Reader
}

func (mr *mockReader) Read(b []byte) (int, error) {
	if mr.wait > 0 {
		fmt.Printf("wait %d ms\n", mr.wait)
		time.Sleep(time.Duration(mr.wait) * time.Millisecond)
	}
	len, err := mr.r.Read(b)
	return len, err
}

func (mr *mockReader) Seek(offset int64, whence int) (int64, error) {
	return mr.r.Seek(offset, whence)
}

var total05 int

// Mock05 : mock stream data by wait and kb
func Mock05(rw http.ResponseWriter, req *http.Request) {
	total05++
	log.Printf("\n===> Mock05, access at %d time\n", total05)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	req.ParseForm()
	wait := GetNumberInReqForm(req, "wait")
	if wait == 0 {
		wait = 3
	}
	kb := GetNumberInReqForm(req, "kb")
	if kb == 0 {
		kb = 1
	}

	b := InitBytesBySize(1024 * kb)
	rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
	rw.WriteHeader(http.StatusOK)
	log.Println("return code => 200")

	io.Copy(rw, &mockReader{wait: wait, r: bytes.NewReader(b)})
	log.Print("===> Mock05, send data done\n\n")
}

var total06 int

// Mock06 : mock httpdns server
func Mock06(rw http.ResponseWriter, req *http.Request) {
	total06++
	log.Printf("===> Mock06, access dns server at %d time\n", total06)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	// ret := `{"errno":-1, "iplist":[]}`
	retIP := `"10.200.20.21"`
	// retIP := `"42.48.232.7", "10.200.20.21"`
	ret := fmt.Sprintf(`{"errno":0, "iplist":[%s]}`, retIP)

	retCode := 200
	if total06%5 == 0 {
		retCode = 500
	}
	rw.Header().Set("Content-Length", strconv.Itoa(len(ret)))
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(retCode)
	log.Printf("return code => %d\n", retCode)

	if total06%4 == 0 {
		wait := 3
		time.Sleep(time.Duration(wait) * time.Second)
		log.Printf("sleep %d seconds\n", wait)
	}

	io.Copy(rw, strings.NewReader(ret))
	log.Print("===> Mock06, send data done\n\n")
}

var total07 int

// Mock07 : mock mirror file server
func Mock07(rw http.ResponseWriter, req *http.Request) {
	total07++
	log.Printf("\n===> Mock07, access mirror at %d time\n", total07)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	rw.WriteHeader(http.StatusOK)
	log.Println("return code => 200")
	io.Copy(rw, strings.NewReader("success"))
	rw.(http.Flusher).Flush()

	if total07%2 == 0 {
		wait := 3
		time.Sleep(time.Duration(wait) * time.Second)
		log.Printf("sleep %d seconds\n", wait)
	}

	// io.Copy(rw, Strings.NewReader("** test content"))
	b := ReadBytesFromFile(testFilePath)
	io.Copy(rw, bytes.NewReader(b))
	log.Print("===> Mock07, data returned\n\n")
}

var total08 = 0

type refreshRet struct {
	Code      int    `json:"code"`
	Error     string `json:"error"`
	RequestID string `json:"requestId"`
}

// Mock08 : handle cdn refresh post request, and return
func Mock08(rw http.ResponseWriter, req *http.Request) {
	total08++
	log.Printf("\n===> Mock08, access mirror at %d time\n", total08)
	reqData, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqData), "\n"))

	// result, _ := ioutil.ReadAll(req.Body)
	// defer req.Body.Close()
	// fmt.Printf("request body: %s\n", string(result))

	JSONObj := refreshRet{
		Code:      http.StatusOK,
		Error:     "null",
		RequestID: "cdnrefresh-test-01",
	}
	b, err := json.Marshal(JSONObj)
	if err != nil {
		log.Fatalln(err.Error())
		b = []byte("json encode error")
	}

	rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
	rw.WriteHeader(http.StatusOK)
	log.Println("return code => 200")

	io.Copy(rw, bytes.NewReader(b))
	log.Print("===> Mock08, data returned\n\n")
}

var total09 int

// Mock09 : mock download file by arg "start"
func Mock09(rw http.ResponseWriter, req *http.Request) {
	total09++
	log.Printf("\n===> Mock09, access at %d time\n", total09)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	if err := req.ParseForm(); err != nil {
		log.Fatalln(err.Error())
	}

	var filepath string
	start := GetNumberInReqForm(req, "start")
	if start < 1000 {
		filepath = "./test1.file"
	} else {
		filepath = "./test2.file"
	}
	b := ReadBytesFromFile(filepath)

	rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
	rw.WriteHeader(http.StatusOK)
	log.Println("return code => 200")

	time.Sleep(500 * time.Millisecond)
	io.Copy(rw, bytes.NewReader(b))
	log.Print("===> mock09, send data done\n\n")
}

var total10 int

// Mock10 : mock server disconnect
func Mock10(rw http.ResponseWriter, req *http.Request) {
	total10++
	log.Printf("\n===> Mock10, access at %d time\n", total10)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	req.ParseForm()
	b := ReadBytesFromFile(testFilePath)
	// set resp header Content-Length
	if GetBoolInReqForm(req, "isSetLen") {
		rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
	}

	rw.WriteHeader(http.StatusOK)
	log.Println("return code => 200")

	wait := GetNumberInReqForm(req, "wait")
	if wait == 0 {
		wait = 3
	}
	go func() {
		time.Sleep(time.Duration(wait) * time.Second)
		if jacker, ok := httpv1.GetHijacker(rw); ok {
			conn, _, err := jacker.Hijack()
			if err != nil {
				fmt.Printf("hijack err: %v\n", err)
			} else {
				log.Println("response can hijack, connection closed")
				conn.Close()
			}
		} else {
			fmt.Printf("http.ResponseWriter not http.Hijacker")
		}
	}()

	len, err := io.Copy(rw, bytes.NewReader(b))
	if err != nil {
		log.Println("copy resp writer error:", err.Error())
		return
	}
	fmt.Printf("copied length: %d\n", len)
	log.Print("===> mock10, send data done\n\n")
}

var total11 int

// Mock11 : mock gzip and chunk
func Mock11(rw http.ResponseWriter, req *http.Request) {
	total11++
	log.Printf("\n===> Mock11, access at %d time\n", total11)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	srcb := ReadBytesFromFile(testFilePath)
	retb, err := zjutils.GzipEncode(srcb)
	if err != nil {
		retErr := fmt.Sprintln("error in gzip encode:", err.Error())
		rw.Header().Set("Content-Length", strconv.Itoa(len(retErr)))
		rw.WriteHeader(599)
		log.Println("return code => 599")
		io.Copy(rw, strings.NewReader(retErr))
		return
	}

	rw.Header().Set("Content-Encoding", "gzip")
	rw.WriteHeader(http.StatusOK)
	log.Println("return code => 200")

	io.Copy(rw, bytes.NewReader(retb))
	log.Print("===> mock11, send data done\n\n")
	// resp headers:
	// Content-Type: application/x-gzip
	// Transfer-Encoding: chunked
}

var total12 int

// Mock12 : mimetype
func Mock12(rw http.ResponseWriter, req *http.Request) {
	total12++
	log.Printf("\n===> Mock12, access at %d time\n", total12)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	mimetypeTable := make(map[string]string)
	mimetypeTable["txt"] = "text/plain"
	mimetypeTable["jpg"] = "image/jpeg"
	mimetypeTable["bin"] = "application/octet-stream"

	req.ParseForm()
	mimetype := GetStringInReqForm(req, "type")
	if len(mimetype) == 0 {
		mimetype = "txt"
	}
	rw.Header().Set("Content-Type", mimetypeTable[mimetype])

	var b []byte
	isLenZero := GetBoolInReqForm(req, "lenzero")
	if !isLenZero {
		path := testFilePath + "." + mimetype
		b = ReadBytesFromFile(path)
	}

	rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
	rw.WriteHeader(http.StatusOK)
	log.Println("return code => 200")

	if len(b) > 0 {
		rw.(http.Flusher).Flush() // write response headers
		time.Sleep(2 * time.Second)
		io.Copy(rw, bytes.NewReader(b))
	}
	log.Print("===> mock12, send data done\n\n")
}

// GetStringInReqForm : return string value from request query form
func GetStringInReqForm(req *http.Request, argName string) string {
	if len(req.Form) > 0 {
		val, ok := req.Form[argName]
		if ok {
			return val[0]
		}
	}
	return ""
}

// GetNumberInReqForm : return int value from request query form
func GetNumberInReqForm(req *http.Request, argName string) int {
	tmp := GetStringInReqForm(req, argName)
	if ret, err := strconv.Atoi(tmp); err == nil {
		return ret
	}
	return 0
}

// GetBoolInReqForm : return bool value from request query form
func GetBoolInReqForm(req *http.Request, argName string) bool {
	tmp := GetStringInReqForm(req, argName)
	if ret, err := strconv.ParseBool(tmp); err == nil {
		return ret
	}
	return false
}

// InitBytesBySize : get mock bytes by size
func InitBytesBySize(size int) []byte {
	buf := make([]byte, size)
	for i := 0; i < len(buf); i++ {
		buf[i] = uint8(i % 10)
	}
	return buf
}

// ReadBytesFromFile : get bytes from file
func ReadBytesFromFile(path string) []byte {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err.Error())
		return []byte("null")
	}
	return buf
}

func isFileExist(filepath string) bool {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}
