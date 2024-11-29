package service

import (
	"github.com/rateitorg/chatrooms/entity"
)

// Hub will hold a set of active clients and broadcast messages.
type Hub struct {
	// Registered clients.
	// TODO: add a map to store clients

	// Inbound messages from the clients.
	Broadcast chan []entity.Message
}

// NewHub creates a new hub.
func NewHub() *Hub {
	return &Hub{
		Broadcast: make(chan []entity.Message),
	}
}