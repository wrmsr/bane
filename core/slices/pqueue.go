package slices

import (
	"container/heap"
)

//

type PqItem[T any] struct {
	Value T

	priority float32
	index    int
}

func (i PqItem[T]) Priority() float32 { return i.priority }

//

type PqInitItem[T any] struct {
	Value    T
	Priority float32
}

//

type PriorityQueue[T any] struct {
	s []PqItem[T]
}

func (pq *PriorityQueue[T]) Init(items []PqInitItem[T]) {
	if len(pq.s) > 0 {
		panic("must only be called on empty instances")
	}
	pq.s = make([]PqItem[T], len(items))
	for i, item := range items {
		pq.s[i] = PqItem[T]{
			Value:    item.Value,
			priority: item.Priority,
			index:    i,
		}
	}
	heap.Init(pq.heap())
}

func (pq *PriorityQueue[T]) Len() int {
	return len(pq.s)
}

func (pq *PriorityQueue[T]) At(index int) *PqItem[T] {
	return &pq.s[index]
}

func (pq *PriorityQueue[T]) Push(value T, priority float32) {
	item := PqItem[T]{
		Value:    value,
		priority: priority,
		index:    len(pq.s),
	}
	heap.Push(pq.heap(), item)
}

func (pq *PriorityQueue[T]) Pop() PqItem[T] {
	return heap.Pop(pq.heap()).(PqItem[T])
}

func (pq *PriorityQueue[T]) UpdatePriority(item *PqItem[T], priority float32) {
	item.priority = priority
	heap.Fix(pq.heap(), item.index)
}

//

type pqHeap[T any] struct {
	pq *PriorityQueue[T]
}

func (pq *PriorityQueue[T]) heap() heap.Interface {
	return pqHeap[T]{pq}
}

var _ heap.Interface = pqHeap[any]{}

func (h pqHeap[T]) Len() int {
	return len(h.pq.s)
}

func (h pqHeap[T]) Less(i, j int) bool {
	return h.pq.s[i].priority > h.pq.s[j].priority
}

func (h pqHeap[T]) Swap(i, j int) {
	h.pq.s[i], h.pq.s[j] = h.pq.s[j], h.pq.s[i]
	h.pq.s[i].index = i
	h.pq.s[j].index = j
}

func (h pqHeap[T]) Push(x any) {
	item := x.(PqItem[T])
	item.index = len(h.pq.s)
	h.pq.s = append(h.pq.s, item)
}

func (h pqHeap[T]) Pop() any {
	old := h.pq.s
	n := len(old)
	item := old[n-1]
	var z T
	h.pq.s[n-1].Value = z
	item.index = -1
	h.pq.s = old[0 : n-1]
	return item
}

//

type PtrPriorityQueue[T any] struct {
	s []*PqItem[T]
}

func (pq *PtrPriorityQueue[T]) Init(items []PqInitItem[T]) {
	if len(pq.s) > 0 {
		panic("must only be called on empty instances")
	}
	pq.s = make([]*PqItem[T], len(items))
	for i, item := range items {
		pq.s[i] = &PqItem[T]{
			Value:    item.Value,
			priority: item.Priority,
			index:    i,
		}
	}
	heap.Init(pq.heap())
}

func (pq *PtrPriorityQueue[T]) Len() int {
	return len(pq.s)
}

func (pq *PtrPriorityQueue[T]) At(index int) *PqItem[T] {
	return pq.s[index]
}

func (pq *PtrPriorityQueue[T]) Push(value T, priority float32) {
	item := PqItem[T]{
		Value:    value,
		priority: priority,
		index:    len(pq.s),
	}
	heap.Push(pq.heap(), item)
}

func (pq *PtrPriorityQueue[T]) Pop() PqItem[T] {
	return heap.Pop(pq.heap()).(PqItem[T])
}

func (pq *PtrPriorityQueue[T]) UpdatePriority(item *PqItem[T], priority float32) {
	item.priority = priority
	heap.Fix(pq.heap(), item.index)
}

//

type ptrPqHeap[T any] struct {
	pq *PtrPriorityQueue[T]
}

func (pq *PtrPriorityQueue[T]) heap() heap.Interface {
	return ptrPqHeap[T]{pq}
}

var _ heap.Interface = ptrPqHeap[any]{}

func (h ptrPqHeap[T]) Len() int {
	return len(h.pq.s)
}

func (h ptrPqHeap[T]) Less(i, j int) bool {
	return h.pq.s[i].priority > h.pq.s[j].priority
}

func (h ptrPqHeap[T]) Swap(i, j int) {
	h.pq.s[i], h.pq.s[j] = h.pq.s[j], h.pq.s[i]
	h.pq.s[i].index = i
	h.pq.s[j].index = j
}

func (h ptrPqHeap[T]) Push(x any) {
	item := x.(*PqItem[T])
	item.index = len(h.pq.s)
	h.pq.s = append(h.pq.s, item)
}

func (h ptrPqHeap[T]) Pop() any {
	old := h.pq.s
	n := len(old)
	item := old[n-1]
	h.pq.s[n-1] = nil
	item.index = -1
	h.pq.s = old[0 : n-1]
	return item
}
