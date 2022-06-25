package os

import (
	"errors"
	"os"

	opt "github.com/wrmsr/bane/pkg/util/optional"
)

func OptStat(name string) (opt.Optional[os.FileInfo], error) {
	stat, err := os.Stat(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return opt.None[os.FileInfo](), nil
		}
		return opt.None[os.FileInfo](), err
	}
	return opt.Just(stat), nil
}

func Exists(name string) (bool, error) {
	st, err := OptStat(name)
	if err != nil {
		return false, err
	}
	return st.Present(), nil
}
