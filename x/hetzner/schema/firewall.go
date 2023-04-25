package schema

import (
	"net"
	"time"
)

//

type FirewallRuleDirection string

const (
	FirewallRuleDirectionIn  FirewallRuleDirection = "in"
	FirewallRuleDirectionOut FirewallRuleDirection = "out"
)

//

type FirewallRuleProtocol string

const (
	FirewallRuleProtocolTCP  FirewallRuleProtocol = "tcp"
	FirewallRuleProtocolUDP  FirewallRuleProtocol = "udp"
	FirewallRuleProtocolICMP FirewallRuleProtocol = "icmp"
	FirewallRuleProtocolESP  FirewallRuleProtocol = "esp"
	FirewallRuleProtocolGRE  FirewallRuleProtocol = "gre"
)

//

type FirewallResourceType string

const (
	FirewallResourceTypeServer        FirewallResourceType = "server"
	FirewallResourceTypeLabelSelector FirewallResourceType = "label_selector"
)

//

type FirewallResourceLabelSelector struct {
	Selector string `json:"selector"`
}

//

type FirewallResourceServer struct {
	Id int `json:"id"`
}

//

type FirewallResource struct {
	Type          FirewallResourceType           `json:"type"`
	Server        *FirewallResourceServer        `json:"server"`
	LabelSelector *FirewallResourceLabelSelector `json:"label_selector"`
}

//

type FirewallRule struct {
	Protocol FirewallRuleProtocol `json:"protocol"`
	Port     *string              `json:"port"`

	Direction FirewallRuleDirection `json:"direction"`

	SourceIps      []net.IPNet `json:"source_ips"`
	DestinationIps []net.IPNet `json:"destination_ips"`

	Description *string `json:"description"`
}

//

type Firewall struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Created time.Time         `json:"created"`
	Labels  map[string]string `json:"labels"`

	Rules     []FirewallRule     `json:"rules"`
	AppliedTo []FirewallResource `json:"applied_to"`
}
