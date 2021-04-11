package auth

import "hydra_gate/web/server"

// SetupRouter -
func (c *controller) SetupRouter(s *server.Server) {
	c.s = s
	//c.s.R.HandleFunc("/", c.health).Methods("POST", "GET", "HEAD")

}
