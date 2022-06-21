package def

import (
	"fmt"
	"reflect"
)

//

type X_structExpectsDef struct {
	Name string

	Fields   []string
	Defaults []string

	Inits int
}

func (sed X_structExpectsDef) isRootDef() {}

func (sed X_structExpectsDef) Register() any {
	addPackageRootDef(globalRegistry, getCallingPackage(), sed)
	return nil
}

//

type X_fieldInit struct {
	Type reflect.Type

	Default any
}

type X_structInit struct {
	Fields map[string]X_fieldInit

	Inits []any
}

//

type structInitBuilder struct {
	reg *packageRegistry
}

func (b *structInitBuilder) buildInit1(sds []StructDef) X_structInit {
	r := X_structInit{
		Fields: make(map[string]X_fieldInit),
	}

	fm := make(map[string][]FieldDef)
	for _, sd := range sds {
		for _, so := range sd.Opts {
			switch so := so.(type) {
			case FieldDef:
				fm[so.Name] = append(fm[so.Name], so)
			case InitOpt:
				r.Inits = append(r.Inits, so.Fn)
			default:
				panic(RegistryError{fmt.Errorf("%T", so)})
			}
		}
	}

	for n, fds := range fm {
		fi := X_fieldInit{}
		hasType := false
		hasDefault := false

		for _, fd := range fds {
			for _, fo := range fd.Opts {
				switch fo := fo.(type) {
				case TypeOpt:
					if hasType {
						panic(RegistryError{fmt.Errorf("duplicate types: %s", n)})
					}
					hasType = true
					fi.Type = fo.Type
				case DefaultOpt:
					if hasDefault {
						panic(RegistryError{fmt.Errorf("duplicate default: %s", n)})
					}
					hasDefault = true
					fi.Default = fo.Val
				default:
					panic(RegistryError{fmt.Errorf("%T", fo)})
				}
			}
		}

		if !hasType {
			if !hasDefault {
				panic(RegistryError{fmt.Errorf("no type: %s", n)})
			}
			fi.Type = reflect.TypeOf(fi.Default)
		}

		r.Fields[n] = fi
	}

	// TODO: expects

	return r
}

func (b *structInitBuilder) buildInit() map[string]X_structInit {
	m := make(map[string][]StructDef)
	for _, d := range b.reg.rootDefsMap[reflect.TypeOf(StructDef{})] {
		sd := d.(StructDef)
		m[sd.Name] = append(m[sd.Name], sd)
	}

	r := make(map[string]X_structInit)
	for n, sds := range m {
		r[n] = b.buildInit1(sds)
	}

	return r
}
