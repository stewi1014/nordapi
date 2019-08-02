package nordapi

// Features are the features a server supports
type Features struct {
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
}

// HasFeatures returns true if the feature set supports the given feature set.
// It simply checks if all features in check are true in f
func (f Features) HasFeatures(check Features) bool {
	if check.Ikev2 && !f.Ikev2 {
		return false
	}
	if check.OpenvpnUDP && !f.OpenvpnUDP {
		return false
	}
	if check.OpenvpnTCP && !f.OpenvpnTCP {
		return false
	}
	if check.Socks && !f.Socks {
		return false
	}
	if check.Proxy && !f.Proxy {
		return false
	}
	if check.Pptp && !f.Pptp {
		return false
	}
	if check.L2Tp && !f.L2Tp {
		return false
	}
	if check.OpenvpnXorUDP && !f.OpenvpnXorUDP {
		return false
	}
	if check.OpenvpnXorTCP && !f.OpenvpnXorTCP {
		return false
	}
	if check.ProxyCybersec && !f.ProxyCybersec {
		return false
	}
	if check.ProxySsl && !f.ProxySsl {
		return false
	}
	if check.ProxySslCybersec && !f.ProxySslCybersec {
		return false
	}
	if check.Ikev2V6 && !f.Ikev2V6 {
		return false
	}
	if check.OpenvpnUDPV6 && !f.OpenvpnUDPV6 {
		return false
	}
	if check.OpenvpnTCPV6 && !f.OpenvpnTCPV6 {
		return false
	}
	if check.WireguardUDP && !f.WireguardUDP {
		return false
	}
	if check.OpenvpnUDPTLSCrypt && !f.OpenvpnUDPTLSCrypt {
		return false
	}
	if check.OpenvpnTCPTLSCrypt && !f.OpenvpnTCPTLSCrypt {
		return false
	}
	return true
}