package nordapi

import "encoding/json"

// Servers fetches a complete server list from NordVPN
// This is a large list, and shouldn't be called more than once.
// I reccomend using MarshalJSON and UnmarshalJSON for caching.
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

// ServerInfo contains a list of NordVPN servers with helper methods for searching them.
type ServerInfo []Server

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
	Load     int `json:"load"`
	Features struct {
		Ikev2              bool `json:"ikev2"`
		OpenvpnUDP         bool `json:"openvpn_udp"`
		OpenvpnTCP         bool `json:"openvpn_tcp"`
		Socks              bool `json:"socks"`
		Proxy              bool `json:"proxy"`
		Pptp               bool `json:"pptp"`
		L2Tp               bool `json:"l2tp"`
		OpenvpnXorUDP      bool `json:"openvpn_xor_udp"`
		OpenvpnXorTCP      bool `json:"openvpn_xor_tcp"`
		ProxyCybersec      bool `json:"proxy_cybersec"`
		ProxySsl           bool `json:"proxy_ssl"`
		ProxySslCybersec   bool `json:"proxy_ssl_cybersec"`
		Ikev2V6            bool `json:"ikev2_v6"`
		OpenvpnUDPV6       bool `json:"openvpn_udp_v6"`
		OpenvpnTCPV6       bool `json:"openvpn_tcp_v6"`
		WireguardUDP       bool `json:"wireguard_udp"`
		OpenvpnUDPTLSCrypt bool `json:"openvpn_udp_tls_crypt"`
		OpenvpnTCPTLSCrypt bool `json:"openvpn_tcp_tls_crypt"`
	} `json:"features"`
}
