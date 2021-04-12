package dispenser

import (
	"hydra_gate/models"
	"hydra_gate/web/middleware"
	"net/http"

	"golang.org/x/xerrors"
)

type Service interface {
	InjectRooms(rooms *map[string]*models.Worker)
	HandleRequest(room, reqID string, w http.ResponseWriter, r *http.Request) error
	ResponseRequest(reqID string, statusCode int, b []byte) error
}

var requests map[string]*responder
var rooms *map[string]*models.Worker

type service struct {
}

func New() Service {
	if requests == nil {
		requests = map[string]*responder{}
	}

	return &service{}
}

func (s *service) InjectRooms(r *map[string]*models.Worker) {
	rooms = r
}

func (s *service) getRooms() map[string]*models.Worker {
	return *rooms
}

func (s *service) HandleRequest(room, reqID string, w http.ResponseWriter, r *http.Request) error {
	sr := s.getRooms()[room]
	if sr == nil {
		return xerrors.Errorf("no workers on room: %w", s.getRooms())
	}

	body, err := middleware.GetBody(r)
	if err != nil {
		return xerrors.Errorf("getting body")
	}

	ch := make(chan bool)

	err = sr.Sokt.Emit("handle", send{
		ID:     reqID,
		Data:   body,
		Path:   r.RequestURI,
		Method: r.Method,
		Header: r.Header,
	})

	requests[reqID] = &responder{
		W:  w,
		Ch: ch,
	}

	if err != nil {
		return xerrors.Errorf("emmiting to socket")
	}

	<-ch

	return nil
}

func (s *service) ResponseRequest(reqID string, statusCode int, b []byte) error {
	req := requests[reqID]
	if req == nil {
		return xerrors.Errorf("invalid request id")
	}

	w := req.W

	w.WriteHeader(statusCode)
	w.Write(b)

	req.Ch <- true

	return nil
}
