package hetzner

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	hsch "github.com/wrmsr/bane/x/hetzner/schema"

	"github.com/wrmsr/bane/core/check"
	ju "github.com/wrmsr/bane/core/json"
)

type ServersResponse struct {
	Servers []hsch.Server `json:"servers"`

	Meta struct {
		Pagination struct {
			Page         int `json:"page"`
			PerPage      int `json:"per_page"`
			PreviousPage any `json:"previous_page"`
			NextPage     any `json:"next_page"`
			LastPage     int `json:"last_page"`
			TotalEntries int `json:"total_entries"`
		} `json:"pagination"`
	} `json:"meta"`
}

func TestHetzner(t *testing.T) {
	// ssh-keygen -b 4096 -t rsa -f /tmp/sshkey -q -N "" && chmod 600 /tmp/sshkey
	// ssh -i /tmp/sshkey root@...

	key := string(check.Must1(os.ReadFile(os.ExpandEnv("$HOME/.hetzner"))))

	url := "https://api.hetzner.cloud/v1/servers"

	ctx := context.Background()

	req := check.Must1(http.NewRequest("GET", url, nil))

	req.Header["Authorization"] = []string{fmt.Sprintf("Bearer %s", strings.TrimSpace(key))}

	req = req.WithContext(ctx)

	client := http.DefaultClient
	res := check.Must1(client.Do(req))

	fmt.Println(res.StatusCode)

	rbo := string(check.Must1(io.ReadAll(res.Body)))
	fmt.Println(rbo)

	check.Equal(res.StatusCode, http.StatusOK)

	var srs ServersResponse
	check.Must(json.Unmarshal([]byte(rbo), &srs))

	fmt.Println(check.Must1(ju.MarshalPretty(srs)))
}
