package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/exp/util/wasm/instrs"
)

type Def struct {
	L string
	N string
	O int16

	T string
	A string
	B string
	C string

	Ma string
	Mz int
}

func main() {
	var sb strings.Builder

	wf := func(k, v string) {
		_, _ = fmt.Fprintf(&sb, ", \"%s\": \"%s\"", k, v)
	}

	ds := instrs.All()
	for i, d := range ds {
		sb.Reset()
		_, _ = fmt.Fprintf(&sb, "{\"L\": \"%s\"", d.L)

		wf("N", d.N)
		ox := strconv.FormatInt(int64(d.O), 16)
		wf("O", d.O)

		_, _ = sb.WriteString("}")
		if i < len(ds)-1 {
			_, _ = sb.WriteString(",")
		}
		_, _ = sb.WriteString("\n")
		fmt.Print(sb.String())
	}
}
