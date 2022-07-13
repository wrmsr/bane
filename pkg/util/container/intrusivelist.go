package container

//

type IntrusiveListNode[T any] struct {
	next, prev *T
}

type intrusiveListRef[T any] struct {
	e *T

	next, prev **T
}

//

type IntrusiveListOps[T any] struct {
	getNode func(*T) *IntrusiveListNode[T]
}

func NewIntrusiveListOps[T any](getNode func(*T) *IntrusiveListNode[T]) IntrusiveListOps[T] {
	return IntrusiveListOps[T]{getNode: getNode}
}

func (o IntrusiveListOps[T]) GetNode(e *T) *IntrusiveListNode[T] {
	return o.getNode(e)
}

func (o IntrusiveListOps[T]) getRef(e *T) intrusiveListRef[T] {
	n := o.getNode(e)
	return intrusiveListRef[T]{e: e, next: &n.next, prev: &n.prev}
}

//

type IntrusiveList[T any] struct {
	ops IntrusiveListOps[T]

	head, tail *T

	l int
}

func NewIntrusiveList[T any](ops IntrusiveListOps[T]) IntrusiveList[T] {
	return IntrusiveList[T]{ops: ops}
}

func (l *IntrusiveList[T]) Len() int {
	return l.l
}

func (l *IntrusiveList[T]) root() intrusiveListRef[T] {
	return intrusiveListRef[T]{next: &l.head, prev: &l.tail}
}

func (l *IntrusiveList[T]) next(r intrusiveListRef[T]) intrusiveListRef[T] {
	if *r.next == nil {
		return l.root()
	}
	return l.ops.getRef(*r.next)
}

func (l *IntrusiveList[T]) prev(r intrusiveListRef[T]) intrusiveListRef[T] {
	if *r.prev == nil {
		return l.root()
	}
	return l.ops.getRef(*r.prev)
}

func (l *IntrusiveList[T]) insert(e, at intrusiveListRef[T]) {
	*e.prev = at.e
	*e.next = *at.next
	*l.prev(e).next = e.e
	*l.next(e).prev = e.e
	l.l++
}

func (l *IntrusiveList[T]) remove(e intrusiveListRef[T]) {
	*l.prev(e).next = *e.next
	*l.next(e).prev = *e.prev
	*e.next = nil
	*e.prev = nil
	l.l--
}

func (l *IntrusiveList[T]) move(e, at intrusiveListRef[T]) {
	if e.e == at.e {
		return
	}
	*l.prev(e).next = *e.next
	*l.next(e).prev = *e.prev

	*e.prev = at.e
	*e.next = *at.next
	*l.prev(e).next = e.e
	*l.next(e).prev = e.e
}

func (l *IntrusiveList[T]) PushFront(v *T) {
	l.insert(l.ops.getRef(v), l.root())
}

func (l *IntrusiveList[T]) PushBack(v *T) {
	l.insert(l.ops.getRef(v), l.prev(l.root()))
}
