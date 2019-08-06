package nordapi

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

// ErrCountryNotFound is returned if a search finds no results.
var ErrCountryNotFound = errors.New("Country not found")

// City is a city in which NordVPN has server(s)
// Unfortunately I cannot figure out how to filter by cities, so you'll just have to make do without
// or submit a pull request.
type City struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	DNSName   string  `json:"dns_name"`
}

// Country is a country in which NordVPN has server(s).
type Country struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Cities []City `json:"cities"`
}

// GetFilter implements Filter
func (c Country) GetFilter() string {
	return "filters[country_id]=" + strconv.Itoa(c.ID)
}

// CountryList is a list of countries.
type CountryList []Country

// Countries returns a list of countries which NordVPN has servers in.
func Countries() (CountryList, error) {
	resp, err := client.Get("https://api.nordvpn.com/v1/servers/countries")
	if err != nil {
		return nil, err
	}

	var countries []Country
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&countries)

	return countries, err
}

// Name returns the country with the given name.
// See ErrCountryNotFound.
func (cl CountryList) Name(name string) (Country, error) {
	name = strings.ToLower(name)
	for i := range cl {
		if strings.ToLower(cl[i].Name) == name {
			return cl[i], nil
		}
	}
	return Country{}, ErrCountryNotFound
}

// Code returns the country with the given country code. All-Caps.
// See ErrCountryNotFound.
func (cl CountryList) Code(code string) (Country, error) {
	for i := range cl {
		if cl[i].Code == code {
			return cl[i], nil
		}
	}
	return Country{}, ErrCountryNotFound
}

// CityName returns the country which has the city with the name name.
// See ErrCountryNotFound.
func (cl CountryList) CityName(name string) (Country, error) {
	for i := range cl {
		for j := range cl[i].Cities {
			if strings.ToLower(cl[i].Cities[j].Name) == name {
				return cl[i], nil
			}
		}
	}
	return Country{}, ErrCountryNotFound
}

// CityID returns the country with the given city ID
// See ErrCountryNotFound.
func (cl CountryList) CityID(id int) (Country, error) {
	for i := range cl {
		for j := range cl[i].Cities {
			if cl[i].Cities[j].ID == id {
				return cl[i], nil
			}
		}
	}
	return Country{}, ErrCountryNotFound
}