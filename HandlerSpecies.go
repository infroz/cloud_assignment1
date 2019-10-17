package assignment1

import (
	"fmt"
	"net/http"
	"strings"
)

func HandlerSpecies(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		status := http.StatusBadRequest
		http.Error(w, "Expecting format .../species", status)
		return
	}
	name := parts[2]
	fmt.Fprintln(w, parts)
}
