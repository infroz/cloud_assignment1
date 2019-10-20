package assignment1

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// StartTime - Set by init() in main.go
var StartTime time.Time

// Uptime - Returns times in seconds since service start
func Uptime() time.Duration {
	return time.Since(StartTime) / 1000000000
}

// HandlerDiag - Handles Get requests for /diag/
func HandlerDiag(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var d Diag
		// Get API Status

		resp, err := http.Get("http://restcountries.eu/rest/v2/")
		if err != nil {
			log.Fatalln(err)
		}

		d.Gbif = resp.StatusCode

		resp, err = http.Get("http://api.gbif.org/v1/")
		if err != nil {
			log.Fatalln(err)
		}

		d.Restcountries = resp.StatusCode // FakeOK
		d.Version = "v1"
		d.Uptime = int(Uptime())

		// Encode new structure to JSON format
		enc, err := json.Marshal(d)
		if err != nil {
			log.Fatalln(err)
		}

		// Gives JSON response for requests
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(enc)
	default:
		// For any other method than GET - undefined
		http.Error(w, "not implemented yet", http.StatusNotImplemented)
		return
	}
}
