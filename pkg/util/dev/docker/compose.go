//go:build !nodev

package docker

import (
	"bytes"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type ComposeService struct {
	Image    string `json:"image"`
	Platform string `json:"platform"`
	Restart  string `json:"restart"`

	Environment map[string]string `json:"environment"`

	Expose []string `json:"expose"`
	Ports  []string `json:"ports"`

	X map[string]any `json:"-"`
}

type ComposeConfig struct {
	Version string `json:"version"`

	Services map[string]ComposeService

	X map[string]any `json:"-"`
}

func ReadComposeConfig(path string) (*ComposeConfig, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg ComposeConfig
	if err := yaml.NewDecoder(bytes.NewReader(b)).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
