package container

import (
	"sync"
	"sync/atomic"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	"github.com/wrmsr/bane/pkg/util/maps"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type CowMap[K comparable, V any] struct {
	mtx sync.Mutex
	r   atomic.Value
}

func NewCowMap[K comparable, V any](it bt.Iterable[bt.Kv[K, V]]) *CowMap[K, V] {
	m := make(map[K]V)
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			c := it.Next()
			m[c.K] = c.V
		}
	}
	cm := &CowMap[K, V]{}
	cm.r.Store(m)
	return cm
}

func (m *CowMap[K, V]) get() map[K]V {
	return m.r.Load().(map[K]V)
}

var _ SyncMutMap[int, any] = &CowMap[int, any]{}

func (m *CowMap[K, V]) isMut()  {}
func (m *CowMap[K, V]) isSync() {}

func (m *CowMap[K, V]) Len() int {
	return len(m.get())
}

func (m *CowMap[K, V]) Contains(k K) bool {
	_, ok := m.get()[k]
	return ok
}

func (m *CowMap[K, V]) Get(k K) V {
	return m.get()[k]
}

func (m *CowMap[K, V]) TryGet(k K) (V, bool) {
	v, ok := m.get()[k]
	return v, ok
}

func (m *CowMap[K, V]) Iterate() bt.Iterator[bt.Kv[K, V]] {
	return its.OfMap[K, V](m.get()).Iterate()
}

func (m *CowMap[K, V]) ForEach(fn func(bt.Kv[K, V]) bool) bool {
	for k, v := range m.get() {
		if !fn(bt.KvOf(k, v)) {
			return false
		}
	}
	return true
}

func (m *CowMap[K, V]) mut(fn func(map[K]V)) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	nm := maps.Clone(m.get())
	fn(nm)
	m.r.Store(nm)
}

func (m *CowMap[K, V]) Put(k K, v V) {
	m.mut(func(m map[K]V) {
		m[k] = v
	})
}

func (m *CowMap[K, V]) Delete(k K) {
	m.mut(func(m map[K]V) {
		delete(m, k)
	})
}

func (m *CowMap[K, V]) Default(k K, v V) bool {
	var r bool
	m.mut(func(m map[K]V) {
		if _, ok := m[k]; !ok {
			m[k] = v
			r = true
		}
	})
	return r
}
