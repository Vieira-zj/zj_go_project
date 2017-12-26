package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func init() {
	fmt.Println("run init")
}

// test qiniu proxy, mock user src server
func index(rw http.ResponseWriter, req *http.Request) {
	log.Println("200")
	rw.Header().Set("Content-Md5", "db742740b369a1c8be6115268c3d358d")
	rw.WriteHeader(200)
	time.Sleep(time.Second * 3)
	log.Println("body")
	io.Copy(rw, bytes.NewReader([]byte("aaaaaaaa")))
}

func index2(rw http.ResponseWriter, req *http.Request) {
	log.Println("200")
	rw.Header().Set("Content-Md5", "db742740b369a1c8be6115268c3d358d")
	rw.Header().Set("Content-Length", "1000000")
	rw.WriteHeader(200)
	for i := 0; i < 100000; i++ {
		time.Sleep(time.Duration(500) * time.Millisecond)
		log.Println("body")
		_, err := io.Copy(rw, bytes.NewReader([]byte("aaaaaaaa")))
		rw.(http.Flusher).Flush()
		if err != nil {
			log.Printf("error: %v\n", err)
			return
		}
	}
}

// build cmd: $ GOOS=linux GOARCH=amd64 go build
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index2/", index2)
	log.Fatal(http.ListenAndServe(":17890", nil))
}
