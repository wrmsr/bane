package container

import (
	"fmt"
	"reflect"

	"github.com/wrmsr/bane/pkg/util/check"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	syncu "github.com/wrmsr/bane/pkg/util/sync"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type hashEqMapNode struct {
	k, v any

	h uintptr

	next, prev *hashEqMapNode
}

func (n *hashEqMapNode) String() string {
	return fmt.Sprintf("%16p{%16p %16p %16x} %v", n, n.next, n.prev, n.h, n.k)
}

var hashEqMapNodePool = syncu.NewDrainPool[*hashEqMapNode](func() *hashEqMapNode {
	return &hashEqMapNode{}
})

//

type hashEqMapImpl[K, V any] struct {
	he bt.HashEqImpl[K]

	head, tail *hashEqMapNode

	m map[uintptr]*hashEqMapNode
	l int
}

func newHashEqMapImpl[K, V any](he bt.HashEqImpl[K], it its.Iterable[bt.Kv[K, V]]) hashEqMapImpl[K, V] {
	m := hashEqMapImpl[K, V]{
		he: he,
		m:  make(map[uintptr]*hashEqMapNode),
	}
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			c := it.Next()
			m.put(c.K, c.V)
			m.print()
		}
	}
	return m
}

func (m *hashEqMapImpl[K, V]) initUnmarshal(a ...any) {
	m.he = check.Single(a).(bt.HashEqImpl[K])
}

func NewHashEqMap[K, V any](he bt.HashEqImpl[K], it its.Iterable[bt.Kv[K, V]]) Map[K, V] {
	return newHashEqMapImpl[K, V](he, it)
}

var _ Map[int, string] = hashEqMapImpl[int, string]{}

func (m hashEqMapImpl[K, V]) ReflectTypeArgs() []reflect.Type {
	return []reflect.Type{rfl.TypeOf[K](), rfl.TypeOf[V]()}
}

func (m hashEqMapImpl[K, V]) Len() int {
	return m.l
}

func (m hashEqMapImpl[K, V]) Contains(k K) bool {
	_, ok := m.TryGet(k)
	return ok
}

func (m hashEqMapImpl[K, V]) Get(k K) V {
	v, _ := m.TryGet(k)
	return v
}

func (m hashEqMapImpl[K, V]) getNode(k K, h uintptr) *hashEqMapNode {
	for cur := m.m[h]; cur != nil && cur.h == h; cur = cur.prev {
		if m.he.Eq(k, cur.k.(K)) {
			return cur
		}
	}
	return nil
}

func (m hashEqMapImpl[K, V]) TryGet(k K) (V, bool) {
	if n := m.getNode(k, m.he.Hash(k)); n != nil {
		return n.v.(V), true
	}
	var z V
	return z, false
}

type hashEqMapIterator[K, V any] struct {
	n *hashEqMapNode
}

var _ its.Iterator[bt.Kv[int, string]] = &hashEqMapIterator[int, string]{}

