package sql

import (
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
)

type RowReader struct {
	r sqb.Rows

	vvs []reflect.Value
	dvs []any
	vs  []any

	shape ctr.MapShape[string]
}

func NewRowReader(r sqb.Rows) (RowReader, error) {
	cs, err := r.Columns()
	if err != nil {
		return RowReader{}, err
	}

	vvs := make([]reflect.Value, len(cs))
	dvs := make([]any, len(cs))
	for i, c := range cs {
		sv := reflect.New(c.Type)
		vvs[i] = sv.Elem()
		dvs[i] = sv.Interface()
	}

	cns := make([]string, len(cs))
	for i, c := range cs {
		cns[i] = c.Name
	}

	return RowReader{
		r: r,

		vvs: vvs,
		dvs: dvs,
		vs:  make([]any, len(cs)),

		shape: ctr.NewMapShape(its.OfSlice(cns)),
	}, nil
}

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
	return Row{m: ctr.NewShapeMapFromSlice[string, any](rr.shape, rr.vs)}, nil
}
