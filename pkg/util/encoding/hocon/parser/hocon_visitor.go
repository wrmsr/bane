// Code generated from Hocon.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Hocon

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// A complete Visitor for a parse tree produced by HoconParser.
type HoconVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by HoconParser#hocon.
	VisitHocon(ctx *HoconContext) any

	// Visit a parse tree produced by HoconParser#prop.
	VisitProp(ctx *PropContext) any

	// Visit a parse tree produced by HoconParser#obj.
	VisitObj(ctx *ObjContext) any

	// Visit a parse tree produced by HoconParser#objectBegin.
	VisitObjectBegin(ctx *ObjectBeginContext) any

	// Visit a parse tree produced by HoconParser#objectEnd.
	VisitObjectEnd(ctx *ObjectEndContext) any

	// Visit a parse tree produced by HoconParser#objectData.
	VisitObjectData(ctx *ObjectDataContext) any

	// Visit a parse tree produced by HoconParser#arrayData.
	VisitArrayData(ctx *ArrayDataContext) any

	// Visit a parse tree produced by HoconParser#stringData.
	VisitStringData(ctx *StringDataContext) any

	// Visit a parse tree produced by HoconParser#referenceData.
	VisitReferenceData(ctx *ReferenceDataContext) any

	// Visit a parse tree produced by HoconParser#numberData.
	VisitNumberData(ctx *NumberDataContext) any

	// Visit a parse tree produced by HoconParser#key.
	VisitKey(ctx *KeyContext) any

	// Visit a parse tree produced by HoconParser#path.
	VisitPath(ctx *PathContext) any

	// Visit a parse tree produced by HoconParser#arrayBegin.
	VisitArrayBegin(ctx *ArrayBeginContext) any

	// Visit a parse tree produced by HoconParser#arrayEnd.
	VisitArrayEnd(ctx *ArrayEndContext) any

	// Visit a parse tree produced by HoconParser#array.
	VisitArray(ctx *ArrayContext) any

	// Visit a parse tree produced by HoconParser#arrayValue.
	VisitArrayValue(ctx *ArrayValueContext) any

	// Visit a parse tree produced by HoconParser#arrayString.
	VisitArrayString(ctx *ArrayStringContext) any

	// Visit a parse tree produced by HoconParser#arrayReference.
	VisitArrayReference(ctx *ArrayReferenceContext) any

	// Visit a parse tree produced by HoconParser#arrayNumber.
	VisitArrayNumber(ctx *ArrayNumberContext) any

	// Visit a parse tree produced by HoconParser#arrayObj.
	VisitArrayObj(ctx *ArrayObjContext) any

	// Visit a parse tree produced by HoconParser#arrayArray.
	VisitArrayArray(ctx *ArrayArrayContext) any

	// Visit a parse tree produced by HoconParser#v_string.
	VisitV_string(ctx *V_stringContext) any

	// Visit a parse tree produced by HoconParser#v_rawstring.
	VisitV_rawstring(ctx *V_rawstringContext) any

	// Visit a parse tree produced by HoconParser#v_reference.
	VisitV_reference(ctx *V_referenceContext) any

	// Visit a parse tree produced by HoconParser#rawstring.
	VisitRawstring(ctx *RawstringContext) any
}
