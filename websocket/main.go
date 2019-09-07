package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)
var(
	upgrader=websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	http.HandleFunc("/ws",wsHandler)
	http.ListenAndServe(":8888",nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var(
		conn *websocket.Conn
		err error

		data []byte
	)
	if conn, err = upgrader.Upgrade(w, r, nil);err!= nil {
		return
	}
	for{
		if _,data,err=conn.ReadMessage();err!= nil {
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, data);err!= nil {
			goto ERR
		}
	}
	ERR:
		conn.Close()
}