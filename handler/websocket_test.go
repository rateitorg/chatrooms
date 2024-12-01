package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/rateitorg/chatrooms/handler"
	"github.com/rateitorg/chatrooms/service"
)

// TestWebSocketHandler tests the WebSocket upgrade logic.
// No need to mock the service, as no service methods are called.
func TestWebSocketHandler(t *testing.T) {
	// Arrange
	hub := service.NewHub()

	// Create a new test server with the WebSocketHandler as the handler
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.WebSocketHandler(hub, w, r)
	}))
	defer server.Close()

	// Get the WebSocket URL
	url := "ws" + server.URL[len("http"):]

	// Act
	// Create a WebSocket client
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
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
