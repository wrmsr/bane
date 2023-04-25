package os

import (
	"os"

	eu "github.com/wrmsr/bane/core/errors"
	"github.com/wrmsr/bane/core/maps"
	bt "github.com/wrmsr/bane/core/types"
)

func GetEnvs(ks []string) map[string]bt.Optional[string] {
	m := make(map[string]bt.Optional[string], len(ks))
	for _, k := range ks {
		if ov, ok := os.LookupEnv(k); ok {
			m[k] = bt.Just(ov)
		} else {
			m[k] = bt.None[string]()
		}
	}
	return m
}

func SetEnvs(m map[string]bt.Optional[string]) (err error) {
	for k, v := range m {
		if v.Present() {
			err = eu.Append(err, os.Setenv(k, v.Value()))
		} else {
			err = eu.Append(err, os.Unsetenv(k))
		}
	}
	return
}

func SwapEnvs(m map[string]bt.Optional[string]) (bt.ErrFn, error) {
	if m == nil || len(m) < 1 {
		return func() error { return nil }, nil
	}

	o := GetEnvs(maps.Keys(m))

	err := SetEnvs(m)
	if err != nil {
		return nil, err
	}

	return func() error {
		return SetEnvs(o)
	}, nil
}
