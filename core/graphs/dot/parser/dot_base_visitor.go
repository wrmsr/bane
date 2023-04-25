// Code generated from Dot.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Dot

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

type BaseDotVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDotVisitor) VisitGraph(ctx *GraphContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitStmtList(ctx *StmtListContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitStmt(ctx *StmtContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitAttrStmt(ctx *AttrStmtContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitAttrList(ctx *AttrListContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitAList(ctx *AListContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitEdgeStmt(ctx *EdgeStmtContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitEdgeRHS(ctx *EdgeRHSContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitEdgeop(ctx *EdgeopContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitNodeStmt(ctx *NodeStmtContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitNodeId(ctx *NodeIdContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitPort(ctx *PortContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitSubgraph(ctx *SubgraphContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitIdent(ctx *IdentContext) any {
	return v.VisitChildren(ctx)
}
