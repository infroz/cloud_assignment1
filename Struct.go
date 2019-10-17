package assignment1

// CountryJSON - Datastructure for JSON from restcountries
type CountryJSON struct {
	Name       string `json:"name"`
	Alpha2Code string `json:"alpha2Code"`
	Flag       string `json:"flag"`
}
