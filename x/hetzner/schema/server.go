package schema

import (
	"time"
)

//

type ServerStatus string

const (
	ServerStatusInitializing ServerStatus = "initializing"
	ServerStatusOff          ServerStatus = "off"
	ServerStatusRunning      ServerStatus = "running"
	ServerStatusStarting     ServerStatus = "starting"
	ServerStatusStopping     ServerStatus = "stopping"
	ServerStatusMigrating    ServerStatus = "migrating"
	ServerStatusRebuilding   ServerStatus = "rebuilding"
	ServerStatusDeleting     ServerStatus = "deleting"
	ServerStatusUnknown      ServerStatus = "unknown"
)

//

type Server struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Created time.Time         `json:"created"`
	Labels  map[string]string `json:"labels"`

	Status ServerStatus `json:"status"`

	PublicNet  ServerPublicNet    `json:"public_net"`
	PrivateNet []ServerPrivateNet `json:"private_net"`

	ServerType *ServerType `json:"server_type"`
	Datacenter *Datacenter `json:"datacenter"`

	IncludedTraffic uint64 `json:"included_traffic"`
	OutgoingTraffic uint64 `json:"outgoing_traffic"`
	IngoingTraffic  uint64 `json:"ingoing_traffic"`

	Iso   *Iso   `json:"iso"`
	Image *Image `json:"image"`

	PrimaryDiskSize int `json:"primary_disk_size"`

	PlacementGroup *PlacementGroup `json:"placement_group"`
}
