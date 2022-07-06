package enums

import (
	"reflect"

	"github.com/wrmsr/bane/pkg/util/maps"
)

type EnumInfo[T comparable] struct {
	ty reflect.Type

	m map[string]T
	r map[T]string
}

func NewEnumInfo[T comparable](m map[string]T) *EnumInfo[T] {
	var z T
	return &EnumInfo[T]{
		ty: reflect.TypeOf(z),

		m: maps.Clone(m),
		r: maps.Invert(m),
	}
}
