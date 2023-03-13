package def

import (
	"fmt"
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
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
	meta ctr.MutMap[reflect.Type, any]
}

func NewFieldSpec(name string, defs []FieldDef) *FieldSpec {
	fs := &FieldSpec{
		name: name,
		defs: defs,
		meta: ctr.NewMutStdMap[reflect.Type, any](nil),
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
				fs.meta.Put(reflect.TypeOf(o.Value), o.Value)

			default:
				panic(RegistryError{fmt.Errorf("%T", o)})
			}
		}
	}

	return fs
}

func (fs FieldSpec) Name() string                     { return fs.name }
func (fs FieldSpec) Defs() []FieldDef                 { return fs.defs }
func (fs FieldSpec) Opts() []FieldOpt                 { return fs.opts }
func (fs FieldSpec) Type() any                        { return fs.ty }
func (fs FieldSpec) Default() any                     { return fs.dfl }
func (fs FieldSpec) Meta() ctr.Map[reflect.Type, any] { return fs.meta }

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
	ty   any
	defs []StructDef
	opts []StructOpt

	fields       []*FieldSpec
	fieldsByName map[string]*FieldSpec

	receiver string
	inits    []any
	meta     ctr.MutMap[reflect.Type, any]
}

func NewStructSpec(ty any, defs []StructDef) *StructSpec {
	ss := &StructSpec{
		ty:   ty,
		defs: defs,

		meta: ctr.NewMutStdMap[reflect.Type, any](nil),
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
				fdm[o.Name] = append(fdm[o.Name], o)

			case InitOpt:
				ss.inits = append(ss.inits, o.Fn)

			case MetaOpt:
				ss.meta.Put(reflect.TypeOf(o.Value), o.Value)

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

func (ss StructSpec) Type() any                        { return ss.ty }
func (ss StructSpec) Defs() []StructDef                { return ss.defs }
func (ss StructSpec) Opts() []StructOpt                { return ss.opts }
func (ss StructSpec) Fields() []*FieldSpec             { return ss.fields }
func (ss StructSpec) Receiver() string                 { return ss.receiver }
func (ss StructSpec) Inits() []any                     { return ss.inits }
func (ss StructSpec) Meta() ctr.Map[reflect.Type, any] { return ss.meta }

func (ss StructSpec) Field(name string) *FieldSpec {
	if ss, ok := ss.fieldsByName[name]; ok {
		return ss
	}
	panic(SpecError{fmt.Errorf("field not found :%s", name)})
}

//

type EnumSpec struct {
	ty   any
	defs []EnumDef
	opts []EnumOpt
}

func (e EnumSpec) Ty() any         { return e.ty }
func (e EnumSpec) Defs() []EnumDef { return e.defs }
func (e EnumSpec) Opts() []EnumOpt { return e.opts }

func NewEnumSpec(ty any, defs []EnumDef) *EnumSpec {
	es := &EnumSpec{
		ty:   ty,
		defs: defs,
	}

	return es
}

//

type PackageSpec struct {
	name string
	defs []PackageDef

	structs     []*StructSpec
	structsByTy map[any]*StructSpec

	inlines     map[any]*InlineDef
	withInlines map[any]*WithInlineDef

	enums     []*EnumSpec
	enumsByTy map[any]*EnumSpec
}

func NewPackageSpec(name string, defs []PackageDef) *PackageSpec {
	ps := &PackageSpec{
		name: name,
		defs: defs,
	}

	var sns []any
	sdm := make(map[any][]StructDef)
	var ets []any
	edm := make(map[any][]EnumDef)

	ilm := make(map[any]*InlineDef)
	wilm := make(map[any]*WithInlineDef)

	for _, d := range defs {
		switch d := d.(type) {

		case StructDef:
			if _, ok := sdm[d.Ty]; !ok {
				sns = append(sns, d.Ty)
			}
			sdm[d.Ty] = append(sdm[d.Ty], d)

		case EnumDef:
			if _, ok := edm[d.Ty]; !ok {
				ets = append(ets, d.Ty)
			}
			edm[d.Ty] = append(edm[d.Ty], d)

		case InlineDef:
			for _, f := range d.Fns {
				ilm[f] = &d
			}

		case WithInlineDef:
			for _, f := range d.Fns {
				wilm[f] = &d
			}

		case SpecializeDef:
			// FIXME:

		default:
			panic(RegistryError{fmt.Errorf("%T", d)})
		}
	}

	ps.structsByTy = make(map[any]*StructSpec, len(sdm))
	for _, sn := range sns {
		ss := NewStructSpec(sn, sdm[sn])
		ps.structs = append(ps.structs, ss)
		ps.structsByTy[sn] = ss
	}

	ps.enumsByTy = make(map[any]*EnumSpec, len(edm))
	for _, et := range ets {
		es := NewEnumSpec(et, edm[et])
		ps.enums = append(ps.enums, es)
		ps.enumsByTy[et] = es
	}

	ps.inlines = ilm
	ps.withInlines = wilm

	return ps
}

func (ps PackageSpec) Name() string { return ps.name }

func (ps PackageSpec) Structs() []*StructSpec { return ps.structs }

func (ps PackageSpec) Struct(ty any) *StructSpec {
	if ss, ok := ps.structsByTy[ty]; ok {
		return ss
	}
	panic(SpecError{fmt.Errorf("struct not found :%s", ty)})
}

func (ps PackageSpec) Enums() []*EnumSpec { return ps.enums }

func (ps PackageSpec) Enum(ty any) *EnumSpec {
	if ss, ok := ps.enumsByTy[ty]; ok {
		return ss
	}
	panic(SpecError{fmt.Errorf("enum not found :%s", ty)})
}

func (ps PackageSpec) Inlines() map[any]*InlineDef         { return ps.inlines }
func (ps PackageSpec) WithInlines() map[any]*WithInlineDef { return ps.withInlines }
