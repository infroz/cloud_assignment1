package assignment1

import (
	"log"
	"net/http"
	"time"
)

var StartTime time.Time

func Uptime() time.Duration {
	return time.Since(StartTime)
}

// HandlerDiag - Handles Get requests for /diag/
func HandlerDiag(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		log.Println(Uptime())
	default:
		// For any other method than GET - undefined
		http.Error(w, "not implemented yet", http.StatusNotImplemented)
		return
	}
}
