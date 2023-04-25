package schema

import (
	"net"
	"time"
)

//

type FloatingIpType string

const (
	FloatingIpTypeIpV4 FloatingIpType = "ipv4"
	FloatingIpTypeIpV6 FloatingIpType = "ipv6"
)

//

type FloatingIp struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Created     time.Time         `json:"created"`
	Description string            `json:"description"`
	Labels      map[string]string `json:"labels"`

	Ip net.IP `json:"ip"`

	Network *net.IPNet `json:"network"`

	Type FloatingIpType `json:"type"`

	Server *Server `json:"server"`

	DnsPtr any `json:"dns_ptr"`

	HomeLocation *Location `json:"home_location"`

	Blocked bool `json:"blocked"`
}
