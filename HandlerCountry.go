package assignment1

// I have used log instead of fmt since I think it is more readable

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// HandlerCountry - This function handles country requests
func HandlerCountry(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 5 {
			status := http.StatusBadRequest
			http.Error(w, "Expecting format .../conservation/v1/country/", status)
			return
		}

		Limit := r.URL.Query().Get("limit")
		if Limit == "" {
			Limit = "30"
		}

		// Getting JSON from GBIF
		getAPI := "http://api.gbif.org/v1/occurrence/search?country=" + parts[4] + "&limit=" + Limit

		Client := http.DefaultClient
		resp := GetRequest(Client, getAPI)

		var s RestGBIFTmp
		err := json.NewDecoder(resp.Body).Decode(&s)
		if err != nil {
			log.Fatalln(err)
		}

		// Getting JSON from Restcountries
		getAPI2 := "http://restcountries.eu/rest/v2/alpha/" + parts[4] + "?fields=name;alpha2Code;flag"
		resp2 := GetRequest(Client, getAPI2)

		var m RestCountryTmp
		err = json.NewDecoder(resp2.Body).Decode(&m)
		if err != nil {
			log.Fatalln(err)
		}

		// Create new structure
		var res SpeciesByCountry
		res.Code = m.Alpha2Code
		res.CountryName = m.Name
		res.CountryFlag = m.Flag
		for i := range s.Results {
			res.Species = append(res.Species, s.Results[i].Species)
			res.SpeciesKey = append(res.SpeciesKey, s.Results[i].SpeciesKey)
		}

		// Encode new structure to JSON format
		enc, err := json.Marshal(res)
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

// GetRequest - Comment
func GetRequest(c *http.Client, s string) *http.Response {
	req, err := http.NewRequest("GET", s, nil)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}
