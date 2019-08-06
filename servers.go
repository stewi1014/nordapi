package nordapi

import (
	"strconv"
)

// ServerList is a list of NordVPN servers
type ServerList []Server

// Servers returns a complete list of NordVPN servers.
func Servers() (ServerList, error) {
	var sl ServerList
	return sl, getAndUnmarshall("https://api.nordvpn.com/v1/servers?limit=16384", &sl)
}

// Reccomended returns the top n recomended servers, filtered by filters.
func Reccomended(n int, filters ...Filter) (ServerList, error) {
	s := strconv.Itoa(n)
	var sl ServerList
	fl := FilterList(filters)
	f := fl.GetFilter()
	url := "https://api.nordvpn.com/v1/servers/recommendations"
	if len(f) > 0 {
		url += "?" + f + "&limit=" + s
	} else {
		url += "?limit=" + s
	}
	return sl, getAndUnmarshall(url, &sl)
}
