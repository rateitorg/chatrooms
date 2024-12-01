package mocks

import (
	"github.com/rateitorg/chatrooms/service"
	"github.com/rateitorg/chatrooms/entity"
	"github.com/stretchr/testify/mock"
)

type Hub struct {
	mock.Mock
}

func (m *Hub) SendToRegisterChannel(client *service.Client) {
	m.Called(client)
}

func (m *Hub) SendToUnregisterChannel(client *service.Client) {
	m.Called(client)
}

func (m *Hub) BroadcastMessageBroadcastMessage(message entity.Message) {
	m.Called(message)
}