// Code generated from Dot.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Dot

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// DotListener is a complete listener for a parse tree produced by DotParser.
type DotListener interface {
	antlr.ParseTreeListener

	// EnterGraph is called when entering the graph production.
	EnterGraph(c *GraphContext)

	// EnterStmtList is called when entering the stmtList production.
	EnterStmtList(c *StmtListContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAttrStmt is called when entering the attrStmt production.
	EnterAttrStmt(c *AttrStmtContext)

	// EnterAttrList is called when entering the attrList production.
	EnterAttrList(c *AttrListContext)

	// EnterAList is called when entering the aList production.
	EnterAList(c *AListContext)

	// EnterEdgeStmt is called when entering the edgeStmt production.
	EnterEdgeStmt(c *EdgeStmtContext)

	// EnterEdgeRHS is called when entering the edgeRHS production.
	EnterEdgeRHS(c *EdgeRHSContext)

	// EnterEdgeop is called when entering the edgeop production.
	EnterEdgeop(c *EdgeopContext)

	// EnterNodeStmt is called when entering the nodeStmt production.
	EnterNodeStmt(c *NodeStmtContext)

	// EnterNodeId is called when entering the nodeId production.
	EnterNodeId(c *NodeIdContext)

	// EnterPort is called when entering the port production.
	EnterPort(c *PortContext)

	// EnterSubgraph is called when entering the subgraph production.
	EnterSubgraph(c *SubgraphContext)

	// EnterIdent is called when entering the ident production.
	EnterIdent(c *IdentContext)

	// ExitGraph is called when exiting the graph production.
	ExitGraph(c *GraphContext)

	// ExitStmtList is called when exiting the stmtList production.
	ExitStmtList(c *StmtListContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAttrStmt is called when exiting the attrStmt production.
	ExitAttrStmt(c *AttrStmtContext)

	// ExitAttrList is called when exiting the attrList production.
	ExitAttrList(c *AttrListContext)

	// ExitAList is called when exiting the aList production.
	ExitAList(c *AListContext)

	// ExitEdgeStmt is called when exiting the edgeStmt production.
	ExitEdgeStmt(c *EdgeStmtContext)

	// ExitEdgeRHS is called when exiting the edgeRHS production.
	ExitEdgeRHS(c *EdgeRHSContext)

	// ExitEdgeop is called when exiting the edgeop production.
	ExitEdgeop(c *EdgeopContext)

	// ExitNodeStmt is called when exiting the nodeStmt production.
	ExitNodeStmt(c *NodeStmtContext)

	// ExitNodeId is called when exiting the nodeId production.
	ExitNodeId(c *NodeIdContext)

	// ExitPort is called when exiting the port production.
	ExitPort(c *PortContext)

	// ExitSubgraph is called when exiting the subgraph production.
	ExitSubgraph(c *SubgraphContext)

	// ExitIdent is called when exiting the ident production.
	ExitIdent(c *IdentContext)
}
