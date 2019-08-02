package nordapi

import (
	"encoding/json"
)

// Countries returns a list of countries which NordVPN has servers in.
func Countries() ([]Country, error) {
	resp, err := client.Get("https://api.nordvpn.com/v1/servers/countries")
	if err != nil {
		return nil, err
	}

	var countries []Country
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&countries)

	return countries, err
}

// Country is a country in which NordVPN has server(s)
type Country struct {
	ID          int
	Name        string
	CountryCode string
	Cities      []City
}

// City is a city in which NordVPN has server(s)
type City struct {
	ID        int
	Name      string
	Latitude  float64
	Longitude float64
	DNSName   string
}
