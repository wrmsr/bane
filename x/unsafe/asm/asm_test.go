package asm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wrmsr/bane/core/check"
)

type Op int16

/*
100003f54: 80 34 80 52  mov     w0, #420
100003f58: c0 03 5f d6  ret
*/

const (
	OpInvalid Op = iota
	OpAdd
	OpMov
	OpRet
)

type Inst struct {
	Op Op
}

//go:noinline
func someFloat() float64 {
	return 420.69
}

func TestFloatBits(t *testing.T) {
	f := someFloat()
	b := math.Float64bits(f)
	fmt.Println(b)
}

//

type InstDef struct {
	Name   string `json:"Name"`
	Bits   string `json:"Bits"`
	Arch   string `json:"Arch"`
	Syntax string `json:"Syntax"`
	Code   string `json:"Code"`
	Alias  string `json:"Alias"`

	X map[string]any `json:"-"`
}

func TestInstrList(t *testing.T) {
	cmd := exec.CommandContext(context.Background(), "go", "list", "-m", "-f", "{{.Dir}}", "golang.org/x/arch")

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	check.Must(cmd.Run())

	dp := strings.TrimSpace(outb.String())
	bs := check.Must1(os.ReadFile(filepath.Join(dp, "arm64/arm64asm/inst.json")))

	var ids []InstDef
	check.Must(json.Unmarshal(bs, &ids))
}
