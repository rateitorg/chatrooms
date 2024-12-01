package handler_test

import (
	"net/http"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/rateitorg/chatrooms/handler"
	"github.com/rateitorg/chatrooms/test"
	"github.com/rateitorg/chatrooms/test/mocks"
	"github.com/stretchr/testify/mock"
)

// TestWebSocketHandler tests the WebSocket upgrade logic.
func TestWebSocketHandler(t *testing.T) {
	// Arrange
	mockHub := new(mocks.Hub)

	// Set Mock Expectations
	mockHub.On("SendToRegisterChannel", mock.AnythingOfType("*service.Client")).Once()

	// Create a new test server with the WebSocketHandler as the handler
	testServer := test.NewTestServerWithCustomHandler(func(w http.ResponseWriter, r *http.Request) {
		handler.WebSocketHandler(mockHub, w, r)
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

	// Verify the mock expectations
	mockHub.AssertExpectations(t)

	// TODO: When origins are stricter, assert for the correct origin
}
