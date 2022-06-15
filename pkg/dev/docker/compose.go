package docker

type ComposeService struct {
	Image    string `json:"image"`
	Platform string `json:"platform"`
	Restart  string `json:"restart"`

	Environment map[string]string `json:"environment"`

	Expose []string `json:"expose"`
	Ports  []string `json:"ports"`

	X map[string]interface{} `json:"-"`
}

type ComposeConfig struct {
	Version string `json:"version"`

	Services map[string]ComposeService

	X map[string]interface{} `json:"-"`
}
