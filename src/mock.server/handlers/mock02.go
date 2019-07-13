package handlers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golib/httprouter"
	"mock.server/common"
	"qbox.us/rpc"
	myutils "tools.app/utils"
)

// MockTestHandler02 router for mock test.
func MockTestHandler02(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	if r.Method == "GET" {
		switch id {
		case 1:
			mockTest0201(w, r, params)
		case 2:
			mockTest0202(w, r, params)
		case 3:
			mockTest0203(w, r, params)
		default:
			common.ErrHandler(w, fmt.Errorf("GET for invalid path: %s", r.URL.Path))
			return
		}
	}
}

// test, mock 403 Forbidden, or return file content => GET /mocktest/two/1
func mockTest0201(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var (
		err   error
		isErr = false
	)
	values := r.URL.Query()
	if val, ok := values["iserr"]; !ok {
		log.Println("isErr as default false.")
	} else {
		isErr, err = strconv.ParseBool(val[0])
		if err != nil {
			common.ErrHandler(w, err)
			return
		}
	}

	var (
		retCode  int
		retBytes []byte
		filePath = "ab_test.out"
	)
	if isErr {
		retCode = http.StatusForbidden
		retBytes = []byte("mock error content.")
		log.Println("mock return error code:", retCode)
	} else {
		retCode = http.StatusOK
		retBytes, err = myutils.ReadFileContentBuf(filePath)
		if err != nil {
			common.ErrHandler(w, err)
			return
		}
	}

	w.Header().Set(common.TextContentLength, strconv.Itoa(len(retBytes)))
	w.WriteHeader(retCode)

	if _, err := io.Copy(w, bufio.NewReader(bytes.NewReader(retBytes))); err != nil {
		common.ErrHandler(w, err)
		return
	}
}

// test, mock return bytes data by flush => GET /mocktest/two/2
func mockTest0202(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)

	const size = 128
	for i := 0; i < 10; i++ {
		log.Println("mock block of bytes body.")
		_, err := io.Copy(w, bufio.NewReader(strings.NewReader(common.CreateMockBytes(size))))
		if err != nil {
			log.Println(err)
			break
		}
		w.(http.Flusher).Flush()
		time.Sleep(time.Duration(200) * time.Millisecond)
	}

	b := []byte("\nmockTest0202, mock bytes body.")
	if _, err := io.Copy(w, bufio.NewReader(bytes.NewReader(b))); err != nil {
		common.ErrHandler(w, err)
		return
	}
}

// test, mock return bytes by range with wait => GET /mocktest/two/3
func mockTest0203(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// data block is set in request header by: Range:bytes=0-4095
	// for qiniuproxy, default range is 4M

	retCode, err := common.GetIntArgFromQuery(r, "code")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	// for 4xx, no connection retry; 5xx, connection retry
	if retCode >= 400 {
		retText := "mockTest0203, mock return error code."
		w.Header().Set(common.TextContentLength, strconv.Itoa(len(retText)))
		w.WriteHeader(retCode)
		log.Println("mock return error code:", retCode)

		if _, err := io.Copy(w, bufio.NewReader(strings.NewReader(retText))); err != nil {
			common.ErrHandler(w, err)
		}
		return
	}

	size := 1024 * 1024 * 10
	fmt.Println("mock bytes body:", size)
	buf := []byte(common.CreateMockBytes(size))

	// send data by range, and return code: 206 Partial Content
	sleepEachRead := 500 // millisecond, sleep between each range request
	mr := createMockReader(buf, sleepEachRead)
	w = rpc.ResponseWriter{w}
	w.(rpc.ResponseWriter).ReplyRange(mr, int64(len(buf)), &rpc.Metas{}, r)
}

func createMockReader(buf []byte, waitForReader int) rpc.ReadSeeker2RangeReader {
	// rpc.ReadSeeker2RangeReader{bytes.NewReader(buf)}
	return rpc.ReadSeeker2RangeReader{&mockReader{wait: waitForReader, r: bytes.NewReader(buf)}}
}

type mockReader struct {
	wait int
	r    *bytes.Reader
}

func (mr *mockReader) Read(b []byte) (int, error) {
	if mr.wait > 0 {
		time.Sleep(time.Duration(mr.wait) * time.Millisecond)
	}
	return mr.r.Read(b)
}

func (mr *mockReader) Seek(offset int64, whence int) (int64, error) {
	return mr.r.Seek(offset, whence)
}
