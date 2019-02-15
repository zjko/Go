package main

import (
	"net/http"
	"time"

	"./impl"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("hello"))
	var (
		wsConn *websocket.Conn
		err    error
		conn   *impl.Connection
		// msgType int
		data []byte
	)
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
	}
	if err = conn.WriteMessage(data); err != nil {
		goto ERR
	}
ERR:
	//TODO: 关闭连接的操作
	conn.Close()
}
func main() {
	// http://localhost:7777/ws
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0:7777", nil)
}
