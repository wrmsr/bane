package schema

import (
	"net"
	"time"
)

//

type NetworkZone string

const (
	NetworkZoneEuCentral NetworkZone = "eu-central"
	NetworkZoneUsEast    NetworkZone = "us-east"
)

//

type NetworkSubnetType string

const (
	NetworkSubnetTypeCloud   NetworkSubnetType = "cloud"
	NetworkSubnetTypeServer  NetworkSubnetType = "server"
	NetworkSubnetTypeVSwitch NetworkSubnetType = "vswitch"
)

//

type NetworkSubnet struct {
	Type NetworkSubnetType `json:"type"`

	IpRange *net.IPNet `json:"ip_range"`

	NetworkZone NetworkZone `json:"network_zone"`

	Gateway net.IP `json:"gateway"`

	VswitchID int `json:"vswitch_id"`
}

//

type NetworkRoute struct {
	Destination *net.IPNet `json:"destination"`

	Gateway net.IP `json:"gateway"`
}

//

type Network struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Created time.Time `json:"created"`
	Labels  map[string]string

	IpRange *net.IPNet `json:"ip_range"`

	Subnets []NetworkSubnet `json:"subnets"`
	Routes  []NetworkRoute  `json:"routes"`

	Servers []*Server `json:"servers"`
}
