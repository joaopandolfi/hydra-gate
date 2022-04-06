package lambda

import "hydra_gate/web/server"

// SetupRouter -
func (c *controller) SetupRouter(s *server.Server) {
	c.s = s
	c.s.R.HandleFunc("/x", c.ServeHTTP).Methods("POST", "GET", "HEAD")

}
