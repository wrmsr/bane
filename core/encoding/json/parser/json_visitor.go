// Code generated from Json.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Json

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

// A complete Visitor for a parse tree produced by JsonParser.
type JsonVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JsonParser#json.
	VisitJson(ctx *JsonContext) any

	// Visit a parse tree produced by JsonParser#object.
	VisitObject(ctx *ObjectContext) any

	// Visit a parse tree produced by JsonParser#pair.
	VisitPair(ctx *PairContext) any

	// Visit a parse tree produced by JsonParser#array.
	VisitArray(ctx *ArrayContext) any

	// Visit a parse tree produced by JsonParser#value.
	VisitValue(ctx *ValueContext) any
}
