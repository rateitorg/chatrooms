package service

import (
	"github.com/rateitorg/chatrooms/entity"
)

// Hub will hold a set of active clients and broadcast messages.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []entity.Message

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

// NewHub creates a new hub.
func NewHub() *Hub {
	return &Hub{
		Broadcast: make(chan []entity.Message),
	}
}

// Hub Event Loop
func (h *Hub) Run() {
	// Continuously listen for events
	for {
		select {
		case client := <-h.Register: // If there is something in the register channel
			h.Clients[client] = true
		case client := <-h.Unregister: // If there is something in the unregister channel
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client) // Remove from map
				close(client.Send) // Close the send channel
			}
		}
	}
}