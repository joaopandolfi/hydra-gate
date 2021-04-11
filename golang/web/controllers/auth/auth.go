package auth

import (
	"hydra_gate/web/controllers"
	"hydra_gate/web/server"
)

// --- Auth ---

type controller struct {
	s *server.Server
}

// New Auth controller
func New() controllers.Controller {
	return &controller{
		s: nil,
	}
}
