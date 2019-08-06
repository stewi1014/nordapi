# nordapi
A simple api for NordVPN

NordAPI tries to use the unfortunately undocumented NordVPN public api for searching NordVPN servers.
[![GoDoc](https://godoc.org/github.com/stewi1014/nordapi?status.svg)](https://godoc.org/github.com/stewi1014/nordapi)

Examples:
```go
// Get top 5 Reccomended servers thare are Peer-To-Peer, Obfuscated and support OpenVPN UDP.
servers, err := nordapi.Reccomended(5,
	nordapi.GroupP2P,
	nordapi.GroupObfuscatedServers,
	nordapi.TechOpenvpnUDP,
)
if err != nil {
  // Couldn't get server list
}
if len(servers) == 0 {
  // No servers matched criterea
}

fmt.Println(servers[0].Hostname)
```

```go
// Get top 5 Reccomended servers in the Netherlands that support OpenVPN UDP
countries, err := nordapi.Countries()
if err != nil {
	// Couldn't get country list
}

country, err := countries.Name("Netherlands")
if err != nil {
	// Country not found
}

servers, err := nordapi.Reccomended(5,
	country,
	nordapi.TechOpenvpnUDP,
)
if err != nil {
	// Couldn't get server list
}
if len(servers) == 0 {
  // No servers matched criterea
}

fmt.Println(servers[0].Hostname)
```
To download OpenVPN config
```go
// Get OpenVPN UDP config for top reccomended server that supports OpenVPN UDP
servers, err := nordapi.Reccomended(5,
	nordapi.TechOpenvpnUDP,
)
if err != nil {
  // Couldn't get server list
}
if len(servers) == 0 {
  // No servers matched criterea
}

configReader, err := servers[0].OpenvpnUDPConfig()
if err != nil {
  // Couldn't read configuration
}

config, _ := ioutil.ReadAll(configReader)
fmt.Println(string(config))
```
