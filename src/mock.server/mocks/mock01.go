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
	"utils.project/encode"
	"utils.project/etag"
)

const (
	mockMd5      = "f900b997e6f8a772994876dff023801e"
	testFilePath = "./test.file"
)

var total int

// MockDefault : default page
func MockDefault(rw http.ResponseWriter, req *http.Request) {
	total++
	log.Printf("access default at %d time\n", total)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	// log.Println("return 404")
	// rw.WriteHeader(http.StatusNotFound)
	log.Println("return 200")
	rw.WriteHeader(http.StatusOK)
	time.Sleep(time.Second)

	req.ParseForm()
	var lines []string
	// lines = append(lines, "not found!")
	lines = append(lines, fmt.Sprint("request uri: ", req.RequestURI))
	if len(req.Form) > 0 {
		lines = append(lines, fmt.Sprintf("request query: %+v", req.Form))
	}

	var b []byte
	if len(lines) > 1 {
		b = []byte(strings.Join(lines, "\n"))
	} else if len(lines) == 1 {
		b = []byte(lines[0])
	} else {
		b = []byte("default")
	}
	io.Copy(rw, bytes.NewReader(b))
	log.Println("send data done")
}

var total01 int

// Mock01 : mock bytes stream, file donwload, with isFile, wait
func Mock01(rw http.ResponseWriter, req *http.Request) {
	total01++
	log.Printf("access at %d time\n", total01)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	req.ParseForm()
	isFile := getQueryValueByName(req, "isFile")
	if isFile == "" {
		isFile = "false"
	}
	wait := getQueryValueByName(req, "wait")
	if wait != "" {
		fmt.Printf("wait %s seconds\n", wait)
		w, _ := strconv.Atoi(wait)
		time.Sleep(time.Duration(w) * time.Second)
	}

	log.Println("return 200")
	// rw.Header().Set("Content-Md5", mockMd5) // mock md5
	rw.WriteHeader(http.StatusOK)
	// time.Sleep(time.Second)

	b := []byte("mock string data")
	// b := initBytesBySize(1024)
	if isFile, _ := strconv.ParseBool(isFile); isFile {
		b = readBytesFromFile(testFilePath)
	}
	io.Copy(rw, bytes.NewReader(b))
	log.Println("send data done")
}

// Mock02 : mock bytes stream by flush
func Mock02(rw http.ResponseWriter, req *http.Request) {
	log.Println("200")
	rw.Header().Set("Content-Md5", mockMd5) // mock md5
	rw.Header().Set("Content-Length", "1000000")
	rw.WriteHeader(http.StatusOK)

	for i := 0; i < 50; i++ {
		log.Println("mock body")
		time.Sleep(time.Duration(200) * time.Millisecond)
		_, err := io.Copy(rw, bytes.NewReader(initBytesBySize(2048)))
		if err != nil {
			log.Fatalf("error: %v\n", err)
			return
		}
		rw.(http.Flusher).Flush()

		// if i == 10 {
		// 	const wait = 15
		// 	time.Sleep(wait * time.Second)
		// }
	}

	b := []byte("mock string data")
	io.Copy(rw, bytes.NewReader(b))
	log.Println("send data done")
}

var total03 int

// Mock03 : mock ret code by "retCode", ex 404, 503
func Mock03(rw http.ResponseWriter, req *http.Request) {
	total03++
	log.Printf("access at %d time\n", total03)

	req.ParseForm()
	retCode := getQueryValueByName(req, "retCode")
	if retCode == "" {
		retCode = "200"
	}

	time.Sleep(time.Second)
	code, _ := strconv.Atoi(retCode)
	log.Printf("return %d\n", code)
	rw.WriteHeader(code)

	b := []byte("mock string data")
	io.Copy(rw, bytes.NewReader(b))
}

func getQueryValueByName(req *http.Request, argName string) string {
	if len(req.Form) > 0 && len(req.Form[argName]) > 0 {
		return req.Form[argName][0]
	}
	return ""
}

var total04 int

