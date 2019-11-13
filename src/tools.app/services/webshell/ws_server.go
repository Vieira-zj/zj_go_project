package webshell

import (
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/websocket"
	myutils "tools.app/utils"
)

// Refer: https://www.cnblogs.com/lanyangsh/p/9822403.html

// Home http handler for home page.
func Home(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(os.Getenv("GOPATH"), "src/tools.app/services/webshell", "ws_home.html")
	htmlContent, err := myutils.ReadFileContent(filePath)
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
