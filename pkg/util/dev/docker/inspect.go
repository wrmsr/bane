//go:build !nodev

package docker

import (
	"context"
	"encoding/json"
	"time"

	"github.com/wrmsr/bane/pkg/util/slices"
)

//

type InspectState struct {
	Status     string    `json:"Status"`
	Running    bool      `json:"Running"`
	Paused     bool      `json:"Paused"`
	Restarting bool      `json:"Restarting"`
	OomKilled  bool      `json:"OOMKilled"`
	Dead       bool      `json:"Dead"`
	Pid        int       `json:"Pid"`
	ExitCode   int       `json:"ExitCode"`
	Error      string    `json:"Error"`
	StartedAt  time.Time `json:"StartedAt"`
	FinishedAt time.Time `json:"FinishedAt"`

	X map[string]any `json:"-"`
}

type InspectHostConfigMount struct {
	Type   string `json:"Type"`
	Source string `json:"Source"`
	Target string `json:"Target"`

	X map[string]any `json:"-"`
}

type InspectHostConfigRestartPolicy struct {
	Name              string `json:"Name"`
	MaximumRetryCount int    `json:"MaximumRetryCount"`

	X map[string]any `json:"-"`
}

type InspectHostConfigPortBinding struct {
	HostIp   string `json:"HostIp"`
	HostPort string `json:"HostPort"`

	X map[string]any `json:"-"`
}

type InspectHostConfig struct {
	Binds       []any  `json:"Binds"`
	NetworkMode string `json:"NetworkMode"`

	PortBindings map[string][]InspectHostConfigPortBinding `json:"PortBindings"`

	RestartPolicy InspectHostConfigRestartPolicy `json:"RestartPolicy"`

	AutoRemove bool `json:"AutoRemove"`

	VolumeDriver string `json:"VolumeDriver"`
	VolumesFrom  any    `json:"VolumesFrom"`

	CapAdd       any    `json:"CapAdd"`
	CapDrop      any    `json:"CapDrop"`
	CgroupnsMode string `json:"CgroupnsMode"`

	Dns        []any `json:"Dns"`
	DnsOptions []any `json:"DnsOptions"`
	DnsSearch  []any `json:"DnsSearch"`

	ExtraHosts      []any  `json:"ExtraHosts"`
	GroupAdd        any    `json:"GroupAdd"`
	IpcMode         string `json:"IpcMode"`
	Cgroup          string `json:"Cgroup"`
	Links           any    `json:"Links"`
	OomScoreAdj     int    `json:"OomScoreAdj"`
	PidMode         string `json:"PidMode"`
	Privileged      bool   `json:"Privileged"`
	PublishAllPorts bool   `json:"PublishAllPorts"`
	ReadonlyRootfs  bool   `json:"ReadonlyRootfs"`
	SecurityOpt     any    `json:"SecurityOpt"`
	UtsMode         string `json:"UTSMode"`
	UsernsMode      string `json:"UsernsMode"`
	ShmSize         int    `json:"ShmSize"`
	Runtime         string `json:"Runtime"`
	ConsoleSize     []int  `json:"ConsoleSize"`
	Isolation       string `json:"Isolation"`
	CpuShares       int    `json:"CpuShares"`
	Memory          int    `json:"Memory"`
	NanoCpus        int    `json:"NanoCpus"`
	CgroupParent    string `json:"CgroupParent"`

	BlkioWeight          int `json:"BlkioWeight"`
	BlkioWeightDevice    any `json:"BlkioWeightDevice"`
	BlkioDeviceReadBps   any `json:"BlkioDeviceReadBps"`
	BlkioDeviceWriteBps  any `json:"BlkioDeviceWriteBps"`
	BlkioDeviceReadIops  any `json:"BlkioDeviceReadIOps"`
	BlkioDeviceWriteIops any `json:"BlkioDeviceWriteIOps"`

	CpuPeriod          int    `json:"CpuPeriod"`
	CpuQuota           int    `json:"CpuQuota"`
	CpuRealtimePeriod  int    `json:"CpuRealtimePeriod"`
	CpuRealtimeRuntime int    `json:"CpuRealtimeRuntime"`
	CpusetCpus         string `json:"CpusetCpus"`
	CpusetMems         string `json:"CpusetMems"`

	Devices           any `json:"Devices"`
	DeviceCgroupRules any `json:"DeviceCgroupRules"`
	DeviceRequests    any `json:"DeviceRequests"`
	KernelMemory      int `json:"KernelMemory"`
	KernelMemoryTcp   int `json:"KernelMemoryTCP"`

	MemoryReservation int `json:"MemoryReservation"`
	MemorySwap        int `json:"MemorySwap"`
	MemorySwappiness  any `json:"MemorySwappiness"`

	OomKillDisable any `json:"OomKillDisable"`

	PidsLimit any `json:"PidsLimit"`
	Ulimits   any `json:"Ulimits"`

	CpuCount   int `json:"CpuCount"`
	CpuPercent int `json:"CpuPercent"`

	IoMaximumIops      int `json:"IOMaximumIOps"`
	IoMaximumBandwidth int `json:"IOMaximumBandwidth"`

	Mounts []InspectHostConfigMount `json:"Mounts"`

	MaskedPaths   []string `json:"MaskedPaths"`
	ReadonlyPaths []string `json:"ReadonlyPaths"`

	X map[string]any `json:"-"`
}

