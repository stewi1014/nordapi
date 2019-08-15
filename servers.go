package nordapi

import (
	"fmt"
	"strconv"
)

// ErrServerNotFound is returned when no servers are found.
type ErrServerNotFound struct {
	Filters []Filter
	URL     string
}

// Error implements error
func (e ErrServerNotFound) Error() string {
	if len(e.Filters) > 0 {
		return fmt.Sprintf("server not found, filters: %v", e.Filters)
	}
	return "server not found"
}

// String implements fmt.Stringer
func (e ErrServerNotFound) String() string {
	if e.URL == "" {
		return "error server not found"
	}
	return fmt.Sprintf("error no servers found at \"%v\"", e.URL)
}

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
	var servers ServerList
	f := FilterList(filters).GetFilter()

	url := "https://api.nordvpn.com/v1/servers/recommendations"
	if f != "" {
		url += "?" + f + "&limit=" + s
	} else {
		url += "?limit=" + s
	}

	err := getAndUnmarshall(url, &servers)
	if err != nil {
		return nil, err
	}
	if len(servers) == 0 {
		return nil, ErrServerNotFound{Filters: filters}
	}
	return servers, nil
}

// Hostname returns the server with the given hostname
func (sl ServerList) Hostname(hostname string) (*Server, error) {
	for i := range sl {
		if sl[i].Hostname == hostname {
			return sl[i], nil
		}
	}
	return nil, ErrServerNotFound{}
}

// Filter filters servers satisfying the given filters.
func (sl ServerList) Filter(filters ...Filter) ServerList {
	fl := FilterList(filters)
	var nsl ServerList
	for _, s := range sl {
		if fl.Satisfies(s) {
			nsl = append(nsl, s)
		}
	}
	return nsl
}
