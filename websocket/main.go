package main

import (
	"github.com/gorilla/websocket"
	"github.com/jsxz/go-tutorial/websocket/impl"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":8888", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		conn   *impl.Connection
		data   []byte
	)
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}
	go func() {
		var (
			err error
		)
		if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}()
	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
