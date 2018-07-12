package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"mock.server/mocks"
)

func init() {
	mocks.InitConfigs()
}

// build cmd: /main$ GOOS=linux GOARCH=amd64 go build
// $ scp main qboxserver@10.200.20.21:~/zhengjin/main
func main() {
	port := flag.Int("p", 17891, "mock server listen port")
	flag.Parse()

	// http://10.200.20.21:17891/test
	http.HandleFunc("/", mocks.MockDefault)

	// curl -v "http://10.200.20.21:17891/index?isFile=false&wait=1"
	http.HandleFunc("/index", mocks.Mock01)
	http.HandleFunc("/index2", mocks.Mock02)
	// curl -v "http://10.200.20.21:17891/error?retCode=206"
	http.HandleFunc("/error", mocks.Mock03)
	// ret 200 => curl -v "http://10.200.20.21:17891/common?isFile=true&md5=true&etag=true"
	// ret 206 => curl -v "http://10.200.20.21:17891/common?isFile=true" -H "Range":"bytes=0-1023"
	http.HandleFunc("/download", mocks.Mock04)
	http.HandleFunc("/index5", mocks.Mock05)

	http.HandleFunc("/httpdns", mocks.Mock06)
	http.HandleFunc("/dirpath/filepath", mocks.Mock07)
	// curl -v "http://10.200.20.21:17891/post/cdnrefresh"
	http.HandleFunc("/post/cdnrefresh", mocks.Mock08)

	http.HandleFunc("/videos/file.ts", mocks.Mock09)
	http.HandleFunc("/videos/other/723c0f1d75a08397/file.m2ts", mocks.Mock09)
	http.HandleFunc("/videos/file.flv", mocks.Mock09)
	http.HandleFunc("/videos/other/723c0f1d75a08397/file.mp4", mocks.Mock09)

	http.HandleFunc("/videos/ots/file1.ts", mocks.Mock09)
	http.HandleFunc("/qpdxv/ots/file2.ts", mocks.Mock09)
	http.HandleFunc("/videos/vts/file3.m2ts", mocks.Mock09)
	http.HandleFunc("/qpdxv/vts/file4.m2ts", mocks.Mock09)

	// curl -v "http://127.0.0.1:17891/disconnect?wait=3&isSetLen=true"
	http.HandleFunc("/disconnect", mocks.Mock10)
	// curl -v "http://127.0.0.1:17891/gzip" -H "Accept-Encoding":"gzip" > /dev/null
	http.HandleFunc("/gzip", mocks.Mock11)
	// curl -v "http://127.0.0.1:17891/mimetype?type=txt&isempty=true&isfaillen=false"
	http.HandleFunc("/mimetype", mocks.Mock12)

	// issue handler
	http.HandleFunc("/kodo/4775", mocks.Mock4775)

	// curl -v "http://127.0.0.1:17891/test/1?userid="xxx"&username="xxx"&url=url1&url=url2"
	http.HandleFunc("/test/1", mocks.MockTest1)
	http.HandleFunc("/test/2", mocks.MockTest2)
	http.HandleFunc("/test/3", mocks.MockTest3)
	// curl -v "http://127.0.0.1:17891/access"
	http.HandleFunc("/access", mocks.MockTest4)

	version := "1.1.7"
	fmt.Printf("mock server start, and listen on %d. version: %s\n", *port, version)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
