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
	Ty   any
	Opts []StructOpt
}

func (sd StructDef) isPackageDef() {}

func Struct[T any](opts ...StructOpt) StructDef {
	var z T
	return addPackageDef(globalRegistry, getCallingPackage(), StructDef{
		Ty:   reflect.TypeOf(z),
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
	Ty  any
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

//

type MetaOpt struct {
	Value any
}

func (o MetaOpt) isStructOpt() {}
func (o MetaOpt) isFieldOpt()  {}

func Meta(value any) MetaOpt {
	return MetaOpt{
		Value: value,
	}
}

//

type EnumOpt interface {
	isEnumOpt()
}

type EnumDef struct {
	Ty   any
	Opts []EnumOpt
}

func (ed EnumDef) isPackageDef() {}

func Enum[T any](opts ...EnumOpt) any {
	var z T
	return addPackageDef(globalRegistry, getCallingPackage(), EnumDef{
		Ty:   reflect.TypeOf(z),
		Opts: opts,
	})
}

//

type InlineDef struct {
	Fns []any
}

func (id InlineDef) isPackageDef() {}

func Inline(fns ...any) any {
	return InlineDef{
		Fns: fns,
	}
}

//

type WithInlineDef struct {
	Fns []any
}

func (id WithInlineDef) isPackageDef() {}

func WithInline(fns ...any) any {
	return WithInlineDef{
		Fns: fns,
	}
}

//

type ConstGenericDef struct {
	Ty  any
	Cvs []any
}

func (ed ConstGenericDef) isPackageDef() {}

func ConstGeneric[T any](cvs ...any) any {
	var z T
	return addPackageDef(globalRegistry, getCallingPackage(), ConstGenericDef{
		Ty:  reflect.TypeOf(z),
		Cvs: cvs,
	})
}
