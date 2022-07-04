package structs

import (
	"errors"
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type FieldInfo struct {
	field reflect.StructField

	name      string
	nameBytes []byte // []byte(name)

	equalFold func(s, t []byte) bool // bytes.EqualFold or equivalent

	index []int

	path     string
	embedded bool

	parent   *FieldInfo
	children []*FieldInfo

	si *StructInfo
}

func (fi FieldInfo) GetValue(v any) (reflect.Value, bool) {
	rv, ok := rfl.UnwrapPointerValue(rfl.AsValue(v))
	if !ok {
		return reflect.Value{}, false
	}

	fv := rv
	for _, i := range fi.index {
		if fv.Kind() == reflect.Pointer {
			if fv.IsNil() {
				return reflect.Value{}, false
			}
			fv = fv.Elem()
		}
		fv = fv.Field(i)
	}

	return fv, true
}

func (fi FieldInfo) SetValue(s, v any) bool {
	subv := reflect.ValueOf(s)
	for _, i := range fi.index {
		if subv.Kind() == reflect.Pointer {
			if subv.IsNil() {
				if !subv.CanSet() {
					return false
				}
				subv.Set(reflect.New(subv.Type().Elem()))
			}
			subv = subv.Elem()
		}
		subv = subv.Field(i)
	}
	subv.Set(reflect.ValueOf(v))
	return true
}

func (fi FieldInfo) Struct() *StructInfo { return fi.si }

func (fi FieldInfo) Field() reflect.StructField        { return fi.field }
func (fi FieldInfo) Name() string                      { return fi.name }
func (fi FieldInfo) NameBytes() []byte                 { return fi.nameBytes }
func (fi FieldInfo) EqualFold() func(s, t []byte) bool { return fi.equalFold }
func (fi FieldInfo) Index() []int                      { return fi.index }
func (fi FieldInfo) Ty() reflect.Type                  { return fi.field.Type }

//

type StructInfo struct {
	name string
	ty   reflect.Type

	fields       []*FieldInfo
	fieldsByName map[string]*FieldInfo
}

func newStructInfo(ty reflect.Type) *StructInfo {
	if ty.Kind() != reflect.Struct {
		panic(errors.New("must be struct"))
	}

	si := &StructInfo{
		name: ty.Name(),
		ty:   ty,
	}

	wfs := walkFields(ty)
	fis := make([]*FieldInfo, len(wfs.flat))
	bn := make(map[string]*FieldInfo)
	for i, wf := range wfs.flat {
		fis[i] = &FieldInfo{
			si: si,
		}
		if wf.name != "" {
			bn[wf.name] = fis[i]
		}
	}

	si.fields = fis
	si.fieldsByName = bn

	return si
}

func (si StructInfo) Ty() reflect.Type                    { return si.ty }
func (si StructInfo) Fields() []*FieldInfo                { return si.fields }
func (si StructInfo) FieldsByName() map[string]*FieldInfo { return si.fieldsByName }

func (si StructInfo) Field(name string) *FieldInfo {
	return si.fieldsByName[name]
}

//

type StructInfoCache struct {
	structInfos map[reflect.Type]*StructInfo
}

func NewStructInfoCache() *StructInfoCache {
	return &StructInfoCache{
		structInfos: make(map[reflect.Type]*StructInfo),
	}
}

func (sic *StructInfoCache) Info(ty any) *StructInfo {
	rty := rfl.AsType(ty)
	if si, ok := sic.structInfos[rty]; ok {
		return si
	}

	si := newStructInfo(rty)
	sic.structInfos[rty] = si
	return si
}
