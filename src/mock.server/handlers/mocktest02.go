package handlers

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"src/mock.server/common"

	httputil "src/mock.server/vendor/qbox.us/httputil.v1"

	"github.com/golib/httprouter"
)

// MockTestHandler02 router for mock test handlers.
func MockTestHandler02(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	if r.Method == "GET" {
		switch id {
		case 1:
			mockTest0201(w, r)
		case 2:
			mockTest0202(w, r)
		case 3:
			mockTest0203(w, r)
		case 4:
			mockTest0204(w, r)
		case 5:
			mockTest0205(w, r)
		default:
			common.ErrHandler(w, fmt.Errorf("GET for invalid path: %s", r.URL.Path))
		}
	}
}

// mock test, returns 403 Forbidden, or file content.
// GET /mocktest/two/1
func mockTest0201(w http.ResponseWriter, r *http.Request) {
	isErr, err := common.GetBoolArgFromQuery(r, "iserr")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	var (
		retCode  int
		retBytes []byte
	)
	if isErr {
		// mock error
		retCode = http.StatusForbidden
		log.Println("mock return error code:", retCode)
		retBytes = []byte("mock error content.")
	} else {
		// read file
		retCode = http.StatusOK
		filePath := "ab_test.out"
		retBytes, err = ioutil.ReadFile(filePath)
		if err != nil {
			common.ErrHandler(w, err)
			return
		}
	}

	w.Header().Set(common.TextContentLength, strconv.Itoa(len(retBytes)))
	w.WriteHeader(retCode)
	if _, err := w.Write(retBytes); err != nil {
		common.ErrHandler(w, err)
	}
}

// mock test, returns chunked of bytes by flush.
// GET /mocktest/two/2
func mockTest0202(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	// mock block of bytes
	const size = 128
	for i := 0; i < 10; i++ {
		log.Println("mock block of bytes:", size)
		_, err := io.Copy(w, bufio.NewReader(strings.NewReader(common.CreateMockString(size))))
		if err != nil {
			log.Println(err)
			break
		}
		w.(http.Flusher).Flush()
		time.Sleep(time.Duration(200) * time.Millisecond)
	}

	b := []byte("\nmockTest0202, mock bytes body.")
	if _, err := w.Write(b); err != nil {
		common.ErrHandler(w, err)
	}
	// Response Headers:
	// < Content-Type: text/plain; charset=utf-8
	// < Transfer-Encoding: chunked
}

// mock test, returns bytes by range with wait.
// GET /mocktest/two/3
func mockTest0203(w http.ResponseWriter, r *http.Request) {
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
	buf := []byte(common.CreateMockString(size))

	// send data by range, and return: 206 Partial Content
	wait := 500 // millisecond, wait between each range response
	// mr := rpc.ReadSeeker2RangeReader{bytes.NewReader(buf)}
	mr := rpc.ReadSeeker2RangeReader{&mockReader{wait: wait, r: bytes.NewReader(buf)}}
	w = rpc.ResponseWriter{w}
	if err := w.(rpc.ResponseWriter).ReplyRange(mr, int64(len(buf)), &rpc.Metas{}, r); err != nil {
		common.ErrHandler(w, err)
	}
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

// mock test, returns kb data with wait in each read.
// GET /mocktest/two/4
func mockTest0204(w http.ResponseWriter, r *http.Request) {
	wait, err := common.GetIntArgFromQuery(r, "wait")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}
	kb, err := common.GetIntArgFromQuery(r, "kb")
	if err != nil {
		common.ErrHandler(w, err)
		return
	}

	limit := 100 // millisecond
	if wait <= limit {
		wait = limit
	}
	limit = 3 // kb
	if kb <= limit {
		kb = limit
	}

	s := []byte(common.CreateMockString(1024 * kb))
	w.Header().Set(common.TextContentLength, strconv.Itoa(len(s)))
	w.WriteHeader(http.StatusOK)

	// custom mock reader, wait in each read()
	reader := &mockReader{wait: wait, r: bytes.NewReader([]byte(s))}
	if _, err := io.Copy(w, reader); err != nil {
		common.ErrHandler(w, err)
	}
}

// mock test, server side close connection.
// Get /mocktest/two/5
func mockTest0205(w http.ResponseWriter, r *http.Request) {
	wait, err := common.GetIntArgFromQuery(r, "wait")
	if err != nil {
		common.ErrHandler(w, err)
	}
	if wait <= 0 {
		wait = 1
	}

	s := common.CreateMockString(1024)
	w.Header().Set(common.TextContentLength, strconv.Itoa(len(s)))
	w.WriteHeader(http.StatusOK)

	// wait and close connection
	go func() {
		defer func() {
			if p := recover(); p != nil {
				common.ErrHandler(w, p.(error))
				return
			}
		}()

		time.Sleep(time.Duration(wait) * time.Second)
		if jacker, ok := httputil.GetHijacker(w); ok {
			conn, _, err := jacker.Hijack()
			if err != nil {
				log.Println("hijack error:", err)
			} else {
				log.Println("response can hijack, connection closed.")
				conn.Close()
			}
		} else {
			log.Println("http.ResponseWriter not http.Hijacker.")
		}
	}()

	time.Sleep(time.Duration(wait) * time.Second)
	if _, err := io.Copy(w, bufio.NewReader(strings.NewReader(s))); err != nil {
		common.ErrHandler(w, err)
	}
	// Http Response:
	// > transfer closed with 1368 bytes remaining to read
	// > stopped the pause stream!
	// > Closing connection 0
}
