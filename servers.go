package nordapi

import (
	"errors"
	"strconv"
)

// ErrServerNotFound is returned if the server cannot be found.
var ErrServerNotFound = errors.New("Server not found")

// ServerList is a list of NordVPN servers
type ServerList []*Server

// Servers returns a complete list of NordVPN servers.
func Servers() (ServerList, error) {
	var servers []Server
	err := getAndUnmarshall("https://api.nordvpn.com/v1/servers?limit=16384", &servers)
	if err != nil {
		return nil, err
	}
	sl := make(ServerList, len(servers))
	for i := range servers {
		sl[i] = &servers[i]
	}
	return sl, nil
}

// Reccomended returns the top n recomended servers, filtered by filters.
func Reccomended(n int, filters ...Filter) (ServerList, error) {
	s := strconv.Itoa(n)
	var servers []Server
	fl := FilterList(filters)
	f := fl.GetFilter()
	url := "https://api.nordvpn.com/v1/servers/recommendations"
	if len(f) > 0 {
		url += "?" + f + "&limit=" + s
	} else {
		url += "?limit=" + s
	}
	err := getAndUnmarshall(url, &servers)
	if err != nil {
		return nil, err
	}
	if len(servers) == 0 {
		return nil, ErrServerNotFound
	}
	sl := make(ServerList, len(servers))
	for i := range servers {
		sl[i] = &servers[i]
	}
	return sl, nil
}

// Hostname returns the server with the given hostname
func (sl ServerList) Hostname(hostname string) (*Server, error) {
	for i := range sl {
		if sl[i].Hostname == hostname {
			return sl[i], nil
		}
	}
	return nil, ErrServerNotFound
}
