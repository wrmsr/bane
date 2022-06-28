package def

import (
	"reflect"
)

//

type PackageDef interface {
	isPackageDef()
}

//

type StructOpt interface {
	isStructOpt()
}

type StructDef struct {
	Name string
	Opts []StructOpt
}

func (sd StructDef) isPackageDef() {}

func Struct(name string, opts ...StructOpt) StructDef {
	return addPackageDef(globalRegistry, getCallingPackage(), StructDef{
		Name: name,
		Opts: opts,
	})
}

//

type ReceiverOpt struct {
	Name string
}

func (o ReceiverOpt) isStructOpt() {}

func Receiver(name string) ReceiverOpt {
	return ReceiverOpt{
		Name: name,
	}
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
	Ty any
}

func (o TypeOpt) isFieldOpt() {}

func Type[T any]() TypeOpt {
	var z T
	return TypeOpt{Ty: reflect.TypeOf(z)}
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
