package lambda

import (
	"hydra_gate/services/dispenser"
	"hydra_gate/utils/logger"
	"hydra_gate/web/controllers"
	"hydra_gate/web/middleware"
	"hydra_gate/web/server"
	"net/http"

	"github.com/google/uuid"
)

// --- Lambda ---

type controller struct {
	s           *server.Server
	dispService dispenser.Service
	http.Handler
}

// New lambda controller
func New() controllers.Controller {
	return &controller{
		s:           nil,
		dispService: dispenser.New(),
	}
}

func (c *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uid := uuid.New()
	room := r.Header.Get("room")
	logger.Debug("> Received request [room]", r.RequestURI, room)
	err := c.dispService.HandleRequest(room, uid.String(), w, r)
	if err != nil {
		middleware.ResponseError(w, err.Error())
	}
}
