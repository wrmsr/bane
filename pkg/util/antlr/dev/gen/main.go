//go:build !nodev

package main

import (
	"context"
	"errors"
	"flag"
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
	flags := flag.NewFlagSet("gen", flag.ExitOnError)

	var force bool
	flags.BoolVar(&force, "f", false, "force generation (ignore mtime)")

	check.Must(flags.Parse(os.Args[1:]))

	if !check.Must1(osu.Exists(".cache")) {
		check.Must(os.Mkdir(".cache", 0744))
	}

	antlrVer := flags.Args()[0]
	jarName := fmt.Sprintf("antlr-%s-complete.jar", antlrVer)
	jarPath := filepath.Join(check.Must1(os.Getwd()), ".cache", jarName)
	if !check.Must1(osu.Exists(jarPath)) {
		jarUrl := fmt.Sprintf("https://www.antlr.org/download/%s", jarName)
		fmt.Println(jarUrl)
		check.Must(httpu.Download(context.Background(), jarUrl, jarPath))
	}

	type g4File struct {
		path string
		info fs.FileInfo
	}

	var g4s []g4File
	for _, arg := range flags.Args()[1:] {
		check.Must(filepath.Walk(arg, func(path string, info fs.FileInfo, err error) error {
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
		if force {
			needsBuildsByDir[dir] = dirG4s
			continue
		}
		var needsBuild bool
		for _, g4 := range dirG4s {
			nr := []rune(g4.info.Name())
			itPath := string(nr[:check.Ok1(slices.FindLast(nr, '.'))]) + ".interp"
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
		if check.Must1(osu.Exists(parserDir)) {
			check.Must(os.RemoveAll(parserDir))
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
				if err := cmd.Run(); err != nil {
					return err
				}

				rp := filepath.Join(dir, "parser")
				check.Must(filepath.Walk(rp, func(path string, info fs.FileInfo, err error) error {
					if !strings.HasSuffix(info.Name(), ".go") {
						return nil
					}

					fp := filepath.Join(rp, info.Name())
					src := string(check.Must1(os.ReadFile(fp)))

					for _, rep := range []struct{ l, r string }{
						{
							`"github.com/antlr/antlr4/runtime/Go/antlr/v4"`,
							`antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"`,
						},
						{"interface{}", "any"},
					} {
						src = strings.ReplaceAll(src, rep.l, rep.r)
					}

					check.Must(os.WriteFile(fp, []byte(src), info.Mode()))

					return nil
				}))

				return nil
			})
		}
	}
	check.Must(eg.Wait())
}
