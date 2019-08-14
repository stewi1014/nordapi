package nordapi

import (
	"io"
)

// Server is a NordVPN server.
type Server struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	Station   string `json:"station"`
	Hostname  string `json:"hostname"`
	Load      int    `json:"load"`
	Status    string `json:"status"`
	Locations []struct {
		ID        int     `json:"id"`
		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Country   Country `json:"country"`
	} `json:"locations"`
	Services []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Identifier string `json:"identifier"`
		CreatedAt  string `json:"created_at"`
		UpdatedAt  string `json:"updated_at"`
	} `json:"services"`
	Technologies []struct {
		ID         Technology `json:"id"`
		Name       string     `json:"name"`
		Identifier string     `json:"identifier"`
		CreatedAt  string     `json:"created_at"`
		UpdatedAt  string     `json:"updated_at"`
		Pivot      struct {
			TechnologyID int    `json:"technology_id"`
			ServerID     int    `json:"server_id"`
			Status       string `json:"status"`
		} `json:"pivot"`
	} `json:"technologies"`
	Groups []struct {
		ID        Group  `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Title     string `json:"title"`
		Type      struct {
			ID         int    `json:"id"`
			CreatedAt  string `json:"created_at"`
			UpdatedAt  string `json:"updated_at"`
			Title      string `json:"title"`
			Identifier string `json:"identifier"`
		} `json:"type"`
	} `json:"groups"`
	Specifications []struct {
		ID         int    `json:"id"`
		Title      string `json:"title"`
		Identifier string `json:"identifier"`
		Values     []struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"values"`
	} `json:"specifications"`
	Ips []struct {
		ID        int    `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		ServerID  int    `json:"server_id"`
		IPID      int    `json:"ip_id"`
		Type      string `json:"type"`
		IP        struct {
			ID      int    `json:"id"`
			IP      string `json:"ip"`
			Version int    `json:"version"`
		} `json:"ip"`
	} `json:"ips"`
}

// Hostname returns the server with the given hostname.
func Hostname(hostname string) (*Server, error) {
	servers, err := Servers()
	if err != nil {
		return nil, err
	}
	return servers.Hostname(hostname)
}

// OpenvpnUDPConfig returns the UDP port 1194 OpenVPN configuration for the server.
func (s *Server) OpenvpnUDPConfig() (io.ReadCloser, error) {
	if s.Hostname == "" {
		return nil, ErrServerNotFound
	}

	resp, err := client.Get(
		"https://downloads.nordcdn.com/configs/files/ovpn_legacy/servers/" +
			s.Hostname +
			".udp1194.ovpn")

	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// OpenvpnTCPConfig returns the TCP port 443 OpenVPN configuration for the server.
func (s *Server) OpenvpnTCPConfig() (io.ReadCloser, error) {
	if s.Hostname == "" {
		return nil, ErrServerNotFound
	}

	resp, err := client.Get(
		"https://downloads.nordcdn.com/configs/files/ovpn_legacy/servers/" +
			s.Hostname +
			".tcp443.ovpn")

	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

// Satisfies returns true if the server satisfies the given filters
func (s *Server) Satisfies(filters ...Filter) bool {
	fl := FilterList(filters)
	return fl.Satisfies(s)
}
