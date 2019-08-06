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
	return "filters[servers_groups][identifier]=" + groupIdentifiers[g]
}