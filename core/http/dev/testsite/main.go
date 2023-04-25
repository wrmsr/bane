//go:build !nodev

package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/http/dev/testsite/services"
	inj "github.com/wrmsr/bane/core/inject"
	iou "github.com/wrmsr/bane/core/io"
)

//

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

//go:embed pagelets/*.gohtml
var pageletsFs embed.FS

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

	for _, de := range check.Must1(pageletsFs.ReadDir("pagelets")) {
		fmt.Println(de)
	}

	fmt.Printf("%#v\n", inj.ProvideAs[*services.StatusService](injector).Status())

	check.Must(http.ListenAndServe(":8090", mux))
}
