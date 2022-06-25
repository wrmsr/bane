package container

import (
	"errors"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type MapShape[K comparable] struct {
	ks []K
	m  map[K]int
}

func NewMapShape[K comparable](ks its.Iterable[K]) MapShape[K] {
	var s []K
	m := make(map[K]int)
	i := 0
	for it := ks.Iterate(); it.HasNext(); {
		k := it.Next()
		if _, ok := m[k]; ok {
			panic(KeyError[K]{k})
		}
		s = append(s, k)
		m[k] = i
		i++
	}
	return MapShape[K]{
		ks: s,
		m:  m,
	}
}

var _ Set[int] = MapShape[int]{}

func (s MapShape[K]) Len() int {
	return len(s.ks)
}

func (s MapShape[K]) Contains(k K) bool {
	_, ok := s.m[k]
	return ok
}

func (s MapShape[K]) Iterate() its.Iterator[K] {
	return its.OfSlice(s.ks).Iterate()
}

func (s MapShape[K]) ForEach(fn func(v K) bool) bool {
	for _, k := range s.ks {
		if !fn(k) {
			return false
		}
	}
	return true
}

//

type ShapeMap[K comparable, V any] struct {
	shape MapShape[K]
	vs    []V
}

func NewShapeMap[K comparable, V any](shape MapShape[K], vs its.Iterable[V]) ShapeMap[K, V] {
	s := make([]V, 0, shape.Len())
	if vs != nil {
		i := 0
		for it := vs.Iterate(); it.HasNext(); {
			s[i] = it.Next()
			i++
		}
		if i != shape.Len() {
			panic(errors.New("length mismatch"))
		}
	}
	return ShapeMap[K, V]{
		shape: shape,
		vs:    s,
	}
}

func (m ShapeMap[K, V]) Shape() MapShape[K] {
	return m.shape
}

var _ Map[int, string] = ShapeMap[int, string]{}

func (m ShapeMap[K, V]) Len() int {
	return len(m.vs)
}

func (m ShapeMap[K, V]) Contains(k K) bool {
	return m.shape.Contains(k)
}

func (m ShapeMap[K, V]) Get(k K) V {
	v, _ := m.TryGet(k)
	return v
}

func (m ShapeMap[K, V]) TryGet(k K) (V, bool) {
	if i, ok := m.shape.m[k]; ok {
		return m.vs[i], true
	}
	var z V
	return z, false
}

func (m ShapeMap[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return its.Map(its.Range(0, m.shape.Len(), 1), func(i int) bt.Kv[K, V] {
		return bt.KvOf(m.shape.ks[i], m.vs[i])
	}).Iterate()
}

func (m ShapeMap[K, V]) ForEach(fn func(kv bt.Kv[K, V]) bool) bool {
	for i, v := range m.vs {
		if !fn(bt.KvOf(m.shape.ks[i], v)) {
			return false
		}
	}
	return true
}

//

type MutShapeMap[K comparable, V any] struct {
	ShapeMap[K, V]
}

func NewMutShapeMap[K comparable, V any](shape MapShape[K], vs its.Iterable[V]) MutShapeMap[K, V] {
	return MutShapeMap[K, V]{NewShapeMap(shape, vs)}
}

var _ MutMap[int, string] = MutShapeMap[int, string]{}

func (m MutShapeMap[K, V]) isMutable() {}

func (m MutShapeMap[K, V]) Put(k K, v V) {
	if i, ok := m.shape.m[k]; ok {
		m.vs[i] = v
		return
	}
	panic(KeyError[K]{k})
}

func (m MutShapeMap[K, V]) Delete(k K) {
	var z V
	m.Put(k, z)
}

func (m MutShapeMap[K, V]) Default(k K, v V) bool {
	return false
}