// Mock04 : mock server for file download by range 4M
func Mock04(rw http.ResponseWriter, req *http.Request) {
	// download: curl -o ./test.file http://127.0.0.1:17890/index4/
	total04++
	log.Printf("access at %d time\n", total04)
	reqHeader, _ := httputil.DumpRequest(req, true)
	log.Println(string(reqHeader))

	req.ParseForm()
	parmRetCode := getQueryValueByName(req, "retCode")
	if parmRetCode == "" {
		parmRetCode = "200"
	}
	// for 5xx, connection retry
	// for 4xx, no connection retry
	if parmRetCode != "200" {
		if errCode, err := strconv.Atoi(parmRetCode); err == nil && total04 >= 3 {
			log.Printf("ret code: %d\n", errCode)
			rw.WriteHeader(errCode)
			return
		}
	}

	parmIsFile := getQueryValueByName(req, "isFile")
	if len(parmIsFile) == 0 {
		parmIsFile = "false"
	}

	// data block is set from request header => [Range]:[bytes=0-4095]
	// for qiniuproxy, default block is 4M
	var buf []byte
	if isFile, err := strconv.ParseBool(parmIsFile); err == nil && isFile {
		fmt.Println("read bytes from file")
		buf = readBytesFromFile(testFilePath)
	} else {
		fmt.Println("mock bytes")
		// buf = initBytesBySize(4096 * 16)
		buf = []byte("mock return string")
	}
	// file size check
	// if total04%3 == 0 {
	// 	buf = initBytesBySize(1024 * 1024 * 20)
	// }

	parmMd5Check := getQueryValueByName(req, "md5")
	if len(parmMd5Check) > 0 {
		md5check, err := strconv.ParseBool(parmMd5Check)
		if err == nil && md5check {
			rw.Header().Set("Content-MD5", encode.GetMd5ForText(string(buf)))
			// rw.Header().Set("Content-MD5", encode.GetURLBasedMd5ForText(string(buf)))
		} else {
			rw.Header().Set("Content-MD5", "0980a9e10670ccc4895432d4b4ae99ff")
		}
	}

	parmEtagCheck := getQueryValueByName(req, "etag")
	if len(parmEtagCheck) > 0 {
		strEtag, err := etag.GetEtagForText(string(buf))
		if err != nil {
			panic(err)
		}
		etagcheck, err := strconv.ParseBool(parmEtagCheck)
		if err == nil && etagcheck {
			rw.Header().Set("ETag", strEtag)
		} else {
			rw.Header().Set("ETag", "FmDc-7ioTJvtbSdoD7X3hHioXCPt")
		}
	}
	// if total04%3 == 0 {
	// 	rw.Header().Set("ETag", "FuujQKlyfG21iOsvBumnJuGNzjp1")
	// }

	// send data
	waitForEachRead := 0
	rw = rpc.ResponseWriter{rw}
	// rr := rpc.ReadSeeker2RangeReader{bytes.NewReader(buf)}
	rr := createMockReader(buf, waitForEachRead)
	rw.(rpc.ResponseWriter).ReplyRange(rr, int64(len(buf)), &rpc.Metas{}, req)
	log.Println("send blocked data done")
}

func initBytesBySize(size int) []byte {
	buf := make([]byte, size)
	for i := 0; i < len(buf); i++ {
		buf[i] = uint8(i % 10)
	}
	return buf
}

func readBytesFromFile(path string) []byte {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return buf
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
	log.Printf("access at %d time\n", total05)
	reqHeader, _ := httputil.DumpRequest(req, true)
	log.Println(string(reqHeader))

	req.ParseForm()
	wait := 20
	contentWait := getQueryValueByName(req, "wait")
	if contentWait != "" {
		wait, _ = strconv.Atoi(contentWait)
	}
	kb := 1
	contentKb := getQueryValueByName(req, "kb")
	if contentWait != "" {
		kb, _ = strconv.Atoi(contentKb)
	}

	rw.WriteHeader(http.StatusOK)

	b := initBytesBySize(1024 * kb)
	io.Copy(rw, &mockReader{wait: wait, r: bytes.NewReader(b)})
	log.Println("send data done")
}

