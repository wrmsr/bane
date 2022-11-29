package graphs

import (
	"fmt"
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type NodeError[T any] struct{ Node T }

func (e NodeError[T]) Error() string { return fmt.Sprintf("duplicate node: %v", e.Node) }

//

type DuplicateNodeError[T any] struct{ NodeError[T] }

func (e DuplicateNodeError[T]) Error() string { return fmt.Sprintf("duplicate node: %v", e.Node) }

//

type UnknownNodeError[T any] struct{ NodeError[T] }

func (e UnknownNodeError[T]) Error() string { return fmt.Sprintf("unknown node: %v", e.Node) }

//

type Tree[T any] struct {
	root T
	walk func(T) bt.Iterable[T]
	he   bt.HashEqImpl[T]

	nodes   ctr.List[T]
	nodeSet ctr.Set[T]

	parentsByNode   ctr.Map[bt.Optional[T], bt.Optional[T]]
	childrenByNode  ctr.Map[bt.Optional[T], ctr.List[T]]
	childSetsByNode ctr.Map[bt.Optional[T], ctr.Set[T]]

	depthsByNode bt.Optional[ctr.Map[T, int]]

	nodeSetsByType bt.Optional[map[reflect.Type]ctr.Set[T]]
}

func NewTree[T any](root T, walk func(T) bt.Iterable[T], he bt.HashEqImpl[T]) (*Tree[T], error) {
	t := &Tree[T]{
		root: root,
		walk: walk,
		he:   he,
	}

	nodes := ctr.NewMutSliceList[T](nil)
	nodeSet := ctr.NewMutHashEqSet[T](he, nil)

	parentsByNode := ctr.NewMutHashEqMap[bt.Optional[T], bt.Optional[T]](bt.OptionalHashEq(he), nil)
	childrenByNode := ctr.NewMutHashEqMap[bt.Optional[T], ctr.List[T]](bt.OptionalHashEq(he), nil)
	childSetsByNode := ctr.NewMutHashEqMap[bt.Optional[T], ctr.Set[T]](bt.OptionalHashEq(he), nil)

	childrenByNode.Put(bt.None[T](), ctr.NewSliceListOf(root))
	childSetsByNode.Put(bt.None[T](), ctr.NewHashEqSet(he, its.Of(root)))

	var rec func(T, bt.Optional[T]) error
	rec = func(cur T, parent bt.Optional[T]) error {
		if nodeSet.Contains(cur) {
			return DuplicateNodeError[T]{NodeError[T]{cur}}
		}

		nodes.Append(cur)
		nodeSet.Add(cur)
		if !parent.Present() {
			if !he.Eq(cur, root) {
				return NodeError[T]{cur}
			}
		} else if !nodeSet.Contains(parent.Value()) {
			return UnknownNodeError[T]{NodeError[T]{parent.Value()}}
		}

		parentsByNode.Put(bt.Just(cur), parent)

		children := ctr.NewSliceList(walk(cur))
		childrenByNode.Put(bt.Just(cur), children)
		childSetsByNode.Put(bt.Just(cur), ctr.NewHashEqSet[T](he, children))

		var err error
		children.ForEach(func(child T) bool {
			err = rec(child, bt.Just(cur))
			return err == nil
		})
		return err
	}
	if err := rec(root, bt.None[T]()); err != nil {
		return nil, err
	}

	t.nodes = nodes
	t.nodeSet = nodeSet

	t.parentsByNode = parentsByNode
	t.childrenByNode = childrenByNode
	t.childSetsByNode = childSetsByNode

	return t, nil
}

func (t Tree[T]) Root() T                      { return t.root }
func (t Tree[T]) Walk() func(T) bt.Iterable[T] { return t.walk }

func (t Tree[T]) Nodes() ctr.List[T]  { return t.nodes }
func (t Tree[T]) NodeSet() ctr.Set[T] { return t.nodeSet }

func (t Tree[T]) ParentsByNode() ctr.Map[bt.Optional[T], bt.Optional[T]] { return t.parentsByNode }
func (t Tree[T]) ChildrenByNode() ctr.Map[bt.Optional[T], ctr.List[T]]   { return t.childrenByNode }
func (t Tree[T]) ChildSetsByNode() ctr.Map[bt.Optional[T], ctr.Set[T]]   { return t.childSetsByNode }

func (t *Tree[T]) DepthsByNode() ctr.Map[T, int] {
	return bt.SetIfAbsent(&t.depthsByNode, func() ctr.Map[T, int] {
		m := ctr.NewMutHashEqMap[T, int](t.he, nil)
		var rec func(T, int)
		rec = func(n T, d int) {
			m.Put(n, d)
			t.childSetsByNode.Get(bt.Just(n)).ForEach(func(c T) bool {
				rec(c, d+1)
				return true
			})
		}
		rec(t.root, 0)
		return m
	})
}

func (t *Tree[T]) NodeSetsByType() map[reflect.Type]ctr.Set[T] {
	return bt.SetIfAbsent(&t.nodeSetsByType, func() map[reflect.Type]ctr.Set[T] {
		m := make(map[reflect.Type]ctr.MutSet[T])
		r := make(map[reflect.Type]ctr.Set[T])
		t.nodes.ForEach(func(n T) bool {
			ty := reflect.TypeOf(n)
			if s, ok := m[ty]; ok {
				s.Add(n)
			} else {
				s := ctr.NewMutHashEqSet(t.he, its.Of(n))
				m[ty] = s
				r[ty] = s
			}
			return true
		})
		return r
	})
}

/*
   def iter_ancestors(self, node: NodeT) -> NodeGenerator[NodeT]:
       cur = node
       while True:
           cur = self.parents_by_node.get(cur)
           if cur is None:
               break
           yield cur

   def get_lineage(self, node: NodeT) -> col.IndexedSeq[NodeT]:
       return self._idx_seq_fac(reversed([node, *self.iter_ancestors(node)]))

   def get_first_parent_of_type(self, node: NodeT, ty: ta.Type[T]) -> ta.Optional[NodeT]:
       for cur in self.iter_ancestors(node):
           if isinstance(cur, ty):
               return cur
       return None

   @properties.cached  # noqa
   @property
   def depths_by_node(self) -> ta.Mapping[NodeT, int]:
       def rec(n, d):
           dct[n] = d
           for c in self._children_by_node[n]:
               rec(c, d + 1)
       dct: ta.MutableMapping[NodeT, int] = self._dict_fac()
       rec(self._root, 0)
       return dct
*/
