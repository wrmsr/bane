package os

import (
	"os"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/maps"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

func GetEnvs(ks []string) map[string]opt.Optional[string] {
	m := make(map[string]opt.Optional[string], len(ks))
	for _, k := range ks {
		if ov, ok := os.LookupEnv(k); ok {
			m[k] = opt.Just(ov)
		} else {
			m[k] = opt.None[string]()
		}
	}
	return m
}

func SetEnvs(m map[string]opt.Optional[string]) {
	for k, v := range m {
		if v.Present() {
			check.Must(os.Setenv(k, v.Value()))
		} else {
			check.Must(os.Unsetenv(k))
		}
	}
}

func SwapEnvs(m map[string]opt.Optional[string]) func() {
	if m == nil || len(m) < 1 {
		return func() {}
	}

	o := GetEnvs(maps.Keys(m))
	SetEnvs(m)
	return func() {
		SetEnvs(o)
	}
}
