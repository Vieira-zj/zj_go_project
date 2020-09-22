package handlers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"src/mock.server/common"
	myutils "src/tools.app/utils"

	"github.com/golib/httprouter"
)

// MockTestHandler01 router for mock test handlers.
func MockTestHandler01(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	if r.Method == "GET" {
		switch id {
		case 1:
			mockTest0101(w, r)
		case 2:
			mockTest0102(w, r)
		case 3:
			mockTest0103(w, r)
		case 4:
			mockTest0104(w, r)
		case 5:
			mockTest0105(w, r)
		case 6:
			mockTest0106(w, r)
		default:
			common.ErrHandler(w, fmt.Errorf("GET for invalid path: %s", r.URL.Path))
		}
	}
}

// mock test, returns bytes body with wait.
// Get /mocktest/one/1
func mockTest0101(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	s := fmt.Sprintf("mockTest01: from Host %s\n", hostname)

	// get query args
	size, err := common.GetIntArgFromQuery(r, "size")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	wait, err := common.GetIntArgFromQuery(r, "wait")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	// mock bytes
	if size > 0 {
		log.Printf("create mock bytes of length %d.\n", size)
		s += common.CreateMockString(size)
	}
	// wait before send header
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

// mock test, returns file content with wait.
// Get /mocktest/one/2
func mockTest0102(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	s := fmt.Sprintf("mockTest02: from Host %s\n", hostname)

	// get query args
	filepath, err := common.GetStringArgFromQuery(r, "file")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	wait, err := common.GetIntArgFromQuery(r, "wait")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	// read file bytes
	if len(filepath) > 0 {
		fContent, err := myutils.ReadFileContent(filepath)
		if err != nil {
			common.ErrHandler(w, err)
			return
		}
		s += fContent
	}

	w.Header().Set(common.TextContentLength, strconv.Itoa(len(s)))
	w.WriteHeader(http.StatusOK)

	// wait before send body
	if wait > 0 {
		fmt.Printf("wait %d seconds before send body.\n", wait)
		time.Sleep(time.Duration(wait) * time.Second)
	}
	if _, err := io.Copy(w, bufio.NewReader(strings.NewReader(s))); err != nil {
		common.ErrHandler(w, err)
	}
}

// mock test, returns custom error code, like 404, 503.
// Get /mocktest/one/3
func mockTest0103(w http.ResponseWriter, r *http.Request) {
	retCode, err := common.GetIntArgFromQuery(r, "code")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	if retCode < http.StatusOK {
		retCode = http.StatusOK
	}

	b := []byte("mockTest0103, mock return error code.")
	w.Header().Set(common.TextContentLength, strconv.Itoa(len(b)))
	w.WriteHeader(retCode)
	log.Println("mock return error code:", retCode)

	if _, err := w.Write(b); err != nil {
		common.ErrHandler(w, err)
	}
}

// mock test, returns httpdns json string.
// Get /mocktest/one/4
func mockTest0104(w http.ResponseWriter, r *http.Request) {
	wait, err := common.GetIntArgFromQuery(r, "wait")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	unit, err := common.GetStringArgFromQuery(r, "unit")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	retIP := `"42.48.232.7", "10.200.20.21"`
	retJSON := fmt.Sprintf(`{"errno":0, "iplist":[%s]}`, retIP)
	// retContent := `{"errno":-1, "iplist":[]}`
	w.Header().Set(common.TextContentLength, strconv.Itoa(len(retJSON)))
	w.Header().Set(common.TextContentType, common.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)

	if wait > 0 {
		if unit == "milli" {
			time.Sleep(time.Duration(wait) * time.Millisecond)
		} else {
			time.Sleep(time.Duration(wait) * time.Second)
		}
	}
	if _, err := io.Copy(w, bufio.NewReader(strings.NewReader(retJSON))); err != nil {
		common.ErrHandler(w, err)
	}
}

// mock test, returns gzip and chunked http response
// Get /mocktest/one/5
func mockTest0105(w http.ResponseWriter, r *http.Request) {
	var (
		b   []byte
		err error
	)
	b, err = ioutil.ReadFile("diskusage")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	// gzip encode
	b, err = myutils.GzipEncode(b)
	if err != nil {
		b = []byte(fmt.Sprintln("gzip encode error:", err))
		w.Header().Set(common.TextContentLength, strconv.Itoa(len(b)))
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set(common.TextContentEncoding, "gzip")
		w.WriteHeader(http.StatusOK)
	}

	if _, err := io.Copy(w, bufio.NewReader(bytes.NewReader(b))); err != nil {
		common.ErrHandler(w, err)
	}
	// Response Header:
	// < Content-Encoding: gzip
	// < Content-Type: application/x-gzip
	// < Transfer-Encoding: chunked
}

// mock test, returns http response with diff mimetype.
// Get /mocktest/one/6
func mockTest0106(w http.ResponseWriter, r *http.Request) {
	mimetypeTable := make(map[string]string)
	mimetypeTable["txt"] = "text/plain"
	mimetypeTable["jpg"] = "image/jpeg"
	mimetypeTable["bin"] = "application/octet-stream"

	// get query args
	mimetype, err := common.GetStringArgFromQuery(r, "type")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	isErrLength, err := common.GetBoolArgFromQuery(r, "errlen")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	// set mimetype
	if len(mimetype) == 0 {
		mimetype = "txt"
	}
	b, err := ioutil.ReadFile(fmt.Sprintf("testfile.%s", mimetype))
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	// set mismatch body length
	contentLen := len(b)
	if isErrLength {
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
