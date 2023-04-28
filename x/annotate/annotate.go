package annotate

import (
	"fmt"
	"reflect"
	"sync"
)

//

//type Target interface {
//	isTarget()
//}

//

type annotationCollection struct {
	s []any
	m map[reflect.Type][]any
}

func (c *annotationCollection) add(vs ...any) {
	if len(vs) < 1 {
		return
	}
	if c.m == nil {
		c.m = make(map[reflect.Type][]any)
	}
	c.s = append(c.s, vs...)
	for _, a := range vs {
		ty := reflect.TypeOf(a)
		c.m[ty] = append(c.m[ty], a)
	}
}

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

type namedCollection[T any] struct {
	s []T
	m map[string]T
}

func (c *namedCollection[T]) getOrMake(n string, f func(string) T) T {
	if c.m == nil {
		c.m = make(map[string]T)
	}

	v, ok := c.m[n]
	if !ok {
		v = f(n)
		c.m[n] = v
		c.s = append(c.s, v)
	}

	return v
}

//

type TypeAnnotations struct {
	mtx sync.Mutex
	ty  reflect.Type

	c  annotationCollection
	fc namedCollection[*FieldAnnotations]
	mc namedCollection[*MethodAnnotations]
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
	defer a.ta.mtx.Unlock()

	fa := a.ta.fc.getOrMake(n, func(s string) *FieldAnnotations {
		f, ok := a.ta.ty.FieldByName(n)
		if !ok {
			panic(fmt.Errorf("type %s has no field %s", a.ta.ty, n))
		}
		return &FieldAnnotations{f: f}
	})
	fa.s = append(fa.s, anns...)

	return a
}

func (a *annotator) Method(n string, anns ...any) Annotator {
	a.ta.mtx.Lock()
	defer a.ta.mtx.Unlock()

	ma := a.ta.mc.getOrMake(n, func(s string) *MethodAnnotations {
		m, ok := a.ta.ty.MethodByName(n)
		if !ok {
			panic(fmt.Errorf("type %s has no method %s", a.ta.ty, n))
		}
		return &MethodAnnotations{m: m}
	})
	ma.s = append(ma.s, anns...)

	return a
}
