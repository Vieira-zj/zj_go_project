package mocks

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
	"time"

	"qbox.us/rpc"
)

// Mock01 : mock user_src_server, qiniuproxy pull file from user_src_server
func Mock01(rw http.ResponseWriter, req *http.Request) {
	log.Println("200")
	rw.Header().Set("Content-Md5", "f900b997e6f8a772994876dff023801e") // mock md5
	rw.WriteHeader(200)

	// b := []byte("stream data mock")
	b := readBytesFromFile("./test.mp3")

	log.Println("mock body")
	time.Sleep(time.Second * 3)
	io.Copy(rw, bytes.NewReader(b))
	log.Println("send data done")
}

// Mock02 : user_src_server, qiniuproxy pull file from user_src_server
func Mock02(rw http.ResponseWriter, req *http.Request) {
	log.Println("200")
	rw.Header().Set("Content-Md5", "db742740b369a1c8be6115268c3d358d") // mock md5
	rw.Header().Set("Content-Length", "1000000")
	rw.WriteHeader(200)

	for i := 0; i < 100000; i++ {
		time.Sleep(time.Duration(500) * time.Millisecond)
		log.Println("mock body")
		_, err := io.Copy(rw, bytes.NewReader([]byte("stream data mock")))
		rw.(http.Flusher).Flush()
		if err != nil {
			log.Printf("error: %v\n", err)
			return
		}
	}
}

var total03 int

// Mock03 : user_src_server, qiniuproxy pull file from user_src_server
func Mock03(rw http.ResponseWriter, req *http.Request) {
	// mock return status code from query "retCode", ex 404, 503
	const keyRetCode = "retCode"
	reqCode := 200
	req.ParseForm()
	if len(req.Form) > 0 {
		for k, v := range req.Form {
			if k == keyRetCode {
				reqCode, _ = strconv.Atoi(v[0])
				break
			}
		}
	}

	total03++
	log.Printf("access at %d time\n", total03)
	log.Printf("%d\n", reqCode)
	rw.WriteHeader(reqCode)
	log.Println("mock body")
	io.Copy(rw, bytes.NewReader([]byte("status code mock")))
}

var total04 int

const waitForReader = 100

// Mock04 : user_src_server, qiniuproxy pull file by range from user_src_server
func Mock04(rw http.ResponseWriter, req *http.Request) {
	// download: curl -o ./test.mp3 http://127.0.0.1:17890/index4/
	total04++
	log.Printf("access at %d time\n", total04)
	reqHeader, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Printf("error: %v\n", err)
		return
	}
	log.Println(string(reqHeader))

	// buf := initBytesBySize(4096 * 1024)
	buf := readBytesFromFile("./test.mp3")

	// send data
	rw = rpc.ResponseWriter{rw}
	// rr := rpc.ReadSeeker2RangeReader{bytes.NewReader(buf)}
	rr := createMockReader(buf)
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
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
	return buf
}

func createMockReader(buf []byte) rpc.ReadSeeker2RangeReader {
	return rpc.ReadSeeker2RangeReader{&mockReader{r: bytes.NewReader(buf)}}
}

type mockReader struct {
	r *bytes.Reader
}

// wait between each block read
func (mr *mockReader) Read(b []byte) (int, error) {
	time.Sleep(time.Duration(waitForReader) * time.Millisecond)
	return mr.r.Read(b)
}

func (mr *mockReader) Seek(offset int64, whence int) (int64, error) {
	return mr.r.Seek(offset, whence)
}
