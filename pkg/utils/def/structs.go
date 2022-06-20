package def

import (
	"reflect"
)

//

type StructOpt interface {
	isStructOpt()
}

type StructDef struct {
	Name string
	Opts []StructOpt
}

func (sd StructDef) isRootDef() {}

func Struct(name string, opts ...StructOpt) StructDef {
	//pkg := getCallingPackage()
	return StructDef{
		Name: name,
		Opts: opts,
	}
}

//

type InitOpt struct {
	Fn any
}

func (o InitOpt) isStructOpt() {}

func Init(fn any) InitOpt {
	return InitOpt{Fn: fn}
}

//

type FieldOpt interface {
	isFieldOpt()
}

type FieldDef struct {
	Name string
	Opts []FieldOpt
}

func (fd FieldDef) isStructOpt() {}

func Field(name string, opts ...FieldOpt) FieldDef {
	return FieldDef{
		Name: name,
		Opts: opts,
	}
}

//

type TypeOpt[T any] struct {
	Type reflect.Type
}

func (o TypeOpt[T]) isFieldOpt() {}

func Type[T any]() TypeOpt[T] {
	var z T
	return TypeOpt[T]{Type: reflect.TypeOf(z)}
}

//

type DefaultOpt[T any] struct {
	Val T
}

func (o DefaultOpt[T]) isFieldOpt() {}

func Default[T any](val T) DefaultOpt[T] {
	return DefaultOpt[T]{Val: val}
}
