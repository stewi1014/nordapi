package nordapi

import "sync"

// Technology is a feature that a NordVPNserver can support.
type Technology struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

// Ikev2/IPSec
// OpenVPN UDP
// OpenVPN TCP
// Socks 5 proxy
// HTTP Proxy
// Pptp
// L2TP/IPSec
// OpenVPN UDP Obfuscated
// OpenVPN TCP Obfuscated
// HTTP CyberSec Proxy
// HTTP Proxy (SSL)
// HTTP CyberSec Proxy (SSL)
// IKEv2/IPSec IPv6
// OpenVPN UDP IPv6
// OpenVPN TCP IPv6
// Wireguard
// OpenVPN UDP TLS Crypt
// OpenVPN TCP TLS Crypt

// NordVPN server feature IDs
var (
	TechIkev2              = &Technology{ID: 1, Name: "IKEv2/IPSec", Identifier: "ikev2"}                             // Ikev2/IPSec
	TechOpenvpnUDP         = &Technology{ID: 3, Name: "OpenVPN UDP", Identifier: "openvpn_udp"}                       // OpenVPN UDP
	TechOpenvpnTCP         = &Technology{ID: 5, Name: "OpenVPN TCP", Identifier: "openvpn_tcp"}                       // OpenVPN TCP
	TechSocks              = &Technology{ID: 7, Name: "Socks 5", Identifier: "socks"}                                 // Socks 5 proxy
	TechProxy              = &Technology{ID: 9, Name: "HTTP Proxy", Identifier: "proxy"}                              // HTTP Proxy
	TechPptp               = &Technology{ID: 11, Name: "PPTP", Identifier: "pptp"}                                    // Pptp
	TechL2Tp               = &Technology{ID: 13, Name: "L2TP/IPSec", Identifier: "l2tp"}                              // L2TP/IPSec
	TechOpenvpnXorUDP      = &Technology{ID: 15, Name: "OpenVPN UDP Obfuscated", Identifier: "openvpn_xor_udp"}       // OpenVPN UDP Obfuscated
	TechOpenvpnXorTCP      = &Technology{ID: 17, Name: "OpenVPN TCP Obfuscated", Identifier: "openvpn_xor_tcp"}       // OpenVPN TCP Obfuscated
	TechProxyCybersec      = &Technology{ID: 19, Name: "HTTP CyberSec Proxy", Identifier: "proxy_cybersec"}           // HTTP CyberSec Proxy
	TechProxySsl           = &Technology{ID: 21, Name: "HTTP Proxy (SSL)", Identifier: "proxy_ssl"}                   // HTTP Proxy (SSL)
	TechProxySslCybersec   = &Technology{ID: 23, Name: "HTTP CyberSec Proxy (SSL)", Identifier: "proxy_ssl_cybersec"} // HTTP CyberSec Proxy (SSL)
	TechIkev2V6            = &Technology{ID: 26, Name: "IKEv2/IPSec IPv6", Identifier: "ikev2_v6"}                    // IKEv2/IPSec IPv6
	TechOpenvpnUDPV6       = &Technology{ID: 29, Name: "OpenVPN UDP IPv6", Identifier: "openvpn_udp_v6"}              // OpenVPN UDP IPv6
	TechOpenvpnTCPV6       = &Technology{ID: 32, Name: "OpenVPN TCP IPv6", Identifier: "openvpn_tcp_v6"}              // OpenVPN TCP IPv6
	TechWireguardUDP       = &Technology{ID: 35, Name: "Wireguard", Identifier: "wireguard_udp"}                      // Wireguard
	TechOpenvpnUDPTLSCrypt = &Technology{ID: 38, Name: "OpenVPN UDP TLS Crypt", Identifier: "openvpn_udp_tls_crypt"}  // OpenVPN UDP TLS Crypt
	TechOpenvpnTCPTLSCrypt = &Technology{ID: 41, Name: "OpenVPN TCP TLS Crypt", Identifier: "openvpn_tcp_tls_crypt"}  // OpenVPN TCP TLS Crypt
)

// Avoid race conditions if two routines acess knownTechnologies
var knownTechnologiesMutex sync.Mutex

// Users should lock knownTechnologiesMutex
var knownTechnologies = []*Technology{
	TechIkev2,
	TechOpenvpnUDP,
	TechOpenvpnTCP,
	TechSocks,
	TechProxy,
	TechPptp,
	TechL2Tp,
	TechOpenvpnXorUDP,
	TechOpenvpnXorTCP,
	TechProxyCybersec,
	TechProxySsl,
	TechProxySslCybersec,
	TechIkev2V6,
	TechOpenvpnUDPV6,
	TechOpenvpnTCPV6,
	TechWireguardUDP,
	TechOpenvpnUDPTLSCrypt,
	TechOpenvpnTCPTLSCrypt,
}

// GetFilter implements Filter
func (t *Technology) GetFilter() string {
	return "filters[servers_technologies][identifier]=" + t.Identifier
}

// Satisfies implements Filter
func (t *Technology) Satisfies(s *Server) bool {
	for _, st := range s.Technologies {
		if st.ID == t.ID {
			return true
		}
	}
	return false
}

// String implements fmt.Stringer
// It returns the identifier string for the technology
func (t *Technology) String() string {
	return t.Identifier
}

// TechnologyIdentifier returns the Technology from its identifying string.
// If the Technology is not found, it will return nil.
func TechnologyIdentifier(identifier string) *Technology {
	knownTechnologiesMutex.Lock()
	defer knownTechnologiesMutex.Unlock()
	for _, t := range knownTechnologies {
		if t.Identifier == identifier {
			return t
		}
	}

	var technologies []*Technology
	err := getAndUnmarshall("https://api.nordvpn.com/v1/technologies", &technologies)
	if err != nil {
		return nil
	}

	var found *Technology

AlreadyHave:
	for i := range technologies {
		for j := range knownTechnologies {
			if technologies[i].Identifier == knownTechnologies[j].Identifier {
				continue AlreadyHave
			}
		}
		knownTechnologies = append(knownTechnologies, technologies[i])
		if technologies[i].Identifier == identifier {
			found = technologies[i]
		}
	}

	return found
}
