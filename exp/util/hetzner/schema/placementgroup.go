package schema

import "time"

//

type PlacementGroupType string

const (
	PlacementGroupTypeSpread PlacementGroupType = "spread"
)

//

type PlacementGroup struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Created time.Time         `json:"created"`
	Labels  map[string]string `json:"labels"`

	Servers []int `json:"servers"`

	Type PlacementGroupType `json:"type"`
}
