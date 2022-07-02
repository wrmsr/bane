package graphs

import (
	"fmt"
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

//

type NodeError[T comparable] struct{ Node T }

func (e NodeError[T]) Error() string { return fmt.Sprintf("duplicate node: %v", e.Node) }

//

type DuplicateNodeError[T comparable] struct{ NodeError[T] }

func (e DuplicateNodeError[T]) Error() string { return fmt.Sprintf("duplicate node: %v", e.Node) }

//

type UnknownNodeError[T comparable] struct{ NodeError[T] }

func (e UnknownNodeError[T]) Error() string { return fmt.Sprintf("unknown node: %v", e.Node) }

//

type Tree[T comparable] struct {
	root T
	walk func(T) its.Iterable[T]

	nodes   ctr.List[T]
	nodeSet ctr.Set[T]

	parentsByNode   ctr.Map[opt.Optional[T], opt.Optional[T]]
	childrenByNode  ctr.Map[opt.Optional[T], ctr.List[T]]
	childSetsByNode ctr.Map[opt.Optional[T], ctr.Set[T]]

	depthsByNode opt.Optional[ctr.Map[T, int]]

	nodeSetsByType opt.Optional[map[reflect.Type]ctr.Set[T]]
}

func NewTree[T comparable](root T, walk func(T) its.Iterable[T]) (*Tree[T], error) {
	t := &Tree[T]{
		root: root,
		walk: walk,
	}

	nodes := ctr.NewMutList[T](nil)
	nodeSet := ctr.NewMutSet[T](nil)

	parentsByNode := ctr.NewMutMap[opt.Optional[T], opt.Optional[T]](nil)
	childrenByNode := ctr.NewMutMap[opt.Optional[T], ctr.List[T]](nil)
	childSetsByNode := ctr.NewMutMap[opt.Optional[T], ctr.Set[T]](nil)

	childrenByNode.Put(opt.None[T](), ctr.NewListOf(root))
	childSetsByNode.Put(opt.None[T](), ctr.NewSetOf(root))

	var rec func(T, opt.Optional[T]) error
	rec = func(cur T, parent opt.Optional[T]) error {
		if nodeSet.Contains(cur) {
			return DuplicateNodeError[T]{NodeError[T]{cur}}
		}

		nodes.Append(cur)
		nodeSet.Add(cur)
		if !parent.Present() {
			if cur != root {
				return NodeError[T]{cur}
			}
		} else if !nodeSet.Contains(parent.Value()) {
			return UnknownNodeError[T]{NodeError[T]{parent.Value()}}
		}

		parentsByNode.Put(opt.Just(cur), parent)

		children := ctr.NewList(walk(cur))
		childrenByNode.Put(opt.Just(cur), children)
		childSetsByNode.Put(opt.Just(cur), ctr.NewSet[T](children))

		var err error
		children.ForEach(func(child T) bool {
			err = rec(child, opt.Just(cur))
			return err == nil
		})
		return err
	}
	if err := rec(root, opt.None[T]()); err != nil {
		return nil, err
	}

	t.nodes = nodes
	t.nodeSet = nodeSet

	t.parentsByNode = parentsByNode
	t.childrenByNode = childrenByNode
	t.childSetsByNode = childSetsByNode

	return t, nil
}

func (t Tree[T]) Root() T                       { return t.root }
func (t Tree[T]) Walk() func(T) its.Iterable[T] { return t.walk }

func (t Tree[T]) Nodes() ctr.List[T]  { return t.nodes }
func (t Tree[T]) NodeSet() ctr.Set[T] { return t.nodeSet }

func (t Tree[T]) ParentsByNode() ctr.Map[opt.Optional[T], opt.Optional[T]] { return t.parentsByNode }
func (t Tree[T]) ChildrenByNode() ctr.Map[opt.Optional[T], ctr.List[T]]    { return t.childrenByNode }
func (t Tree[T]) ChildSetsByNode() ctr.Map[opt.Optional[T], ctr.Set[T]]    { return t.childSetsByNode }

func (t *Tree[T]) DepthsByNode() ctr.Map[T, int] {
	return opt.SetIfAbsent(&t.depthsByNode, func() ctr.Map[T, int] {
		m := make(map[T]int, t.nodes.Len())
		var rec func(T, int)
		rec = func(n T, d int) {
			m[n] = d
			t.childSetsByNode.Get(opt.Just(n)).ForEach(func(c T) bool {
				rec(c, d+1)
				return true
			})
		}
		rec(t.root, 0)
		return ctr.WrapMap(m)
	})
}

func (t *Tree[T]) NodeSetsByType() map[reflect.Type]ctr.Set[T] {
	return opt.SetIfAbsent(&t.nodeSetsByType, func() map[reflect.Type]ctr.Set[T] {
		m := make(map[reflect.Type]ctr.MutSet[T])
		r := make(map[reflect.Type]ctr.Set[T])
		t.nodes.ForEach(func(n T) bool {
			ty := reflect.TypeOf(n)
			if s, ok := m[ty]; ok {
				s.Add(n)
			} else {
				s := ctr.NewMutSetOf(n)
				m[ty] = s
				r[ty] = s
			}
			return true
		})
		return r
	})
}
