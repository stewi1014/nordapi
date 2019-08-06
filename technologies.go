package nordapi

// Technology is a feature that a NordVPNserver can support.
type Technology int

// NordVPN server feature IDs
const (
	FeatIkev2              Technology = 1  // Ikev2/IPSec
	FeatOpenvpnUDP         Technology = 3  // OpenVPN UDP
	FeatOpenvpnTCP         Technology = 5  // OpenVPN TCP
	FeatSocks              Technology = 7  // Socks 5 proxy
	FeatProxy              Technology = 9  // HTTP Proxy
	FeatPptp               Technology = 11 // Pptp
	FeatL2Tp               Technology = 13 // L2TP/IPSec
	FeatOpenvpnXorUDP      Technology = 15 // OpenVPN UDP Obfuscated
	FeatOpenvpnXorTCP      Technology = 17 // OpenVPN TCP Obfuscated
	FeatProxyCybersec      Technology = 19 // HTTP CyberSec Proxy
	FeatProxySsl           Technology = 21 // HTTP Proxy (SSL)
	FeatProxySslCybersec   Technology = 23 // HTTP CyberSec Proxy (SSL)
	FeatIkev2V6            Technology = 26 // IKEv2/IPSec IPv6
	FeatOpenvpnUDPV6       Technology = 29 // OpenVPN UDP IPv6
	FeatOpenvpnTCPV6       Technology = 32 // OpenVPN TCP IPv6
	FeatWireguardUDP       Technology = 35 // Wireguard
	FeatOpenvpnUDPTLSCrypt Technology = 38 // OpenVPN UDP TLS Crypt
	FeatOpenvpnTCPTLSCrypt Technology = 41 // OpenVPN TCP TLS Crypt
)

var featureIdentifiers = map[Technology]string{
	1:  "ikev2",
	3:  "openvpn_udp",
	5:  "openvpn_tcp",
	7:  "socks",
	9:  "proxy",
	11: "pptp",
	13: "l2tp",
	15: "openvpn_xor_udp",
	17: "openvpn_xor_tcp",
	19: "proxy_cybersec",
	21: "proxy_ssl",
	23: "proxy_ssl_cybersec",
	26: "ikev2_v6",
	29: "openvpn_udp_v6",
	32: "openvpn_tcp_v6",
	35: "wireguard_udp",
	38: "openvpn_udp_tls_crypt",
	41: "openvpn_tcp_tls_crypt",
}

// GetFilter implements Filter
func (f Technology) GetFilter() string {
	return "filters[servers_technologies][identifier]=" + featureIdentifiers[f]
}
