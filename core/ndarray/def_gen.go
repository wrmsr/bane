package ndarray

import (
	"sync"

	"github.com/wrmsr/bane/core/def"
)

var _def_init_once sync.Once

func _def_init() {
	_def_init_once.Do(func() {
		spec := def.X_getPackageSpec()
		_ = spec
	})
}

func init() {
	_def_init()
}

func _def_inl_foo(v View) Dim {
	var __def_inl_0 Dim
	__def_inl_1 := v
	__def_inl_2 := DimsOf(2, 2)
	{
		var __def_inl_3 Dim
		__def_inl_4 := __def_inl_1.
			st
		__def_inl_5 := __def_inl_2
		{
			__def_inl_4.CheckEqualLen(__def_inl_5)
			var __def_inl_6 Dim
			__def_inl_7 := __def_inl_4
			__def_inl_8 := __def_inl_5
			{
				l := __def_inl_7.Len()

				o := Dim(0)
				for i := 0; i < _dimsWidth; i++ {
					o += __def_inl_8._a[i] * __def_inl_7._a[i]
				}
				for i := 0; i < l-_dimsWidth; i++ {
					o += __def_inl_8._s[i] * __def_inl_7._s[i]
				}
				{
					__def_inl_6 = o
					goto __def_inl_9
				}
			__def_inl_9:
			}
			{
				__def_inl_3 = __def_inl_6
				goto __def_inl_10
			}
		__def_inl_10:
		}
		{
			__def_inl_0 = __def_inl_3 + __def_inl_1.o
			goto __def_inl_11
		}
	__def_inl_11:
	}

	return __def_inl_0
}
