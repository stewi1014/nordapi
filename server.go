package nordapi

import (
	"github.com/golang/geo/s2"
)

// Server is a NordVPN server
type Server struct {
	ID             int      `json:"id"`
	IPAddress      string   `json:"ip_address"`
	SearchKeywords []string `json:"search_keywords"`
	Categories     []struct {
		Name string `json:"name"`
	} `json:"categories"`
	Name     string `json:"name"`
	Domain   string `json:"domain"`
	Price    int    `json:"price"`
	Flag     string `json:"flag"`
	Country  string `json:"country"`
	Location struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	} `json:"location"`
	Load     int      `json:"load"`
	Features Features `json:"features"`
}

// LatLng returns the servers location (in Server.Location) as a s2.LatLng type
func (s Server) LatLng() s2.LatLng {
	return s2.LatLngFromDegrees(s.Location.Lat, s.Location.Long)
}
