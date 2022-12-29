// Code generated from Minml.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Minml

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// A complete Visitor for a parse tree produced by MinmlParser.
type MinmlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MinmlParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by MinmlParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by MinmlParser#obj.
	VisitObj(ctx *ObjContext) interface{}

	// Visit a parse tree produced by MinmlParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by MinmlParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by MinmlParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by MinmlParser#stringValue.
	VisitStringValue(ctx *StringValueContext) interface{}

	// Visit a parse tree produced by MinmlParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by MinmlParser#true.
	VisitTrue(ctx *TrueContext) interface{}

	// Visit a parse tree produced by MinmlParser#false.
	VisitFalse(ctx *FalseContext) interface{}

	// Visit a parse tree produced by MinmlParser#null.
	VisitNull(ctx *NullContext) interface{}
}
