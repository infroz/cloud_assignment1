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

// HandlerCountry - This function handles country requests
func HandlerCountry(w http.ResponseWriter, r *http.Request) {
	// Splits url into readable parts
	// url/[1]Country/[2]{CountryCode}
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		status := http.StatusBadRequest
		http.Error(w, "Expecting format .../countryName", status)
		return
	}
	country := parts[2] // Sets variable country to country code from GET

	// API link for alpha2Code
	countryAPI := "https://restcountries.eu/rest/v2/alpha/"

	// Print to Console for visualization of what is happening
	fmt.Println("Running GET: " + countryAPI + country)

	// Preparing GET statement
	req, err := http.NewRequest(http.MethodGet, countryAPI+country+"?fields=name;alpha2Code;flag", nil)
	if err != nil {
		panic(err) // Error handling
	}

	// Executing GET statement
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err) // Error handling
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body) // Sets JSON into body
	if err != nil {
		panic(err)
	}
	// Print out values
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
