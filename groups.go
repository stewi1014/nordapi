package nordapi

import (
	"sync"
)

// Group types.
var (
	GroupDoubleVPN                   = &Group{ID: 1, Identifier: "legacy_double_vpn", Title: "Double VPN"}
	GroupOnionOverVPN                = &Group{ID: 3, Identifier: "legacy_onion_over_vpn", Title: "Onion Over VPN"}
	GroupUltraFastTV                 = &Group{ID: 5, Identifier: "legacy_ultra_fast_tv", Title: "Ultra fast TV"}
	GroupAntiDDOS                    = &Group{ID: 7, Identifier: "legacy_anti_ddos", Title: "Anti DDoS"}
	GroupDedicatedIP                 = &Group{ID: 9, Identifier: "legacy_dedicated_ip", Title: "Dedicated IP"}
	GroupStandard                    = &Group{ID: 11, Identifier: "legacy_standard", Title: "Standard VPN servers"}
	GroupNetflixUSA                  = &Group{ID: 13, Identifier: "legacy_netflix_usa", Title: "Netflix USA"}
	GroupP2P                         = &Group{ID: 15, Identifier: "legacy_p2p", Title: "P2P"}
	GroupObfuscatedServers           = &Group{ID: 17, Identifier: "legacy_obfuscated_servers", Title: "Obfuscated Servers"}
	GroupEurope                      = &Group{ID: 19, Identifier: "europe", Title: "Europe"}
	GroupTheAmericas                 = &Group{ID: 21, Identifier: "the_americas", Title: "The Americas"}
	GroupAsiaPacific                 = &Group{ID: 23, Identifier: "asia_pacific", Title: "Asia Pacific"}
	GroupAfricaTheMiddleEastAndIndia = &Group{ID: 25, Identifier: "africa_the_middle_east_and_india", Title: "Africa, the Middle East and India"}
)

// Avoid race condition if two routines want to acess knownGroups
var knownGroupsMutex sync.Mutex

// Users should lock knownGroupsMutex!!
var knownGroups = []*Group{
	GroupDoubleVPN,
	GroupOnionOverVPN,
	GroupUltraFastTV,
	GroupAntiDDOS,
	GroupDedicatedIP,
	GroupStandard,
	GroupNetflixUSA,
	GroupP2P,
	GroupObfuscatedServers,
	GroupEurope,
	GroupTheAmericas,
	GroupAsiaPacific,
	GroupAfricaTheMiddleEastAndIndia,
}

// Group is a NordVPN server group
type Group struct {
	ID         int    `json:"id"`
	Identifier string `json:"identifier"`
	Title      string `json:"title"`
}

// GroupIdentifier returns a group given its identifier
// If the group is not found, it returns nil
func GroupIdentifier(identifier string) *Group {
	knownGroupsMutex.Lock()
	defer knownGroupsMutex.Unlock()
	for _, g := range knownGroups {
		if g.Identifier == identifier {
			return g
		}
	}

	var groups []*Group
	err := getAndUnmarshall("https://api.nordvpn.com/v1/servers/groups", &groups)
	if err != nil {
		return nil
	}

	var found *Group

AlreadyHave:
	for i := range groups {
		for j := range knownGroups {
			if groups[i].Identifier == knownGroups[j].Identifier {
				continue AlreadyHave
			}
		}
		knownGroups = append(knownGroups, groups[i])
		if groups[i].Identifier == identifier {
			found = groups[i]
		}
	}

	return found
}

// Groups returns a list of groups
func Groups() ([]*Group, error) {
	var groups []*Group
	err := getAndUnmarshall("https://api.nordvpn.com/v1/servers/groups", &groups)
	if err != nil {
		return nil, err
	}

AlreadyHave:
	for i := range groups {
		for j := range knownGroups {
			if groups[i].Identifier == knownGroups[j].Identifier {
				continue AlreadyHave
			}
		}
		knownGroups = append(knownGroups, groups[i])
	}

	return groups, nil
}

// GetFilter implements Filter
func (g *Group) GetFilter() string {
	return "filters[servers_groups][identifier]=" + g.Identifier
}

// Satisfies implements Filter
func (g *Group) Satisfies(s *Server) bool {
	for _, sg := range s.Groups {
		if sg.ID == g.ID {
			return true
		}
	}
	return false
}

// String implements fmt.Stringer
// It returns the identifier string for the group
func (g *Group) String() string {
	return g.Identifier
}
