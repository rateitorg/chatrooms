package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/websocket"
	"github.com/rateitorg/chatrooms/entity"
	"github.com/rateitorg/chatrooms/service"
)

type TestServer struct {
	Server *httptest.Server
	URL    string
}

// NewTestServer creates a new test server with custom handler.
func NewTestServerWithCustomHandler(handlerFunc http.HandlerFunc) *TestServer {

	// Create a test server with the handler function
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(w, r)
	}))

	return &TestServer{
		Server: server,
		URL:    "ws" + server.URL[len("http"):],
	}
}

// NewTestServer creates a new test server with client write running.
func NewTestServerWithClientWrite(hub service.HubInterface) *TestServer {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		client := &service.Client{
			Hub:  hub,
			Conn: conn,
			Send: make(chan entity.Message),
		}

		// Register the client
		client.Hub.SendToRegisterChannel(client)

		// Start the client's write goroutines
		go client.Write()
	}))

	return &TestServer{
		Server: server,
		URL:    "ws" + server.URL[len("http"):],
	}
}

// Close closes the test server.
func (testServer *TestServer) Close() {
	testServer.Server.Close()
}
