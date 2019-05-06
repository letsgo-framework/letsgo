package controllers

import (
	"github.com/gorilla/websocket"
	"github.com/letsgo-framework/letsgo/types"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeWebsocket(hub *types.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &types.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()

}

func NewHub() *types.Hub {
	return &types.Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *types.Client),
		Unregister: make(chan *types.Client),
		Clients:    make(map[*types.Client]bool),
	}
}

