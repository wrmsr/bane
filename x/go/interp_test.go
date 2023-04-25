package _go

import (
	"fmt"
	"go/types"
	"testing"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/interp"
	"golang.org/x/tools/go/ssa/ssautil"

	"github.com/wrmsr/bane/core/check"
)

func TestInterp(t *testing.T) {
	conf := loader.Config{}
	check.Must1(conf.FromArgs([]string{"foo.go"}, false))
	conf.Import("runtime")
	iprog := check.Must1(conf.Load())

	bmode := ssa.InstantiateGenerics | ssa.SanityCheckFunctions
	// bmode |= ssa.PrintFunctions // enable for debugging
	prog := ssautil.CreateProgram(iprog, bmode)
	prog.Build()

	mainPkg := prog.Package(iprog.Created[0].Pkg)

	sizes := types.SizesFor("gc", conf.Build.GOARCH)

	input := `foo`

	var imode interp.Mode // default mode
	// imode |= interp.DisableRecover // enable for debugging
	// imode |= interp.EnableTracing // enable for debugging
	exitCode := interp.Interpret(mainPkg, imode, sizes, input, []string{})
	if exitCode != 0 {
		t.Fatalf("interpreting %s: exit code was %d", input, exitCode)
	}
	fmt.Println(mainPkg)
}
