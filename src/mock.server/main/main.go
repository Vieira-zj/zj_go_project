package main

import (
	"fmt"
	"log"
	"net/http"

	"mock.server/mocks"
)

func init() {
	fmt.Println("mock server init")
}

// build cmd: /main$ GOOS=linux GOARCH=amd64 go build
// $ scp main qboxserver@10.200.20.21:~/zhengjin/main
// http://10.200.20.21:17890/index1/
func main() {
	http.HandleFunc("/index1/", mocks.Mock01)
	http.HandleFunc("/index2/", mocks.Mock02)
	http.HandleFunc("/index3/", mocks.Mock03)
	http.HandleFunc("/index4/", mocks.Mock04)
	http.HandleFunc("/index5/", mocks.Mock05)

	http.HandleFunc("/httpdns", mocks.Mock06)
	http.HandleFunc("/dirpath/filepath", mocks.Mock07)

	http.HandleFunc("/post/cdnrefresh", mocks.Mock08)

	http.HandleFunc("/mock1", mocks.Mock21)
	http.HandleFunc("/mock2", mocks.Mock22)
	http.HandleFunc("/mock3", mocks.Mock23)

	const port = 17890
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
