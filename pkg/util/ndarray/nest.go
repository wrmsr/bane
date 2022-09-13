package ndarray

import "reflect"

func (a NdArray[T]) NestArray() any {
	sh := a.View().Shape()
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
