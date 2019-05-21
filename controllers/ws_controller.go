package controllers

import (
	"github.com/gorilla/websocket"
	letslog "github.com/letsgo-framework/letsgo/log"
	"github.com/letsgo-framework/letsgo/types"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// ServeWebsocket starts websocket
func ServeWebsocket(hub *types.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		letslog.Error(err.Error())
		return
	}
	client := &types.Client{Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()

}

// NewHub created new Hub
func NewHub() *types.Hub {
	return &types.Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *types.Client),
		Unregister: make(chan *types.Client),
		Clients:    make(map[*types.Client]bool),
	}
}
