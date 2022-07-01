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

type OrderedMap[K, V any] interface {
	Map[K, V]
	Ordered
}

type MutOrderedMap[K, V any] interface {
	OrderedMap[K, V]
	MutMap[K, V]
}

//

type orderedMapImpl[K comparable, V any] struct {
	s []bt.Kv[K, V]
	m map[K]int
}

func newOrderedMapImpl[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) orderedMapImpl[K, V] {
	m := orderedMapImpl[K, V]{
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

func NewOrderedMap[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) OrderedMap[K, V] {
	return newOrderedMapImpl[K, V](it)
}

func (m orderedMapImpl[K, V]) string(tn string) string {
	var sb strings.Builder
	sb.WriteString("ordMap")
	for i, kv := range m.s {
		if i > 0 {
			sb.WriteRune(' ')
		}
		iou.FprintfDiscard(&sb, "%v:%v", kv.K, kv.V)
	}
	return sb.String()
}

func (m orderedMapImpl[K, V]) String() string {
	return m.string("ordMap")
}

func (m orderedMapImpl[K, V]) format(tn string, f fmt.State, c rune) {
	iou.WriteStringDiscard(f, tn)
	if f.Flag('#') {
		iou.WriteStringDiscard(f, "{")
	} else {
		iou.WriteStringDiscard(f, "[")
	}
	for i, kv := range m.s {
		if i > 0 {
			if f.Flag('#') {
				iou.WriteStringDiscard(f, ", ")
			} else {
				iou.WriteStringDiscard(f, " ")
			}
		}
		if f.Flag('#') {
			iou.FprintfDiscard(f, "%#v:%#v", kv.K, kv.V)
		} else {
			iou.FprintfDiscard(f, "%v:%v", kv.K, kv.V)
		}
	}
	if f.Flag('#') {
		iou.WriteStringDiscard(f, "}")
	} else {
		iou.WriteStringDiscard(f, "]")
	}
}

func (m orderedMapImpl[K, V]) Format(f fmt.State, c rune) {
	m.format("ordMap", f, c)
}

var _ OrderedMap[int, any] = orderedMapImpl[int, any]{}

func (m orderedMapImpl[K, V]) isOrdered() {}

func (m orderedMapImpl[K, V]) Len() int {
	return len(m.s)
}

func (m orderedMapImpl[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m orderedMapImpl[K, V]) Get(k K) V {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V]()
	}
	return m.s[i].V
}

func (m orderedMapImpl[K, V]) TryGet(k K) (V, bool) {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V](), false
	}
	return m.s[i].V, true
}

func (m orderedMapImpl[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return its.OfSlice(m.s).Iterate()
}

func (m orderedMapImpl[K, V]) ForEach(fn func(bt.Kv[K, V]) bool) bool {
	for _, kv := range m.s {
		if !fn(kv) {
			return false
		}
	}
	return true
}

func (m *orderedMapImpl[K, V]) put(k K, v V) {
	i, ok := m.m[k]
	if ok {
		m.s[i].V = v
		return
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
}

func (m *orderedMapImpl[K, V]) delete(k K) {
	i, ok := m.m[k]
	if !ok {
		return
	}
	m.s = slices.DeleteAt(m.s, i)
}

func (m *orderedMapImpl[K, V]) default_(k K, v V) bool {
	_, ok := m.m[k]
	if !ok {
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}

//

type mutOrderedMapImpl[K comparable, V any] struct {
	orderedMapImpl[K, V]
}

func NewMutOrderedMap[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) MutOrderedMap[K, V] {
	return &mutOrderedMapImpl[K, V]{
		orderedMapImpl: newOrderedMapImpl[K, V](it),
	}
}

func (m mutOrderedMapImpl[K, V]) String() string {
	return m.string("mutOrdMap")
}

func (m mutOrderedMapImpl[K, V]) Format(f fmt.State, c rune) {
	m.format("mutOrdMap", f, c)
}

var _ MutMap[int, any] = &mutOrderedMapImpl[int, any]{}

func (m *mutOrderedMapImpl[K, V]) isMutable() {}

func (m *mutOrderedMapImpl[K, V]) Put(k K, v V) {
	m.put(k, v)
}

func (m *mutOrderedMapImpl[K, V]) Delete(k K) {
	m.delete(k)
}

func (m *mutOrderedMapImpl[K, V]) Default(k K, v V) bool {
	return m.default_(k, v)
}
