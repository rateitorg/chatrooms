package handler_test

import (
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/rateitorg/chatrooms/handler"
	"github.com/rateitorg/chatrooms/service"
	"github.com/rateitorg/chatrooms/test"
)

// TestWebSocketHandler tests the WebSocket upgrade logic.
// No need to mock the service, as no service methods are called.
func TestWebSocketHandler(t *testing.T) {
	// Arrange
	hub := service.NewHub()

	// Create a new test server with the WebSocketHandler as the handler
	testServer := test.NewTestServer(func(w http.ResponseWriter, r *http.Request) {
		handler.WebSocketHandler(hub, w, r)
	})
	defer testServer.Close()

	// Act
	// Create a WebSocket client
	ws, _, err := websocket.DefaultDialer.Dial(testServer.URL, nil)
	if err != nil {
		t.Fatalf("Could not connect to WebSocket: %v", err)
	}
	defer ws.Close()

	// Assert
	// Verify the WebSocket connection is successful
	if ws == nil || ws.RemoteAddr() == nil {
		t.Fatal("WebSocket connection failed")
	}

	// TODO: When origins are stricter, assert for the correct origin
}
