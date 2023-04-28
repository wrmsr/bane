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

func (k MemberKind) String() string {
	switch k {
	case FieldMember:
		return "field"
	case MethodMember:
		return "method"
	}
	panic("<invalid>")
}

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
	mc keyedCollection[Member, *memberAnnotations]
}

func (ta *TypeAnnotations) Type() reflect.Type { return ta.ty }

//

type TypeAnnotationsMap struct {
	mtx sync.Mutex

	m map[reflect.Type]*TypeAnnotations
}

func NewTypeAnnotationsMap() *TypeAnnotationsMap {
	return &TypeAnnotationsMap{}
}

func (tam *TypeAnnotationsMap) Get(ty reflect.Type) *TypeAnnotations {
	return tam.get(ty, false)
}

func (tam *TypeAnnotationsMap) get(ty reflect.Type, orMake bool) *TypeAnnotations {
	tam.mtx.Lock()

	if tam.m == nil {
		tam.m = make(map[reflect.Type]*TypeAnnotations)
	}

	ta := tam.m[ty]
	if ta == nil && orMake {
		ta = &TypeAnnotations{ty: ty}
		tam.m[ty] = ta
	}

	tam.mtx.Unlock()
	return ta
}

//

var _global = NewTypeAnnotationsMap()

func Global() *TypeAnnotationsMap {
	return _global
}
