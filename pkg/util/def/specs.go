package def

import (
	"fmt"
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
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
	opts []FieldOpt

	ty   any
	dfl  any
	meta ctr.MutTypeMap[any]
}

func NewFieldSpec(name string, defs []FieldDef) *FieldSpec {
	fs := &FieldSpec{
		name: name,
		defs: defs,
		meta: ctr.NewMutTypeMap[any](nil),
	}

	for _, d := range defs {
		for _, o := range d.Opts {
			fs.opts = append(fs.opts, o)

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

			case MetaOpt:
				fs.meta.Put(o.Value)

			default:
				panic(RegistryError{fmt.Errorf("%T", o)})
			}
		}
	}

	return fs
}

func (fs FieldSpec) Name() string           { return fs.name }
func (fs FieldSpec) Defs() []FieldDef       { return fs.defs }
func (fs FieldSpec) Opts() []FieldOpt       { return fs.opts }
func (fs FieldSpec) Type() any              { return fs.ty }
func (fs FieldSpec) Default() any           { return fs.dfl }
func (fs FieldSpec) Meta() ctr.TypeMap[any] { return fs.meta }

func (fs FieldSpec) RuntimeType() reflect.Type {
	if fs.ty != nil {
		return fs.ty.(reflect.Type)
	}
	if fs.dfl != nil {
		return reflect.TypeOf(fs.dfl)
	}
	panic("no type or default set")
}

//

type StructSpec struct {
	name string
	defs []StructDef
	opts []StructOpt

	fields       []*FieldSpec
	fieldsByName map[string]*FieldSpec

	receiver string
	inits    []any
	meta     ctr.MutTypeMap[any]
}

func NewStructSpec(name string, defs []StructDef) *StructSpec {
	ss := &StructSpec{
		name: name,
		defs: defs,

		meta: ctr.NewMutTypeMap[any](nil),
	}

	var fns []string
	fdm := make(map[string][]FieldDef)

	for _, d := range defs {
		for _, o := range d.Opts {
			ss.opts = append(ss.opts, o)

			switch o := o.(type) {

			case ReceiverOpt:
				ss.receiver = o.Name

			case FieldDef:
				if _, ok := fdm[o.Name]; !ok {
					fns = append(fns, o.Name)
				}
				fdm[o.Name] = append(fdm[d.Name], o)

			case InitOpt:
				ss.inits = append(ss.inits, o.Fn)

			case MetaOpt:
				ss.meta.Put(o.Value)

			default:
				panic(RegistryError{fmt.Errorf("%T", o)})
			}
		}
	}

	ss.fieldsByName = make(map[string]*FieldSpec, len(fdm))
	for _, fn := range fns {
		fs := NewFieldSpec(fn, fdm[fn])
		ss.fields = append(ss.fields, fs)
		ss.fieldsByName[fn] = fs
	}

	return ss
}

func (ss StructSpec) Name() string           { return ss.name }
func (ss StructSpec) Defs() []StructDef      { return ss.defs }
func (ss StructSpec) Opts() []StructOpt      { return ss.opts }
func (ss StructSpec) Fields() []*FieldSpec   { return ss.fields }
func (ss StructSpec) Receiver() string       { return ss.receiver }
func (ss StructSpec) Inits() []any           { return ss.inits }
func (ss StructSpec) Meta() ctr.TypeMap[any] { return ss.meta }

func (ss StructSpec) Field(name string) *FieldSpec {
	if ss, ok := ss.fieldsByName[name]; ok {
		return ss
	}
	panic(SpecError{fmt.Errorf("field not found :%s", name)})
}

//

type PackageSpec struct {
	name string
	defs []PackageDef

	structs       []*StructSpec
	structsByName map[string]*StructSpec
}

func NewPackageSpec(name string, defs []PackageDef) *PackageSpec {
	ps := &PackageSpec{
		name: name,
		defs: defs,
	}

	var sns []string
	sdm := make(map[string][]StructDef)
	var ex opt.Optional[X_packageExpect]

	for _, d := range defs {
		switch d := d.(type) {

		case StructDef:
			if _, ok := sdm[d.Name]; !ok {
				sns = append(sns, d.Name)
			}
			sdm[d.Name] = append(sdm[d.Name], d)

		case X_packageExpect:
			ex.IfPresent(func() { panic(RegistryError{fmt.Errorf("duplicate expect: %s", name)}) })
			ex = opt.Just(d)

		default:
			panic(RegistryError{fmt.Errorf("%T", d)})
		}
	}

	ps.structsByName = make(map[string]*StructSpec, len(sdm))
	for _, sn := range sns {
		ss := NewStructSpec(sn, sdm[sn])
		ps.structs = append(ps.structs, ss)
		ps.structsByName[sn] = ss
	}

	if ex.Present() {
		ex.Value().check(ps)
	}

	return ps
}

func (ps PackageSpec) Name() string           { return ps.name }
func (ps PackageSpec) Structs() []*StructSpec { return ps.structs }

func (ps PackageSpec) Struct(name string) *StructSpec {
	if ss, ok := ps.structsByName[name]; ok {
		return ss
	}
	panic(SpecError{fmt.Errorf("struct not found :%s", name)})
}
