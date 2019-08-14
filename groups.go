package nordapi

// Group is a NordVPN server group
type Group int

// Group types.
const (
	GroupDoubleVPN                   Group = 1
	GroupOnionOverVPN                Group = 3
	GroupUltraFastTV                 Group = 5
	GroupAntiDDOS                    Group = 7
	GroupDedicatedIP                 Group = 9
	GroupStandard                    Group = 11
	GroupNetflixUSA                  Group = 13
	GroupP2P                         Group = 15
	GroupObfuscatedServers           Group = 17
	GroupEurope                      Group = 19
	GroupTheAmericas                 Group = 21
	GroupAsiaPacific                 Group = 23
	GroupAfricaTheMiddleEastAndIndia Group = 25
)

var groupIdentifiers = map[Group]string{
	1:  "legacy_double_vpn",
	3:  "legacy_onion_over_vpn",
	5:  "legacy_ultra_fast_tv",
	7:  "legacy_anti_ddos",
	9:  "legacy_dedicated_ip",
	11: "legacy_standard",
	13: "legacy_netflix_usa",
	15: "legacy_p2p",
	17: "legacy_obfuscated_servers",
	19: "europe",
	21: "the_americas",
	23: "asia_pacific",
	25: "africa_the_middle_east_and_india",
}

// GetFilter implements Filter
func (g Group) GetFilter() string {
	s, ok := groupIdentifiers[g]
	if !ok {
		return ""
	}
	return "filters[servers_groups][identifier]=" + s
}

// Satisfies implements Filter
func (g Group) Satisfies(s *Server) bool {
	for _, sg := range s.Groups {
		if sg.ID == g {
			return true
		}
	}
	return false
}

// String implements fmt.Stringer
// It returns the identifier string for the group
func (g Group) String() string {
	s, ok := groupIdentifiers[g]
	if !ok {
		return "unknown_group"
	}
	return s
}

// GroupIdentifier returns the group from its identifying string.
// If it is not found, Group == 0 will be true.
func GroupIdentifier(identifier string) Group {
	for key, ident := range groupIdentifiers {
		if ident == identifier {
			return key
		}
	}
	return Group(0)
}
