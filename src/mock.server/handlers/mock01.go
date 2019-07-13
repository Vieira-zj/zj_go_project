package handlers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golib/httprouter"
	"mock.server/common"
	myutils "tools.app/utils"
)

// MockTestHandler01 router for mock test.
func MockTestHandler01(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	if r.Method == "GET" {
		switch id {
		case 1:
			mockTest0101(w, r, params)
		case 2:
			mockTest0102(w, r, params)
		case 3:
			mockTest0103(w, r, params)
		default:
			common.ErrHandler(w, fmt.Errorf("GET for invalid path: %s", r.URL.Path))
			return
		}
	}
}

// MockNotFound mocks server not found page.
func MockNotFound(w http.ResponseWriter, r *http.Request) {
	retContent := fmt.Sprintf("Default Not Found Page.\nPage not found for path: %s\n", r.RequestURI)
	w.Header().Set(common.TextContentLength, strconv.Itoa(len([]byte(retContent))))
	w.WriteHeader(http.StatusNotFound)
	log.Printf("page not found: %s.\n", r.URL.Path)

	io.Copy(w, strings.NewReader(retContent))
}

// test, mock return bytes body with wait => Get /mocktest/one/1
func mockTest0101(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	hostname, err := os.Hostname()
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	s := fmt.Sprintf("mockTest01, from Host: %s\n", hostname)

	// mock bytes
	size, err := common.GetIntArgFromQuery(r, "size")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	if size > 0 {
		log.Printf("create mock bytes of length %d.\n", size)
		s += common.CreateMockBytes(size)
	}

	// wait before send header
	wait, err := common.GetIntArgFromQuery(r, "wait")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	if wait > 0 {
		log.Printf("wait %d seconds before send header.\n", wait)
		time.Sleep(time.Duration(wait) * time.Second)
	}

	w.Header().Set(common.TextContentLength, strconv.Itoa(len(s)))
	// mockMD5 := "f900b997e6f8a772994876dff023801e"
	// w.Header().Set("Content-Md5", mockMD5)
	w.WriteHeader(http.StatusOK)

	if _, err := io.Copy(w, bufio.NewReader(strings.NewReader(s))); err != nil {
		common.ErrHandler(w, err)
		return
	}
}

// test, mock return file content with wait => Get /mocktest/one/2
func mockTest0102(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var b []byte
	hostname, err := os.Hostname()
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	b = []byte(fmt.Sprintf("mockTest02, from Host: %s\n", hostname))

	// read file bytes
	filepath, err := common.GetStringArgFromQuery(r, "file")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	if len(filepath) > 0 {
		b, err = myutils.ReadFileContentBuf(filepath)
		if err != nil {
			common.ErrHandler(w, err)
			return
		}
	}

	w.Header().Set(common.TextContentLength, strconv.Itoa(len(b)))
	w.WriteHeader(http.StatusOK)

	// wait before send body
	wait, err := common.GetIntArgFromQuery(r, "wait")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	if wait > 0 {
		fmt.Printf("wait %d seconds before send body.\n", wait)
		time.Sleep(time.Duration(wait) * time.Second)
	}

	if _, err := io.Copy(w, bufio.NewReader(bytes.NewReader(b))); err != nil {
		common.ErrHandler(w, err)
		return
	}
}

// test, mock return custom error code, ex 404, 503 => Get /mocktest/one/3
func mockTest0103(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		retCode = http.StatusOK
		err     error
	)

	if retCode, err = common.GetIntArgFromQuery(r, "code"); err != nil {
		common.ErrHandler(w, err)
		return
	}

	b := []byte("mockTest0103, mock return error code.")
	w.Header().Set(common.TextContentLength, strconv.Itoa(len(b)))
	w.WriteHeader(retCode)
	log.Println("mock return error code:", retCode)

	if _, err := io.Copy(w, bufio.NewReader(bytes.NewReader(b))); err != nil {
		common.ErrHandler(w, err)
		return
	}
}

// var total05 int

// // Mock05 : mock stream data by wait and kb
// func Mock05(rw http.ResponseWriter, req *http.Request) {
// 	total05++
// 	log.Printf("\n===> Mock05, access at %d time\n", total05)
// 	reqHeader, _ := httputil.DumpRequest(req, true)
// 	fmt.Println(strings.Trim(string(reqHeader), "\n"))

// 	req.ParseForm()
// 	wait := GetNumberInReqForm(req, "wait")
// 	if wait == 0 {
// 		wait = 3
// 	}
// 	kb := GetNumberInReqForm(req, "kb")
// 	if kb == 0 {
// 		kb = 1
// 	}

// 	b := InitBytesBySize(1024 * kb)
// 	rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
// 	rw.WriteHeader(http.StatusOK)
// 	log.Println("return code => 200")

