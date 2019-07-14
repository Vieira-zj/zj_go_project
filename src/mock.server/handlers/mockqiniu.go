package handlers

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golib/httprouter"
	"mock.server/common"
	myutils "tools.app/utils"
)

// MockQiNiuHandler router for mock qiniu test.
func MockQiNiuHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	if r.Method == "GET" {
		switch id {
		case 1:
			mockQiNiu01(w, r, params)
		case 2:
			mockQiNiu02(w, r, params)
		case 3:
			mockQiNiu03(w, r, params)
		default:
			common.ErrHandler(w, fmt.Errorf("GET for invalid path: %s", r.URL.Path))
		}
	}
}

// mock mirror file server handler => Get /mockqiniu/1
func mockQiNiu01(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	io.Copy(w, strings.NewReader("success"))
	w.(http.Flusher).Flush()

	wait, err := common.GetIntArgFromQuery(r, "wait")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	if wait > 0 {
		time.Sleep(time.Duration(wait) * time.Second)
		log.Printf("sleep for %d seconds before send file body.\n", wait)
	}

	b, err := myutils.ReadFileContentBuf("ab_test.out")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	io.Copy(w, bufio.NewReader(bytes.NewReader(b)))
}

// cdn refresh request handler => Get /mockqiniu/2
type refreshResp struct {
	Code      int    `json:"code"`
	Error     string `json:"error"`
	RequestID string `json:"requestId"`
}

func mockQiNiu02(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	retJSON := refreshResp{
		Code:      http.StatusOK,
		Error:     "null",
		RequestID: "cdnrefresh-test-001",
	}
	b, err := json.Marshal(retJSON)
	if err != nil {
		log.Println(err)
		b = []byte("json Marshal error.")
	}

	w.Header().Set(common.TextContentLength, strconv.Itoa(len(b)))
	w.WriteHeader(http.StatusOK)
	if _, err := io.Copy(w, bufio.NewReader(bytes.NewReader(b))); err != nil {
		common.ErrHandler(w, err)
	}
}

// mock return diff file content by arg "start" => Get /mockqiniu/3
func mockQiNiu03(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	start, err := common.GetIntArgFromQuery(r, "start")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	var filepath string
	if start < 1000 {
		filepath = "./test1.file"
	} else {
		filepath = "./test2.file"
	}
	b, err := myutils.ReadFileContentBuf(filepath)
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	w.Header().Set(common.TextContentLength, strconv.Itoa(len(b)))
	w.WriteHeader(http.StatusOK)

	time.Sleep(500 * time.Millisecond)
	if _, err := io.Copy(w, bufio.NewReader(bytes.NewReader(b))); err != nil {
		common.ErrHandler(w, err)
	}
}
