package models

import (
	"time"

	gosocketio "github.com/ambelovsky/gosf-socketio"
)

type Worker struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ConenctedAt time.Time `json:"timestamp"`
	Sokt        *gosocketio.Channel
	Room        string
}

func removeWorker(s []Worker, i int) []Worker {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}
