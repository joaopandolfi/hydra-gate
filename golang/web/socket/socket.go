package socket

import (
	"context"
	"fmt"
	"net/http"

	"hydra_gate/config"
	"hydra_gate/utils/logger"

	gosocketio "github.com/ambelovsky/gosf-socketio"
	"github.com/ambelovsky/gosf-socketio/transport"
	"github.com/gorilla/mux"
)

// Socket is a public interface to start sockets
type Socket interface {
	Setup(port, cors, path string)
	Start()
	Handlers()
	InjectEvent(event string, f interface{})
	Shutdown(ctx context.Context) error
	GetServer() gosocketio.Server
	InjectOnRouter(r *mux.Router)
}

// New Socket
func New() Socket {

	s := &socket{
		clients: map[string]client{},
	}

	s.Setup(config.Get().Socket.Port, config.Get().Socket.CORS, config.Get().Socket.Path)
	return s
}

type client struct {
	c *gosocketio.Channel
	//token utils.Token
}

type socket struct {
	port    string
	cors    string
	path    string
	server  *gosocketio.Server
	clients map[string]client
	websrv  *http.Server
}

func (s *socket) Setup(port, cors, path string) {
	s.server = gosocketio.NewServer(transport.GetDefaultWebsocketTransport())
	s.port = port
	s.cors = cors
	s.path = path

	//handle connected
	s.server.On(gosocketio.OnConnection, func(c *gosocketio.Channel) {
		logger.Debug("[Socket] A wild user appears", c.Id())

		// vars := handlers.GetQueryes(c.Request())
		// token, err := utils.CheckJwtToken(vars.Get("token"))
		// if err != nil {
		// 	utils.Debug("[Socket] - Removing beacuse token is not valid", c.Id())
		// 	c.Close()
		// }

		// s.clients[c.Id()] = client{c: c, token: token}
		c.Emit("welcome", "take your pills and sit")

		c.Emit("welcome2", "take your pills and sit")
		c.Emit("welcome3", "take your pills and sit")
	})

	s.Handlers()
	//setup http server
}

// GetServer returns socket server to inject in a mux webserver
func (s *socket) GetServer() gosocketio.Server {
	return *s.server
}

func (s *socket) Start() {
	serveMux := http.NewServeMux()
	serveMux.Handle(fmt.Sprintf("%s/socket.io/", s.path), s.server)
	srv := &http.Server{
		Handler: serveMux,
		Addr:    ":" + s.port,
	}
	s.websrv = srv
	go srv.ListenAndServe()
	logger.Info("[SOCKET][START] listenning on: ", s.port, s.path)

}

// Shutdown server
func (s *socket) Shutdown(ctx context.Context) error {
	return s.websrv.Shutdown(ctx)
}

func (s *socket) InjectEvent(event string, f interface{}) {
	s.server.On(event, f)
}

func (s *socket) InjectOnRouter(r *mux.Router) {
	r.Handle(fmt.Sprintf("%s/socket.io/", config.Get().Socket.Path), s.server)
}

func (s *socket) Handlers() {
	//handle custom event
	s.server.On(".ping", func(c *gosocketio.Channel, msg interface{}) string {
		//send event to all in room
		logger.Debug("[.PING] received", msg)
		//c.BroadcastTo(".pong", "", msg)
		c.Emit(".pong", msg)
		return fmt.Sprint(msg)
	})

}
