// Copyright (C) 2018 Ramesh Vyaghrapuri. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

// Package treap implements a persistent treap (tree/heap combination).
//
// https://en.wikipedia.org/wiki/Treap
//
// A treap is a binary search tree for storing ordered distinct values (duplicates not allowed). In addition, each node
// actually a random priority field which is stored in heap order (i.e. all children have lower priority than the
// parent)
//
// This provides the basis for efficient immutable ordered Set operations.  See the ordered map example for how this can
// be used as an ordered map
//
// Much of this is based on "Fast Set Operations Using Treaps" by Guy E Blelloch and Margaret Reid-Miller:
// https://www.cs.cmu.edu/~scandal/papers/treaps-spaa98.pdf
//
// Benchmark
//
// The most interesting benchmark is the performance of insert where a single random key is inserted into a 5k sized
// map.  As the example shows, the treap structure does well here as opposed to a regular persistent map (which involves
// full copying).  This benchmark does not take into account the fact that the regular maps are not sorted unlike
// treaps.
//
// The intersection benchmark compares the case where two 10k sets with 5k in common being interesected. The regular
// persistent array is about 30% faster but this is still respectable showing for treaps.
//
//    $ go test --bench=. -benchmem
//    goos: darwin
//    goarch: amd64
//    pkg: github.com/perdata/treap
//    BenchmarkInsert-4                   	 1000000	      2347 ns/op	    1719 B/op	      36 allocs/op
//    BenchmarkInsertRegularMap-4         	    2000	    890745 ns/op	  336311 B/op	       8 allocs/op
//    BenchmarkIntersection-4             	     500	   3125772 ns/op	 1719838 B/op	   35836 allocs/op
//    BenchmarkIntersectionRegularMap-4   	     500	   2436519 ns/op	  718142 B/op	     123 allocs/op
//    BenchmarkUnion-4                    	    1000	   1451047 ns/op	  939846 B/op	   19580 allocs/op
//    BenchmarkDiff-4                     	     500	   3280823 ns/op	 1742080 B/op	   36298 allocs/op
//    PASS
//
package containers

import (
	brnd "github.com/wrmsr/bane/pkg/utils/rand"
	bt "github.com/wrmsr/bane/pkg/utils/types"
)

//

// Comparer compares two values. The return value is zero if the values are equal, negative if the first is smaller and
// positive otherwise.
type Comparer[T any] interface {
	Compare(left, right T) int
}

// Node is the basic recursive treap data structure
type TreapNode[T any] struct {
	Value       T
	Priority    int
	Left, Right *TreapNode[T]
}

// ForEach does inorder traversal of the treap
func (n *TreapNode[T]) ForEach(fn func(v T)) {
	if n != nil {
		n.Left.ForEach(fn)
		fn(n.Value)
		n.Right.ForEach(fn)
	}
}

// Find finds the node in the treap with matching value
func (n *TreapNode[T]) Find(v T, c Comparer[T]) *TreapNode[T] {
	for {
		if n == nil {
			return nil
		}
		diff := c.Compare(n.Value, v)
		switch {
		case diff == 0:
			return n
		case diff < 0:
			n = n.Right
		case diff > 0:
			n = n.Left
		}
	}
}

// Union combines any two treaps. In case of duplicates, the overwrite field controls whether the union keeps the
// original value or whether it is updated based on value in the "other" arg
func (n *TreapNode[T]) Union(other *TreapNode[T], c Comparer[T], overwrite bool) *TreapNode[T] {
	if n == nil {
		return other
	}
	if other == nil {
		return n
	}

	if n.Priority < other.Priority {
		other, n, overwrite = n, other, !overwrite
	}

	left, dupe, right := other.Split(n.Value, c)
	value := n.Value
	if overwrite && dupe != nil {
		value = dupe.Value
	}
	left = n.Left.Union(left, c, overwrite)
	right = n.Right.Union(right, c, overwrite)
	return &TreapNode[T]{value, n.Priority, left, right}
}

// Split splits the treap into all nodes that compare less-than, equal and greater-than the provided value.  The
// resulting values are properly formed treaps or nil if they contain no values.
func (n *TreapNode[T]) Split(v T, c Comparer[T]) (left, mid, right *TreapNode[T]) {
	leftp, rightp := &left, &right
	for {
		if n == nil {
			*leftp = nil
			*rightp = nil
			return left, nil, right
		}

		root := &TreapNode[T]{n.Value, n.Priority, nil, nil}
		diff := c.Compare(n.Value, v)
		switch {
		case diff < 0:
			*leftp = root
			root.Left = n.Left
			leftp = &root.Right
			n = n.Right
		case diff > 0:
			*rightp = root
			root.Right = n.Right
			rightp = &root.Left
			n = n.Left
		default:
			*leftp = n.Left
			*rightp = n.Right
			return left, root, right
		}
	}
}

