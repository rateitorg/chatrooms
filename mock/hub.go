package mock

import (
	"github.com/rateitorg/chatrooms/service"
	"github.com/stretchr/testify/mock"
)

type Hub struct {
	mock.Mock
}

func (m *Hub) RegisterClient(client *service.Client) {
	m.Called(client)
}

func (m *Hub) UnregisterClient(client *service.Client) {
	m.Called(client)
}