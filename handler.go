package assignment1

import (
	"fmt"
	"net/http"
	"strings"
)

// HandlerNil - comment - - - - -
func HandlerNil(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Default Handler: Invalid request received.")
	http.Error(w, "Invalid request", http.StatusBadRequest)
}

// HandlerCountry - - -s
func HandlerCountry(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		status := http.StatusBadRequest
		http.Error(w, "Expecting format .../firstname/lastname", status)
		return
	}
	country := parts[2]

	fmt.Fprintf(w, country)

}
