package sql

import (
	"reflect"

	bt "github.com/wrmsr/bane/pkg/util/types"
	sqb "github.com/wrmsr/bane/sql/base"
)

type RowReader struct {
	r sqb.Rows

	cs *sqb.Columns

	vvs []reflect.Value
	dvs []any
	vs  []any
}

func NewRowReader(r sqb.Rows) (RowReader, error) {
	cs, err := r.Columns()
	if err != nil {
		return RowReader{}, err
	}

	vvs := make([]reflect.Value, cs.Len())
	dvs := make([]any, cs.Len())
	cns := make([]string, cs.Len())
	i := 0
	cs.ForEach(func(kv bt.Kv[string, sqb.Column]) bool {
		sv := reflect.New(kv.V.Type)
		vvs[i] = sv.Elem()
		dvs[i] = sv.Interface()
		cns[i] = kv.V.Name
		i++
		return true
	})

	return RowReader{
		r: r,

		cs: cs,

		vvs: vvs,
		dvs: dvs,
		vs:  make([]any, cs.Len()),
	}, nil
}

func (rr RowReader) Columns() *sqb.Columns { return rr.cs }

func (rr RowReader) Read() (Row, error) {
	if err := rr.r.Scan(rr.dvs...); err != nil {
		for i := range rr.vs {
			rr.vs[i] = nil
		}
		return Row{}, err
	}

	for i, vv := range rr.vvs {
		rr.vs[i] = vv.Interface()
	}
	return Row{cs: rr.cs, vs: rr.vs}, nil
}
