package ndarray

import (
	"sync"

	"github.com/wrmsr/bane/pkg/util/def"
)

var _def_init_once sync.Once

func _def_init() {
	_def_init_once.Do(func() {
		spec := def.X_getPackageSpec()
		_ = spec
	})
}
