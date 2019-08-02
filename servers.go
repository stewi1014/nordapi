package nordapi

import (
	"encoding/json"
	"sort"

	"github.com/golang/geo/s2"
)

// ServerInfo contains a list of NordVPN servers with helper methods for searching them.
type ServerInfo []Server

// Servers fetches a complete server list from NordVPN
func Servers() (ServerInfo, error) {
	resp, err := client.Get("https://api.nordvpn.com/server")
	if err != nil {
		return nil, err
	}

	var si ServerInfo
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&si)
	return si, err
}

// Country returns the servers in a given country by it's name; e.g. "United Kingdom", "Germany" etc...
func (si ServerInfo) Country(name string) ServerInfo {
	var nsi ServerInfo
	for i := range si {
		if si[i].Country == name {
			nsi = append(nsi, si[i])
		}
	}
	return nsi
}

// CountryCode returns the servers in a given country; e.g. "UK", "DE" etc...
func (si ServerInfo) CountryCode(code string) ServerInfo {
	var nsi ServerInfo
	for i := range si {
		if si[i].Flag == code {
			nsi = append(nsi, si[i])
		}
	}
	return nsi
}

// Features returns the servers supporting the specified features.
func (si ServerInfo) Features(features Features) ServerInfo {
	var nsi ServerInfo
	for i := range si {
		if si[i].Features.HasFeatures(features) {
			nsi = append(nsi, si[i])
		}
	}
	return nsi
}

// SortByLoad sorts the server list from least-loaded to most-loaded
func (si ServerInfo) SortByLoad() ServerInfo {
	sort.Slice(si, func(i, j int) bool {
		return si[i].Load < si[j].Load
	})
	return si
}

// SortByDistance sorts the server list by distance to the given coordinate
func (si ServerInfo) SortByDistance(position s2.LatLng) ServerInfo {
	distances := make([]float64, len(si))
	for i := range si {
		distances[i] = si[i].LatLng().Distance(position).Degrees()
	}

	sort.Slice(si, func(i, j int) bool {
		return distances[i] < distances[j]
	})

	return si
}
