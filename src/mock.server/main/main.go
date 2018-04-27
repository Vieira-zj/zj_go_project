package main

import (
	"fmt"
	"log"
	"net/http"

	"mock.server/mocks"
)

func init() {
	fmt.Println("mock server started.")
}

// build cmd: /main$ GOOS=linux GOARCH=amd64 go build
// $ scp main qboxserver@10.200.20.21:~/zhengjin/main
// http://10.200.20.21:17890/index1/
func main() {
	http.HandleFunc("/", mocks.MockDefault)

	http.HandleFunc("/index1/", mocks.Mock01)
	http.HandleFunc("/index2/", mocks.Mock02)
	http.HandleFunc("/index3/", mocks.Mock03)
	// curl -v "http://10.200.20.21:17890/index4/?isFile=true" -H "Range":"bytes=16-32"
	http.HandleFunc("/index4/", mocks.Mock04)
	http.HandleFunc("/index5/", mocks.Mock05)

	http.HandleFunc("/httpdns", mocks.Mock06)
	http.HandleFunc("/dirpath/filepath", mocks.Mock07)
	http.HandleFunc("/post/cdnrefresh", mocks.Mock08)

	http.HandleFunc("/videos/file.ts", mocks.Mock09)
	http.HandleFunc("/videos/other/723c0f1d75a08397/file.m2ts", mocks.Mock09)
	http.HandleFunc("/videos/file.flv", mocks.Mock09)
	http.HandleFunc("/videos/other/723c0f1d75a08397/file.mp4", mocks.Mock09)

	http.HandleFunc("/videos/ots/file1.ts", mocks.Mock09)
	http.HandleFunc("/qpdxv/ots/file2.ts", mocks.Mock09)
	http.HandleFunc("/videos/vts/file3.m2ts", mocks.Mock09)
	http.HandleFunc("/qpdxv/vts/file4.m2ts", mocks.Mock09)

	http.HandleFunc("/mock1", mocks.Mock21)
	http.HandleFunc("/mock2", mocks.Mock22)
	http.HandleFunc("/mock3", mocks.Mock23)

	const port = 17890
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
