package service

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/rateitorg/chatrooms/entity"
)

const (
	// Understanding the below constants
	// The server will send a ping, every pingPeriod.
	// The client has pongWait time to respond to the ping.
	// If the client does not respond within pingPeriod time, the connection is closed.

	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from the client
	maxMessageSize = 512
)

type Client struct {
	// The hub that the client belongs to
	Hub *Hub

	// The WebSocket connection
	Conn *websocket.Conn

	// Buffered channel of outbound messages
	Send chan entity.Message
}

// Takes a message from client and sends it to the hub
func (client *Client) Write() {
	defer func() {
		// On exit, unregister the client
		client.Hub.Unregister <- client
		client.Conn.Close()
	}()

	client.Conn.SetReadLimit(maxMessageSize)              // Set the maximum message size
	client.Conn.SetReadDeadline(time.Now().Add(pongWait)) // Set the read deadline

	client.Conn.SetPongHandler(func(string) error {
		client.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// Continuously listen for messages from the client
	for {
		var message entity.Message
		err := client.Conn.ReadJSON(&message)
		if err != nil {
			// TODO: Handler error gracefully
			break
		}
		client.Hub.Broadcast <- message
	}
}

// Takes a message from the hub and sends it to the client
func (client *Client) Read() {
	// Create a ticker that sends a ping to the client every pingPeriod
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		// On exit, stop the ticker and close the connection
		ticker.Stop()
		client.Conn.Close()
	}()

	// Continuously listen for messages from the hub
	for {
		select {
		case messages, ok := <-client.Send: // If there is something in the send channel
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel
				client.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			err := client.Conn.WriteJSON(messages)
			if err != nil {
				return
			}

		case <-ticker.C: // If the ticker sends a ping
			client.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
