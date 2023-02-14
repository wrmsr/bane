package annotate

import (
	"fmt"
	"reflect"
	"sync"
)

//

type typeAnnotations struct {
	mtx sync.Mutex
	ty  reflect.Type
	s   []any
	fm  map[string][]any
	mm  map[string][]any
}

type typeAnnotationsMap struct {
	mtx sync.Mutex
	m   map[reflect.Type]*typeAnnotations
}

func newTypeAnnotationsMap() *typeAnnotationsMap {
	return &typeAnnotationsMap{
		m: make(map[reflect.Type]*typeAnnotations),
	}
}

var _typeAnnotationsMap = newTypeAnnotationsMap()

//

type Annotator interface {
	Field(n string, anns ...any) Annotator
	Method(n string, anns ...any) Annotator
}

type annotator struct {
	ta *typeAnnotations
}

var _ Annotator = &annotator{}

func Annotate[T any](anns ...any) Annotator {
	var z T
	ty := reflect.TypeOf(z)
	if ty.Kind() != reflect.Struct {
		panic(fmt.Errorf("must be struct: %v", ty))
	}

	_typeAnnotationsMap.mtx.Lock()
	ta := _typeAnnotationsMap.m[ty]
	if ta == nil {
		ta = &typeAnnotations{
			ty: ty,
		}
		_typeAnnotationsMap.m[ty] = ta
	}
	_typeAnnotationsMap.mtx.Unlock()

	ta.mtx.Lock()
	ta.s = append(ta.s, anns...)
	ta.mtx.Unlock()

	return &annotator{
		ta: ta,
	}
}

func (a *annotator) Field(n string, anns ...any) Annotator {
	a.ta.mtx.Lock()
	if a.ta.fm == nil {
		a.ta.fm = make(map[string][]any)
	}
	a.ta.fm[n] = append(a.ta.fm[n], anns...)
	a.ta.mtx.Unlock()
	return a
}

func (a *annotator) Method(n string, anns ...any) Annotator {
	a.ta.mtx.Lock()
	if a.ta.mm == nil {
		a.ta.mm = make(map[string][]any)
	}
	a.ta.mm[n] = append(a.ta.mm[n], anns...)
	a.ta.mtx.Unlock()
	return a
}
