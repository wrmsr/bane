package schema

//

type DatacenterServerTypes struct {
	Supported []int `json:"supported"`
	Available []int `json:"available"`
}

//

type Datacenter struct {
	Id   int    `json:"id"`
	Name string `json:"name"`

	Description string `json:"description"`

	Location *Location `json:"location"`

	ServerTypes DatacenterServerTypes `json:"server_types"`
}
