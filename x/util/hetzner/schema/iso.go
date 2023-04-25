package schema

import "time"

//

type IsoType string

const (
	IsoTypePublic  IsoType = "public"
	IsoTypePrivate IsoType = "private"
)

//

type Iso struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Description string `json:"description"`

	Type IsoType `json:"type"`

	Deprecated time.Time `json:"deprecated"`
}
