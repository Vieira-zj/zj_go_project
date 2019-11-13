package main

import (
	"flag"
	"log"
	"net/http"

	mysvc "tools.app/services/webshell"
)

func main() {
	addr := flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()

	http.HandleFunc("/", mysvc.Home)
	http.HandleFunc("/echo", mysvc.EchoMsg)

	log.Println("http server (websocket) is started...")
	log.Fatal(http.ListenAndServe(*addr, nil))
}
