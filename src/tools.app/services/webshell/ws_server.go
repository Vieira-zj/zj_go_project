package webshell

import (
	"fmt"
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

// EchoMsg echos message by websocket.
func EchoMsg(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			if r.Method != "GET" {
				fmt.Println("websocket not support GET")
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
	log.Println("remote connect:", conn.RemoteAddr())

	for {
		msgType, message, err := conn.ReadMessage()
		if err != nil {
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() {
					log.Printf("ws read message timeout error, remote [%v]\n", conn.RemoteAddr())
					return
				}
			}
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				log.Printf("ws read message remote [%v] unexpected close error: %v\n", conn.RemoteAddr(), err)
				return
			}
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				log.Printf("ws read message remote [%v] close: %v\n", conn.RemoteAddr(), err)
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

// Home http handler for home page.
func Home(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(os.Getenv("GOPATH"), "src/tools.app/services/webshell", "home.html")
	htmlContent, err := myutils.ReadFileContent(filePath)
	if err != nil {
		panic(err)
	}

	homeTemplate := template.Must(template.New("").Parse(htmlContent))
	if err := homeTemplate.Execute(w, "ws://"+r.Host+"/echo"); err != nil {
		panic(err)
	}
}
