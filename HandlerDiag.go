package assignment1

import "net/http"

func HandlerDiag(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		break
	default:
		// For any other method than GET - undefined
		http.Error(w, "not implemented yet", http.StatusNotImplemented)
		return
	}
}
