package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"

	"github.com/wrmsr/bane/core/check"
	log "github.com/wrmsr/bane/core/slog"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("\tastdump [filename]")
		os.Exit(1)
	}

	fname := os.Args[1]
	file := check.Must1(os.Open(fname))
	defer log.OrError(file.Close)

	fset := token.NewFileSet()
	f := check.Must1(parser.ParseFile(fset, fname, file, 0))
	check.Must(ast.Print(fset, f))
}
