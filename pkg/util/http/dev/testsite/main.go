//go:build !nodev

package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/http/dev/testsite/services"
	inj "github.com/wrmsr/bane/pkg/util/inject"
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

//func main() {
//	http.HandleFunc("/hello", hello)
//	http.HandleFunc("/headers", headers)
//
//	check.Must(http.ListenAndServe(":8090", nil))
//}

//

type serverFunc = func(mux *http.ServeMux)

func main() {
	mux := http.NewServeMux()

	injector := inj.NewInjector(inj.Bind(
		inj.Singleton{services.NewUptimeService},
		inj.Singleton{services.NewStatusService},

		inj.BindArrayOf[serverFunc]{},

		inj.As(inj.ArrayOf[serverFunc]{}, inj.Const{func(mux *http.ServeMux) {
			mux.HandleFunc("/hello", hello)
		}}),

		inj.As(inj.ArrayOf[serverFunc]{}, inj.Const{func(mux *http.ServeMux) {
			mux.HandleFunc("/headers", headers)
		}}),
	))

	arr := injector.Provide(inj.ArrayOf[serverFunc]{}).([]serverFunc)
	for _, fn := range arr {
		fn(mux)
	}

	fmt.Printf("%#v\n", inj.ProvideAs[*services.StatusService](injector).Status())

	check.Must(http.ListenAndServe(":8090", mux))
}
