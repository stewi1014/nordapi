package nordapi

import (
	"encoding/json"
	"fmt"
	"sort"
	"sync"

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
		if si[i].HasFeatures(features) {
			nsi = append(nsi, si[i])
		}
	}
	return nsi
}

// Domain returns the server with the given domain
func (si ServerInfo) Domain(domain string) (Server, error) {
	for i := range si {
		if si[i].Domain == domain {
			return si[i], nil
		}
	}
	return Server{}, fmt.Errorf("Couldn't find server with domain %v", domain)
}

// SortByLoad sorts the server list from least-loaded to most-loaded
func (si ServerInfo) SortByLoad() ServerInfo {
	sort.SliceStable(si, func(i, j int) bool {
		return si[i].Load < si[j].Load
	})
	return si
}

// SortByDistance sorts the server list by distance to the given coordinate
// If you have a Latitude and Longitude in degrees, do SortByDistance(s2.LatLngFromDegrees(Latitude, Longitude))
func (si ServerInfo) SortByDistance(position s2.LatLng) ServerInfo {
	dist := make([]float64, len(si))
	for i := range si {
		dist[i] = si[i].LatLng().Distance(position).Degrees()
	}
	sort.Stable(paralellSorter{attr: dist, si: si})
	return si
}

// SortByPing sorts the server list by ping times, pinging each server n times
// and taking the average. Servers that do not reply are removed from the list.
func (si ServerInfo) SortByPing(n int) ServerInfo {
	pings := make([]float64, len(si)) //Negative ping value means we remove later
	wg := sync.WaitGroup{}
	wg.Add(len(si))
	for i := range si {
		go func(i int) {
			ping, err := si[i].PingTime(4)
			if err != nil {
				pings[i] = -1
			} else {
				pings[i] = ping.Seconds()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	sort.Stable(paralellSorter{attr: pings, si: si})
	return si
}

// Instead of recalculating server attributes all the time,
// I cache them in another slice, and sort it instead, swapping the server list
// alongside it. If the Sort function is even better than I think it is, this might just be a waste of a
// few KB of memory. If so, sue me (or submit a pull request).
type paralellSorter struct {
	attr []float64
	si   ServerInfo
}

func (d paralellSorter) Len() int           { return len(d.attr) }
func (d paralellSorter) Less(i, j int) bool { return d.attr[i] < d.attr[j] }
func (d paralellSorter) Swap(i, j int) {
	d.attr[i], d.attr[j] = d.attr[j], d.attr[i]
	d.si[i], d.si[j] = d.si[j], d.si[i]
}
