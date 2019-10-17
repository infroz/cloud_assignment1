package assignment1

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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

	countryAPI := "https://restcountries.eu/rest/v2/alpha/"
	fmt.Println("Running GET: " + countryAPI + country)
	req, err := http.NewRequest(http.MethodGet, countryAPI+country+"?fields=name;alpha2Code;flag", nil)
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
	fmt.Fprintf(w, string(body)+"\n\nDecoded Data:\n")

	dec := json.NewDecoder(strings.NewReader(string(body)))

	var m CountryJSON
	if err := dec.Decode(&m); err == io.EOF {

	} else if err != nil {
		log.Fatal(err)
	}

	// Prints to console
	fmt.Println("GET \"" + country + "\"=\"" + m.Name + "\", Alpha=\"" + m.Alpha2Code + "\" FlagURL=\"" + m.Flag + "\"")
	fmt.Fprintf(w, "<img href=\""+m.Flag+"\"> </img><h1>test</h1>")
}
