package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/exp/util/wasm/instrs"
	wt "github.com/wrmsr/bane/exp/util/wasm/types"
	"github.com/wrmsr/bane/pkg/util/check"
)

func render(ds []instrs.Def) string {
	var sb strings.Builder
	_, _ = sb.WriteString("[\n")

	for i, d := range ds {
		_, _ = fmt.Fprintf(&sb, "  {\"L\": \"%s\"", d.Class)

		wfr := func(k, v string) {
			_, _ = fmt.Fprintf(&sb, ", \"%s\": %s", k, v)
		}
		wf := func(k, v string) {
			wfr(k, fmt.Sprintf("\"%s\"", v))
		}
		wf("N", d.Name)

		ox := strconv.FormatInt(int64(d.Op), 16)
		if len(ox)%2 != 0 {
			ox = "0" + ox
		}
		var oy string
		for j := 0; j < len(ox); j += 2 {
			if j > 0 {
				oy = "_" + oy
			}
			oy = ox[j:j+2] + oy
		}
		wf("O", oy)

		wt := func(k string, t wt.Type) {
			if t != nil {
				wf(k, t.String())
			}
		}
		wt("T", d.T)
		wt("A", d.A)
		wt("B", d.B)
		wt("C", d.C)

		if d.Ma != 0 {
			wf("Ma", d.Ma.String())
		}
		if d.Mz != 0 {
			wfr("Mz", strconv.Itoa(d.Mz))
		}

		_, _ = sb.WriteString("}")
		if i < len(ds)-1 {
			_, _ = sb.WriteString(",")
		}
		_, _ = sb.WriteString("\n")
	}

	_, _ = sb.WriteString("]\n")
	return sb.String()
}

type JsonDef struct {
	L string
	N string
	O string

	T string
	A string
	B string
	C string

	Ma string
	Mz int
}

func main() {
	ds := instrs.All()

	j := render(ds)
	fmt.Print(j)

	var jds []JsonDef
	check.Must(json.Unmarshal([]byte(j), &jds))

	fmt.Println(jds)
}
