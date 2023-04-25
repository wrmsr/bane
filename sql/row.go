package sql

import (
	"database/sql"
	"fmt"
	"strings"

	ctr "github.com/wrmsr/bane/core/container"
	iou "github.com/wrmsr/bane/core/io"
	its "github.com/wrmsr/bane/core/iterators"
	"github.com/wrmsr/bane/core/slices"
	bt "github.com/wrmsr/bane/core/types"
	sqb "github.com/wrmsr/bane/sql/base"
)

type Row struct {
	cs *sqb.Columns
	vs []any
}

func (r Row) Columns() *sqb.Columns { return r.cs }

func (r Row) Values() []any   { return slices.Clone(r.vs) }
func (r Row) Value(i int) any { return r.vs[i] }

func (r Row) Clone() Row {
	vs := slices.Clone(r.vs)
	for i, v := range vs {
		if v, ok := v.(sql.RawBytes); ok {
			vs[i] = slices.Clone(v)
		}
	}
	return Row{cs: r.cs, vs: vs}
}

var _ ctr.Map[string, any] = Row{}

func (r Row) Len() int {
	return len(r.vs)
}

func (r Row) Contains(k string) bool {
	return r.cs.Contains(k)
}

func (r Row) Get(k string) any {
	v, _ := r.TryGet(k)
	return v
}

func (r Row) TryGet(k string) (any, bool) {
	i, ok := r.cs.Index(k)
	if !ok {
		return nil, false
	}
	return r.vs[i], true
}

func (r Row) Iterate() bt.Iterator[bt.Kv[string, any]] {
	return its.Map(its.Range(0, len(r.vs), 1), func(i int) bt.Kv[string, any] {
		return bt.KvOf(r.cs.At(i).Name, r.vs[i])
	}).Iterate()
}

func (r Row) ForEach(fn func(kv bt.Kv[string, any]) bool) bool {
	for i, v := range r.vs {
		if !fn(bt.KvOf(r.cs.At(i).Name, v)) {
			return false
		}
	}
	return true
}

func (r Row) MarshalJSON() ([]byte, error) {
	return ctr.MapMarshalJson[string, any](r)
}

func (r Row) Format(f fmt.State, c rune) {
	iou.WriteStringDiscard(f, "row")
	ctr.MapFormat[string, any](f, r)
}

func (r Row) string(tn string) string {
	var sb strings.Builder
	sb.WriteString(tn)
	ctr.MapString[string, any](&sb, r)
	return sb.String()
}
