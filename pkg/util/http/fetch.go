package http

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	eu "github.com/wrmsr/bane/pkg/util/errors"
)

func Fetch(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
		err = eu.Append(err, res.Body.Close())
		return nil, err
	}

	return res, err
}

func FetchBytes(ctx context.Context, url string) ([]byte, error) {
	res, err := Fetch(ctx, url)
	if err != nil {
		return nil, err
	}
	defer eu.AppendInvoke(&err, eu.Close(res.Body))

	var out bytes.Buffer
	_, err = io.Copy(&out, res.Body)
	return out.Bytes(), err
}

func Download(ctx context.Context, url string, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer eu.AppendInvoke(&err, eu.Close(out))

	res, err := Fetch(ctx, url)
	if err != nil {
		return err
	}
	defer eu.AppendInvoke(&err, eu.Close(res.Body))

	_, err = io.Copy(out, res.Body)
	return err
}
