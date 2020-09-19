package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"src/mock.server/common"
	"src/mock.server/handlers"
)

func init() {
	if err := common.InitConfigs(); err != nil {
		log.Println(err)
	}
}

func main() {
	help := flag.Bool("h", false, "help.")
	port := flag.Int("p", 17891, "mock server listening port.")

	flag.Parse()
	if *help {
		flag.Usage()
	}

	router := handlers.NewHTTPRouter()
	log.Printf("Mock Server start, and listen on %d.\n", *port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), router))
}
