package assignment1

import (
	"fmt"
	"net/http"
	"strings"
)

// HandlerSpecies Handler for species GET
func HandlerSpecies(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		status := http.StatusBadRequest
		http.Error(w, "Expecting format .../species", status)
		return
	}

	//speciesAPI := ""

	fmt.Fprintln(w, parts)
}
