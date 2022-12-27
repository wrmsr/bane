//go:build !nodev

package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
	stru "github.com/wrmsr/bane/pkg/util/strings"
)

//func render(ds []instrs.Def) string {
//	var sb strings.Builder
//	_, _ = sb.WriteString("[\n")
//
//	for i, d := range ds {
//		_, _ = fmt.Fprintf(&sb, "  {\"Class\": \"%s\"", d.Class)
//
//		wfr := func(k, v string) {
//			_, _ = fmt.Fprintf(&sb, ", \"%s\": %s", k, v)
//		}
//		wf := func(k, v string) {
//			wfr(k, fmt.Sprintf("\"%s\"", v))
//		}
//		wf("Name", d.Name)
//
//		if ox := (d.Op & 0xFF00) >> 8; ox != 0 {
//			wf("OpPfx", fmt.Sprintf("%02X", ox))
//		}
//		wf("Op", fmt.Sprintf("%02X", d.Op&0xFF))
//
//		wtf := func(k string, t wt.Type) {
//			if t != nil {
//				wf(k, t.String())
//			}
//		}
//		wtf("T", d.T)
//		wtf("A", d.A)
//		wtf("B", d.B)
//		wtf("C", d.C)
//
//		if d.Ma != 0 {
//			wf("Ma", d.Ma.String())
//		}
//		if d.Mz != 0 {
//			wfr("Mz", strconv.Itoa(d.Mz))
//		}
//
//		_, _ = sb.WriteString("}")
//		if i < len(ds)-1 {
//			_, _ = sb.WriteString(",")
//		}
//		_, _ = sb.WriteString("\n")
//	}
//
//	_, _ = sb.WriteString("]\n")
//	return sb.String()
//}

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
	dir := filepath.Join(paths.FindProjectRoot(), "exp/util/wasm/instrs/dev/gen")
	jb := check.Must1(os.ReadFile(filepath.Join(dir, "defs.json")))

	var jds []JsonDef
	check.Must(json.Unmarshal(jb, &jds))

	//

	var (
		sbInsts strings.Builder
		sbDefs  strings.Builder
	)

	_, _ = sbInsts.WriteString("const (\n")
	_, _ = sbInsts.WriteString("\t_ Instr = 0\n\n")
	_, _ = sbDefs.WriteString("var defs = []Def{\n")
	_, _ = sbDefs.WriteString("\t{Name: \"invalid\"},\n")

	title := cases.Title(language.Und).String

	mkName := func(jd JsonDef) string {
		n := stru.ToCamel(jd.Name)
		if strings.Contains(n, ".") {
			ps := strings.Split(n, ".")
			check.Equal(len(ps), 2)
			n = fmt.Sprintf("%s_%s", ps[0], title(ps[1]))
		}
		return n
	}

	for j, jd := range jds {
		i := j + 1

		_, _ = sbInsts.WriteString("\t")
		_, _ = sbDefs.WriteString("\t")

		n := mkName(jd)

		_, _ = fmt.Fprintf(&sbInsts, "%s = %d\n", n, i)

		_, _ = sbDefs.WriteString("{")

		_, _ = fmt.Fprintf(&sbDefs, "I: %s", n)
		_, _ = fmt.Fprintf(&sbDefs, ", Class: %s", title(jd.Class))
		_, _ = fmt.Fprintf(&sbDefs, ", Name: \"%s\"", jd.Name)
		_, _ = fmt.Fprintf(&sbDefs, ", Name2: \"%s\"", n)

		if jd.OpPfx != "" {
			_, _ = fmt.Fprintf(&sbDefs, ", OpPfx: uint8(0x%s)", strings.ToUpper(jd.OpPfx))
		}
		_, _ = fmt.Fprintf(&sbDefs, ", Op: uint8(0x%s)", strings.ToUpper(jd.Op))

		wtf := func(k, v string) {
			if v != "" {
				_, _ = fmt.Fprintf(&sbDefs, ", %s: wt.%s", k, strings.ToUpper(v))
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

	//

	var om [256]*[256]int

	for j, jd := range jds {
		var ox uint8
		if jd.OpPfx != "" {
			ox = uint8(check.Must1(strconv.ParseInt(jd.OpPfx, 16, 32)))
		}

		om2 := om[ox]
		if om2 == nil {
			om2 = new([256]int)
			om[ox] = om2
		}

		op := uint8(check.Must1(strconv.ParseInt(jd.Op, 16, 32)))
		check.Equal(om2[op], 0)
		i := j + 1
		om2[op] = i
	}

	var sbOpMap strings.Builder

	_, _ = sbOpMap.WriteString("var opMap = [256]*[256]Instr{\n")
	for i, om2 := range om {
		if om2 == nil {
			continue
		}
		_, _ = fmt.Fprintf(&sbOpMap, "\t0x%02X: {\n", i)
		for j, inst := range om2 {
			if inst == 0 {
				continue
			}
			_, _ = fmt.Fprintf(&sbOpMap, "\t\t0x%02X: %s,\n", j, mkName(jds[inst-1]))
		}
		_, _ = sbOpMap.WriteString("\t},\n")
	}
	_, _ = sbOpMap.WriteString("}\n")

	//

	var (
		sbNameMap  strings.Builder
		sbName2Map strings.Builder
	)

	_, _ = sbNameMap.WriteString("var nameMap = map[string]Instr{\n")
	_, _ = sbName2Map.WriteString("var name2Map = map[string]Instr{\n")
	for _, jd := range jds {
		n := mkName(jd)
		_, _ = fmt.Fprintf(&sbNameMap, "\t\"%s\": %s,\n", jd.Name, n)
		_, _ = fmt.Fprintf(&sbName2Map, "\t\"%s\": %s,\n", n, n)
	}
	_, _ = sbNameMap.WriteString("}\n")
	_, _ = sbName2Map.WriteString("}\n")

	//

	var sbSrc strings.Builder
	_, _ = sbSrc.WriteString("package instrs\n\n")
	_, _ = sbSrc.WriteString("import wt \"github.com/wrmsr/bane/exp/util/wasm/types\"\n\n")
	for _, sb := range []*strings.Builder{
		&sbInsts,
		&sbDefs,
		&sbOpMap,
		&sbNameMap,
		&sbName2Map,
	} {
		_, _ = sbSrc.WriteString(sb.String())
		_, _ = sbSrc.WriteString("\n\n")
	}
	src := sbSrc.String()

	fmt.Println(src)

	src = string(check.Must1(format.Source([]byte(src))))

	//

	fp := filepath.Join(dir, "../../defs_gen.go")
	check.Must(os.WriteFile(fp, []byte(src), 0644))
}
