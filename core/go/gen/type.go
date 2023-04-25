package gen

import (
	bt "github.com/wrmsr/bane/core/types"
)

//

type Type interface {
	Node
	isType()
}

type type_ struct {
	node
}

func (t type_) isType() {}

//

func TypeOf(o any) Type {
	if o, ok := o.(Type); ok {
		return o
	}

	switch o.(type) {
	case Ident:
		return NameTypeOf(o)
	case string:
		return NameTypeOf(IdentOf(o))
	}

	panic(o)
}

//

type Array struct {
	Len  bt.Optional[int]
	Elem Type

	type_
}

//

type FuncType struct {
	Func Func

	type_
}

func FuncTypeOf(func_ Func) FuncType {
	return FuncType{
		Func: func_,
	}
}

//

type Map struct {
	Key   Type
	Value Type

	type_
}

func MapOf(key, value any) Map {
	return Map{
		Key:   TypeOf(key),
		Value: TypeOf(value),
	}
}

//

type Ptr struct {
	Elem Type

	type_
}

func PtrOf(elem any) Ptr {
	return Ptr{
		Elem: TypeOf(elem),
	}
}

//

type NameType struct {
	Name Ident

	type_
}

func NameTypeOf(name any) NameType {
	return NameType{
		Name: IdentOf(name),
	}
}

//

type Slice struct {
	Elem Type

	type_
}

func SliceOf(elem any) Slice {
	return Slice{
		Elem: TypeOf(elem),
	}
}
