package assignment1

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// HandlerCountry - - -s
func HandlerCountry(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		status := http.StatusBadRequest
		http.Error(w, "Expecting format .../countryName", status)
		return
	}
	country := parts[2]

	fmt.Fprintf(w, country)

	countryAPI := "https://restcountries.eu/rest/v2/name/"
	fmt.Println("Running GET: " + countryAPI + country)
	req, err := http.NewRequest(http.MethodGet, countryAPI+country, nil)
	if err != nil {
		panic(err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(body))
}
