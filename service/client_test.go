package service_test

import (
	"github.com/gorilla/websocket"
	"github.com/rateitorg/chatrooms/entity"
	"github.com/rateitorg/chatrooms/test"
	"github.com/rateitorg/chatrooms/test/mocks"
	"github.com/stretchr/testify/mock"
	"sync"
	"testing"
)

// Test Client Write
// Verifies that the client can read message from the WebSocket connection and broadcast it to the hub.
func TestClientWrite_MessageBroadcast(t *testing.T) {
	// Arrange
	mockHub := new(mocks.Hub)
	var wg sync.WaitGroup // WaitGroup to synchronize goroutines

	// Increment the WaitGroup counter for each expected call
	wg.Add(2)

	// Set Mock Expectations with Done signaling
	mockHub.On("SendToRegisterChannel", mock.AnythingOfType("*service.Client")).Once().Run(func(args mock.Arguments) {
		wg.Done() // Signal this expectation has been met
	})
	mockHub.On("SendToBroadcastChannel", mock.AnythingOfType("entity.Message")).Once().Run(func(args mock.Arguments) {
		wg.Done() // Signal this expectation has been met
	})

	// Create a new test server
	testServer := test.NewTestServerWithClientWrite(mockHub)
	defer testServer.Close()

	// Create a WebSocket client
	ws, _, err := websocket.DefaultDialer.Dial(testServer.URL, nil)
	if err != nil {
		t.Fatalf("Could not connect to WebSocket: %v", err)
	}
	defer ws.Close()

	// Act
	// Write a message to the WebSocket connection
	testMessage := entity.NewMessage("Message Content", "User", "Timestamp")
	err = ws.WriteJSON(testMessage)
	if err != nil {
		t.Fatalf("Could not write message to WebSocket: %v", err)
	}

	// Wait for all expectations to be met
	wg.Wait()

	// Assert
	// Verify the mock expectations
	mockHub.AssertExpectations(t)
}
