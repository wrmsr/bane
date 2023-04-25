package schema

type Location struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Description string `json:"description"`

	Country string `json:"country"`
	City    string `json:"city"`

	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	NetworkZone NetworkZone `json:"network_zone"`
}
