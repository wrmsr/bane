// Code generated from Expr.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Expr

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#singleExpr.
	VisitSingleExpr(ctx *SingleExprContext) interface{}

	// Visit a parse tree produced by ExprParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by ExprParser#orTest.
	VisitOrTest(ctx *OrTestContext) interface{}

	// Visit a parse tree produced by ExprParser#andTest.
	VisitAndTest(ctx *AndTestContext) interface{}

	// Visit a parse tree produced by ExprParser#notTest.
	VisitNotTest(ctx *NotTestContext) interface{}

	// Visit a parse tree produced by ExprParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by ExprParser#compOp.
	VisitCompOp(ctx *CompOpContext) interface{}

	// Visit a parse tree produced by ExprParser#exprMain.
	VisitExprMain(ctx *ExprMainContext) interface{}

	// Visit a parse tree produced by ExprParser#exprCont.
	VisitExprCont(ctx *ExprContContext) interface{}

	// Visit a parse tree produced by ExprParser#xorExpr.
	VisitXorExpr(ctx *XorExprContext) interface{}

	// Visit a parse tree produced by ExprParser#xorExprCont.
	VisitXorExprCont(ctx *XorExprContContext) interface{}

	// Visit a parse tree produced by ExprParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by ExprParser#andExprCont.
	VisitAndExprCont(ctx *AndExprContContext) interface{}

	// Visit a parse tree produced by ExprParser#shiftExpr.
	VisitShiftExpr(ctx *ShiftExprContext) interface{}

	// Visit a parse tree produced by ExprParser#shiftExprCont.
	VisitShiftExprCont(ctx *ShiftExprContContext) interface{}

	// Visit a parse tree produced by ExprParser#arithExpr.
	VisitArithExpr(ctx *ArithExprContext) interface{}

	// Visit a parse tree produced by ExprParser#arithExprCont.
	VisitArithExprCont(ctx *ArithExprContContext) interface{}

	// Visit a parse tree produced by ExprParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by ExprParser#termCont.
	VisitTermCont(ctx *TermContContext) interface{}

	// Visit a parse tree produced by ExprParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by ExprParser#power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by ExprParser#atomExpr.
	VisitAtomExpr(ctx *AtomExprContext) interface{}

	// Visit a parse tree produced by ExprParser#parenAtom.
	VisitParenAtom(ctx *ParenAtomContext) interface{}

	// Visit a parse tree produced by ExprParser#bracketAtom.
	VisitBracketAtom(ctx *BracketAtomContext) interface{}

	// Visit a parse tree produced by ExprParser#dictOrSetAtom.
	VisitDictOrSetAtom(ctx *DictOrSetAtomContext) interface{}

	// Visit a parse tree produced by ExprParser#constAtom.
	VisitConstAtom(ctx *ConstAtomContext) interface{}

	// Visit a parse tree produced by ExprParser#const.
	VisitConst(ctx *ConstContext) interface{}

	// Visit a parse tree produced by ExprParser#testListComp.
	VisitTestListComp(ctx *TestListCompContext) interface{}

	// Visit a parse tree produced by ExprParser#exprOrStarExpr.
	VisitExprOrStarExpr(ctx *ExprOrStarExprContext) interface{}

	// Visit a parse tree produced by ExprParser#starExpr.
	VisitStarExpr(ctx *StarExprContext) interface{}

	// Visit a parse tree produced by ExprParser#trailer.
	VisitTrailer(ctx *TrailerContext) interface{}

	// Visit a parse tree produced by ExprParser#subscriptList.
	VisitSubscriptList(ctx *SubscriptListContext) interface{}

	// Visit a parse tree produced by ExprParser#exprSubscript.
	VisitExprSubscript(ctx *ExprSubscriptContext) interface{}

	// Visit a parse tree produced by ExprParser#sliceSubscript.
	VisitSliceSubscript(ctx *SliceSubscriptContext) interface{}

	// Visit a parse tree produced by ExprParser#sliceOp.
	VisitSliceOp(ctx *SliceOpContext) interface{}

	// Visit a parse tree produced by ExprParser#dictMaker.
	VisitDictMaker(ctx *DictMakerContext) interface{}

	// Visit a parse tree produced by ExprParser#setMaker.
	VisitSetMaker(ctx *SetMakerContext) interface{}

	// Visit a parse tree produced by ExprParser#kvDictItem.
	VisitKvDictItem(ctx *KvDictItemContext) interface{}

	// Visit a parse tree produced by ExprParser#starsDictItem.
	VisitStarsDictItem(ctx *StarsDictItemContext) interface{}
}
