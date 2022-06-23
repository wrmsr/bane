package container

import (
	"fmt"
	"strings"

	iou "github.com/wrmsr/bane/pkg/util/io"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type OrdMap[K comparable, V any] interface {
	Map[K, V]
	Ordered
}

type MutOrdMap[K comparable, V any] interface {
	OrdMap[K, V]
	MutMap[K, V]
}

//

type ordMapImpl[K comparable, V any] struct {
	s []bt.Kv[K, V]
	m map[K]int
}

func newOrdMapImpl[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) ordMapImpl[K, V] {
	m := ordMapImpl[K, V]{
		m: make(map[K]int),
	}
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			c := it.Next()
			m.put(c.K, c.V)
		}
	}
	return m
}

func NewOrdMap[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) OrdMap[K, V] {
	return newOrdMapImpl[K, V](it)
}

func (m ordMapImpl[K, V]) string(tn string) string {
	var sb strings.Builder
	sb.WriteString("ordMap")
	for i, kv := range m.s {
		if i > 0 {
			sb.WriteRune(' ')
		}
		iou.FprintfX(&sb, "%v:%v", kv.K, kv.V)
	}
	return sb.String()
}

func (m ordMapImpl[K, V]) String() string {
	return m.string("ordMap")
}

func (m ordMapImpl[K, V]) format(tn string, f fmt.State, c rune) {
	iou.WriteStringX(f, tn)
	if f.Flag('#') {
		iou.WriteStringX(f, "{")
	} else {
		iou.WriteStringX(f, "[")
	}
	for i, kv := range m.s {
		if i > 0 {
			if f.Flag('#') {
				iou.WriteStringX(f, ", ")
			} else {
				iou.WriteStringX(f, " ")
			}
		}
		if f.Flag('#') {
			iou.FprintfX(f, "%#v:%#v", kv.K, kv.V)
		} else {
			iou.FprintfX(f, "%v:%v", kv.K, kv.V)
		}
	}
	if f.Flag('#') {
		iou.WriteStringX(f, "}")
	} else {
		iou.WriteStringX(f, "]")
	}
}

func (m ordMapImpl[K, V]) Format(f fmt.State, c rune) {
	m.format("ordMap", f, c)
}

var _ OrdMap[int, any] = ordMapImpl[int, any]{}

func (m ordMapImpl[K, V]) isOrdered() {}

func (m ordMapImpl[K, V]) Len() int {
	return len(m.s)
}

func (m ordMapImpl[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m ordMapImpl[K, V]) Get(k K) V {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V]()
	}
	return m.s[i].V
}

func (m ordMapImpl[K, V]) TryGet(k K) (V, bool) {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V](), false
	}
	return m.s[i].V, true
}

func (m ordMapImpl[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return its.OfSlice(m.s).Iterate()
}

func (m ordMapImpl[K, V]) ForEach(fn func(bt.Kv[K, V]) bool) bool {
	for _, kv := range m.s {
		if !fn(kv) {
			return false
		}
	}
	return true
}

func (m *ordMapImpl[K, V]) put(k K, v V) bool {
	i, ok := m.m[k]
	if ok {
		m.s[i].V = v
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}

func (m *ordMapImpl[K, V]) delete(k K) bool {
	i, ok := m.m[k]
	if !ok {
		return false
	}
	m.s = slices.DeleteAt(m.s, i)
	return true
}

func (m *ordMapImpl[K, V]) default_(k K, v V) bool {
	_, ok := m.m[k]
	if !ok {
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}

//

type mutOrdMapImpl[K comparable, V any] struct {
	ordMapImpl[K, V]
}

func NewMutOrdMap[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) MutOrdMap[K, V] {
	return &mutOrdMapImpl[K, V]{
		ordMapImpl: newOrdMapImpl[K, V](it),
	}
}

func (m mutOrdMapImpl[K, V]) String() string {
	return m.string("mutOrdMap")
}

func (m mutOrdMapImpl[K, V]) Format(f fmt.State, c rune) {
	m.format("mutOrdMap", f, c)
}

var _ MutMap[int, any] = &mutOrdMapImpl[int, any]{}

func (m *mutOrdMapImpl[K, V]) isMutable() {}

func (m *mutOrdMapImpl[K, V]) Put(k K, v V) bool {
	return m.put(k, v)
}

func (m *mutOrdMapImpl[K, V]) Delete(k K) bool {
	return m.delete(k)
}

func (m *mutOrdMapImpl[K, V]) Default(k K, v V) bool {
	return m.default_(k, v)
}
