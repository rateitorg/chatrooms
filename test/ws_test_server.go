package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/websocket"
)

type TestServer struct {
	Server   *httptest.Server
	Upgrader websocket.Upgrader
	URL      string
}

// NewTestServer creates a new test server.
func NewTestServer(handlerFunc http.HandlerFunc) *TestServer {
	upgrader := websocket.Upgrader{}

	// Create a test server with the handler function
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(w, r)
	}))

	return &TestServer{
		Server:   server,
		Upgrader: upgrader,
		URL:      "ws" + server.URL[len("http"):],
	}
}

// Close closes the test server.
func (testServer *TestServer) Close() {
	testServer.Server.Close()
}
