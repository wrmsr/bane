package unsafe

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os/exec"
	"testing"
	"time"

	"golang.org/x/arch/arm64/arm64asm"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestAsm(t *testing.T) {
	fmt.Printf("%x\n", __addints(3, 4))
}

func TestGoTool(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "env", "-json")
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	check.Must(cmd.Run())

	var env map[string]string
	check.Must(json.Unmarshal(outb.Bytes(), &env))
	fmt.Println(env)
}

func TestArchPkg(t *testing.T) {
	var s string

	//s = "0a011f1a"
	s = "80348052"
	//s = "c0035fd6"

	inst := check.Must1(arm64asm.Decode(check.Must1(hex.DecodeString(s))))
	fmt.Println(inst)
}
