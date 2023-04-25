package schema

//

type StorageType string

const (
	StorageTypeLocal StorageType = "local"
	StorageTypeCeph  StorageType = "ceph"
)

//

type CpuType string

const (
	CpuTypeShared    CpuType = "shared"
	CpuTypeDedicated CpuType = "dedicated"
)

//

type ServerTypePriceValue struct {
	Net   float64 `json:"net"`
	Gross float64 `json:"gross"`
}

type ServerTypePrice struct {
	Location string `json:"location"`

	PriceHourly  ServerTypePriceValue `json:"price_hourly"`
	PriceMonthly ServerTypePriceValue `json:"price_monthly"`
}

//

type ServerType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Description string `json:"description"`

	Cores  int     `json:"cores"`
	Memory float32 `json:"memory"`
	Disk   int     `json:"disk"`

	StorageType StorageType `json:"storage_type"`

	CpuType CpuType `json:"cpu_type"`
}
