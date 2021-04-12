package router

import (
	"hydra_gate/config"
	"hydra_gate/web/controllers/health"
	"hydra_gate/web/controllers/lambda"
	socController "hydra_gate/web/controllers/socket"
	"hydra_gate/web/server"
	"hydra_gate/web/socket"

	"github.com/unrolled/secure"
)

// Router public struct
type Router struct {
	s   *server.Server
	skt socket.Socket
}

// New Router
func New(s *server.Server) Router {
	return Router{s: s}
}

// Setup router
func (r *Router) Setup() {
	r.secure()

	health.New().SetupRouter(r.s)

	lambda.New().SetupRouter(r.s)

	//api := r.createSubRouter("/api")

	// Socket
	socController.New().SetupListenner(r.skt)
}

// HandleSocket gestor
func (r *Router) HandleSocket(socket socket.Socket) {
	r.skt = socket
	//r.s.R.Handle(fmt.Sprintf("%s/socket.io/", config.Get().Socket.Path), socket.GetServer())
	socket.InjectOnRouter(r.s.R)
}

// CreateSubRouter with path
func (r *Router) createSubRouter(path string) *server.Server {
	return &server.Server{
		R:      r.s.R.PathPrefix(path).Subrouter(),
		Config: r.s.Config,
	}
}

func (r *Router) secure() {
	secureMiddleware := secure.New(config.Get().Server.Security.Opsec)
	r.s.R.Use(secureMiddleware.Handler)
}
