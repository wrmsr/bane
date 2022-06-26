package gogen

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
	type_
	Len  opt.Optional[int]
	Elem Type
}

func NewArray(len_ opt.Optional[int], elem Type) Array {
	return Array{
		Len:  len_,
		Elem: elem,
	}
}

//

type Map struct {
	type_
	Key   Type
	Value Type
}

func NewMap(key, value Type) Map {
	return Map{
		Key:   key,
		Value: value,
	}
}

//

type Ptr struct {
	type_
	Elem Type
}

func NewPtr(elem Type) Ptr {
	return Ptr{
		Elem: elem,
	}
}

//

type NameType struct {
	type_
	Name Ident
}

func NewNameType(name Ident) NameType {
	return NameType{
		Name: name,
	}
}

//

type Slice struct {
	type_
	Elem Type
}

func NewSlice(elem Type) Slice {
	return Slice{
		Elem: elem,
	}
}
