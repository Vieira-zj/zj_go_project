package main

import (
	"fmt"
	"log"
	"net/http"

	mysvc "tools.app/services/httprouter"
)

func main() {
	// test:
	// curl -v "http://127.0.0.1:17890/"

	// curl -v "http://127.0.0.1:17890/books"
	// curl -v "http://127.0.0.1:17890/books/001"

	// curl -v "http://127.0.0.1:17890/templates/1"
	// curl -v "http://127.0.0.1:17890/templates/2"

	router := mysvc.NewBooksSvr().GetHTTPRouter()
	fmt.Println("Books http router started, and listen at 17890.")
	log.Fatal(http.ListenAndServe(":17890", router))
}
