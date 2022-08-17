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

func MapOf(key, value Type) Map {
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

func PtrOf(elem Type) Ptr {
	return Ptr{
		Elem: elem,
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

func SliceOf(elem Type) Slice {
	return Slice{
		Elem: elem,
	}
}
