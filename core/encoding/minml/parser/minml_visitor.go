// Code generated from Minml.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Minml

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

// A complete Visitor for a parse tree produced by MinmlParser.
type MinmlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MinmlParser#root.
	VisitRoot(ctx *RootContext) any

	// Visit a parse tree produced by MinmlParser#value.
	VisitValue(ctx *ValueContext) any

	// Visit a parse tree produced by MinmlParser#obj.
	VisitObj(ctx *ObjContext) any

	// Visit a parse tree produced by MinmlParser#pair.
	VisitPair(ctx *PairContext) any

	// Visit a parse tree produced by MinmlParser#array.
	VisitArray(ctx *ArrayContext) any

	// Visit a parse tree produced by MinmlParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) any

	// Visit a parse tree produced by MinmlParser#stringValue.
	VisitStringValue(ctx *StringValueContext) any

	// Visit a parse tree produced by MinmlParser#number.
	VisitNumber(ctx *NumberContext) any

	// Visit a parse tree produced by MinmlParser#true.
	VisitTrue(ctx *TrueContext) any

	// Visit a parse tree produced by MinmlParser#false.
	VisitFalse(ctx *FalseContext) any

	// Visit a parse tree produced by MinmlParser#null.
	VisitNull(ctx *NullContext) any
}
