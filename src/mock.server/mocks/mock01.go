package mocks

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"

	"qbox.us/rpc"
)

const (
	mockMd5      = "f900b997e6f8a772994876dff023801e"
	testFilePath = "./test.file"
)

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
	time.Sleep(time.Second)

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
	retCode := getQueryValueByName(req, "retCode")
	if retCode == "" {
		retCode = "200"
	}
	code, _ := strconv.Atoi(retCode)

	// check error msg "unexpected status"
	// for 5xx, connection retry
	// for 4xx, no connection retry
	if total04 >= 3 && code != http.StatusOK {
		log.Printf("ret code: %d\n", code)
		rw.WriteHeader(code)
		return
	}

	// md5 check
	// rw.Header().Set("Content-MD5", "314398b1025a0d6a522fbdc1fb456a00")

	// etag check
	// rw.Header().Set("Etag", "f900b997e6f8a772994876dff023801e")
	// if total04%3 == 0 {
	// 	rw.Header().Set("Etag", "f900b997e6f8a772994876dff0238000")
	// }

	// buf := initBytesBySize(4096 * 1024)
	buf := readBytesFromFile(testFilePath)
	// file size check
	// if total04%3 == 0 {
	// 	buf = initBytesBySize(1024 * 1024 * 20)
	// }

	// send data
	rw = rpc.ResponseWriter{rw}
	// rr := rpc.ReadSeeker2RangeReader{bytes.NewReader(buf)}
	rr := createMockReader(buf, 20)
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
	fmt.Printf("wait %d ms\n", mr.wait)
	time.Sleep(time.Duration(mr.wait) * time.Millisecond)
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