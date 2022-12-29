// Code generated from Dot.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Dot

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// BaseDotListener is a complete listener for a parse tree produced by DotParser.
type BaseDotListener struct{}

var _ DotListener = &BaseDotListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDotListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDotListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDotListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDotListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterGraph is called when production graph is entered.
func (s *BaseDotListener) EnterGraph(ctx *GraphContext) {}

// ExitGraph is called when production graph is exited.
func (s *BaseDotListener) ExitGraph(ctx *GraphContext) {}

// EnterStmtList is called when production stmtList is entered.
func (s *BaseDotListener) EnterStmtList(ctx *StmtListContext) {}

// ExitStmtList is called when production stmtList is exited.
func (s *BaseDotListener) ExitStmtList(ctx *StmtListContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseDotListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseDotListener) ExitStmt(ctx *StmtContext) {}

// EnterAttrStmt is called when production attrStmt is entered.
func (s *BaseDotListener) EnterAttrStmt(ctx *AttrStmtContext) {}

// ExitAttrStmt is called when production attrStmt is exited.
func (s *BaseDotListener) ExitAttrStmt(ctx *AttrStmtContext) {}

// EnterAttrList is called when production attrList is entered.
func (s *BaseDotListener) EnterAttrList(ctx *AttrListContext) {}

// ExitAttrList is called when production attrList is exited.
func (s *BaseDotListener) ExitAttrList(ctx *AttrListContext) {}

// EnterAList is called when production aList is entered.
func (s *BaseDotListener) EnterAList(ctx *AListContext) {}

// ExitAList is called when production aList is exited.
func (s *BaseDotListener) ExitAList(ctx *AListContext) {}

// EnterEdgeStmt is called when production edgeStmt is entered.
func (s *BaseDotListener) EnterEdgeStmt(ctx *EdgeStmtContext) {}

// ExitEdgeStmt is called when production edgeStmt is exited.
func (s *BaseDotListener) ExitEdgeStmt(ctx *EdgeStmtContext) {}

// EnterEdgeRHS is called when production edgeRHS is entered.
func (s *BaseDotListener) EnterEdgeRHS(ctx *EdgeRHSContext) {}

// ExitEdgeRHS is called when production edgeRHS is exited.
func (s *BaseDotListener) ExitEdgeRHS(ctx *EdgeRHSContext) {}

// EnterEdgeop is called when production edgeop is entered.
func (s *BaseDotListener) EnterEdgeop(ctx *EdgeopContext) {}

// ExitEdgeop is called when production edgeop is exited.
func (s *BaseDotListener) ExitEdgeop(ctx *EdgeopContext) {}

// EnterNodeStmt is called when production nodeStmt is entered.
func (s *BaseDotListener) EnterNodeStmt(ctx *NodeStmtContext) {}

// ExitNodeStmt is called when production nodeStmt is exited.
func (s *BaseDotListener) ExitNodeStmt(ctx *NodeStmtContext) {}

// EnterNodeId is called when production nodeId is entered.
func (s *BaseDotListener) EnterNodeId(ctx *NodeIdContext) {}

// ExitNodeId is called when production nodeId is exited.
func (s *BaseDotListener) ExitNodeId(ctx *NodeIdContext) {}

// EnterPort is called when production port is entered.
func (s *BaseDotListener) EnterPort(ctx *PortContext) {}

// ExitPort is called when production port is exited.
func (s *BaseDotListener) ExitPort(ctx *PortContext) {}

// EnterSubgraph is called when production subgraph is entered.
func (s *BaseDotListener) EnterSubgraph(ctx *SubgraphContext) {}

// ExitSubgraph is called when production subgraph is exited.
func (s *BaseDotListener) ExitSubgraph(ctx *SubgraphContext) {}

// EnterIdent is called when production ident is entered.
func (s *BaseDotListener) EnterIdent(ctx *IdentContext) {}

// ExitIdent is called when production ident is exited.
func (s *BaseDotListener) ExitIdent(ctx *IdentContext) {}
