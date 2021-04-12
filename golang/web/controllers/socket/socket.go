package socket

import (
	"net/http"
	"sync"

	b64 "encoding/base64"
	"hydra_gate/models"
	"hydra_gate/services/dispenser"
	"hydra_gate/web/controllers"
	"hydra_gate/web/socket"

	"hydra_gate/utils/format"
	"hydra_gate/utils/logger"

	gosocketio "github.com/ambelovsky/gosf-socketio"
)

type controller struct {
	s         socket.Socket
	workers   map[string]*models.Worker
	rooms     map[string]*models.Worker
	mu        sync.RWMutex
	dispenser dispenser.Service
}

// New socket controller
func New() controllers.SocketController {
	rooms := map[string]*models.Worker{}
	disp := dispenser.New()
	disp.InjectRooms(&rooms)

	return &controller{
		s:         nil,
		rooms:     rooms,
		workers:   map[string]*models.Worker{},
		dispenser: disp,
	}
}

func (c *controller) test(ch *gosocketio.Channel, msg interface{}) string {
	logger.Debug("[SOCKET][TEST] - handle socket event", msg)
	//ch.Emit()
	return ""
}

// Register new user
func (c *controller) register(ch *gosocketio.Channel, registerPayload registerPayload) string {
	logger.Debug("> Worker [1] is conencted -> [2]", registerPayload.Room, ch.Id())

	w := c.rooms[registerPayload.Room]

	if w != nil {
		logger.Error("Disconnecting worker because alrealdy have one in this room")
		ch.Emit("error", "Room is full")
		ch.Close()
		return "closed"
	}

	newWorker := models.Worker{
		ID:          ch.Id(),
		Name:        registerPayload.Name,
		ConenctedAt: format.CurrentDate(),
		Sokt:        ch,
	}

	c.rooms[registerPayload.Room] = &newWorker

	c.mu.Lock()
	c.workers[ch.Id()] = &newWorker
	c.mu.Unlock()

	ch.Emit("registered", map[string]string{"sid": ch.Id()})
	//ch.BroadcastTo(registerPayload.Room, "dash_new_map_marker", registerPayload.Metadata)

	return "ok"
}

func (c *controller) response(ch *gosocketio.Channel, data response) string {

	statusCode := 200
	if !data.Success {
		statusCode = http.StatusInternalServerError
	}

	b, err := b64.StdEncoding.DecodeString(data.Data)
	if err != nil {
		logger.Error(" Respondig request: ", err.Error())
		return "error"
	}

	err = c.dispenser.ResponseRequest(data.ID, statusCode, b, data.Header)
	if err != nil {
		logger.Error("invalid response: %w", err)
		return "error"
	}

	logger.Debug(" Respond: callerID [1] worker [2]", data.ID, ch.Id())

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
