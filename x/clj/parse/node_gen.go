// FIXME: gen lol
package parse

//

//func (n *VectorNode) copy() Node {
//	c := *n
//	c.Nodes = copyNodes(c.Nodes)
//	return &c
//}
//func (n *VectorNode) doChildren(do func(Node) bool) bool {
//	if !doNodes(n.Nodes, do) {
//		return false
//	}
//	return true
//}
//func (n *VectorNode) editChildren(edit func(Node) Node) {
//	editNodes(n.Nodes, edit)
//}

//

//func copyNodes(list []Node) []Node {
//	if list == nil {
//		return nil
//	}
//	c := make([]Node, len(list))
//	copy(c, list)
//	return c
//}
//func doNodes(list []Node, do func(Node) bool) bool {
//	for _, x := range list {
//		if x != nil && !do(x) {
//			return false
//		}
//	}
//	return true
//}
//func editNodes(list []Node, edit func(Node) Node) {
//	for i, x := range list {
//		if x != nil {
//			list[i] = edit(x).(Node)
//		}
//	}
//}
