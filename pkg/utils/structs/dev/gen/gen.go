//go:build !nodev

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/tools/go/packages"

	"github.com/wrmsr/bane/pkg/utils/check"
	eu "github.com/wrmsr/bane/pkg/utils/errors"
	"github.com/wrmsr/bane/pkg/utils/log"
)

func findDirWithFile(cd, fn string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(cd, fn)); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return "", err
			}
		} else {
			return cd, nil
		}

		nd := filepath.Dir(cd)
		if nd == cd {
			return "", fmt.Errorf("file %s not found from dir %s", fn, cd)
		}
		cd = nd
	}
}

func readFirstLine(fp string) (s string, err error) {
	f, err := os.Open(fp)
	if err != nil {
		return
	}
	defer eu.AppendInvoke(&err, eu.Close(f))

	r := bufio.NewReader(f)
	l, p, err := r.ReadLine()
	if err != nil {
		return
	}
	if p {
		err = errors.New("line too big")
		return
	}

	s = string(l)
	return
}

func main() {
	cd := check.Must(os.Getwd())
	rd := check.Must(findDirWithFile(cd, "go.mod"))
	mod := check.Must(readFirstLine(filepath.Join(rd, "go.mod")))

	fmt.Println(mod)

	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes,
	}

	pkgs, err := packages.Load(cfg, "github.com/wrmsr/bane/pkg/utils/inject")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(pkgs)
}