var total06 int

// Mock06 : mock httpdns server
func Mock06(rw http.ResponseWriter, req *http.Request) {
	total06++
	log.Printf("access dns server at %d time\n", total06)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	retCode := 200
	if total06%5 == 0 {
		retCode = 500
	}
	log.Printf("return %d\n", retCode)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(retCode)

	wait := 5
	if total06%4 == 0 {
		log.Printf("sleep %d seconds\n", wait)
		time.Sleep(time.Duration(wait) * time.Second)
	}

	// ret := `{"errno":-1, "iplist":[]}`
	retIP := `"10.200.20.21"`
	// retIP := `"42.48.232.7", "10.200.20.21"`
	ret := fmt.Sprintf(`{"errno":0, "iplist":[%s]}`, retIP)
	io.Copy(rw, strings.NewReader(ret))
	log.Printf("return %s\n", ret)
}

var total07 int

// Mock07 : mock mirror file server
func Mock07(rw http.ResponseWriter, req *http.Request) {
	total07++
	log.Printf("access mirror at %d time\n", total07)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	log.Println("return 200")
	rw.WriteHeader(http.StatusOK)
	io.Copy(rw, strings.NewReader("success"))
	rw.(http.Flusher).Flush()

	wait := 3
	if total07%2 == 0 {
		log.Printf("sleep %d seconds\n", wait)
		time.Sleep(time.Duration(wait) * time.Second)
	}

	// io.Copy(rw, Strings.NewReader("** test content"))
	f, _ := os.Open(testFilePath)
	defer f.Close()
	io.Copy(rw, f)

	log.Println("data returned")
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
	log.Printf("\n***** access mirror at %d time\n", total08)
	reqData, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqData), "\n"))

	// result, _ := ioutil.ReadAll(req.Body)
	// defer req.Body.Close()
	// fmt.Printf("request body: %s\n", string(result))

	rw.WriteHeader(http.StatusOK)
	retJSONObj := refreshRet{
		Code:      http.StatusOK,
		Error:     "null",
		RequestID: "cdn-refresh-test",
	}
	if retBytes, err := json.Marshal(retJSONObj); err == nil {
		fmt.Fprintf(rw, string(retBytes))
	} else {
		fmt.Println(err.Error())
	}
}

var total09 int

// Mock09 : mock download file by arg "start"
func Mock09(rw http.ResponseWriter, req *http.Request) {
	total09++
	log.Printf("access at %d time\n", total09)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	var (
		filepath string
		tmp      int
		err      error
	)
	if err = req.ParseForm(); err != nil {
		log.Fatalln(err.Error())
	}
	if start := getQueryValueByName(req, "start"); len(start) > 0 {
		if tmp, err = strconv.Atoi(start); err != nil {
			log.Fatalln(err.Error())
		}
	}
	if tmp < 1000 {
		filepath = "./test1.file"
	} else {
		filepath = "./test2.file"
	}

	log.Println("return 200")
	rw.WriteHeader(http.StatusOK)
	time.Sleep(500 * time.Millisecond)

	b := readBytesFromFile(filepath)
	io.Copy(rw, bytes.NewReader(b))
	log.Println("mock09 => send data done")
}

var total10 int

// Mock10 : mock server disconnect
func Mock10(rw http.ResponseWriter, req *http.Request) {
	total10++
	log.Printf("access at %d time\n", total10)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))

	req.ParseForm()
	wait := 3
	parmWait := getQueryValueByName(req, "wait")
	if parmWait != "" {
		wait, _ = strconv.Atoi(parmWait)
	}
	isSetLen := false
	parmIsSetLen := getQueryValueByName(req, "isSetLen")
	if parmIsSetLen != "" {
		isSetLen, _ = strconv.ParseBool(parmIsSetLen)
	}

	b := readBytesFromFile(testFilePath)
	log.Println("return 200")
	if isSetLen {
		rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
	}
	rw.WriteHeader(http.StatusOK)

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
		fmt.Printf("copied length: %d\n", len)
		return
	}
	log.Println("mock10 => send data done")
}
