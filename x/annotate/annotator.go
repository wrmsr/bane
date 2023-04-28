package annotate

import (
	"fmt"
	"reflect"
)

type Annotator interface {
	Field(n string, anns ...any) Annotator
	Method(n string, anns ...any) Annotator
}

type annotator struct {
	ta *TypeAnnotations
}

var _ Annotator = &annotator{}

func On[T any](anns ...any) Annotator {
	var z T
	ty := reflect.TypeOf(z)
	if ty.Kind() != reflect.Struct {
		panic(fmt.Errorf("must be struct: %v", ty))
	}

	ta := _global.get(ty, true)

	if len(anns) > 0 {
		ta.mtx.Lock()
		ta.c.add(anns...)
		ta.mtx.Unlock()
	}

	return &annotator{ta: ta}
}

func (a *annotator) Member(m Member, anns ...any) Annotator {
	a.ta.mtx.Lock()
	defer a.ta.mtx.Unlock()

	ma := a.ta.mc.getOrMake(m, func() *memberAnnotations {
		r := lookupMember(a.ta.ty, m)
		return &memberAnnotations{m: m, r: r}
	})
	ma.c.add(anns...)

	return a

}

func (a *annotator) Field(n string, anns ...any) Annotator  { return a.Member(Field(n), anns...) }
func (a *annotator) Method(n string, anns ...any) Annotator { return a.Member(Method(n), anns...) }
