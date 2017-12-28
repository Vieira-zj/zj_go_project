package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var total int

func init() {
	fmt.Println("run init")
}

// test qiniu proxy, mock user src server
func index(rw http.ResponseWriter, req *http.Request) {
	log.Println("200")
	rw.Header().Set("Content-Md5", "db742740b369a1c8be6115268c3d358d")
	rw.WriteHeader(200)
	time.Sleep(time.Second * 3)
	log.Println("index body")
	io.Copy(rw, bytes.NewReader([]byte("stream data mock")))
}

func index2(rw http.ResponseWriter, req *http.Request) {
	log.Println("200")
	rw.Header().Set("Content-Md5", "db742740b369a1c8be6115268c3d358d")
	rw.Header().Set("Content-Length", "1000000")
	rw.WriteHeader(200)
	for i := 0; i < 100000; i++ {
		time.Sleep(time.Duration(500) * time.Millisecond)
		log.Println("index2 body")
		_, err := io.Copy(rw, bytes.NewReader([]byte("stream data mock")))
		rw.(http.Flusher).Flush()
		if err != nil {
			log.Printf("error: %v\n", err)
			return
		}
	}
}

func index3(rw http.ResponseWriter, req *http.Request) {
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

	total++
	log.Printf("access at %d time\n", total)
	log.Printf("%d\n", reqCode)
	rw.WriteHeader(reqCode)
	log.Println("index3 body")
	io.Copy(rw, bytes.NewReader([]byte("return code mock")))
}

func getContentMd5(content string, md5Type string) string {
	md5hash := md5.New()
	md5hash.Write([]byte(content))

	if md5Type == "hex" {
		return hex.EncodeToString(md5hash.Sum(nil))
	} else if md5Type == "std64" {
		return base64.StdEncoding.EncodeToString(md5hash.Sum(nil))
	}
	//url64
	return base64.URLEncoding.EncodeToString(md5hash.Sum(nil))
}

// build cmd: $ GOOS=linux GOARCH=amd64 go build
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index2/", index2)
	http.HandleFunc("/index3/", index3)
	log.Fatal(http.ListenAndServe(":17890", nil))
}
