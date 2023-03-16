// Code generated from Dot.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Dot

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

type BaseDotVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDotVisitor) VisitGraph(ctx *GraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitStmtList(ctx *StmtListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitAttrStmt(ctx *AttrStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitAttrList(ctx *AttrListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitAList(ctx *AListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitEdgeStmt(ctx *EdgeStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitEdgeRHS(ctx *EdgeRHSContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitEdgeop(ctx *EdgeopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitNodeStmt(ctx *NodeStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitNodeId(ctx *NodeIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitPort(ctx *PortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitSubgraph(ctx *SubgraphContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDotVisitor) VisitIdent(ctx *IdentContext) interface{} {
	return v.VisitChildren(ctx)
}
