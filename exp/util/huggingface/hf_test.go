package huggingface

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
)

type DatasetInfo struct {
	PrivId string `json:"_id"`
	Id     string `json:"id"`
	Sha    string `json:"sha"`

	LastModified string   `json:"lastModified"`
	Tags         []string `json:"tags"`
	Downloads    int      `json:"downloads"`

	Author      string `json:"author"`
	Description string `json:"description"`
	Citation    string `json:"citation"`

	Private  bool `json:"private"`
	Disabled bool `json:"disabled"`
	Gated    bool `json:"gated"`

	X map[string]any `json:"-"`
}

func TestHf(t *testing.T) {
	tok := string(check.Must1(os.ReadFile(filepath.Join(check.Must1(os.UserHomeDir()), ".huggingface/token"))))

	url := "https://huggingface.co/api/datasets"
	req := check.Must1(http.NewRequest("GET", url, nil))

	req.Header["user-agent"] = []string{"unknown/None"}
	req.Header["authorization"] = []string{fmt.Sprintf("Bearer %s", strings.TrimSpace(tok))}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	req = req.WithContext(ctx)

	client := http.DefaultClient
	res := check.Must1(client.Do(req))

	rb := check.Must1(io.ReadAll(res.Body))

	check.Equal(res.StatusCode, http.StatusOK)

	var dsis []DatasetInfo
	check.Must(json.Unmarshal(rb, &dsis))

	fmt.Println(dsis)
}
