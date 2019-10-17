package assignment1

import (
	"fmt"
	"net/http"
)

// HandlerNil - comment - - - - -
func HandlerNil(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Default Handler: Invalid request received.")
	http.Error(w, "Invalid request", http.StatusBadRequest)
}
