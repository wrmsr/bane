// Code generated from Dot.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dot

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by DotParser.
type DotVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DotParser#graph.
	VisitGraph(ctx *GraphContext) interface{}

	// Visit a parse tree produced by DotParser#stmtList.
	VisitStmtList(ctx *StmtListContext) interface{}

	// Visit a parse tree produced by DotParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by DotParser#attrStmt.
	VisitAttrStmt(ctx *AttrStmtContext) interface{}

	// Visit a parse tree produced by DotParser#attrList.
	VisitAttrList(ctx *AttrListContext) interface{}

	// Visit a parse tree produced by DotParser#aList.
	VisitAList(ctx *AListContext) interface{}

	// Visit a parse tree produced by DotParser#edgeStmt.
	VisitEdgeStmt(ctx *EdgeStmtContext) interface{}

	// Visit a parse tree produced by DotParser#edgeRHS.
	VisitEdgeRHS(ctx *EdgeRHSContext) interface{}

	// Visit a parse tree produced by DotParser#edgeop.
	VisitEdgeop(ctx *EdgeopContext) interface{}

	// Visit a parse tree produced by DotParser#nodeStmt.
	VisitNodeStmt(ctx *NodeStmtContext) interface{}

	// Visit a parse tree produced by DotParser#nodeId.
	VisitNodeId(ctx *NodeIdContext) interface{}

	// Visit a parse tree produced by DotParser#port.
	VisitPort(ctx *PortContext) interface{}

	// Visit a parse tree produced by DotParser#subgraph.
	VisitSubgraph(ctx *SubgraphContext) interface{}

	// Visit a parse tree produced by DotParser#ident.
	VisitIdent(ctx *IdentContext) interface{}
}
