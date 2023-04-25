package sql

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	sqb "github.com/wrmsr/bane/sql/base"
)

type RowScanner func(dst ...any) error

func ScanMap(cs []sqb.Column, scan RowScanner) (map[string]any, error) {
	vvs := make([]reflect.Value, len(cs))
	dvs := make([]any, len(cs))
	for i, c := range cs {
		sv := reflect.New(c.Type)
		vvs[i] = sv.Elem()
		dvs[i] = sv.Interface()
	}

	if err := scan(dvs...); err != nil {
		return nil, err
	}

	m := make(map[string]any, len(cs))
	for i, c := range cs {
		m[c.Name] = vvs[i].Interface()
	}

	return m, nil
}

func MakeAnyColumns(names ...string) []sqb.Column {
	cs := make([]sqb.Column, len(names))
	for i, n := range names {
		cs[i] = sqb.Column{Name: n, Type: rfl.Any()}
	}
	return cs
}
