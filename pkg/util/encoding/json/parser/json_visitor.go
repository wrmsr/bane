// Code generated from Json.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Json

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by JsonParser.
type JsonVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JsonParser#json.
	VisitJson(ctx *JsonContext) interface{}

	// Visit a parse tree produced by JsonParser#object.
	VisitObject(ctx *ObjectContext) interface{}

	// Visit a parse tree produced by JsonParser#pair.
	VisitPair(ctx *PairContext) interface{}

	// Visit a parse tree produced by JsonParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by JsonParser#value.
	VisitValue(ctx *ValueContext) interface{}
}
