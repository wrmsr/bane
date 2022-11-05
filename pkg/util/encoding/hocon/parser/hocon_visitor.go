// Code generated from Hocon.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hocon

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by HoconParser.
type HoconVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by HoconParser#hocon.
	VisitHocon(ctx *HoconContext) interface{}

	// Visit a parse tree produced by HoconParser#prop.
	VisitProp(ctx *PropContext) interface{}

	// Visit a parse tree produced by HoconParser#obj.
	VisitObj(ctx *ObjContext) interface{}

	// Visit a parse tree produced by HoconParser#objectBegin.
	VisitObjectBegin(ctx *ObjectBeginContext) interface{}

	// Visit a parse tree produced by HoconParser#objectEnd.
	VisitObjectEnd(ctx *ObjectEndContext) interface{}

	// Visit a parse tree produced by HoconParser#objectData.
	VisitObjectData(ctx *ObjectDataContext) interface{}

	// Visit a parse tree produced by HoconParser#arrayData.
	VisitArrayData(ctx *ArrayDataContext) interface{}

	// Visit a parse tree produced by HoconParser#stringData.
	VisitStringData(ctx *StringDataContext) interface{}

	// Visit a parse tree produced by HoconParser#referenceData.
	VisitReferenceData(ctx *ReferenceDataContext) interface{}

	// Visit a parse tree produced by HoconParser#numberData.
	VisitNumberData(ctx *NumberDataContext) interface{}

	// Visit a parse tree produced by HoconParser#key.
	VisitKey(ctx *KeyContext) interface{}

	// Visit a parse tree produced by HoconParser#path.
	VisitPath(ctx *PathContext) interface{}

	// Visit a parse tree produced by HoconParser#arrayBegin.
	VisitArrayBegin(ctx *ArrayBeginContext) interface{}

	// Visit a parse tree produced by HoconParser#arrayEnd.
	VisitArrayEnd(ctx *ArrayEndContext) interface{}

	// Visit a parse tree produced by HoconParser#array.
	VisitArray(ctx *ArrayContext) interface{}

	// Visit a parse tree produced by HoconParser#arrayValue.
	VisitArrayValue(ctx *ArrayValueContext) interface{}

	// Visit a parse tree produced by HoconParser#arrayString.
	VisitArrayString(ctx *ArrayStringContext) interface{}

	// Visit a parse tree produced by HoconParser#arrayReference.
	VisitArrayReference(ctx *ArrayReferenceContext) interface{}

	// Visit a parse tree produced by HoconParser#arrayNumber.
	VisitArrayNumber(ctx *ArrayNumberContext) interface{}

	// Visit a parse tree produced by HoconParser#arrayObj.
	VisitArrayObj(ctx *ArrayObjContext) interface{}

	// Visit a parse tree produced by HoconParser#arrayArray.
	VisitArrayArray(ctx *ArrayArrayContext) interface{}

	// Visit a parse tree produced by HoconParser#v_string.
	VisitV_string(ctx *V_stringContext) interface{}

	// Visit a parse tree produced by HoconParser#v_rawstring.
	VisitV_rawstring(ctx *V_rawstringContext) interface{}

	// Visit a parse tree produced by HoconParser#v_reference.
	VisitV_reference(ctx *V_referenceContext) interface{}

	// Visit a parse tree produced by HoconParser#rawstring.
	VisitRawstring(ctx *RawstringContext) interface{}
}
