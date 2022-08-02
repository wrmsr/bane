package sql

import (
	"database/sql"
	"fmt"
	"strings"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	iou "github.com/wrmsr/bane/pkg/util/io"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type Row struct {
	m ctr.ShapeMap[string, any]
}

func (r Row) Values() []any   { return r.m.Values() }
func (r Row) Value(i int) any { return r.m.Value(i) }

func (r Row) Clone() Row {
	vs := r.m.Values()
	for i, v := range vs {
		if v, ok := v.(sql.RawBytes); ok {
			vs[i] = slices.Clone(v)
		}
	}
	return Row{m: ctr.NewShapeMapFromSlice[string, any](r.m.Shape(), vs)}
}

var _ ctr.Map[string, any] = Row{}

func (r Row) Len() int                                         { return r.m.Len() }
func (r Row) Contains(k string) bool                           { return r.m.Contains(k) }
func (r Row) Get(k string) any                                 { return r.m.Get(k) }
func (r Row) TryGet(k string) (any, bool)                      { return r.m.TryGet(k) }
func (r Row) Iterate() bt.Iterator[bt.Kv[string, any]]         { return r.m.Iterate() }
func (r Row) ForEach(fn func(kv bt.Kv[string, any]) bool) bool { return r.m.ForEach(fn) }

func (r Row) MarshalJSON() ([]byte, error) {
	return r.m.MarshalJSON()
}

func (r Row) Format(f fmt.State, c rune) {
	iou.WriteStringDiscard(f, "row")
	ctr.MapFormat[string, any](f, r.m)
}

func (r Row) string(tn string) string {
	var sb strings.Builder
	sb.WriteString(tn)
	ctr.MapString[string, any](&sb, r.m)
	return sb.String()
}
