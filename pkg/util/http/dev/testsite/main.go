//go:build !nodev

package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/wrmsr/bane/pkg/util/check"
	iou "github.com/wrmsr/bane/pkg/util/io"
)

//go:embed pagelets/*.gohtml
var pageletsFs embed.FS

func hello(w http.ResponseWriter, req *http.Request) {
	iou.Discard(fmt.Fprintf(w, "hello\n"))
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			iou.Discard(fmt.Fprintf(w, "%v: %v\n", name, h))
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	check.Must(http.ListenAndServe(":8090", nil))
}
