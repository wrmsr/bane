package def

import (
	"fmt"
	"reflect"

	opt "github.com/wrmsr/bane/pkg/utils/optional"
)

//

type SpecError struct {
	err error
}

func (e SpecError) Error() string { return e.err.Error() }
func (e SpecError) Unwrap() error { return e.err }

//

type FieldSpec struct {
	name string
	defs []FieldDef

	ty  reflect.Type
	dfl opt.Optional[any]
}

func newFieldSpec(name string, defs []FieldDef) *FieldSpec {
	var ty opt.Optional[reflect.Type]
	var dfl opt.Optional[any]

	for _, d := range defs {
		for _, o := range d.Opts {
			switch o := o.(type) {

			case TypeOpt:
				ty.IfPresent(func() { panic(RegistryError{fmt.Errorf("duplicate types: %s", name)}) })
				ty = opt.Just(o.Type)

			case DefaultOpt:
				dfl.IfPresent(func() { panic(RegistryError{fmt.Errorf("duplicate defaults: %s", name)}) })
				dfl = opt.Just(o.Val)

			default:
				panic(RegistryError{fmt.Errorf("%T", o)})
			}
		}
	}

	if !ty.Present() {
		if !dfl.Present() {
			panic(RegistryError{fmt.Errorf("no type: %s", name)})
		}
		ty = opt.Just(reflect.TypeOf(dfl.Value()))
	}

	return &FieldSpec{
		name: name,
		defs: defs,

		ty:  ty.Value(),
		dfl: dfl,
	}
}

func (fs FieldSpec) Default() any {
	return fs.dfl.Value()
}

//

type StructSpec struct {
	name string
	defs []StructDef

	fields map[string]*FieldSpec
	inits  []any
}

func newStructSpec(name string, defs []StructDef) *StructSpec {
	fdm := make(map[string][]FieldDef)
	var inits []any

	for _, d := range defs {
		for _, o := range d.Opts {
			switch o := o.(type) {

			case FieldDef:
				fdm[o.Name] = append(fdm[d.Name], o)

			case InitOpt:
				inits = append(inits, o.Fn)

			default:
				panic(RegistryError{fmt.Errorf("%T", o)})
			}
		}
	}

	fields := make(map[string]*FieldSpec, len(fdm))
	for fn, fds := range fdm {
		fields[fn] = newFieldSpec(fn, fds)
	}

	return &StructSpec{
		name: name,
		defs: defs,

		fields: fields,
		inits:  inits,
	}
}

func (ss StructSpec) Field(name string) *FieldSpec {
	if ss, ok := ss.fields[name]; ok {
		return ss
	}
	panic(SpecError{fmt.Errorf("field not found :%s", name)})
}

func (ss StructSpec) Inits() []any {
	return ss.inits
}

//

type PackageSpec struct {
	name string
	defs []PackageDef

	structs map[string]*StructSpec
}

func newPackageSpec(name string, defs []PackageDef) *PackageSpec {
	sdm := make(map[string][]StructDef)
	var ex opt.Optional[X_packageExpect]

	for _, d := range defs {
		switch d := d.(type) {

		case StructDef:
			sdm[d.Name] = append(sdm[d.Name], d)

		case X_packageExpect:
			ex.IfPresent(func() { panic(RegistryError{fmt.Errorf("duplicate expect: %s", name)}) })
			ex = opt.Just(d)

		default:
			panic(RegistryError{fmt.Errorf("%T", d)})
		}
	}

	structs := make(map[string]*StructSpec, len(sdm))
	for sn, sds := range sdm {
		structs[sn] = newStructSpec(sn, sds)
	}

	ps := &PackageSpec{
		name: name,
		defs: defs,

		structs: structs,
	}

	if ex.Present() {
		ex.Value().check(ps)
	}

	return ps
}

func (ps PackageSpec) Struct(name string) *StructSpec {
	if ss, ok := ps.structs[name]; ok {
		return ss
	}
	panic(SpecError{fmt.Errorf("struct not found :%s", name)})
}
