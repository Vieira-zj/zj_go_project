package main

import (
	"flag"
	"log"
	"net/http"

	"src/mock.server/common"
	"src/mock.server/handlers"
)

func init() {
	if err := common.InitConfigs(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	help := flag.Bool("h", false, "help.")
	port := flag.String("p", "17891", "mock server listening port.")

	flag.Parse()
	if *help {
		flag.Usage()
	}

	router := handlers.NewHTTPRouter()
	log.Printf("Mock Server start, and listen on %s.\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, router))
}
