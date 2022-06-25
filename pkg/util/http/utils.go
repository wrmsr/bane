package http

import (
	"fmt"
	"io"
	"net/http"
	"os"

	eu "github.com/wrmsr/bane/pkg/util/errors"
)

func DownloadFile(url string, filepath string) (err error) {
	out, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer eu.AppendInvoke(&err, eu.Close(out))

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer eu.AppendInvoke(&err, eu.Close(resp.Body))

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", resp.Status)
		return
	}

	_, err = io.Copy(out, resp.Body)
	return
}
