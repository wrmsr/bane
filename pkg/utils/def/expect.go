package def

import "reflect"

//

type X_fieldExpect struct {
	Ty reflect.Type

	HasDefault bool
}

func (ex X_fieldExpect) check(fs *FieldSpec) {

}

//

type X_structExpect struct {
	Fields map[string]X_fieldExpect

	NumInits int
}

func (ex X_structExpect) check(ss *StructSpec) {

}

//

type X_packageExpect struct {
	Structs map[string]X_structExpect
}

func (ex X_packageExpect) check(ps *PackageSpec) {

}

func (ex X_packageExpect) isPackageDef() {}
