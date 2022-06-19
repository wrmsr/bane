package structs

import "reflect"

type RootDef interface {
	isRootDef()
}

func Def(defs ...RootDef) any {
	return nil
}

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
	return StructDef{
		Name: name,
		Opts: opts,
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

type TypeOpt[T any] struct {
	Type reflect.Type
}

func (o TypeOpt[T]) isFieldOpt() {}

func Type[T any]() TypeOpt[T] {
	var z T
	return TypeOpt[T]{Type: reflect.TypeOf(z)}
}
