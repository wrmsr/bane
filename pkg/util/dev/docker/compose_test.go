//go:build !nodev

package docker

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
)

func TestCompose(t *testing.T) {
	b := check.Must1(ioutil.ReadFile(filepath.Join(paths.FindProjectRoot(), "docker", "docker-compose.yml")))
	var cfg ComposeConfig
	check.Must(yaml.NewDecoder(bytes.NewReader(b)).Decode(&cfg))
	fmt.Println(cfg)
}
