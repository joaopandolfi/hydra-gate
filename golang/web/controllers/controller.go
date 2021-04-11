package controllers

import (
	"hydra_gate/web/server"
	"hydra_gate/web/socket"
)

// Controller public contract
type Controller interface {
	SetupRouter(s *server.Server)
}

// SocketController public contract
type SocketController interface {
	SetupListenner(s socket.Socket)
}
