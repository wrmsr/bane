//go:build !nodev

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
)

func main() {
	check.Condition(check.Must(os.Stat(".cache")).IsDir())

	type g4File struct {
		path string
		info fs.FileInfo
	}

	var g4s []g4File
	for _, arg := range os.Args[1:] {
		check.NoErr(filepath.Walk(arg, func(path string, info fs.FileInfo, err error) error {
			if strings.HasSuffix(info.Name(), ".g4") {
				g4s = append(g4s, g4File{path, info})
			}
			return nil
		}))
	}

	if len(g4s) < 1 {
		return
	}

	g4sByDir := make(map[string][]g4File)
	for _, g4 := range g4s {
		dir := filepath.Dir(g4.path)
		g4sByDir[dir] = append(g4sByDir[dir], g4)
	}

	needsBuildsByDir := make(map[string][]g4File)
	for dir, dirG4s := range g4sByDir {
		var needsBuild bool
		for _, g4 := range dirG4s {
			nr := []rune(g4.info.Name())
			itPath := string(nr[:check.MustOk(slices.FindLast(nr, '.'))]) + ".interp"
			itStat, err := os.Stat(filepath.Join(dir, "parser", itPath))
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					needsBuild = true
				} else {
					panic(err)
				}
			} else {
				needsBuild = needsBuild || g4.info.ModTime().After(itStat.ModTime())
			}
		}
		if needsBuild {
			needsBuildsByDir[dir] = dirG4s
		}
	}

	for dir, dirG4s := range needsBuildsByDir {
		fmt.Println(dir)
		fmt.Println(dirG4s)
	}
}
