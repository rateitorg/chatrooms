package handler

import (
	"log"
	"net/http"

	"github.com/rateitorg/chatrooms/service"
	"github.com/gorilla/websocket"
)

// WebSocketHandler handles WebSocket connections.
func WebSocketHandler(hub *service.Hub, w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Close the connection when the function returns
	defer conn.Close()
}


// Upgrade upgrades the HTTP connection to a WebSocket connection.
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	// FIX: Should not allow all origins
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}