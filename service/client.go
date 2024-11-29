package service

import (
	"github.com/rateitorg/chatrooms/entity"
	"github.com/gorilla/websocket"
)

type Client struct {
	// The hub that the client belongs to
	Hub *Hub

	// The WebSocket connection
	Conn *websocket.Conn

	// Buffered channel of outbound messages
	Send chan []entity.Message
}



