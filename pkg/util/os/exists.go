package os

import (
	"errors"
	"os"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

func OptStat(name string) (bt.Optional[os.FileInfo], error) {
	stat, err := os.Stat(name)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return bt.None[os.FileInfo](), nil
		}
		return bt.None[os.FileInfo](), err
	}
	return bt.Just(stat), nil
}

func Exists(name string) (bool, error) {
	st, err := OptStat(name)
	if err != nil {
		return false, err
	}
	return st.Present(), nil
}
