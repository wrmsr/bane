package schema

import "net"

//

type ServerPublicNetIpV4 struct {
	Id int    `json:"id"`
	Ip net.IP `json:"ip"`

	Blocked bool `json:"blocked"`

	DnsPtr any `json:"dns_ptr"`
}

//

type ServerPublicNetIpV6 struct {
	Id int    `json:"id"`
	Ip string `json:"ip"`

	Network *net.IPNet `json:"network"`

	Blocked bool `json:"blocked"`

	DnsPtr any `json:"dns_ptr"`
}

//

type ServerFirewallStatusEnum string

const (
	ServerFirewallStatusPending ServerFirewallStatusEnum = "pending"
	ServerFirewallStatusApplied ServerFirewallStatusEnum = "applied"
)

//

type ServerFirewallStatus struct {
	Firewall Firewall                 `json:"firewall"`
	Status   ServerFirewallStatusEnum `json:"status"`
}

//

type ServerPublicNet struct {
	IpV4 ServerPublicNetIpV4 `json:"ipv4"`
	IpV6 ServerPublicNetIpV6 `json:"ipv6"`

	FloatingIps []*FloatingIp `json:"floating_ips"`

	Firewalls []*ServerFirewallStatus `json:"firewalls"`
}

//

type ServerPrivateNet struct {
	Network    *Network `json:"network"`
	Ip         net.IP   `json:"ip"`
	Aliases    []net.IP `json:"aliases"`
	MacAddress string   `json:"mac_address"`
}
