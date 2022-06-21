package def

import (
	"fmt"
	"reflect"

	"github.com/wrmsr/bane/pkg/utils/maps"
	"github.com/wrmsr/bane/pkg/utils/slices"
)

//

type ExpectError struct {
	err error
}

func (e ExpectError) Error() string { return e.err.Error() }
func (e ExpectError) Unwrap() error { return e.err }

func checkExpectedStrs(l, r []string, msg string) {
	slices.Sort(l)
	slices.Sort(r)
	if !slices.Equal(l, r) {
		panic(ExpectError{fmt.Errorf("%s: %v != %v", msg, l, r)})
	}
}

//

type X_fieldExpect struct {
	Ty reflect.Type

	HasDefault bool
}

func (fx X_fieldExpect) check(fs *FieldSpec) {
	if fs.dfl.Present() != fx.HasDefault {
		panic(ExpectError{fmt.Errorf("%s defaults differ", fs.name)})
	}
}

//

type X_structExpect struct {
	Fields map[string]X_fieldExpect

	NumInits int
}

func (sx X_structExpect) check(ss *StructSpec) {
	checkExpectedStrs(
		maps.Keys(ss.fields),
		maps.Keys(sx.Fields),
		fmt.Sprintf("%s fields differ", ss.name),
	)
	for n, fx := range sx.Fields {
		fx.check(ss.Field(n))
	}

	if len(ss.inits) != sx.NumInits {
		panic(ExpectError{fmt.Errorf("%s num inits differ:  %d != %d", ss.name, len(ss.inits), sx.NumInits)})
	}
}

//

type X_packageExpect struct {
	Structs map[string]X_structExpect
}

func (px X_packageExpect) check(ps *PackageSpec) {
	checkExpectedStrs(
		maps.Keys(ps.structs),
		maps.Keys(px.Structs),
		fmt.Sprintf("%s structs differ", ps.name),
	)
	for n, sx := range px.Structs {
		sx.check(ps.Struct(n))
	}
}

func (px X_packageExpect) isPackageDef() {}
