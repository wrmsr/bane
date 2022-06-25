//go:build !nodev

package main

import (
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
	numModified := 0
	for _, arg := range os.Args[1:] {
		check.NoErr(filepath.Walk(arg, func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() || !strings.HasSuffix(info.Name(), ".go") {
				return nil
			}

			shouldDev := info.Name() == "dev.go" ||
				strings.HasSuffix(info.Name(), "_dev.go") ||
				slices.Contains(strings.Split(path, string(filepath.Separator)), "dev")

			src := string(check.Must(ioutil.ReadFile(path)))
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
			check.NoErr(ioutil.WriteFile(path, []byte(newSrc), info.Mode().Perm()))

			numModified++
			return nil
		}))
	}

	if numModified > 0 {
		os.Exit(1)
	}
	os.Exit(0)
}
