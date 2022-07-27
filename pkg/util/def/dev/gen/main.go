//go:build !nodev

package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/def/dev/gen/impl"
)

func main() {
	flags := flag.NewFlagSet("gen", flag.ExitOnError)

	var noWrite bool
	flags.BoolVar(&noWrite, "p", false, "do not write file, just print")

	check.Must(flags.Parse(os.Args[1:]))
	if flags.NArg() > 0 {
		panic(errors.New("unexpected args"))
	}

	cwd := check.Must1(os.Getwd())

	s := impl.Run(cwd)

	if noWrite {
		fmt.Println(s)
	} else {
		fp := filepath.Join(cwd, "def_gen.go")
		check.Must(ioutil.WriteFile(fp, []byte(s), 0644))
	}
}
