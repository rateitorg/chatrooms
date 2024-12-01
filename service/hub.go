package service

import (
	"github.com/rateitorg/chatrooms/entity"
)

type HubInterface interface {
	SendToRegisterChannel(client *Client)
	SendToUnregisterChannel(client *Client)
	SendToBroadcastChannel(message entity.Message)
}

// Hub will hold a set of active clients and broadcast messages.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan entity.Message

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

// NewHub creates a new hub.
func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan entity.Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

// Hub Event Loop
func (h *Hub) Run() {
	// Continuously listen for events
	for {
		select {
		case client := <-h.Register: // If there is something in the register channel
			h.registerClient(client)
		case client := <-h.Unregister: // If there is something in the unregister channel
			h.unregisterClient(client)
		case message := <-h.Broadcast: // If there is something in the broadcast channel
			for client := range h.Clients {
				h.broadcastMessage(client, message)
			}
		}
	}
}

// Channel data senders
func (h *Hub) SendToRegisterChannel(client *Client) {
	h.Register <- client
}

func (h *Hub) SendToUnregisterChannel(client *Client) {
	h.Unregister <- client
}

func (h *Hub) SendToBroadcastChannel(message entity.Message) {
	h.Broadcast <- message
}

// Event Handlers
func (h *Hub) registerClient(client *Client) {
	h.Clients[client] = true
}

func (h *Hub) unregisterClient(client *Client) {
	if _, ok := h.Clients[client]; ok {
		delete(h.Clients, client) // Remove from map
		close(client.Send)        // Close the send channel
	}
}

func (h *Hub) broadcastMessage(client *Client, message entity.Message) {
	// TODO: performance optimization. If the client is the sender, don't send the message back to the client
	select {
		case client.Send <- message: // If the send channel is not full send the message
		default:
			delete(h.Clients, client) // Remove the client if the send channel is full
			close(client.Send)        // Close the send channel
		}
}

