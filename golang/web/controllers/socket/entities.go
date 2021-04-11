package socket

import "time"

type registerPayload struct {
	Room  string `json:"id"`
	Token string `json:"token"`
	Name  string `json:"name"`
}

type worker struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ConenctedAt time.Time `json:"timestamp"`
}
