package schema

import (
	"net"
	"time"
)

//

type LoadBalancerType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Description string `json:"description"`

	MaxConnections          int `json:"max_connections"`
	MaxServices             int `json:"max_services"`
	MaxTargets              int `json:"max_targets"`
	MaxAssignedCertificates int `json:"max_assigned_certificates"`
}

//

type LoadBalancerPublicNet struct {
	Enabled bool `json:"enabled"`

	IpV4 LoadBalancerPublicNetIpV4 `json:"ipv4"`
	IpV6 LoadBalancerPublicNetIpV6 `json:"ipv6"`
}

//

type LoadBalancerPublicNetIpV4 struct {
	Ip     net.IP `json:"ip"`
	DnsPtr any    `json:"dns_ptr"`
}

//

type LoadBalancerPublicNetIpV6 struct {
	Ip     net.IP `json:"ip"`
	DnsPtr any    `json:"dns_ptr"`
}

//

type LoadBalancerPrivateNet struct {
	Ip      net.IP   `json:"ip"`
	Network *Network `json:"network"`
}

//

type LoadBalancerService struct {
	Protocol LoadBalancerServiceProtocol `json:"protocol"`

	ListenPort      int `json:"listen_port"`
	DestinationPort int `json:"destination_port"`

	ProxyProtocol bool `json:"proxy_protocol"`

	Http LoadBalancerServiceHTTP `json:"http"`

	HealthCheck LoadBalancerServiceHealthCheck `json:"health_check"`
}

//

type LoadBalancerServiceHTTP struct {
	CookieName     string        `json:"cookie_name"`
	CookieLifetime time.Duration `json:"cookie_lifetime"`

	Certificates []*Certificate `json:"certificates"`

	RedirectHTTP bool `json:"redirect_http"`

	StickySessions bool `json:"sticky_sessions"`
}

//

type LoadBalancerServiceHealthCheck struct {
	Protocol LoadBalancerServiceProtocol `json:"protocol"`

	Port int `json:"port"`

	Interval time.Duration `json:"interval"`
	Timeout  time.Duration `json:"timeout"`

	Retries int `json:"retries"`

	Http *LoadBalancerServiceHealthCheckHTTP `json:"http"`
}

//

type LoadBalancerServiceHealthCheckHTTP struct {
	Domain      string   `json:"domain"`
	Path        string   `json:"path"`
	Response    string   `json:"response"`
	StatusCodes []string `json:"status_codes"`
	Tls         bool     `json:"tls"`
}

//

type LoadBalancerAlgorithmType string

const (
	LoadBalancerAlgorithmTypeRoundRobin       LoadBalancerAlgorithmType = "round_robin"
	LoadBalancerAlgorithmTypeLeastConnections LoadBalancerAlgorithmType = "least_connections"
)

//

type LoadBalancerAlgorithm struct {
	Type LoadBalancerAlgorithmType `json:"type"`
}

//

type LoadBalancerTargetType string

const (
	LoadBalancerTargetTypeServer        LoadBalancerTargetType = "server"
	LoadBalancerTargetTypeLabelSelector LoadBalancerTargetType = "label_selector"
	LoadBalancerTargetTypeIP            LoadBalancerTargetType = "ip"
)

//

type LoadBalancerServiceProtocol string

const (
	LoadBalancerServiceProtocolTCP   LoadBalancerServiceProtocol = "tcp"
	LoadBalancerServiceProtocolHTTP  LoadBalancerServiceProtocol = "http"
	LoadBalancerServiceProtocolHTTPS LoadBalancerServiceProtocol = "https"
)

//

type LoadBalancerTarget struct {
	Type LoadBalancerTargetType `json:"type"`

	Server *LoadBalancerTargetServer `json:"server"`
	Ip     *LoadBalancerTargetIP     `json:"ip"`

	LabelSelector *LoadBalancerTargetLabelSelector `json:"label_selector"`

	HealthStatus []LoadBalancerTargetHealthStatus `json:"health_status"`
	Targets      []LoadBalancerTarget             `json:"targets"`

	UsePrivateIp bool `json:"use_private_ip"`
}

//

type LoadBalancerTargetServer struct {
	Server *Server `json:"server"`
}

//

type LoadBalancerTargetLabelSelector struct {
	Selector string `json:"selector"`
}

//

type LoadBalancerTargetIP struct {
	Ip string `json:"ip"`
}

//

type LoadBalancerTargetHealthStatusStatus string

const (
	LoadBalancerTargetHealthStatusStatusUnknown   LoadBalancerTargetHealthStatusStatus = "unknown"
	LoadBalancerTargetHealthStatusStatusHealthy   LoadBalancerTargetHealthStatusStatus = "healthy"
	LoadBalancerTargetHealthStatusStatusUnhealthy LoadBalancerTargetHealthStatusStatus = "unhealthy"
)

//

type LoadBalancerTargetHealthStatus struct {
	ListenPort int                                  `json:"listen_port"`
	Status     LoadBalancerTargetHealthStatusStatus `json:"status"`
}

//

type LoadBalancer struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Created time.Time         `json:"created"`
	Labels  map[string]string `json:"labels"`

	PublicNet  LoadBalancerPublicNet    `json:"public_net"`
	PrivateNet []LoadBalancerPrivateNet `json:"private_net"`

	Location *Location `json:"location"`

	LoadBalancerType *LoadBalancerType `json:"load_balancer_type"`

	Algorithm LoadBalancerAlgorithm `json:"algorithm"`

	Services []LoadBalancerService `json:"services"`
	Targets  []LoadBalancerTarget  `json:"targets"`

	IncludedTraffic uint64 `json:"included_traffic"`
	OutgoingTraffic uint64 `json:"outgoing_traffic"`
	IngoingTraffic  uint64 `json:"ingoing_traffic"`
}
