//go:build !nodev

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
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

var dontFixRetract modfile.VersionFixer = func(_, vers string) (string, error) {
	return vers, nil
}

func main() {
	cd := check.Must(os.Getwd())
	rd := check.Must(findDirWithFile(cd, "go.mod"))
	if !strings.HasPrefix(cd, rd) {
		panic(fmt.Errorf("can't find path %s from root %s", cd, rd))
	}
	do := cd[len(rd)+1:]
	fmt.Println(do)

	mf := filepath.Join(rd, "go.mod")
	mc := check.Must(ioutil.ReadFile(mf))
	mo := check.Must(modfile.Parse(mf, mc, dontFixRetract))
	mp := mo.Module.Mod.Path
	fmt.Println(mp)

	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes,
	}

	pn := fmt.Sprintf("%s/%s", mp, do)
	pkgs := check.Must(packages.Load(cfg, pn))

	log.Println(pkgs)
}
