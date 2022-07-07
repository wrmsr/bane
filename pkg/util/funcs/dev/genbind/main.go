//go:build !nodev

/*
.a_ab
.a_ba
.a_bar
.a_p

.n_a
.n_b
.n_r

.p_a
.p_b

.pt_r

.t_a
.t_ab
.t_abr
.t_ba
.t_p
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/wrmsr/bane/pkg/util/check"
	its "github.com/wrmsr/bane/pkg/util/iterators"
)

func main() {
	if len(os.Args) != 4 {
		panic("invalid args")
	}

	bn := check.Must1(strconv.Atoi(os.Args[1]))
	an := check.Must1(strconv.Atoi(os.Args[2]))
	rn := check.Must1(strconv.Atoi(os.Args[3]))

	tmpl := check.Must1(template.New("?").Parse(`
func Bind{{- .n_b -}}x{{- .n_a -}}x{{- .n_r -}}[{{.a_bar}} any](fn func({{.t_ba}}) {{.pr_r}}, {{.p_b}}) func({{.t_a}}) {{.pt_r}} {
	return func({{.a_p}}) {{.pt_r}} {
		return fn({{.a_ba}})
	}
}

func BindR{{- .n_a -}}x{{- .n_b -}}x{{- .n_r -}}[{{.t_abr}} any](fn func({{.t_ab}}) {{.pt_r}}, {{.p_b}}) func({{.t_a}}) {{.pt_r}} {
	return func({{.p_a}}) {{.pt_r}} {
		return fn({{.a_ab}})
	}
}
`))

	pop := func(m map[string]string, p string, n int) {
		ts := its.Seq(its.Map(its.Range1(0, n), func(i int) string { return fmt.Sprintf("%s%d", strings.ToUpper(p), i) }))
		ps := its.Seq(its.Map(its.Range1(0, n), func(i int) string { return fmt.Sprintf("%s%d %s%d", p, i, strings.ToUpper(p), i) }))
		as := its.Seq(its.Map(its.Range1(0, n), func(i int) string { return fmt.Sprintf("%s%d", p, i) }))
		m[p+"n"] = strconv.Itoa(n)
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

	for r := 1; r < rn+1; r++ {
		for a := 1; a < an+1; a++ {
			for b := 1; b < bn+1; b++ {
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
