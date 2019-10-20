package assignment1

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
	Results []SpeciesStruct
}

// SpeciesByCountry Contains formated data for Species by Country
type SpeciesByCountry struct {
	Code        string
	CountryName string
	CountryFlag string
	Species     []string
	SpeciesKey  []int
}

// Diag - Stores data for HandlerDiag
type Diag struct {
	Gbif          int
	Restcountries int
	Version       string
	Uptime        int
}

// SpecificSpecies Contains formated data for /species/
type SpecificSpecies struct {
	Key            string
	Kingdom        string
	Phylum         string
	Order          string
	Family         string
	Genus          string
	ScientificName string
	CanonicalName  string
	Year           string
}
