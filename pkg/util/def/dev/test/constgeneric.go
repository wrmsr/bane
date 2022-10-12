//go:build !nodev

package test

import "github.com/wrmsr/bane/pkg/util/def"

//

const _N = 4

var _ = def.ConstGeneric[DenseVec_N[any]](_N)

type DenseVec_N[T any] struct {
	a [_N]T
	s []T
	l int
}

func (v *DenseVec_N[T]) Append(e T) {
	if v.l < _N {
		v.a[v.l] = e
	} else {
		v.s = append(v.s, e)
	}
	v.l++
}

func (v *DenseVec_N[T]) Get(i int) T {
	if i >= v.l {
		panic("index out of bounds")
	}
	if i < _N {
		return v.a[i]
	}
	return v.s[i-_N]
}
