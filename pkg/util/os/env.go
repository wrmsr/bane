package os

import (
	"os"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/maps"
	bt "github.com/wrmsr/bane/pkg/util/types"
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

func SetEnvs(m map[string]bt.Optional[string]) {
	for k, v := range m {
		if v.Present() {
			check.Must(os.Setenv(k, v.Value()))
		} else {
			check.Must(os.Unsetenv(k))
		}
	}
}

func SwapEnvs(m map[string]bt.Optional[string]) func() {
	if m == nil || len(m) < 1 {
		return func() {}
	}

	o := GetEnvs(maps.Keys(m))
	SetEnvs(m)
	return func() {
		SetEnvs(o)
	}
}
