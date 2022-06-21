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
	return addPackageRootDef(globalRegistry, getCallingPackage(), StructDef{
		Name: name,
		Opts: opts,
	})
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

type TypeOpt struct {
	Type reflect.Type
}

func (o TypeOpt) isFieldOpt() {}

func Type[T any]() TypeOpt {
	var z T
	return TypeOpt{Type: reflect.TypeOf(z)}
}

//

type DefaultOpt struct {
	Val any
}

func (o DefaultOpt) isFieldOpt() {}

func Default(val any) DefaultOpt {
	return DefaultOpt{Val: val}
}

//

type InitOpt struct {
	Fn any
}

func (o InitOpt) isStructOpt() {}

func Init(fn any) InitOpt {
	return InitOpt{Fn: fn}
}
