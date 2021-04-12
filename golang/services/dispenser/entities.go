package dispenser

import "net/http"

type send struct {
	ID     string      `json:"id"`
	Data   interface{} `json:"data"`
	Path   string      `json:"path"`
	Method string      `json:"method"`
	Header interface{} `json:"header"`
}

type responder struct {
	W  http.ResponseWriter
	Ch chan bool
}
