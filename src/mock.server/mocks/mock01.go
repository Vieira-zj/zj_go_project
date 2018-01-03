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

const (
	mockMd5     = "f900b997e6f8a772994876dff023801e"
	filePathMp3 = "./test.mp3"
)

// Mock01 : mock short bytes stream / file donwload, diff md5
func Mock01(rw http.ResponseWriter, req *http.Request) {
	log.Println("200")
	rw.Header().Set("Content-Md5", mockMd5) // mock md5
	rw.WriteHeader(200)

	// b := []byte("stream data mock")
	b := readBytesFromFile(filePathMp3)

	log.Println("mock body")
	time.Sleep(time.Second * 3)
	io.Copy(rw, bytes.NewReader(b))
	log.Println("send data done")
}

// Mock02 : mock long bytes stream, diff md5
func Mock02(rw http.ResponseWriter, req *http.Request) {
	log.Println("200")
	rw.Header().Set("Content-Md5", mockMd5) // mock md5
	rw.Header().Set("Content-Length", "1000000")
	rw.WriteHeader(200)

	for i := 0; i < 100000; i++ {
		log.Println("mock body")
		time.Sleep(time.Duration(500) * time.Millisecond)
		_, err := io.Copy(rw, bytes.NewReader([]byte("stream data mock")))
		rw.(http.Flusher).Flush()
		if err != nil {
			log.Printf("error: %v\n", err)
			return
		}
	}
}

var total03 int

// Mock03 : mock diff ret code from query "retCode", ex 404, 503
func Mock03(rw http.ResponseWriter, req *http.Request) {
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
	io.Copy(rw, bytes.NewReader([]byte("stram data mock")))
}

var total04 int

const waitForReader = 20

// Mock04 : mock server for file download by range 4M
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

	// connection retry test
	// if total04%3 == 0 {
	// 	rw.WriteHeader(500)
	// 	return
	// }

	// etag check
	// rw.Header().Set("Etag", "f900b997e6f8a772994876dff023801e")
	// if total04%3 == 0 {
	// 	rw.Header().Set("Etag", "f900b997e6f8a772994876dff0238000")
	// }

	// md5 check
	// rw.Header().Set("Content-MD5", "314398b1025a0d6a522fbdc1fb456a00")

	// buf := initBytesBySize(4096 * 1024)
	buf := readBytesFromFile("./test.mp3")
	// file size check
	// if total04%3 == 0 {
	// 	buf = initBytesBySize(1024 * 1024 * 20)
	// }

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

func (mr *mockReader) Read(b []byte) (int, error) {
	time.Sleep(time.Duration(waitForReader) * time.Millisecond) // custom wait
	len, err := mr.r.Read(b)
	return len, err
}

func (mr *mockReader) Seek(offset int64, whence int) (int64, error) {
	return mr.r.Seek(offset, whence)
}