type InspectMount struct {
	Type        string `json:"Type"`
	Name        string `json:"Name"`
	Source      string `json:"Source"`
	Destination string `json:"Destination"`
	Driver      string `json:"Driver"`
	Mode        string `json:"Mode"`
	Rw          bool   `json:"RW"`
	Propagation string `json:"Propagation"`

	X map[string]any `json:"-"`
}

type InspectConfig struct {
	Hostname   string `json:"Hostname"`
	Domainname string `json:"Domainname"`

	User string `json:"User"`

	AttachStdin  bool `json:"AttachStdin"`
	AttachStdout bool `json:"AttachStdout"`
	AttachStderr bool `json:"AttachStderr"`

	ExposedPorts map[string]any `json:"ExposedPorts"`

	Tty       bool `json:"Tty"`
	OpenStdin bool `json:"OpenStdin"`
	StdinOnce bool `json:"StdinOnce"`

	Env   []string `json:"Env"`
	Cmd   []string `json:"Cmd"`
	Image string   `json:"Image"`

	Volumes    map[string]any `json:"Volumes"`
	WorkingDir string         `json:"WorkingDir"`
	Entrypoint []string       `json:"Entrypoint"`

	OnBuild any `json:"OnBuild"`

	Labels map[string]string `json:"Labels"`

	StopSignal string `json:"StopSignal"`

	X map[string]any `json:"-"`
}

type InspectNetwork struct {
	IpamConfig          any      `json:"IPAMConfig"`
	Links               any      `json:"Links"`
	Aliases             []string `json:"Aliases"`
	NetworkId           string   `json:"NetworkID"`
	EndpointId          string   `json:"EndpointID"`
	Gateway             string   `json:"Gateway"`
	IpAddress           string   `json:"IPAddress"`
	IpPrefixLen         int      `json:"IPPrefixLen"`
	Ipv6Gateway         string   `json:"IPv6Gateway"`
	GlobalIpv6Address   string   `json:"GlobalIPv6Address"`
	GlobalIpv6PrefixLen int      `json:"GlobalIPv6PrefixLen"`
	MacAddress          string   `json:"MacAddress"`
	DriverOpts          any      `json:"DriverOpts"`

	X map[string]any `json:"-"`
}

type InspectNetworkSettingsPort struct {
	HostIp   string `json:"HostIp"`
	HostPort string `json:"HostPort"`

	X map[string]any `json:"-"`
}

type InspectNetworkSettings struct {
	Bridge      string `json:"Bridge"`
	SandboxID   string `json:"SandboxID"`
	HairpinMode bool   `json:"HairpinMode"`

	LinkLocalIpv6Address   string `json:"LinkLocalIPv6Address"`
	LinkLocalIpv6PrefixLen int    `json:"LinkLocalIPv6PrefixLen"`

	Ports map[string][]InspectNetworkSettingsPort `json:"Ports"`

	SandboxKey string `json:"SandboxKey"`

	SecondaryIpAddresses   any `json:"SecondaryIPAddresses"`
	SecondaryIpv6Addresses any `json:"SecondaryIPv6Addresses"`

	EndpointID string `json:"EndpointID"`

	Gateway string `json:"Gateway"`

	GlobalIpv6Address   string `json:"GlobalIPv6Address"`
	GlobalIpv6PrefixLen int    `json:"GlobalIPv6PrefixLen"`

	IpAddress   string `json:"IPAddress"`
	IpPrefixLen int    `json:"IPPrefixLen"`
	Ipv6Gateway string `json:"IPv6Gateway"`

	MacAddress string `json:"MacAddress"`

	Networks map[string]InspectNetwork `json:"Networks"`

	X map[string]any `json:"-"`
}

type Inspect struct {
	Id      string    `json:"Id"`
	Created time.Time `json:"Created"`

	Path string   `json:"Path"`
	Args []string `json:"Args"`

	State InspectState `json:"State"`

	Image string `json:"Image"`

	ResolvConfPath string `json:"ResolvConfPath"`
	HostnamePath   string `json:"HostnamePath"`
	HostsPath      string `json:"HostsPath"`
	LogPath        string `json:"LogPath"`

	Name string `json:"Name"`

	RestartCount int `json:"RestartCount"`

	Driver   string `json:"Driver"`
	Platform string `json:"Platform"`

	MountLabel   string `json:"MountLabel"`
	ProcessLabel string `json:"ProcessLabel"`

	AppArmorProfile string `json:"AppArmorProfile"`

	ExecIDs any `json:"ExecIDs"`

	HostConfig InspectHostConfig `json:"HostConfig"`

	Mounts []InspectMount `json:"Mounts"`

	Config InspectConfig `json:"Config"`

	NetworkSettings InspectNetworkSettings `json:"NetworkSettings"`

	X map[string]any `json:"-"`
}

//

type Inspects []Inspect

//

func CliInspect(ctx context.Context, ids []string) (Inspects, error) {
	out, err := CliCmd(ctx, slices.Join([]string{"inspect"}, ids)...)
	if err != nil {
		return nil, err
	}

	var inspects Inspects
	if err := json.Unmarshal(out, &inspects); err != nil {
		return nil, err
	}

	return inspects, nil
}

func CliInspectAll(ctx context.Context) (Inspects, error) {
	pss, err := CliPs(ctx)
	if err != nil {
		return nil, err
	}

	return CliInspect(ctx, pss.Ids())
}
