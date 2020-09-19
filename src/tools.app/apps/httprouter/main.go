package main

import (
	"fmt"
	"log"
	"net/http"

	mysvc "src/tools.app/services/httprouter"
)

func main() {
	// hello:
	// curl -v "http://127.0.0.1:17890/"
	// book demo:
	// curl -v "http://127.0.0.1:17890/books"
	// curl -v "http://127.0.0.1:17890/books/001"
	// html template demo:
	// curl -v "http://127.0.0.1:17890/templates/{1-4}"

	router := mysvc.NewBooksSvr().GetHTTPRouter()
	fmt.Println("Books http router started, and listen at 17890.")
	log.Fatal(http.ListenAndServe(":17890", router))
}
