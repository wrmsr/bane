//go:build !nodev

package docker

import (
	"bytes"
	"encoding/json"
	"os"
	"regexp"

	"gopkg.in/yaml.v3"

	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/slices"
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

func ReadComposeConfig(paths []string) (*ComposeConfig, error) {
	ms := slices.Map(func(p string) map[string]any {
		b := check.Must1(os.ReadFile(p))
		var o map[string]any
		check.Must(yaml.NewDecoder(bytes.NewReader(b)).Decode(&o))
		return o
	}, paths)

	var merge func(l, r any) any
	merge = func(l, r any) any {
		lm, lok := l.(map[string]any)
		rm, rok := r.(map[string]any)
		if !lok || !rok {
			return r
		}
		for k, rv := range rm {
			if lv, ok := lm[k]; ok {
				lm[k] = merge(lv, rv)
			} else {
				lm[k] = rv
			}
		}
		return lm
	}

	m := ms[0]
	for _, r := range ms[1:] {
		if _, ok := r["services"]; ok {
			m["services"] = merge(m["services"], r["services"])
		}
	}

	b := check.Must1(json.Marshal(m))

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
