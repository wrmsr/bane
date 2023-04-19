// Code generated from Dot.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Dot

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// A complete Visitor for a parse tree produced by DotParser.
type DotVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DotParser#graph.
	VisitGraph(ctx *GraphContext) any

	// Visit a parse tree produced by DotParser#stmtList.
	VisitStmtList(ctx *StmtListContext) any

	// Visit a parse tree produced by DotParser#stmt.
	VisitStmt(ctx *StmtContext) any

	// Visit a parse tree produced by DotParser#attrStmt.
	VisitAttrStmt(ctx *AttrStmtContext) any

	// Visit a parse tree produced by DotParser#attrList.
	VisitAttrList(ctx *AttrListContext) any

	// Visit a parse tree produced by DotParser#aList.
	VisitAList(ctx *AListContext) any

	// Visit a parse tree produced by DotParser#edgeStmt.
	VisitEdgeStmt(ctx *EdgeStmtContext) any

	// Visit a parse tree produced by DotParser#edgeRHS.
	VisitEdgeRHS(ctx *EdgeRHSContext) any

	// Visit a parse tree produced by DotParser#edgeop.
	VisitEdgeop(ctx *EdgeopContext) any

	// Visit a parse tree produced by DotParser#nodeStmt.
	VisitNodeStmt(ctx *NodeStmtContext) any

	// Visit a parse tree produced by DotParser#nodeId.
	VisitNodeId(ctx *NodeIdContext) any

	// Visit a parse tree produced by DotParser#port.
	VisitPort(ctx *PortContext) any

	// Visit a parse tree produced by DotParser#subgraph.
	VisitSubgraph(ctx *SubgraphContext) any

	// Visit a parse tree produced by DotParser#ident.
	VisitIdent(ctx *IdentContext) any
}
