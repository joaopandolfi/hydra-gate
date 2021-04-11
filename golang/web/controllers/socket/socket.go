package socket

import (
	"net/http"
	"sync"

	"hydra_gate/web/controllers"
	"hydra_gate/web/socket"

	"hydra_gate/utils/format"
	"hydra_gate/utils/logger"

	gosocketio "github.com/ambelovsky/gosf-socketio"
)

type controller struct {
	s       socket.Socket
	workers map[string]*worker
	rooms   map[string]*worker
	mu      sync.RWMutex
}

var caller map[string]*http.ResponseWriter

// New socket controller
func New() controllers.SocketController {
	return &controller{
		s:       nil,
		rooms:   map[string]*worker{},
		workers: map[string]*worker{},
	}
}

func (c *controller) test(ch *gosocketio.Channel, msg interface{}) string {
	logger.Debug("[SOCKET][TEST] - handle socket event")
	//ch.Emit()
	return ""
}

func removeWorker(s []worker, i int) []worker {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}

// Register new user
func (c *controller) register(ch *gosocketio.Channel, registerPayload registerPayload) string {
	logger.Debug("Worker [1] is conencted -> [2]", registerPayload.Room, ch.Id())

	w := c.rooms[registerPayload.Room]

	if w != nil {
		logger.Error("Disconnecting worker because alrealdy have one in this room")
		ch.Emit("error", "Room is full")
		ch.Close()
		return "closed"
	}

	newWorker := worker{
		ID:          ch.Id(),
		Name:        registerPayload.Name,
		ConenctedAt: format.CurrentDate(),
	}

	c.rooms[registerPayload.Room] = &newWorker

	c.mu.Lock()
	c.workers[ch.Id()] = &newWorker
	c.mu.Unlock()

	ch.Emit("registered", map[string]string{"sid": ch.Id()})
	//ch.BroadcastTo(registerPayload.Room, "dash_new_map_marker", registerPayload.Metadata)

	return "ok"
}

func (c *controller) response(ch *gosocketio.Channel, data map[string]interface{}) string {
	callerID, ok := data["id"].(string)
	if !ok {
		logger.Debug(" >> Caller ID not found on response")
		return "error"
	}

	logger.Debug(" Respond: callerID [1] worker [2]", callerID, ch.Id())

	//caller
	return "ok"
}

func (c *controller) disconnect(ch *gosocketio.Channel, m interface{}) string {
	room := ""
	c.mu.Lock()
	w := c.workers[ch.Id()]
	if w != nil {
		room = w.ID
		delete(c.rooms, w.ID)
		delete(c.workers, ch.Id())
	}
	c.mu.Unlock()

	logger.Debug("[-] Worker [1] disconnected -> [2]", ch.Id(), room)

	return "ok"
}
