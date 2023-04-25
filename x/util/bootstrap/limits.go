package bootstrap

import (
	"fmt"
	"syscall"

	eu "github.com/wrmsr/bane/pkg/util/errors"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

// FIXME: crossplat
var limitNames = map[string]int{
	"as":                syscall.RLIMIT_AS,
	"core":              syscall.RLIMIT_CORE,
	"cpu":               syscall.RLIMIT_CPU,
	"cpu_usage_monitor": syscall.RLIMIT_CPU_USAGE_MONITOR,
	"data":              syscall.RLIMIT_DATA,
	"fsize":             syscall.RLIMIT_FSIZE,
	"nofile":            syscall.RLIMIT_NOFILE,
	"stack":             syscall.RLIMIT_STACK,
}

type LimitsConfig struct {
	Limits map[string]bt.Optional[uint64]
	Reset  bool
}

func LimitsBootstrap(cfg LimitsConfig, ds *eu.DeferStack) error {
	var r map[int]syscall.Rlimit
	if cfg.Reset {
		r = make(map[int]syscall.Rlimit)
		ds.DeferErr(func() (re error) {
			for k, v := range r {
				eu.AppendInto(&re, syscall.Setrlimit(k, &v))
			}
			return
		})
	}

	if cfg.Limits != nil {
		for k, v := range cfg.Limits {
			i, ok := limitNames[k]
			if !ok {
				return fmt.Errorf("unknown limit name: %s", k)
			}

			var l syscall.Rlimit
			if err := syscall.Getrlimit(i, &l); err != nil {
				return err
			}
			if r != nil {
				r[i] = l
			}

			if !v.Present() {
				l.Cur = syscall.RLIM_INFINITY
			} else {
				l.Cur = v.Value()
			}
			if err := syscall.Setrlimit(i, &l); err != nil {
				return err
			}
		}
	}

	return nil
}
