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
		case 4:
			mockTest0104(w, r, params)
		case 5:
			mockTest0105(w, r, params)
		case 6:
			mockTest0106(w, r, params)
		default:
			common.ErrHandler(w, fmt.Errorf("GET for invalid path: %s", r.URL.Path))
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
	}
}

// test, mock httpdns server which returns json string => Get /mocktest/one/4
func mockTest0104(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	wait, err := common.GetIntArgFromQuery(r, "wait")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	retIP := `"42.48.232.7", "10.200.20.21"`
	// retContent := `{"errno":-1, "iplist":[]}`
	retContent := fmt.Sprintf(`{"errno":0, "iplist":[%s]}`, retIP)
	w.Header().Set(common.TextContentLength, strconv.Itoa(len(retContent)))
	w.Header().Set(common.TextContentType, common.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)

	if wait > 0 {
		time.Sleep(time.Duration(wait) * time.Second)
	}
	if _, err := io.Copy(w, bufio.NewReader(strings.NewReader(retContent))); err != nil {
		common.ErrHandler(w, err)
	}
}

// test, mock gzip and chunk http response => Get /mocktest/one/5
func mockTest0105(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		b   []byte
		err error
	)
	b, err = myutils.ReadFileContentBuf("ab_test.out")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	// gzip encode
	b, err = myutils.GzipEncode(b)
	if err != nil {
		b = []byte(fmt.Sprintln("error in gzip encode:", err))
		w.Header().Set(common.TextContentLength, strconv.Itoa(len(b)))
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set(common.TextContentEncoding, "gzip")
		w.WriteHeader(http.StatusOK)
	}

	if _, err := io.Copy(w, bufio.NewReader(bytes.NewReader(b))); err != nil {
		common.ErrHandler(w, err)
	}
	// response headers:
	// Content-Type: application/x-gzip
	// Transfer-Encoding: chunked
}

// test, mock http response diff mimetype => Get /mocktest/one/6
func mockTest0106(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	mimetypeTable := make(map[string]string)
	mimetypeTable["txt"] = "text/plain"
	mimetypeTable["jpg"] = "image/jpeg"
	mimetypeTable["bin"] = "application/octet-stream"

	var (
		b   []byte
		err error
	)
	mimetype, err := common.GetStringArgFromQuery(r, "type")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	if len(mimetype) == 0 {
		mimetype = "txt"
	}

	b, err = myutils.ReadFileContentBuf(fmt.Sprintf("testfile.%s", mimetype))
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	isMismatchLen, err := common.GetBoolArgFromQuery(r, "errlen")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	contentLen := len(b)
	if isMismatchLen {
		contentLen += 10
	}

	w.Header().Set(common.TextContentType, mimetypeTable[mimetype])
	w.Header().Set(common.TextContentLength, strconv.Itoa(contentLen))
	w.WriteHeader(http.StatusOK)

	w.(http.Flusher).Flush() // write response headers
	time.Sleep(time.Second)
	if _, err := io.Copy(w, bufio.NewReader(bytes.NewReader(b))); err != nil {
		common.ErrHandler(w, err)
	}
}