// 	io.Copy(rw, &mockReader{wait: wait, r: bytes.NewReader(b)})
// 	log.Print("===> Mock05, send data done\n\n")
// }

// var total06 int

// // Mock06 : mock httpdns server
// func Mock06(rw http.ResponseWriter, req *http.Request) {
// 	total06++
// 	log.Printf("===> Mock06, access dns server at %d time\n", total06)
// 	reqHeader, _ := httputil.DumpRequest(req, true)
// 	fmt.Println(strings.Trim(string(reqHeader), "\n"))

// 	// ret := `{"errno":-1, "iplist":[]}`
// 	retIP := `"10.200.20.21"`
// 	// retIP := `"42.48.232.7", "10.200.20.21"`
// 	ret := fmt.Sprintf(`{"errno":0, "iplist":[%s]}`, retIP)

// 	retCode := 200
// 	if total06%5 == 0 {
// 		retCode = 500
// 	}
// 	rw.Header().Set("Content-Length", strconv.Itoa(len(ret)))
// 	rw.Header().Set("Content-Type", "application/json")
// 	rw.WriteHeader(retCode)
// 	log.Printf("return code => %d\n", retCode)

// 	if total06%4 == 0 {
// 		wait := 3
// 		time.Sleep(time.Duration(wait) * time.Second)
// 		log.Printf("sleep %d seconds\n", wait)
// 	}

// 	io.Copy(rw, strings.NewReader(ret))
// 	log.Print("===> Mock06, send data done\n\n")
// }

// var total07 int

// // Mock07 : mock mirror file server
// func Mock07(rw http.ResponseWriter, req *http.Request) {
// 	total07++
// 	log.Printf("\n===> Mock07, access mirror at %d time\n", total07)
// 	reqHeader, _ := httputil.DumpRequest(req, true)
// 	fmt.Println(strings.Trim(string(reqHeader), "\n"))

// 	rw.WriteHeader(http.StatusOK)
// 	log.Println("return code => 200")
// 	io.Copy(rw, strings.NewReader("success"))
// 	rw.(http.Flusher).Flush()

// 	if total07%2 == 0 {
// 		wait := 3
// 		time.Sleep(time.Duration(wait) * time.Second)
// 		log.Printf("sleep %d seconds\n", wait)
// 	}

// 	// io.Copy(rw, Strings.NewReader("** test content"))
// 	b := ReadBytesFromFile(testFilePath)
// 	io.Copy(rw, bytes.NewReader(b))
// 	log.Print("===> Mock07, data returned\n\n")
// }

// var total08 = 0

// type refreshRet struct {
// 	Code      int    `json:"code"`
// 	Error     string `json:"error"`
// 	RequestID string `json:"requestId"`
// }

// // Mock08 : handle cdn refresh post request, and return
// func Mock08(rw http.ResponseWriter, req *http.Request) {
// 	total08++
// 	log.Printf("\n===> Mock08, access mirror at %d time\n", total08)
// 	reqData, _ := httputil.DumpRequest(req, true)
// 	fmt.Println(strings.Trim(string(reqData), "\n"))

// 	// result, _ := ioutil.ReadAll(req.Body)
// 	// defer req.Body.Close()
// 	// fmt.Printf("request body: %s\n", string(result))

// 	JSONObj := refreshRet{
// 		Code:      http.StatusOK,
// 		Error:     "null",
// 		RequestID: "cdnrefresh-test-01",
// 	}
// 	b, err := json.Marshal(JSONObj)
// 	if err != nil {
// 		log.Fatalln(err.Error())
// 		b = []byte("json encode error")
// 	}

// 	rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
// 	rw.WriteHeader(http.StatusOK)
// 	log.Println("return code => 200")

// 	io.Copy(rw, bytes.NewReader(b))
// 	log.Print("===> Mock08, data returned\n\n")
// }

// var total09 int

// // Mock09 : mock download file by arg "start"
// func Mock09(rw http.ResponseWriter, req *http.Request) {
// 	total09++
// 	log.Printf("\n===> Mock09, access at %d time\n", total09)
// 	reqHeader, _ := httputil.DumpRequest(req, true)
// 	fmt.Println(strings.Trim(string(reqHeader), "\n"))

// 	if err := req.ParseForm(); err != nil {
// 		log.Fatalln(err.Error())
// 	}

// 	var filepath string
// 	start := GetNumberInReqForm(req, "start")
// 	if start < 1000 {
// 		filepath = "./test1.file"
// 	} else {
// 		filepath = "./test2.file"
// 	}
// 	b := ReadBytesFromFile(filepath)

// 	rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
// 	rw.WriteHeader(http.StatusOK)
// 	log.Println("return code => 200")

