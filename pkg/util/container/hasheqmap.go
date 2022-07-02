package container

import (
	"fmt"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type hashEqMapNode[K, V any] struct {
	bt.Kv[K, V]
	h uintptr

	next, prev *hashEqMapNode[K, V]
}

func (n *hashEqMapNode[K, V]) String() string {
	return fmt.Sprintf("%16p{%16p %16p %x %v %v}", n, n.next, n.prev, n.h, n.K, n.V)
}

type hashEqMapImpl[K, V any] struct {
	he bt.HashEqImpl[K]

	head, tail *hashEqMapNode[K, V]

	m map[uintptr]*hashEqMapNode[K, V]
	l int
}

func newHashEqMapImpl[K comparable, V any](he bt.HashEqImpl[K], it its.Iterable[bt.Kv[K, V]]) hashEqMapImpl[K, V] {
	m := hashEqMapImpl[K, V]{
		he: he,
		m:  make(map[uintptr]*hashEqMapNode[K, V]),
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

func NewHashEqMap[K comparable, V any](he bt.HashEqImpl[K], it its.Iterable[bt.Kv[K, V]]) Map[K, V] {
	return newHashEqMapImpl[K, V](he, it)
}

var _ Map[int, string] = hashEqMapImpl[int, string]{}

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

func (m hashEqMapImpl[K, V]) getNode(k K, h uintptr) *hashEqMapNode[K, V] {
	for cur := m.m[h]; cur != nil && cur.h == h; cur = cur.prev {
		if m.he.Eq(k, cur.K) {
			return cur
		}
	}
	return nil
}

func (m hashEqMapImpl[K, V]) TryGet(k K) (V, bool) {
	if n := m.getNode(k, m.he.Hash(k)); n != nil {
		return n.V, true
	}
	var z V
	return z, false
}

type hashEqMapIterator[K, V any] struct {
	n *hashEqMapNode[K, V]
}

var _ its.Iterator[bt.Kv[int, string]] = &hashEqMapIterator[int, string]{}

func (i *hashEqMapIterator[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return i
}

func (i *hashEqMapIterator[K, V]) HasNext() bool {
	return i.n.next != nil
}

func (i *hashEqMapIterator[K, V]) Next() bt.Kv[K, V] {
	kv := i.n.Kv
	i.n = i.n.next
	return kv
}

func (m hashEqMapImpl[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return &hashEqMapIterator[K, V]{m.head}
}

func (m hashEqMapImpl[K, V]) ForEach(fn func(v bt.Kv[K, V]) bool) bool {
	for cur := m.head; cur != nil; cur = cur.next {
		if !fn(cur.Kv) {
			return false
		}
	}
	return true
}

func (m *hashEqMapImpl[K, V]) put(k K, v V) {
	h := m.he.Hash(k)
	if n := m.getNode(k, h); n != nil {
		n.V = v
		return
	}

	n := &hashEqMapNode[K, V]{
		Kv: bt.KvOf(k, v),
		h:  h,
	}

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
	var prev *hashEqMapNode[K, V]
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
}

func (m *hashEqMapImpl[K, V]) default_(k K, v V) bool {
	//_, ok := m.m[k]
	//if !ok {
	//	return false
	//}
	//m.m[k] = len(m.s)
	//m.s = append(m.s, bt.KvOf(k, v))
	//return true
	panic("nyi")
}

//

type mutHashEqMapImpl[K, V any] struct {
	hashEqMapImpl[K, V]
}
