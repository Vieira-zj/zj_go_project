package handlers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
)

var total4775 int

// Mock4775 : mock data stream and file download
func Mock4775(rw http.ResponseWriter, req *http.Request) {
	total4775++
	log.Printf("\n===> Mock4775, access at %d time\n", total4775)
	reqHeader, _ := httputil.DumpRequest(req, true)
	fmt.Println(strings.Trim(string(reqHeader), "\n"))
	req.ParseForm()

	var b []byte
	retCode := 200
	if total4775%2 == 0 {
		b = []byte("mock error content")
		retCode = 403
	} else {
		b = ReadBytesFromFile(testFilePath)
	}

	rw.Header().Set("Content-Length", strconv.Itoa(len(b)))
	rw.WriteHeader(retCode)
	log.Println("return code =>", retCode)

	io.Copy(rw, bytes.NewReader(b))
	log.Print("===> Mock4775, send data done\n\n")
}
