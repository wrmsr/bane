//go:build !nodev

package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
)

const devPrefix = "//go:build !nodev"

func main() {
	flags := flag.NewFlagSet("checkdev", flag.ExitOnError)

	var noFail bool
	flags.BoolVar(&noFail, "q", false, "do not return failure")

	check.Must(flags.Parse(os.Args[1:]))

	numModified := 0
	for _, arg := range flags.Args() {
		check.Must(filepath.Walk(arg, func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
				return nil
			}

			shouldDev := info.Name() == "dev.go" ||
				strings.HasSuffix(info.Name(), "_dev.go") ||
				slices.Contains(strings.Split(path, string(filepath.Separator)), "dev")

			src := string(check.Must1(ioutil.ReadFile(path)))
			hasDevPrefix := strings.HasPrefix(src, devPrefix)

			var newSrc string
			if shouldDev {
				if !hasDevPrefix {
					newSrc = devPrefix + "\n\n" + src
				}
			} else {
				if hasDevPrefix {
					newSrc = strings.TrimLeftFunc(src[len(devPrefix):], unicode.IsSpace)
				}
			}

			if newSrc == "" {
				return nil
			}

			fmt.Println(path)
			check.Must(ioutil.WriteFile(path, []byte(newSrc), info.Mode().Perm()))

			numModified++
			return nil
		}))
	}

	if !noFail {
		if numModified > 0 {
			os.Exit(1)
		}
		os.Exit(0)
	}
}
