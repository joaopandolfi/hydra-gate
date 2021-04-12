package lambda

import "hydra_gate/web/server"

// SetupRouter -
func (c *controller) SetupRouter(s *server.Server) {
	c.s = s
	s.R.PathPrefix("/").Handler(c)
}
