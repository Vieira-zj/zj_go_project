package handlers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"src/mock.server/common"

	"github.com/golib/httprouter"
)

// MockQiNiuHandler router for mock qiniu test handlers.
func MockQiNiuHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	if r.Method == "GET" {
		switch id {
		case 1:
			mockQiNiu01(w, r)
		case 2:
			mockQiNiu02(w, r)
		case 3:
			mockQiNiu03(w, r)
		default:
			common.ErrHandler(w, fmt.Errorf("GET for invalid path: %s", r.URL.Path))
		}
	} else {
		common.ErrHandler(w, fmt.Errorf("Method not support: %s", r.Method))
	}
}

// mock mirror file server handler.
// Get /mockqiniu/1
func mockQiNiu01(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("success")); err != nil {
		common.ErrHandler(w, err)
		return
	}
	w.(http.Flusher).Flush()

	b, err := ioutil.ReadFile("ab_test.out")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	if _, err := io.Copy(w, bufio.NewReader(bytes.NewReader(b))); err != nil {
		common.ErrHandler(w, err)
	}
}

// mock CDN refresh request handler.
type refreshResp struct {
	Code      int    `json:"code"`
	Error     string `json:"error"`
	RequestID string `json:"requestId"`
}

// Get /mockqiniu/2
func mockQiNiu02(w http.ResponseWriter, r *http.Request) {
	retJSON := refreshResp{
		Code:      http.StatusOK,
		Error:     "null",
		RequestID: "cdnrefresh-test-001",
	}

	w.Header().Set(common.TextContentType, common.ContentTypeJSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&retJSON); err != nil {
		common.ErrHandler(w, err)
	}
}

// mock return diff file content by arg "start".
// Get /mockqiniu/3
func mockQiNiu03(w http.ResponseWriter, r *http.Request) {
	start, err := common.GetIntArgFromQuery(r, "start")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	var filepath string
	if start < 1000 {
		filepath = "./testfile1.txt"
	} else {
		filepath = "./testfile2.txt"
	}
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	w.Header().Set(common.TextContentLength, strconv.Itoa(len(b)))
	w.WriteHeader(http.StatusOK)
	time.Sleep(time.Duration(500) * time.Millisecond)
	if _, err := io.Copy(w, bufio.NewReader(bytes.NewReader(b))); err != nil {
		common.ErrHandler(w, err)
	}
}
