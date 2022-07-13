package container

//

type IntrusiveListNode[T any] struct {
	next, prev *T
}

type intrusiveListPair[T any] struct {
	v *T
	n *IntrusiveListNode[T]
}

//

type IntrusiveListOps[T any] struct {
	getNode func(*T) *IntrusiveListNode[T]
}

func NewIntrusiveListOps[T any](getNode func(*T) *IntrusiveListNode[T]) IntrusiveListOps[T] {
	return IntrusiveListOps[T]{getNode: getNode}
}

func (l IntrusiveListOps[T]) GetNode(e *T) *IntrusiveListNode[T] {
	return l.getNode(e)
}

func (l IntrusiveListOps[T]) pair(v *T) intrusiveListPair[T] {
	return intrusiveListPair[T]{v: v, n: l.getNode(v)}
}

//

type IntrusiveList[T any] struct {
	ops IntrusiveListOps[T]

	root IntrusiveListNode[T]

	l int
}

func NewIntrusiveList[T any](ops IntrusiveListOps[T]) IntrusiveList[T] {
	return IntrusiveList[T]{ops: ops}
}

func (l *IntrusiveList[T]) Len() int {
	return l.l
}

func (l *IntrusiveList[T]) insert(e, at intrusiveListPair[T]) {
	e.n.prev = at.v
	e.n.next = at.n.next
	if e.n.prev != nil {
		l.ops.getNode(e.n.prev).next = e.v
	} else {

	}
	if e.n.next != nil {
		l.ops.getNode(e.n.next).prev = e.v
	}
}

func (l *IntrusiveList[T]) remove(e intrusiveListPair[T]) {
	if e.n.prev != nil {
		l.ops.getNode(e.n.prev).next = e.n.next
	}
	if e.n.next != nil {
		l.ops.getNode(e.n.next).prev = e.n.prev
	}
	e.n.next = nil
	e.n.prev = nil
}

func (l *IntrusiveList[T]) move(e, at intrusiveListPair[T]) {
	if e.n == at.n {
		return
	}
	if e.n.prev != nil {
		l.ops.getNode(e.n.prev).next = e.n.next
	}
	if e.n.next != nil {
		l.ops.getNode(e.n.next).prev = e.n.prev
	}

	e.n.prev = at.v
	e.n.next = at.n.next
	if e.n.prev != nil {
		l.ops.getNode(e.n.prev).next = e.v
	}
	if e.n.next != nil {
		l.ops.getNode(e.n.next).prev = e.v
	}
}

func (l *IntrusiveList[T]) PushFront(v *T) {
	l.insert(l.ops.pair(v), intrusiveListPair[T]{n: &l.root})
}

//func (l *IntrusiveList[T]) PushBack(v *T) {
//	l.insert(v, l.tail, l.ops.pair(v))
//}