// 	time.Sleep(500 * time.Millisecond)
// 	io.Copy(rw, bytes.NewReader(b))
// 	log.Print("===> mock09, send data done\n\n")
// }

// var total10 int

// // Mock10 : mock server disconnect
// func Mock10(rw http.ResponseWriter, req *http.Request) {
// 	total10++
// 	log.Printf("\n===> Mock10, access at %d time\n", total10)
// 	reqHeader, _ := httputil.DumpRequest(req, true)
// 	fmt.Println(strings.Trim(string(reqHeader), "\n"))

// 	req.ParseForm()
// 	b := ReadBytesFromFile(testFilePath)
// 	// set resp header Content-Length
// 	if GetBoolInReqForm(req, "isSetLen") {
// 		rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
// 	}

// 	rw.WriteHeader(http.StatusOK)
// 	log.Println("return code => 200")

// 	wait := GetNumberInReqForm(req, "wait")
// 	if wait == 0 {
// 		wait = 3
// 	}
// 	go func() {
// 		time.Sleep(time.Duration(wait) * time.Second)
// 		if jacker, ok := httpv1.GetHijacker(rw); ok {
// 			conn, _, err := jacker.Hijack()
// 			if err != nil {
// 				fmt.Printf("hijack err: %v\n", err)
// 			} else {
// 				log.Println("response can hijack, connection closed")
// 				conn.Close()
// 			}
// 		} else {
// 			fmt.Printf("http.ResponseWriter not http.Hijacker")
// 		}
// 	}()

// 	len, err := io.Copy(rw, bytes.NewReader(b))
// 	if err != nil {
// 		log.Println("copy resp writer error:", err.Error())
// 		return
// 	}
// 	fmt.Printf("copied length: %d\n", len)
// 	log.Print("===> mock10, send data done\n\n")
// }

// var total11 int

// // Mock11 : mock gzip and chunk
// func Mock11(rw http.ResponseWriter, req *http.Request) {
// 	total11++
// 	log.Printf("\n===> Mock11, access at %d time\n", total11)
// 	reqHeader, _ := httputil.DumpRequest(req, true)
// 	fmt.Println(strings.Trim(string(reqHeader), "\n"))

// 	srcb := ReadBytesFromFile(testFilePath)
// 	retb, err := zjutils.GzipEncode(srcb)
// 	if err != nil {
// 		retErr := fmt.Sprintln("error in gzip encode:", err.Error())
// 		rw.Header().Set("Content-Length", strconv.Itoa(len(retErr)))
// 		rw.WriteHeader(599)
// 		log.Println("return code => 599")
// 		io.Copy(rw, strings.NewReader(retErr))
// 		return
// 	}

// 	rw.Header().Set("Content-Encoding", "gzip")
// 	rw.WriteHeader(http.StatusOK)
// 	log.Println("return code => 200")

// 	io.Copy(rw, bytes.NewReader(retb))
// 	log.Print("===> mock11, send data done\n\n")
// 	// resp headers:
// 	// Content-Type: application/x-gzip
// 	// Transfer-Encoding: chunked
// }

// var total12 int

// // Mock12 : mimetype
// func Mock12(rw http.ResponseWriter, req *http.Request) {
// 	total12++
// 	log.Printf("\n===> Mock12, access at %d time\n", total12)
// 	reqHeader, _ := httputil.DumpRequest(req, true)
// 	fmt.Println(strings.Trim(string(reqHeader), "\n"))

// 	mimetypeTable := make(map[string]string)
// 	mimetypeTable["txt"] = "text/plain"
// 	mimetypeTable["jpg"] = "image/jpeg"
// 	mimetypeTable["bin"] = "application/octet-stream"

// 	req.ParseForm()
// 	mimetype := GetStringInReqForm(req, "type")
// 	if len(mimetype) == 0 {
// 		mimetype = "txt"
// 	}
// 	rw.Header().Set("Content-Type", mimetypeTable[mimetype])

// 	var b []byte
// 	isEmpty := GetBoolInReqForm(req, "isempty")
// 	if !isEmpty {
// 		path := testFilePath + "." + mimetype
// 		b = ReadBytesFromFile(path)
// 	}

// 	isMismatchContentLen := GetBoolInReqForm(req, "isfaillen")
// 	contentLen := len(b)
// 	if isMismatchContentLen {
// 		contentLen += 10
// 	}

// 	rw.Header().Set("Content-Length", strconv.Itoa(contentLen))
// 	rw.WriteHeader(http.StatusOK)
// 	log.Println("return code => 200")

// 	if len(b) > 0 {
// 		rw.(http.Flusher).Flush() // write response headers
// 		time.Sleep(2 * time.Second)
// 		io.Copy(rw, bytes.NewReader(b))
// 	}
// 	log.Print("===> mock12, send data done\n\n")
// }