// Intersection returns a new treap with all the common values in the two treaps.
//
// see https://www.cs.cmu.edu/~scandal/papers/treaps-spaa98.pdf
// "Fast Set Operations Using Treaps"
//   by Guy E Blelloch and Margaret Reid-Miller.
//
// The algorithm is a very slight variation on that.
func (n *TreapNode[T]) Intersection(other *TreapNode[T], c Comparer[T]) *TreapNode[T] {
	if n == nil || other == nil {
		return nil
	}

	if n.Priority < other.Priority {
		n, other = other, n
	}

	left, found, right := other.Split(n.Value, c)
	left = n.Left.Intersection(left, c)
	right = n.Right.Intersection(right, c)

	if found == nil {
		// TODO: use a destructive join as both left/right are copies
		return left.join(right)
	}

	return &TreapNode[T]{n.Value, n.Priority, left, right}
}

// Delete removes a node if it exists.
func (n *TreapNode[T]) Delete(v T, c Comparer[T]) *TreapNode[T] {
	left, _, right := n.Split(v, c)
	return left.join(right)
}

// Diff finds all elements of current treap which aren't present in the other heap
func (n *TreapNode[T]) Diff(other *TreapNode[T], c Comparer[T]) *TreapNode[T] {
	if n == nil || other == nil {
		return n
	}

	// TODO -- use  count
	if n.Priority >= other.Priority {
		left, dupe, right := other.Split(n.Value, c)
		left, right = n.Left.Diff(left, c), n.Right.Diff(right, c)
		if dupe != nil {
			return left.join(right)
		}
		return &TreapNode[T]{n.Value, n.Priority, left, right}
	}

	left, _, right := n.Split(other.Value, c)
	left = left.Diff(other.Left, c)
	right = right.Diff(other.Right, c)
	return left.join(right)
}

// see https://www.cs.cmu.edu/~scandal/papers/treaps-spaa98.pdf
// "Fast Set Operations Using Treaps"
//   by Guy E Blelloch and Margaret Reid-Miller.
//
// The algorithm is a very slight variation on that provided there.
//
// Note that all nodes in n have priority <= that of "other" for this call to work correctly.  It traverses  the right
// spine of n and left-spine of other, merging things along the way
//
// The algorithm is not that  different from zipping up a spine
func (n *TreapNode[T]) join(other *TreapNode[T]) *TreapNode[T] {
	var result *TreapNode[T]
	resultp := &result
	for {
		if n == nil {
			*resultp = other
			return result
		}
		if other == nil {
			*resultp = n
			return result
		}

		if n.Priority <= other.Priority {
			root := &TreapNode[T]{n.Value, n.Priority, n.Left, nil}
			*resultp = root
			resultp = &root.Right
			n = n.Right
		} else {
			root := &TreapNode[T]{other.Value, other.Priority, nil, other.Right}
			*resultp = root
			resultp = &root.Left
			other = other.Left
		}
	}
}

//

type treapMapComparer[K comparable, V any] func(kv1, kv2 bt.Kv[K, V]) int

func (c treapMapComparer[K, V]) Compare(kv1, kv2 bt.Kv[K, V]) int {
	return c(kv1, kv2)
}

type TreapMap[K comparable, V any] struct {
	n *TreapNode[bt.Kv[K, V]]
	c Comparer[bt.Kv[K, V]]
}

func NewTreapMap[K comparable, V any](keyCompare Comparer[K]) TreapMap[K, V] {
	return TreapMap[K, V]{
		c: treapMapComparer[K, V](func(v1, v2 bt.Kv[K, V]) int {
			return keyCompare.Compare(v1.K, v2.K)
		}),
	}
}

func (m TreapMap[K, V]) Get(key K) (any, bool) {
	n := m.n.Find(bt.KvOf[K, V](key, bt.Zero[V]()), m.c)
	if n == nil {
		return nil, false
	}
	return n.Value.V, true
}

func (m TreapMap[K, V]) Set(key K, value V) TreapMap[K, V] {
	node := &TreapNode[bt.Kv[K, V]]{
		bt.KvOf(key, value),
		int(brnd.FastUint32()),
		nil,
		nil,
	}
	n := m.n.Union(node, m.c, true)
	return TreapMap[K, V]{n, m.c}
}

func (m TreapMap[K, V]) Delete(key K) TreapMap[K, V] {
	n := m.n.Delete(bt.KvOf[K, V](key, bt.Zero[V]()), m.c)
	return TreapMap[K, V]{n, m.c}
}

func (m TreapMap[K, V]) ForEach(fn func(key K, value V)) {
	m.n.ForEach(func(kv bt.Kv[K, V]) {
		fn(kv.K, kv.V)
	})
}

func (m TreapMap[K, V]) Count() int {
	result := 0
	m.n.ForEach(func(_ bt.Kv[K, V]) {
		result++
	})
	return result
}