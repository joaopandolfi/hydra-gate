package socket

import "net/http"

type registerPayload struct {
	Room  string `json:"id"`
	Token string `json:"token"`
	Name  string `json:"name"`
}

type response struct {
	ID      string      `json:"id"`
	Success bool        `json:"success"`
	Data    string      `json:"data"`
	Header  http.Header `json:"header"`
}
