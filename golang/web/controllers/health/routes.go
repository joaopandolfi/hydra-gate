package health

import (
	"hydra_gate/web/server"
)

// SetupRouter -
func (c *controller) SetupRouter(s *server.Server) {
	c.s = s

	c.s.R.HandleFunc("/_hydra_gate/health", c.health).Methods("POST", "GET", "HEAD")
	c.s.R.HandleFunc("/_hydra_gate/config", c.config).Methods("POST", "GET", "HEAD")
}
