//go:build !nodev

package docker

import (
	"context"
	"encoding/json"
	"strings"

	stru "github.com/wrmsr/bane/pkg/util/strings"
)

//

type Ps struct {
	Command      string `json:"Command"`
	CreatedAt    string `json:"CreatedAt"`
	ID           string `json:"ID"`
	Image        string `json:"Image"`
	Labels       string `json:"Labels"`
	LocalVolumes string `json:"LocalVolumes"`
	Mounts       string `json:"Mounts"`
	Names        string `json:"Names"`
	Networks     string `json:"Networks"`
	Ports        string `json:"Ports"`
	RunningFor   string `json:"RunningFor"`
	Size         string `json:"Size"`
	State        string `json:"State"`
	Status       string `json:"Status"`

	X map[string]any `json:"-"`
}

//

type Pss []*Ps

func (pss Pss) Ids() []string {
	s := make([]string, len(pss))
	for i, ps := range pss {
		s[i] = ps.ID
	}
	return s
}

//

func CliPs(ctx context.Context) (Pss, error) {
	out, err := CliCmd(ctx, "ps", "--format", "{{json .}}")
	if err != nil {
		return nil, err
	}

	lines, err := stru.ScanAllLines(strings.NewReader(string(out)), true)
	if err != nil {
		return nil, err
	}

	var pss Pss
	for _, l := range lines {
		var ps Ps
		if err := json.Unmarshal([]byte(l), &ps); err != nil {
			return nil, err
		}
		pss = append(pss, &ps)
	}

	return pss, nil
}
