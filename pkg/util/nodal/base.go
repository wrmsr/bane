package nodal

type Node interface {
	Copy() Node

	DoChildren(func(Node) bool) bool
	EditChildren(func(Node) Node)
}
