package slices

import (
	"reflect"
	"sort"

	"golang.org/x/exp/constraints"

	bt "github.com/wrmsr/bane/core/types"
)

func Clone[T any](s []T) []T {
	r := make([]T, len(s))
	copy(r, s)
	return r
}

func DeleteAt[T any](s []T, i int) []T {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = bt.Zero[T]()
	return s[:len(s)-1]
}

func Join[T any](ss ...[]T) []T {
	var r []T
	for _, s := range ss {
		r = append(r, s...)
	}
	return r
}

func Reverse[T any](s []T) []T {
	x := len(s)
	mx := x / 2
	r := x - 1
	for l := 0; l < mx; l++ {
		s[l], s[r] = s[r], s[l]
		r--
	}
	return s
}

func Reversed[T any](s []T) []T {
	r := make([]T, len(s))
	n := len(s) - 1
	for _, e := range s {
		r[n] = e
		n--
	}
	return r
}

func Pop[T any](s *[]T) T {
	l := len(*s) - 1
	r := (*s)[l]
	*s = (*s)[:l]
	return r
}

func Sort[T constraints.Ordered](s []T) []T {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

func Sorted[T constraints.Ordered](s []T) []T {
	return Sort(Clone(s))
}

func Equal[T comparable](l, r []T) bool {
	if len(l) != len(r) {
		return false
	}
	for i := range l {
		if l[i] != r[i] {
			return false
		}
	}
	return true
}

func Map[T, R any](fn func(t T) R, s []T) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = fn(v)
	}
	return r
}

func Cast[T, R any](s []T) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = bt.As[T, R](v)
	}
	return r
}

func Apply[T any](fn func(t T), s []T) {
	for _, e := range s {
		fn(e)
	}
}

func AsAny[T any](s []T) []any {
	r := make([]any, len(s))
	for i, e := range s {
		r[i] = e
	}
	return r
}

func IndexMap[T, R any](fn func(i int, t T) R, s []T) []R {
	r := make([]R, len(s))
	for i, v := range s {
		r[i] = fn(i, v)
	}
	return r
}

func ZipMap2[T, U, R any](fn func(t T, u U) R, ts []T, us []U) []R {
	l := len(ts)
	if len(us) < l {
		l = len(us)
	}
	r := make([]R, l)
	for i := 0; i < l; i++ {
		r[i] = fn(ts[i], us[i])
	}
	return r
}

func ZipMap3[T, U, V, R any](fn func(t T, u U, v V) R, ts []T, us []U, vs []V) []R {
	l := len(ts)
	if len(us) < l {
		l = len(us)
	}
	if len(vs) < l {
		l = len(vs)
	}
	r := make([]R, l)
	for i := 0; i < l; i++ {
		r[i] = fn(ts[i], us[i], vs[i])
	}
	return r
}

func Filter[T any](fn func(t T) bool, s []T) []T {
	var r []T
	for _, v := range s {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}

func Partition[T any](fn func(t T) bool, s []T) (t []T, f []T) {
	for _, c := range s {
		if fn(c) {
			t = append(t, c)
		} else {
			f = append(f, c)
		}
	}
	return
}

func Flatten[T any](s [][]T) []T {
	var r []T
	for _, v := range s {
		r = append(r, v...)
	}
	return r
}

func DeepFlatten[T any](vs ...any) []T {
	var r []T

	var urfl func(reflect.Value)
	urfl = func(rv reflect.Value) {
		if rv.Kind() == reflect.Slice {
			l := rv.Len()
			for i := 0; i < l; i++ {
				urfl(rv.Index(i))
			}
			return
		}

		r = append(r, rv.Interface().(T))
	}

	for _, v := range vs {
		if v == nil {
			continue
		}
		switch v := v.(type) {
		case T:
			r = append(r, v)
		case []T:
			r = append(r, v...)
		default:
			urfl(reflect.ValueOf(v))
		}
	}

	return r
}

func FlatMap[T, R any](fn func(t T) []R, s []T) []R {
	var r []R
	for _, v := range s {
		r = append(r, fn(v)...)
	}
	return r
}

func Repeat[T any](s []T, n int) []T {
	r := make([]T, n*len(s))
	p := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r[p] = s[j]
			p++
		}
	}
	return r
}

func Interleave[T any](s []T, e T) []T {
	r := make([]T, 0, (len(s)*2)-1)
	for i, c := range s {
		if i > 0 {
			r = append(r, e)
		}
		r = append(r, c)
	}
	return r
}

func Reduce[T any](fn func(l, r T) T, d T, s []T) T {
	r := d
	for _, v := range s {
		r = fn(r, v)
	}
	return r
}

func MakeAppend[T any](s []T, a ...T) []T {
	r := make([]T, len(s)+len(a))
	copy(r, s)
	copy(r[len(s):], a)
	return r
}

func MakePrepend[T any](s []T, a ...T) []T {
	r := make([]T, len(s)+len(a))
	copy(r, a)
	copy(r[len(a):], s)
	return r
}

func KvsOf[K, V any](o ...any) []bt.Kv[K, V] {
	if len(o)%2 != 0 {
		panic("must pass even number of args")
	}
	n := len(o) / 2
	r := make([]bt.Kv[K, V], n)
	for i, j := 0, 0; i < n; i, j = i+1, j+2 {
		r[i] = bt.KvOf(o[j].(K), o[j+1].(V))
	}
	return r
}

//

func Get[T any](s []T, i int) T {
	return s[i]
}

func Set[T any](s []T, i int, t T) {
	s[i] = t
}

func PtrTo[T any](s []T, i int) *T {
	return &s[i]
}
