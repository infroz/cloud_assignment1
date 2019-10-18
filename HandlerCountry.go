package assignment1

import (
	"log"
	"net/http"
	"strings"
	"encoding/json"
	"strconv"
)

// HandlerCountry - This function handles country requests
func HandlerCountry(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		status := http.StatusBadRequest
		http.Error(w, "Expecting format .../country/", status)
		return
	}
	getAPI := "http://api.gbif.org/v1/occurrence/search?country=" + parts[2] + "&limit=" + strconv.Itoa(Limit)

	Client := http.DefaultClient
	resp := GetRequest(Client, getAPI)

	var s RestGBIFTmp
	err := json.NewDecoder(resp.Body).Decode(&s)
	if err != nil {
		log.Fatalln(err)
	}
	for i:= 0; i < Limit; i++ {
		log.Println(s.Results[i].Species)
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
