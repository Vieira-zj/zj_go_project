package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	var addr = flag.String("addr", "localhost:8080", "http service (websocket) address")
	flag.Parse()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{
		Scheme: "ws",
		Host:   *addr,
		Path:   "/echo",
	}
	log.Println("connecting to ws server:", u.String())
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// read message
	done := make(chan struct{})
	go func() {
		closeErrors := []int{websocket.CloseGoingAway, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived}
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, closeErrors...) {
					log.Printf("(remote server [%v]) ws client read msg close error: %v\n", conn.RemoteAddr(), err)
					return
				}
				if websocket.IsUnexpectedCloseError(err, closeErrors...) {
					log.Printf("(remote server [%v]) ws client read msg unexpected close error: %v\n", conn.RemoteAddr(), err)
					return
				}
				panic(err)
			}
			log.Println("recv:", string(message))
		}
	}()

	// write message
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-done:
			log.Println("write msg done and exit ws client")
			return
		case t := <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				panic(err)
			}
		case <-interrupt:
			if err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
				panic(err)
			}
			select {
			case <-done:
				log.Println("interrupt, write msg done and exit ws client")
			case <-time.After(time.Second):
				log.Println("interrupt, wait 1 sec and exit ws client")
			}
			return
		}
	}
}
