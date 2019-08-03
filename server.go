package nordapi

import (
	"errors"
	"time"

	"github.com/golang/geo/s2"
	ping "github.com/sparrc/go-ping"
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

// LatLng returns the servers location (in Server.Location) as a s2.LatLng type.
func (s Server) LatLng() s2.LatLng {
	return s2.LatLngFromDegrees(s.Location.Lat, s.Location.Long)
}

// HasFeatures returns true if the server supports the passed features.
func (s Server) HasFeatures(features Features) bool {
	return s.Features.HasFeatures(features)
}

// ErrNoResponce is returned when no ping reply is received
var ErrNoResponce = errors.New("No ping reply was received")

// PingTime pings the server, calculating the average time from n pings.
func (s Server) PingTime(n int) (time.Duration, error) {
	pinger, err := ping.NewPinger(s.IPAddress)
	if err != nil {
		return 0, err
	}

	pinger.Count = n
	pinger.Timeout = time.Duration(n) * time.Second * 5
	pinger.Run()
	stats := pinger.Statistics()
	if stats.PacketsRecv == 0 {
		return 0, ErrNoResponce
	}

	return stats.AvgRtt, nil
}
