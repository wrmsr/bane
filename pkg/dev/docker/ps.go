package docker

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"strings"
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

	X map[string]interface{} `json:"-"`
}

//

type Pss []Ps

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

	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	var pss Pss
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

		var ps Ps
		fmt.Printf(s)
		if err := json.Unmarshal([]byte(s), &ps); err != nil {
			return nil, err
		}
		pss = append(pss, ps)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return pss, nil
}
