package container

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type mapSet[K comparable, V any] struct {
	m Map[K, V]
}

func MapSet[K comparable, V any](m Map[K, V]) Set[K] {
	return mapSet[K, V]{m: m}
}

var _ Set[int] = mapSet[int, any]{}

func (s mapSet[K, V]) Len() int {
	return s.m.Len()
}

func (s mapSet[K, V]) Contains(k K) bool {
	return s.m.Contains(k)
}

func (s mapSet[K, V]) Iterate() its.Iterator[K] {
	return its.Keys[K, V](s.m).Iterate()
}

func (s mapSet[K, V]) ForEach(fn func(k K) bool) bool {
	return s.m.ForEach(func(kv bt.Kv[K, V]) bool {
		return fn(kv.K)
	})
}

//

type mutMapSet[K comparable, V any] struct {
	mapSet[K, V]
	m MutMap[K, V]
}

func MutMapSet[K comparable, V any](m MutMap[K, V]) MutSet[K] {
	return mutMapSet[K, V]{mapSet: mapSet[K, V]{m: m}, m: m}
}

var _ MutSet[int] = mutMapSet[int, any]{}

func (s mutMapSet[K, V]) isMutable() {}

func (s mutMapSet[K, V]) Add(k K) {
	var z V
	s.m.Put(k, z)
}

func (s mutMapSet[K, V]) TryAdd(k K) bool {
	if s.m.Contains(k) {
		return false
	}
	var z V
	s.m.Put(k, z)
	return true
}

func (s mutMapSet[K, V]) Remove(k K) {
	s.m.Delete(k)
}

func (s mutMapSet[K, V]) TryRemove(k K) bool {
	if s.m.Contains(k) {
		return false
	}
	s.m.Delete(k)
	return true
}

//

type setMap[K comparable, V any] struct {
	s Set[K]
}

func SetMap[K comparable, V any](s Set[K]) Map[K, V] {
	return setMap[K, V]{s: s}
}

var _ Map[int, string] = setMap[int, string]{}

func (m setMap[K, V]) Len() int {
	return m.s.Len()
}

func (m setMap[K, V]) Contains(k K) bool {
	return m.s.Contains(k)
}

func (m setMap[K, V]) Get(k K) V {
	var z V
	return z
}

func (m setMap[K, V]) TryGet(k K) (V, bool) {
	var z V
	return z, m.s.Contains(k)
}

func (m setMap[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return its.Map[K, bt.Kv[K, V]](m.s, func(k K) bt.Kv[K, V] {
		var z V
		return bt.KvOf(k, z)
	}).Iterate()
}

func (m setMap[K, V]) ForEach(fn func(bt.Kv[K, V]) bool) bool {
	return m.s.ForEach(func(k K) bool {
		var z V
		return fn(bt.KvOf(k, z))
	})
}

//
