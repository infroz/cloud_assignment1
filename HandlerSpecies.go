package assignment1

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// HandlerSpecies Handler for species GET
func HandlerSpecies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 5 {
			status := http.StatusBadRequest
			http.Error(w, "Expecting format .../conservation/v1/species/", status)
			return
		}
		// API URL
		getAPI := "http://api.gbif.org/v1/species/" + parts[4]

		Client := http.DefaultClient
		resp := GetRequest(Client, getAPI)
		log.Println("Get Request: " + getAPI)

		var s SpecificSpeciesTmp
		err := json.NewDecoder(resp.Body).Decode(&s)
		if err != nil {
			log.Fatalln(err)
		}
		// API URL - gets year
		getAPI = "http://api.gbif.org/v1/species/" + parts[4] + "/name"
		respYear := GetRequest(Client, getAPI)
		log.Println("Get Request: " + getAPI)

		var sYear SpecificSpeciesYearTmp
		err = json.NewDecoder(respYear.Body).Decode(&sYear)
		if err != nil {
			log.Fatalln(err)
		}

		// Sets values for SpecificSpecies
		var response SpecificSpecies
		response.Key = s.Key
		response.Kingdom = s.Kingdom
		response.Phylum = s.Phylum
		response.Order = s.Order
		response.Family = s.Family
		response.Genus = s.Genus
		response.ScientificName = s.ScientificName
		response.CanonicalName = s.CanonicalName
		response.Year = sYear.BracketYear

		// Encode new structure to JSON format
		enc, err := json.Marshal(response)
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
