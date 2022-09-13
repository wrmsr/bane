package ndarray

import (
	"reflect"
)

func (a NdArray[T]) NestArray() any {
	sh := a.v.sh
	tys := make([]reflect.Type, sh.Len()+1)
	var z T
	tys[len(tys)-1] = reflect.TypeOf(z)
	for i := len(tys) - 2; i >= 0; i-- {
		tys[i] = reflect.ArrayOf(int(sh.Get(i)), tys[i+1])
	}

	sl := make([]Dim, sh.Len())
	var rec func(int) reflect.Value
	rec = func(dim int) reflect.Value {
		l := sh.Get(dim)
		ret := reflect.New(tys[dim]).Elem()
		for i := Dim(0); i < l; i++ {
			sl[dim] = i
			var o reflect.Value
			if dim < sh.Len()-1 {
				o = rec(dim + 1)
			} else {
				o = reflect.ValueOf(a.Get(sl...))
			}
			ret.Index(int(i)).Set(o)
		}
		return ret
	}
	return rec(0).Interface()
}

func (a NdArray[T]) NestSlice() any {
	sh := a.v.sh
	atys := make([]reflect.Type, sh.Len())
	stys := make([]reflect.Type, sh.Len())
	var sz []T
	atys[len(atys)-1] = reflect.TypeOf(sz)
	stys[len(atys)-1] = reflect.TypeOf(sz)
	for i := len(atys) - 2; i >= 0; i-- {
		atys[i] = reflect.ArrayOf(int(sh.Get(i)), stys[i+1])
		stys[i] = reflect.SliceOf(atys[i+1])
	}

	fl := a.Flat()
	p := 0
	var rec func(int) reflect.Value
	rec = func(dim int) reflect.Value {
		ln := sh.Get(dim)
		if dim < sh.Len()-1 {
			ret := reflect.New(atys[dim]).Elem()
			for i := Dim(0); i < ln; i++ {
				o := rec(dim + 1)
				ret.Index(int(i)).Set(o)
			}
			return ret.Slice(0, int(ln))
		} else {
			c := fl[p : p+int(ln)]
			p += int(ln)
			return reflect.ValueOf(c)
		}
	}
	return rec(0).Interface()
}
