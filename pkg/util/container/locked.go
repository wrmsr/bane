package container

import (
	"reflect"
	"sync"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type locked struct {
	mtx sync.Mutex
}

var _ Locked = &locked{}

func (l *locked) isLocked() {}

//

type LockedMap[K, V any] interface {
	Map[K, V]
	Locked
}

type lockedMapImpl[K, V any] struct {
	locked
	m Map[K, V]
}

func NewLockedMap[K, V any](m Map[K, V]) LockedMap[K, V] {
	return &lockedMapImpl[K, V]{m: m}
}

var _ LockedMap[int, string] = &lockedMapImpl[int, string]{}

func (m *lockedMapImpl[K, V]) ReflectTypeArgs() []reflect.Type {
	return []reflect.Type{rfl.TypeOf[K](), rfl.TypeOf[V]()}
}

func (m *lockedMapImpl[K, V]) Len() int {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.m.Len()
}

func (m *lockedMapImpl[K, V]) Contains(k K) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.m.Contains(k)
}

func (m *lockedMapImpl[K, V]) Get(k K) V {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.m.Get(k)
}

func (m *lockedMapImpl[K, V]) TryGet(k K) (V, bool) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.m.TryGet(k)
}

func (m *lockedMapImpl[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return its.OfSlice(its.Seq[bt.Kv[K, V]](m.m)).Iterate()
}

func (m *lockedMapImpl[K, V]) ForEach(fn func(v bt.Kv[K, V]) bool) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.m.ForEach(fn)
}

//

type LockedMutMap[K, V any] interface {
	MutMap[K, V]
	Locked
}

type lockedMutMapImpl[K, V any] struct {
	lockedMapImpl[K, V]
	m MutMap[K, V]
}

func NewLockedMutMap[K, V any](m MutMap[K, V]) LockedMutMap[K, V] {
	return &lockedMutMapImpl[K, V]{lockedMapImpl: lockedMapImpl[K, V]{m: m}, m: m}
}

var _ LockedMutMap[int, string] = &lockedMutMapImpl[int, string]{}

func (m *lockedMutMapImpl[K, V]) isMutable() {}

func (m *lockedMutMapImpl[K, V]) Put(k K, v V) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.m.Put(k, v)
}

func (m *lockedMutMapImpl[K, V]) Delete(k K) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.m.Delete(k)
}

func (m *lockedMutMapImpl[K, V]) Default(k K, v V) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	return m.m.Default(k, v)
}
