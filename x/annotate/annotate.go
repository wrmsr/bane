package annotate

import (
	"fmt"
	"reflect"
	"sync"
)

//

type MemberKind int8

const (
	FieldMember MemberKind = iota + 1
	MethodMember
)

func Field(n string) Member  { return Member{K: FieldMember, N: n} }
func Method(n string) Member { return Member{K: MethodMember, N: n} }

type Member struct {
	K MemberKind
	N string
}

func lookupMember(ty reflect.Type, m Member) any {
	switch m.K {
	case FieldMember:
		r, ok := ty.FieldByName(m.N)
		if !ok {
			panic(fmt.Errorf("type %s has no field %s", ty, m.N))
		}
		return r
	case MethodMember:
		r, ok := ty.MethodByName(m.N)
		if !ok {
			panic(fmt.Errorf("type %s has no method %s", ty, m.N))
		}
		return r
	}
	panic(m)
}

//

type memberAnnotations struct {
	m Member
	r any
	c typeIndexedCollection
}

type TypeAnnotations struct {
	ty reflect.Type

	mtx sync.Mutex

	c  typeIndexedCollection
	mm keyedCollection[Member, *memberAnnotations]
}

func (ta *TypeAnnotations) Type() reflect.Type { return ta.ty }

//

type TypeAnnotationsMap struct {
	mtx sync.Mutex

	m map[reflect.Type]*TypeAnnotations
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
		ta = &TypeAnnotations{ty: ty}
		_global.m[ty] = ta
	}
	_global.mtx.Unlock()

	if len(anns) > 0 {
		ta.mtx.Lock()
		ta.c.add(anns...)
		ta.mtx.Unlock()
	}

	return &annotator{ta: ta}
}

func (a *annotator) Member(m Member, anns ...any) Annotator {
	a.ta.mtx.Lock()
	defer a.ta.mtx.Unlock()

	ma := a.ta.mm.getOrMake(m, func() *memberAnnotations {
		r := lookupMember(a.ta.ty, m)
		return &memberAnnotations{m: m, r: r}
	})
	ma.c.add(anns...)

	return a

}

func (a *annotator) Field(n string, anns ...any) Annotator  { return a.Member(Field(n), anns...) }
func (a *annotator) Method(n string, anns ...any) Annotator { return a.Member(Method(n), anns...) }
