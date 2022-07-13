package container

//

type IntrusiveListNode[T any] struct {
	next, prev *T
}

//

type IntrusiveListOps[T any] struct {
	getNode func(*T) *IntrusiveListNode[T]
}

func NewIntrusiveListOps[T any](getNode func(*T) *IntrusiveListNode[T]) IntrusiveListOps[T] {
	return IntrusiveListOps[T]{getNode: getNode}
}

func (l *IntrusiveListOps[T]) GetNode(e *T) *IntrusiveListNode[T] {
	return l.getNode(e)
}

func (l *IntrusiveListOps[T]) InsertNode(e *T, en *IntrusiveListNode[T], at *T, atn *IntrusiveListNode[T]) {
	en.prev = at
	en.next = atn.next
	l.getNode(en.prev).next = e
	l.getNode(en.next).prev = e
}

func (l *IntrusiveListOps[T]) Insert(e, at *T) {
	l.InsertNode(e, l.getNode(e), at, l.getNode(at))
}

func (l *IntrusiveListOps[T]) RemoveNode(en *IntrusiveListNode[T]) {
	l.getNode(en.prev).next = en.next
	l.getNode(en.next).prev = en.prev
	en.next = nil
	en.prev = nil
}

func (l *IntrusiveListOps[T]) Remove(e *T) {
	l.RemoveNode(l.getNode(e))
}

func (l *IntrusiveListOps[T]) MoveNode(e *T, en *IntrusiveListNode[T], at *T, atn *IntrusiveListNode[T]) {
	if en == atn {
		return
	}
	l.getNode(en.prev).next = en.next
	l.getNode(en.next).prev = en.prev

	en.prev = at
	en.next = atn.next
	l.getNode(en.prev).next = e
	l.getNode(en.next).prev = e
}

func (l *IntrusiveListOps[T]) Move(e, at *T) {
	l.MoveNode(e, l.getNode(e), at, l.getNode(at))
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

func (l *IntrusiveList[T]) PushFront(v *T) {
	l.ops.InsertNode(v, l.ops.getNode(v), nil, &l.root)
}
