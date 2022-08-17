package gen

import opt "github.com/wrmsr/bane/pkg/util/optional"

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

type Array struct {
	Len  opt.Optional[int]
	Elem Type

	type_
}

func NewArray(len_ opt.Optional[int], elem Type) Array {
	return Array{
		Len:  len_,
		Elem: elem,
	}
}

//

type FuncType struct {
	Func Func

	type_
}

func NewFuncType(func_ Func) FuncType {
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

func NewMap(key, value Type) Map {
	return Map{
		Key:   key,
		Value: value,
	}
}

//

type Ptr struct {
	Elem Type

	type_
}

func NewPtr(elem Type) Ptr {
	return Ptr{
		Elem: elem,
	}
}

//

type NameType struct {
	Name Ident

	type_
}

func NewNameType(name Ident) NameType {
	return NameType{
		Name: name,
	}
}

//

type Slice struct {
	Elem Type

	type_
}

func NewSlice(elem Type) Slice {
	return Slice{
		Elem: elem,
	}
}
