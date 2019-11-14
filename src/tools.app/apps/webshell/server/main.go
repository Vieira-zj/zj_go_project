package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
	myutils "tools.app/utils"
)

// Refer: https://www.cnblogs.com/lanyangsh/p/9822403.html

func main() {
	addr := flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()

	http.HandleFunc("/", Home)
	http.HandleFunc("/echo", EchoMsg)

	log.Println("http server (websocket) is started...")
	log.Fatal(http.ListenAndServe(*addr, nil))
}

// Home http handler for home page.
func Home(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := myutils.ReadFileContent("home.html")
	if err != nil {
		panic(err)
	}

	homeTemplate := template.Must(template.New("").Parse(htmlContent))
	if err := homeTemplate.Execute(w, "ws://"+r.Host+"/echo"); err != nil {
		panic(err)
	}
}

// EchoMsg echos message by websocket.
func EchoMsg(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			if r.Method != "GET" {
				log.Println("websocket not support GET")
				return false
			}
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	log.Println("remote client connect:", conn.RemoteAddr())

	closeErrors := []int{websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived}
	for {
		log.Println("ws server pending for read message...")
		msgType, message, err := conn.ReadMessage()
		if err != nil {
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() {
					log.Printf("(remote client [%v]) ws server read msg timeout error\n", conn.RemoteAddr())
					return
				}
			}
			if websocket.IsCloseError(err, closeErrors...) {
				log.Printf("(remote client [%v]) ws server read msg close error: %v\n", conn.RemoteAddr(), err)
				return
			}
			if websocket.IsUnexpectedCloseError(err, closeErrors...) {
				log.Printf("(remote client [%v]) ws server read msg unexpected close error: %v\n", conn.RemoteAddr(), err)
				return
			}
			panic(err)
		}
		log.Println("recv:", string(message))

		if err := conn.WriteMessage(msgType, message); err != nil {
			panic(err)
		}
	}
}
