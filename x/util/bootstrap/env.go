//
/*
TODO:
 - expansion - {foo}
  - 'EnvSetter' ..
*/
package bootstrap

import (
	"os"

	eu "github.com/wrmsr/bane/pkg/util/errors"
)

//

type SetEnvEntry struct {
	Key   string
	Value string
}

type UnsetEnvEntry struct {
	Key string
}

type FileEnvEntry struct {
	Name string
}

type EnvEntry interface {
	isEnvEntry()
}

func (e SetEnvEntry) isEnvEntry()   {}
func (e UnsetEnvEntry) isEnvEntry() {}
func (e FileEnvEntry) isEnvEntry()  {}

//

type EnvConfig struct {
	Entries []EnvEntry
}

func EnvBootstrap(cfg EnvConfig, ds *eu.DeferStack) error {
	if len(cfg.Entries) < 1 {
		return nil
	}

	for _, e := range cfg.Entries {
		switch e := e.(type) {
		case SetEnvEntry:
			old, ok := os.LookupEnv(e.Key)
			if err := os.Setenv(e.Key, e.Value); err != nil {
				return err
			}
			ds.DeferErr(func() error {
				if ok {
					return os.Setenv(e.Key, old)
				} else {
					return os.Unsetenv(e.Key)
				}
			})

		case UnsetEnvEntry:
			old, ok := os.LookupEnv(e.Key)
			if err := os.Unsetenv(e.Key); err != nil {
				return err
			}
			ds.DeferErr(func() error {
				if ok {
					return os.Setenv(e.Key, old)
				} else {
					return os.Unsetenv(e.Key)
				}
			})

		case FileEnvEntry:
			panic("NYI")

		default:
			panic(e)
		}
	}

	return nil
}
