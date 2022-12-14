package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/exp/util/wasm/instrs"
	wt "github.com/wrmsr/bane/exp/util/wasm/types"
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
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
			wf("OpPfx", strconv.FormatInt(int64(ox), 16))
		}
		wf("Op", strconv.FormatInt(int64(d.Op&0xFF), 16))

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

	var sb strings.Builder

	for i, jd := range jds {
		_, _ = fmt.Fprintf(&sb, "Def{I: %d, N: \"%s\"}\n", i+1, jd.Name)
	}

	fmt.Println(sb.String())
}
