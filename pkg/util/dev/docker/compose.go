//go:build !nodev

package docker

import (
	"bytes"
	"io/ioutil"
	"regexp"

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

	Services map[string]*ComposeService

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

var composeNameRegexp = regexp.MustCompile(`^docker-(?P<n>[\w\-]+)-\d+$`)

func GetComposeName(n string) string {
	s := composeNameRegexp.FindStringSubmatch(n)
	if len(s) < 1 {
		return ""
	}
	return s[1]
}

func MatchComposeName(h, n string) bool {
	if h == n {
		return true
	}
	cn := GetComposeName(h)
	return cn != "" && cn == n
}
