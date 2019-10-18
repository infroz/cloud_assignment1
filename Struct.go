package assignment1

// Limit is the limit on searches from GBIF
const Limit = 100

// SpeciesStruct Contains data for Species
type SpeciesStruct struct {
	SpeciesKey int
	Species    string
}

// RestCountryTmp Collects temporary data from restCountry
type RestCountryTmp struct {
	Name       string
	Alpha2Code string
	Flag       string
}

// RestGBIFTmp Stores Results
type RestGBIFTmp struct {
	Results [Limit]SpeciesStruct
}

// SpeciesByCountry Contains formated data for Species by Country
type SpeciesByCountry struct {
	Code        string
	CountryName string
	CountryFlag string
	Species     [Limit]string
	SpeciesKey  [Limit]int
}

// Diag - Stores data for HandlerDiag
type Diag struct {
	Gbif          int
	Restcountries int
	version       string
	uptime        int
}
