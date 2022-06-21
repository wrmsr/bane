package def

import (
	"reflect"

	opt "github.com/wrmsr/bane/pkg/utils/optional"
)

//

type FieldSpec struct {
	name string
	defs []FieldDef

	ty  reflect.Type
	dfl opt.Optional[any]
}

func newFieldSpec(name string, defs []FieldDef) *FieldSpec {
	return &FieldSpec{
		name: name,
		defs: defs,
	}
}

//

type StructSpec struct {
	name string
	defs []StructDef

	fields map[string]*FieldSpec
	inits  []any
}

func newStructSpec(name string, defs []StructDef) *StructSpec {
	return &StructSpec{
		name: name,
	}
}

//

type PackageSpec struct {
	name string
	defs []PackageDef

	structs map[string]*StructSpec
}

func newPackageSpec(name string, defs []PackageDef) *PackageSpec {
	return &PackageSpec{
		name: name,
		defs: defs,
	}
}
