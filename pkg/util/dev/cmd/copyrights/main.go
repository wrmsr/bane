package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/wrmsr/bane/pkg/util/check"
)

func main() {
	flags := flag.NewFlagSet("checkdev", flag.ExitOnError)

	check.Must(flags.Parse(os.Args[1:]))

	for _, arg := range flags.Args() {
		check.Must(filepath.Walk(arg, func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
				return nil
			}

			fmt.Println(path)

			src := check.Must1(os.ReadFile(path))
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "", src, parser.ParseComments)

			fmt.Println(f)

			return nil
		}))
	}
}
