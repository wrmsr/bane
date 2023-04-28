package annotate

import (
	"fmt"
	"reflect"
	"sync"
)

//

type FieldAnnotations struct {
	f reflect.StructField
	s []any
}

func (fa *FieldAnnotations) Field() reflect.StructField { return fa.f }
func (fa *FieldAnnotations) All() []any                 { return fa.s }

//

type MethodAnnotations struct {
	m reflect.Method
	s []any
}

func (ma *MethodAnnotations) Method() reflect.Method { return ma.m }
func (ma *MethodAnnotations) All() []any             { return ma.s }

//

type TypeAnnotations struct {
	mtx sync.Mutex
	ty  reflect.Type

	s []any

	fm map[string]*FieldAnnotations
	fs []*FieldAnnotations

	mm map[string]*MethodAnnotations
	ms []*MethodAnnotations
}

func (ta *TypeAnnotations) Type() reflect.Type { return ta.ty }
func (ta *TypeAnnotations) All() []any         { return ta.s }

//

type TypeAnnotationsMap struct {
	mtx sync.Mutex
	m   map[reflect.Type]*TypeAnnotations
}

func NewTypeAnnotationsMap() *TypeAnnotationsMap {
	return &TypeAnnotationsMap{
		m: make(map[reflect.Type]*TypeAnnotations),
	}
}

func (tam *TypeAnnotationsMap) Get(ty reflect.Type) *TypeAnnotations {
	tam.mtx.Lock()
	ta := tam.m[ty]
	tam.mtx.Unlock()
	return ta
}

//

var _global = NewTypeAnnotationsMap()

func Global() *TypeAnnotationsMap {
	return _global
}

//

type Annotator interface {
	Field(n string, anns ...any) Annotator
	Method(n string, anns ...any) Annotator
}

type annotator struct {
	ta *TypeAnnotations
}

var _ Annotator = &annotator{}

func On[T any](anns ...any) Annotator {
	var z T
	ty := reflect.TypeOf(z)
	if ty.Kind() != reflect.Struct {
		panic(fmt.Errorf("must be struct: %v", ty))
	}

	_global.mtx.Lock()
	ta := _global.m[ty]
	if ta == nil {
		ta = &TypeAnnotations{
			ty: ty,
		}
		_global.m[ty] = ta
	}
	_global.mtx.Unlock()

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
		a.ta.fm = make(map[string]*FieldAnnotations)
	}

	fa := a.ta.fm[n]
	if fa == nil {
		f, ok := a.ta.ty.FieldByName(n)
		if !ok {
			panic(fmt.Errorf("type %s has no field %s", a.ta.ty, n))
		}
		fa = &FieldAnnotations{f: f}
		a.ta.fm[n] = fa
		a.ta.fs = append(a.ta.fs, fa)
	}

	fa.s = append(fa.s, anns...)

	a.ta.mtx.Unlock()
	return a
}

func (a *annotator) Method(n string, anns ...any) Annotator {
	a.ta.mtx.Lock()

	if a.ta.mm == nil {
		a.ta.mm = make(map[string]*MethodAnnotations)
	}

	ma := a.ta.mm[n]
	if ma == nil {
		m, ok := a.ta.ty.MethodByName(n)
		if !ok {
			panic(fmt.Errorf("type %s has no method %s", a.ta.ty, n))
		}
		ma = &MethodAnnotations{m: m}
		a.ta.mm[n] = ma
		a.ta.ms = append(a.ta.ms, ma)
	}

	ma.s = append(ma.s, anns...)

	a.ta.mtx.Unlock()
	return a
}
