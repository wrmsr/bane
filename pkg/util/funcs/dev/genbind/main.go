//go:build !nodev

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/wrmsr/bane/pkg/util/check"
)

func main() {
	if len(os.Args) != 4 {
		panic("invalid args")
	}

	bn := check.Must1(strconv.Atoi(os.Args[1]))
	an := check.Must1(strconv.Atoi(os.Args[2]))
	rn := check.Must1(strconv.Atoi(os.Args[3]))

	tmpl := check.Must1(template.New("?").Parse(`
func Bind{{- .bn -}}x{{- .an -}}x{{- .rn -}}[{{.bts}}, {{.ats}}, {{.rts}} any](fn func({{.bts}}, {{.ats}}) {{.prts}}, {{.bps}}) func({{.ats}}) {{.prts}} {
	return func({{.aps}}) {{.prts}} {
		return fn({{.bas}}, {{.aas}})
	}
}
`))

	pop := func(m map[string]string, p string, n int) {
		m[p+"n"] = strconv.Itoa(n)
		var tsb strings.Builder
		var psb strings.Builder
		var asb strings.Builder
		for i := 0; i < n; i++ {
			if i > 0 {
				tsb.WriteString(", ")
				psb.WriteString(", ")
				asb.WriteString(", ")
			}
			tsb.WriteString(fmt.Sprintf("%s%d", strings.ToUpper(p), i))
			psb.WriteString(fmt.Sprintf("%s%d %s%d", p, i, strings.ToUpper(p), i))
			asb.WriteString(fmt.Sprintf("%s%d", p, i))
		}
		m[p+"ts"] = tsb.String()
		m[p+"ps"] = psb.String()
		m[p+"as"] = asb.String()
		if p == "r" {
			if n > 1 {
				m["prts"] = "(" + tsb.String() + ")"
			} else {
				m["prts"] = tsb.String()
			}
		}
	}

	var out strings.Builder
	out.WriteString("package funcs\n")

	for b := 1; b < bn+1; b++ {
		for a := 1; a < an+1; a++ {
			for r := 1; r < rn+1; r++ {
				m := make(map[string]string)
				pop(m, "b", b)
				pop(m, "a", a)
				pop(m, "r", r)

				check.Must(tmpl.Execute(&out, m))
			}
		}
	}

	check.Must(ioutil.WriteFile("bind_gen.go", []byte(out.String()), 0644))
}
