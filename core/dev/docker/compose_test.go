//go:build !nodev

package docker

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/dev/paths"
	ju "github.com/wrmsr/bane/core/json"
	"github.com/wrmsr/bane/core/slices"
)

func TestCompose(t *testing.T) {
	b := check.Must1(os.ReadFile(filepath.Join(paths.FindProjectRoot(), "docker", "docker-compose.yml")))
	var cfg ComposeConfig
	check.Must(yaml.NewDecoder(bytes.NewReader(b)).Decode(&cfg))
	fmt.Println(cfg)
}

func TestComposeMerge(t *testing.T) {
	bs := slices.Map(func(n string) map[string]any {
		b := check.Must1(os.ReadFile(filepath.Join(paths.FindProjectRoot(), "docker", n)))
		var o map[string]any
		check.Must(yaml.NewDecoder(bytes.NewReader(b)).Decode(&o))
		return o
	}, []string{
		"docker-compose-ext.yml",
		"docker-compose.yml",
	})

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

	l := bs[0]
	for _, r := range bs[1:] {
		if _, ok := r["services"]; ok {
			l["services"] = merge(l["services"], r["services"])
		}
	}

	fmt.Println(check.Must1(ju.MarshalPretty(l)))
}
