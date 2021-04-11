package socket

import "hydra_gate/web/socket"

func (c *controller) SetupListenner(s socket.Socket) {
	c.s = s

	s.InjectEvent("teste", c.test)
	s.InjectEvent("register", c.register)
	s.InjectEvent("response", c.response)
	s.InjectEvent("disconnection", c.disconnect)
}
