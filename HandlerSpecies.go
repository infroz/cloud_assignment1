package assignment1

import (
	"fmt"
	"net/http"
	"strings"
)

// HandlerSpecies Handler for species GET
func HandlerSpecies(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 5 {
		status := http.StatusBadRequest
		http.Error(w, "Expecting format .../conservation/v1/species/", status)
		return
	}
	// Get limit on results
	Limit := r.URL.Query().Get("limit")
	if Limit == "" {
		Limit = "30" // Default limit
	}
	// API URL
	getAPI := "http://api.gbif.org/v1/occurrence/search?country=" + parts[4] + "&limit=" + Limit

	Client := http.DefaultClient
	resp := GetRequest(Client, getAPI)

	fmt.Fprintln(w, parts)
}
