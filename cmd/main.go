package main

import (
	"assignment1"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// -----------------
// Sets system start time - used in uptime
func init() {
	assignment1.StartTime = time.Now()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	/* -- Handlers -- */

	http.HandleFunc("/", assignment1.HandlerNil)
	http.HandleFunc("/conservation/v1/country/", assignment1.HandlerCountry)
	http.HandleFunc("/conservation/v1/species/", assignment1.HandlerSpecies)
	http.HandleFunc("/conservation/v1/diag/", assignment1.HandlerDiag)

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
