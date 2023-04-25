//go:build !nodev

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

	"github.com/wrmsr/bane/core/check"
)

func main() {
	flags := flag.NewFlagSet("checkdev", flag.ExitOnError)

	check.Must(flags.Parse(os.Args[1:]))

	for _, arg := range flags.Args() {
		check.Must(filepath.Walk(arg, func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
				return nil
			}

			src := check.Must1(os.ReadFile(path))
			fset := token.NewFileSet()
			f, err := parser.ParseFile(fset, "", src, parser.ParseComments)

			var cls []string
			if f.Doc != nil {
				for _, d := range f.Doc.List {
					if strings.Contains(strings.ToLower(d.Text), "copyright") {
						cls = append(cls, d.Text)
					}
				}
			}

			if len(cls) > 0 {
				fmt.Println(path)
				for _, cl := range cls {
					fmt.Println(strings.TrimSpace(cl))
				}
				fmt.Println()
			}

			return nil
		}))
	}
}
