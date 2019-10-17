package main

import (
	"assignment1"
	"fmt"
	"log"
	"net/http"
	"os"
)

// -----------------

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	/* -- Handlers -- */

	http.HandleFunc("/", assignment1.HandlerNil)
	http.HandleFunc("/country/", assignment1.HandlerCountry)
	http.HandleFunc("/species/", assignment1.HandlerSpecies)

	fmt.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
