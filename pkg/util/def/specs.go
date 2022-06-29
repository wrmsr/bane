package def

import (
	"fmt"
	"reflect"

	opt "github.com/wrmsr/bane/pkg/util/optional"
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

	ty  any
	dfl any
}

func NewFieldSpec(name string, defs []FieldDef) *FieldSpec {
	fs := &FieldSpec{
		name: name,
		defs: defs,
	}

	for _, d := range defs {
		for _, o := range d.Opts {
			switch o := o.(type) {

			case TypeOpt:
				if fs.ty != nil {
					panic(RegistryError{fmt.Errorf("duplicate types: %s", name)})
				}
				fs.ty = o.Ty

			case DefaultOpt:
				if fs.dfl != nil {
					panic(RegistryError{fmt.Errorf("duplicate defaults: %s", name)})
				}
				fs.dfl = o.Val

			default:
				panic(RegistryError{fmt.Errorf("%T", o)})
			}
		}
	}

	if fs.ty == nil {
		if fs.dfl == nil {
			panic(RegistryError{fmt.Errorf("no type: %s", name)})
		}
		fs.ty = reflect.TypeOf(fs.dfl)
	}

	return fs
}

func (fs FieldSpec) Name() string { return fs.name }
func (fs FieldSpec) Type() any    { return fs.ty }
func (fs FieldSpec) Default() any { return fs.dfl }

//

type StructSpec struct {
	name string
	defs []StructDef

	receiver string
	fields   map[string]*FieldSpec
	inits    []any
}

func NewStructSpec(name string, defs []StructDef) *StructSpec {
	ss := &StructSpec{
		name: name,
		defs: defs,
	}

	fdm := make(map[string][]FieldDef)

	for _, d := range defs {
		for _, o := range d.Opts {
			switch o := o.(type) {

			case ReceiverOpt:
				ss.receiver = o.Name

			case FieldDef:
				fdm[o.Name] = append(fdm[d.Name], o)

			case InitOpt:
				ss.inits = append(ss.inits, o.Fn)

			default:
				panic(RegistryError{fmt.Errorf("%T", o)})
			}
		}
	}

	ss.fields = make(map[string]*FieldSpec, len(fdm))
	for fn, fds := range fdm {
		ss.fields[fn] = NewFieldSpec(fn, fds)
	}

	return ss
}

func (ss StructSpec) Name() string                  { return ss.name }
func (ss StructSpec) Receiver() string              { return ss.receiver }
func (ss StructSpec) Fields() map[string]*FieldSpec { return ss.fields }
func (ss StructSpec) Inits() []any                  { return ss.inits }

func (ss StructSpec) Field(name string) *FieldSpec {
	if ss, ok := ss.fields[name]; ok {
		return ss
	}
	panic(SpecError{fmt.Errorf("field not found :%s", name)})
}

//

type PackageSpec struct {
	name string
	defs []PackageDef

	structs map[string]*StructSpec
}

func NewPackageSpec(name string, defs []PackageDef) *PackageSpec {
	ps := &PackageSpec{
		name: name,
		defs: defs,
	}

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

	ps.structs = make(map[string]*StructSpec, len(sdm))
	for sn, sds := range sdm {
		ps.structs[sn] = NewStructSpec(sn, sds)
	}

	if ex.Present() {
		ex.Value().check(ps)
	}

	return ps
}

func (ps PackageSpec) Name() string                    { return ps.name }
func (ps PackageSpec) Structs() map[string]*StructSpec { return ps.structs }

func (ps PackageSpec) Struct(name string) *StructSpec {
	if ss, ok := ps.structs[name]; ok {
		return ss
	}
	panic(SpecError{fmt.Errorf("struct not found :%s", name)})
}
