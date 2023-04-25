package chm

import (
	"runtime"
	"sync/atomic"
	"testing"

	bt "github.com/wrmsr/bane/core/types"
)

type concHashMapNode struct {
	h uintptr
	k any
	v any
}

type concHashMap struct {
	he  bt.HashEqImpl[any]
	tav atomic.Value // []atomic.Value -> *concHashMapNode
	szc int64
}

func (m *concHashMap) getTable() []atomic.Value {
	if tv := m.tav.Load(); tv != nil {
		return tv.([]atomic.Value)
	}
	return nil
}

func (m *concHashMap) init() {
	for {
		table := m.getTable()
		if len(table) > 0 {
			break
		}

		szc := atomic.LoadInt64(&m.szc)
		if szc < 0 {
			runtime.Gosched()
			continue
		}

		if !atomic.CompareAndSwapInt64(&m.szc, szc, -1) {
			continue
		}

		table = m.getTable()
		if len(table) < 1 {
			n := szc
			if n < 1 {
				n = 16
			}
			m.tav.Store(make([]atomic.Value, n))
			szc = n - (n >> 2)
		}

		atomic.StoreInt64(&m.szc, szc)
		break
	}
}

const chmHashMask = ^uintptr(0) & (^uintptr(0) >> 1)

func (m *concHashMap) Put(k, v any) {
	h := m.he.Hash(k) & chmHashMask

	for {
		table := m.getTable()
		if len(table) < 1 {
			m.init()
			continue
		}

		i := (len(table) - 1) & int(h)
		nav := &table[i]
		nv := nav.Load()
		if nv == nil {
			node := &concHashMapNode{h: h, k: k, v: v}
			if nav.CompareAndSwap(nil, node) {
				break
			}
			continue
		}
		node := nv.(*concHashMapNode)

		nh := node.h
		if nh == ^uintptr(0) {

		}
	}
}

func TestConcHashMap(t *testing.T) {
	m := &concHashMap{
		he: bt.HashEqOf[any](
			func(o any) uintptr { return uintptr(o.(int) % 10) },
			func(l, r any) bool { return l.(int) == r.(int) },
		),
	}

	m.init()
	m.Put(420, 100)
}
