//go:build !nodev

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/wrmsr/bane/pkg/util/check"
	its "github.com/wrmsr/bane/pkg/util/iterators"
)

type dim struct {
	name string
	num  int

	types  []string
	params []string
	args   []string
}

func newDim(name string, num int) dim {
	return dim{
		name: name,
		num:  num,

		types:  its.Seq(its.Map(its.Range1(0, num), func(i int) string { return fmt.Sprintf("%s%d", strings.ToUpper(name), i) })),
		params: its.Seq(its.Map(its.Range1(0, num), func(i int) string { return fmt.Sprintf("%s%d %s%d", name, i, strings.ToUpper(name), i) })),
		args:   its.Seq(its.Map(its.Range1(0, num), func(i int) string { return fmt.Sprintf("%s%d", name, i) })),
	}
}

func joinDims(ds ...dim) dim {
	var j dim
	for _, d := range ds {
		j.name += d.name
		j.num += d.num

		j.types = append(j.types, d.types...)
		j.params = append(j.params, d.params...)
		j.args = append(j.args, d.args...)
	}
	return j
}

func (d dim) load(m map[string]any) {
	m["n_"+d.name] = strconv.Itoa(d.num)

	t := strings.Join(d.types, ", ")

	m["t_"+d.name] = t
	if d.num > 1 {
		m["pt_"+d.name] = "(" + t + ")"
	} else {
		m["pt_"+d.name] = t
	}

	m["p_"+d.name] = strings.Join(d.params, ", ")

	m["a_"+d.name] = strings.Join(d.args, ", ")
}

func main() {
	flags := flag.NewFlagSet("gen", flag.ExitOnError)

	var noWrite bool
	flags.BoolVar(&noWrite, "p", false, "do not write file, just print")

	check.Must(flags.Parse(os.Args[1:]))
	if flags.NArg() != 3 {
		panic(errors.New("expected 3 args"))
	}

	bn := check.Must1(strconv.Atoi(flags.Args()[0]))
	an := check.Must1(strconv.Atoi(flags.Args()[1]))
	rn := check.Must1(strconv.Atoi(flags.Args()[2]))

	tmpl := check.Must1(template.New("?").Parse(`
func Bind{{- .n_b -}}x{{- .n_a -}}x{{- .n_r -}}{{if .t_bar}}[{{.t_bar}} any]{{end}}(fn func({{.t_ba}}) {{.pt_r}}{{if .p_b}}, {{.p_b}}{{end -}}) func({{.t_a}}) {{.pt_r}} {
	return func({{.p_a}}) {{.pt_r}} {
		{{if .pt_r}}return {{end}}fn({{.a_ba}})
	}
}

func BindR{{- .n_a -}}x{{- .n_b -}}x{{- .n_r -}}{{if .t_abr}}[{{.t_abr}} any]{{end}}(fn func({{.t_ab}}) {{.pt_r}}{{if .p_b}}, {{.p_b}}{{end -}}) func({{.t_a}}) {{.pt_r}} {
	return func({{.p_a}}) {{.pt_r}} {
		{{if .pt_r}}return {{end}}fn({{.a_ab}})
	}
}
`))

	var out strings.Builder
	out.WriteString("package funcs\n")

	for cr := 0; cr < rn+1; cr++ {
		for ca := 0; ca < an+1; ca++ {
			for cb := 1; cb < bn+1; cb++ {
				b := newDim("b", cb)
				a := newDim("a", ca)
				r := newDim("r", cr)

				m := make(map[string]any)
				for _, ds := range [][]dim{
					{a},
					{b},
					{r},
					{a, b},
					{b, a},
					{a, b, r},
					{b, a, r},
				} {
					joinDims(ds...).load(m)
				}

				check.Must(tmpl.Execute(&out, m))
			}
		}
	}

	s := out.String()
	s = strings.ReplaceAll(s, " , ", ", ")
	s = strings.ReplaceAll(s, "  {", " {")
	s = strings.ReplaceAll(s, " ) ", ") ")

	if noWrite {
		fmt.Println(s)
	} else {
		check.Must(os.WriteFile("bind_gen.go", []byte(s), 0644))
	}
}
