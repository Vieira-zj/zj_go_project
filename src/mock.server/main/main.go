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

// build cmd: $ GOOS=linux GOARCH=amd64 go build
// $ scp main qboxserver@10.200.20.21:~/zhengjin/main
func main() {
	http.HandleFunc("/index1/", mocks.Mock01)
	http.HandleFunc("/index2/", mocks.Mock02)
	http.HandleFunc("/index3/", mocks.Mock03)
	http.HandleFunc("/index4/", mocks.Mock04)
	log.Fatal(http.ListenAndServe(":17890", nil))
}