func (i *hashEqMapIterator[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return i
}

func (i *hashEqMapIterator[K, V]) HasNext() bool {
	return i.n.next != nil
}

func (i *hashEqMapIterator[K, V]) Next() bt.Kv[K, V] {
	kv := bt.KvOf(i.n.k.(K), i.n.v.(V))
	i.n = i.n.next
	return kv
}

func (m hashEqMapImpl[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return &hashEqMapIterator[K, V]{m.head}
}

func (m hashEqMapImpl[K, V]) ForEach(fn func(v bt.Kv[K, V]) bool) bool {
	for cur := m.head; cur != nil; cur = cur.next {
		if !fn(bt.KvOf(cur.k.(K), cur.v.(V))) {
			return false
		}
	}
	return true
}

func (m *hashEqMapImpl[K, V]) put(k K, v V) {
	h := m.he.Hash(k)
	if n := m.getNode(k, h); n != nil {
		n.v = v
		return
	}

	n := hashEqMapNodePool.Get()
	n.k = k
	n.v = v
	n.h = h

	p := m.m[h]

	if p != nil {
		n.next = p
		n.prev = p.prev

		if p.prev != nil {
			p.prev.next = n
		}
		p.prev = n

		if m.tail == nil {
			m.tail = n
		}

	} else {
		n.prev = m.tail

		if m.tail != nil {
			m.tail.next = n
		}
		m.tail = n

		m.m[h] = n
	}

	if m.head == p {
		m.head = n
	}

	m.l++
}

func (m *hashEqMapImpl[K, V]) verify() {
	i := 0
	var prev *hashEqMapNode
	for cur := m.head; cur != nil; prev, cur = cur, cur.next {
		if cur.prev != prev {
			panic(cur)
		}
		i++
	}
	if prev != m.tail {
		panic(prev)
	}
	if i != m.l {
		panic(m.l)
	}
}

func (m *hashEqMapImpl[K, V]) print() {
	for cur := m.head; cur != nil; cur = cur.next {
		fmt.Println(cur)
	}
	fmt.Println()
}

func (m *hashEqMapImpl[K, V]) delete(k K) {
	h := m.he.Hash(k)
	n := m.getNode(k, h)
	if n == nil {
		return
	}

	if n.next != nil {
		n.next.prev = n.prev
	}
	if n.prev != nil {
		n.prev.next = n.next
	}

	if m.head == n {
		m.head = n.next
	}
	if m.tail == n {
		m.tail = n.prev
	}

	if mn := m.m[h]; mn == n {
		if n.prev != nil {
			m.m[h] = n.prev
		} else {
			delete(m.m, h)
		}
	}

	m.l--

	*n = hashEqMapNode{}
	hashEqMapNodePool.Put(n)
}

func (m *hashEqMapImpl[K, V]) default_(k K, v V) bool {
	if _, ok := m.TryGet(k); ok {
		return false
	}
	m.put(k, v)
	return true
}

func (m *hashEqMapImpl[K, V]) clear() {
	for cur := m.head; cur != nil; cur = cur.next {
		*cur = hashEqMapNode{}
		hashEqMapNodePool.Put(cur)
	}
	m.head = nil
	m.tail = nil
	m.m = nil
	m.l = 0
}

//

type mutHashEqMapImpl[K, V any] struct {
	hashEqMapImpl[K, V]
}

func NewMutHashEqMap[K, V any](he bt.HashEqImpl[K], it its.Iterable[bt.Kv[K, V]]) MutMap[K, V] {
	return &mutHashEqMapImpl[K, V]{newHashEqMapImpl[K, V](he, it)}
}

func (m *mutHashEqMapImpl[K, V]) initUnmarshal(a ...any) {
	m.hashEqMapImpl.initUnmarshal(a...)
}

var _ MutMap[int, string] = &mutHashEqMapImpl[int, string]{}

func (m *mutHashEqMapImpl[K, V]) isMutable() {}

func (m *mutHashEqMapImpl[K, V]) Put(k K, v V) {
	m.put(k, v)
}

func (m *mutHashEqMapImpl[K, V]) Delete(k K) {
	m.delete(k)
}

func (m *mutHashEqMapImpl[K, V]) Default(k K, v V) bool {
	return m.default_(k, v)
}

//

func NewHashEqSet[K any](he bt.HashEqImpl[K], it its.Iterable[K]) Set[K] {
	var kvs its.Iterable[bt.Kv[K, struct{}]]
	if it != nil {
		kvs = its.StubKvs(it)
	}
	return MapSet(NewHashEqMap[K, struct{}](he, kvs))
}

func NewMutHashEqSet[K any](he bt.HashEqImpl[K], it its.Iterable[K]) MutSet[K] {
	var kvs its.Iterable[bt.Kv[K, struct{}]]
	if it != nil {
		kvs = its.StubKvs(it)
	}
	return MutMapSet(NewMutHashEqMap[K, struct{}](he, kvs))
}
