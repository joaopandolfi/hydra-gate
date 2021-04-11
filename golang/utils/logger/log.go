package logger

import (
	"encoding/json"
	"fmt"
	"time"
)

// Info log
func Info(message string, data ...interface{}) {
	dispatch(fmt.Sprintf("[+][%s]%s", time.Now().UTC().Format(time.RFC3339), message), data)
}

// Error log
func Error(message string, data ...interface{}) {
	dispatch(fmt.Sprintf("[-][%s][Error]%s", time.Now().UTC().Format(time.RFC3339), message), data)
}

// CriticalError log
func CriticalError(message string, data ...interface{}) {
	dispatch(fmt.Sprintf("[X][%s][Error]%s", time.Now().UTC().Format(time.RFC3339), message), data)
}

// Debug log
func Debug(message string, data ...interface{}) {
	dispatch(fmt.Sprintf("[.][%s]%s", time.Now().UTC().Format(time.RFC3339), message), data)
}

func dispatch(m string, data interface{}) {
	r, _ := json.Marshal(data)
	fmt.Println(m, string(r))
}
