package schema

import "time"

//

type ImageType string

const (
	ImageTypeSnapshot ImageType = "snapshot"
	ImageTypeBackup   ImageType = "backup"
	ImageTypeSystem   ImageType = "system"
	ImageTypeApp      ImageType = "app"
)

//

type ImageStatus string

const (
	ImageStatusCreating  ImageStatus = "creating"
	ImageStatusAvailable ImageStatus = "available"
)

//

type Image struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Type   ImageType   `json:"type"`
	Status ImageStatus `json:"status"`

	Created     time.Time         `json:"created"`
	Description string            `json:"description"`
	Labels      map[string]string `json:"labels"`

	ImageSize float32 `json:"image_size"`
	DiskSize  float32 `json:"disk_size"`

	CreatedFrom *Server `json:"created_from"`
	BoundTo     *Server `json:"bound_to"`

	RapidDeploy bool `json:"rapid_deploy"`

	OsFlavor  string `json:"os_flavor"`
	OsVersion string `json:"os_version"`

	Deprecated time.Time `json:"deprecated"`
}
