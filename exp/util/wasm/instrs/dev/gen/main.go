package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/wrmsr/bane/exp/util/wasm/instrs"
	wt "github.com/wrmsr/bane/exp/util/wasm/types"
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
	stru "github.com/wrmsr/bane/pkg/util/strings"
)

func render(ds []instrs.Def) string {
	var sb strings.Builder
	_, _ = sb.WriteString("[\n")

	for i, d := range ds {
		_, _ = fmt.Fprintf(&sb, "  {\"Class\": \"%s\"", d.Class)

		wfr := func(k, v string) {
			_, _ = fmt.Fprintf(&sb, ", \"%s\": %s", k, v)
		}
		wf := func(k, v string) {
			wfr(k, fmt.Sprintf("\"%s\"", v))
		}
		wf("Name", d.Name)

		if ox := (d.Op & 0xFF00) >> 8; ox != 0 {
			wf("OpPfx", fmt.Sprintf("%02X", ox))
		}
		wf("Op", fmt.Sprintf("%02X", d.Op&0xFF))

		wtf := func(k string, t wt.Type) {
			if t != nil {
				wf(k, t.String())
			}
		}
		wtf("T", d.T)
		wtf("A", d.A)
		wtf("B", d.B)
		wtf("C", d.C)

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
	Class string
	Name  string
	OpPfx string
	Op    string

	T string
	A string
	B string
	C string

	Ma string
	Mz int
}

func main() {
	//jb := []byte(render(instrs.All()))
	jb := check.Must1(os.ReadFile(filepath.Join(paths.FindProjectRoot(), "exp/util/wasm/instrs/dev/gen/defs.json")))

	var jds []JsonDef
	check.Must(json.Unmarshal(jb, &jds))

	var (
		sbInsts strings.Builder
		sbDefs  strings.Builder
	)

	_, _ = sbInsts.WriteString("const (\n")
	_, _ = sbInsts.WriteString("\t_ Instr = 0\n\n")
	_, _ = sbDefs.WriteString("var defs = []Def{\n")

	title := cases.Title(language.Und).String

	for j, jd := range jds {
		i := j + 1

		_, _ = sbInsts.WriteString("\t")
		_, _ = sbDefs.WriteString("\t")

		n := stru.ToCamel(jd.Name)
		if strings.Contains(n, ".") {
			ps := strings.Split(n, ".")
			check.Equal(len(ps), 2)
			n = fmt.Sprintf("%s_%s", ps[0], title(ps[1]))
		}

		_, _ = fmt.Fprintf(&sbInsts, "%s = %d\n", n, i)

		_, _ = sbDefs.WriteString("{")

		_, _ = fmt.Fprintf(&sbDefs, "I: %s", n)
		_, _ = fmt.Fprintf(&sbDefs, ", Class: %s", title(jd.Class))
		_, _ = fmt.Fprintf(&sbDefs, ", Name: \"%s\"", jd.Name)

		if jd.OpPfx != "" {
			_, _ = fmt.Fprintf(&sbDefs, ", OpPfx: int8(0x%s)", strings.ToUpper(jd.OpPfx))
		}
		_, _ = fmt.Fprintf(&sbDefs, ", Op: int8(0x%s)", strings.ToUpper(jd.Op))

		wtf := func(k, v string) {
			if v != "" {
				_, _ = fmt.Fprintf(&sbDefs, ", %s: %s", k, strings.ToUpper(v))
			}
		}
		wtf("T", jd.T)
		wtf("A", jd.A)
		wtf("B", jd.B)
		wtf("C", jd.C)

		if jd.Ma != "" {
			_, _ = fmt.Fprintf(&sbDefs, ", Ma: %s", title(jd.Ma))
		}
		if jd.Mz != 0 {
			_, _ = fmt.Fprintf(&sbDefs, ", Mz: %d", jd.Mz)
		}

		_, _ = sbDefs.WriteString("},\n")
	}

	_, _ = sbInsts.WriteString(")")
	_, _ = sbDefs.WriteString("}")

	fmt.Println(sbInsts.String())
	fmt.Println(sbDefs.String())
}
