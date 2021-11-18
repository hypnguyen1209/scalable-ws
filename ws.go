package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/index.html")
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		log.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}
