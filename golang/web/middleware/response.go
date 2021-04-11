package middleware

import (
	"encoding/json"
	"hydra_gate/config"
	"net/http"

	"hydra_gate/utils/logger"
)

// marshaler
var marshaler func(v interface{}) ([]byte, error) = json.Marshal

// header -
func header(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", config.Get().Server.Security.CORS)
	w.Header().Add("Content-Type", "application/json")
}

// Response - Make default generic response
func Response(w http.ResponseWriter, resp interface{}, status int) {
	// set Header
	header(w)
	w.WriteHeader(status)
	b, err := marshaler(resp)

	if err == nil {
		// Responde
		w.Write(b)
	} else {
		logger.Error("Error on convert response to JSON", err)
		ResponseError(w, "Error on convert response to JSON")
	}
}

// ResponseError - Make default generic response
func ResponseError(w http.ResponseWriter, resp interface{}) {
	// set Header
	header(w)
	b, _ := marshaler(resp)
	responseError(w, string(b))
}

// responseError - Private function to make response
func responseError(w http.ResponseWriter, message string) {
	b, _ := json.Marshal(map[string]string{"message": message})
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(b)
}
