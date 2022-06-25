//go:build !nodev

package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/wrmsr/bane/pkg/util/check"
	httpu "github.com/wrmsr/bane/pkg/util/http"
	osu "github.com/wrmsr/bane/pkg/util/os"
	"github.com/wrmsr/bane/pkg/util/slices"
)

func main() {
	check.Condition(check.Must(os.Stat(".cache")).IsDir())

	antlrVer := os.Args[1]
	jarName := fmt.Sprintf("antlr-%s-complete.jar", antlrVer)
	jarPath := filepath.Join(check.Must(os.Getwd()), ".cache", jarName)
	if !check.Must(osu.Exists(jarPath)) {
		jarUrl := fmt.Sprintf("https://www.antlr.org/download/%s", jarName)
		fmt.Println(jarUrl)
		check.NoErr(httpu.DownloadFile(jarUrl, jarPath))
	}

	type g4File struct {
		path string
		info fs.FileInfo
	}

	var g4s []g4File
	for _, arg := range os.Args[2:] {
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)
	for dir, dirG4s := range needsBuildsByDir {
		dir := dir

		parserDir := filepath.Join(dir, "parser")
		if check.Must(osu.Exists(parserDir)) {
			check.NoErr(os.RemoveAll(parserDir))
		}

		for _, g4 := range dirG4s {
			g4 := g4
			eg.Go(func() error {
				fmt.Println(g4.path)

				cmd := exec.CommandContext(
					ctx,
					"java",
					"-jar", jarPath,
					"-Dlanguage=Go",
					g4.info.Name(),
					"-visitor",
					"-o", "parser",
				)
				cmd.Dir = dir

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				return cmd.Run()
			})
		}
	}
	check.NoErr(eg.Wait())
}
