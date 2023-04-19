// Code generated from Expr.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Expr

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#singleExpr.
	VisitSingleExpr(ctx *SingleExprContext) any

	// Visit a parse tree produced by ExprParser#expr.
	VisitExpr(ctx *ExprContext) any

	// Visit a parse tree produced by ExprParser#orTest.
	VisitOrTest(ctx *OrTestContext) any

	// Visit a parse tree produced by ExprParser#andTest.
	VisitAndTest(ctx *AndTestContext) any

	// Visit a parse tree produced by ExprParser#notTest.
	VisitNotTest(ctx *NotTestContext) any

	// Visit a parse tree produced by ExprParser#comparison.
	VisitComparison(ctx *ComparisonContext) any

	// Visit a parse tree produced by ExprParser#compOp.
	VisitCompOp(ctx *CompOpContext) any

	// Visit a parse tree produced by ExprParser#exprMain.
	VisitExprMain(ctx *ExprMainContext) any

	// Visit a parse tree produced by ExprParser#exprCont.
	VisitExprCont(ctx *ExprContContext) any

	// Visit a parse tree produced by ExprParser#xorExpr.
	VisitXorExpr(ctx *XorExprContext) any

	// Visit a parse tree produced by ExprParser#xorExprCont.
	VisitXorExprCont(ctx *XorExprContContext) any

	// Visit a parse tree produced by ExprParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) any

	// Visit a parse tree produced by ExprParser#andExprCont.
	VisitAndExprCont(ctx *AndExprContContext) any

	// Visit a parse tree produced by ExprParser#shiftExpr.
	VisitShiftExpr(ctx *ShiftExprContext) any

	// Visit a parse tree produced by ExprParser#shiftExprCont.
	VisitShiftExprCont(ctx *ShiftExprContContext) any

	// Visit a parse tree produced by ExprParser#arithExpr.
	VisitArithExpr(ctx *ArithExprContext) any

	// Visit a parse tree produced by ExprParser#arithExprCont.
	VisitArithExprCont(ctx *ArithExprContContext) any

	// Visit a parse tree produced by ExprParser#term.
	VisitTerm(ctx *TermContext) any

	// Visit a parse tree produced by ExprParser#termCont.
	VisitTermCont(ctx *TermContContext) any

	// Visit a parse tree produced by ExprParser#factor.
	VisitFactor(ctx *FactorContext) any

	// Visit a parse tree produced by ExprParser#power.
	VisitPower(ctx *PowerContext) any

	// Visit a parse tree produced by ExprParser#atomExpr.
	VisitAtomExpr(ctx *AtomExprContext) any

	// Visit a parse tree produced by ExprParser#parenAtom.
	VisitParenAtom(ctx *ParenAtomContext) any

	// Visit a parse tree produced by ExprParser#bracketAtom.
	VisitBracketAtom(ctx *BracketAtomContext) any

	// Visit a parse tree produced by ExprParser#dictOrSetAtom.
	VisitDictOrSetAtom(ctx *DictOrSetAtomContext) any

	// Visit a parse tree produced by ExprParser#constAtom.
	VisitConstAtom(ctx *ConstAtomContext) any

	// Visit a parse tree produced by ExprParser#constVal.
	VisitConstVal(ctx *ConstValContext) any

	// Visit a parse tree produced by ExprParser#testListComp.
	VisitTestListComp(ctx *TestListCompContext) any

	// Visit a parse tree produced by ExprParser#exprOrStarExpr.
	VisitExprOrStarExpr(ctx *ExprOrStarExprContext) any

	// Visit a parse tree produced by ExprParser#starExpr.
	VisitStarExpr(ctx *StarExprContext) any

	// Visit a parse tree produced by ExprParser#trailer.
	VisitTrailer(ctx *TrailerContext) any

	// Visit a parse tree produced by ExprParser#subscriptList.
	VisitSubscriptList(ctx *SubscriptListContext) any

	// Visit a parse tree produced by ExprParser#exprSubscript.
	VisitExprSubscript(ctx *ExprSubscriptContext) any

	// Visit a parse tree produced by ExprParser#sliceSubscript.
	VisitSliceSubscript(ctx *SliceSubscriptContext) any

	// Visit a parse tree produced by ExprParser#sliceOp.
	VisitSliceOp(ctx *SliceOpContext) any

	// Visit a parse tree produced by ExprParser#dictMaker.
	VisitDictMaker(ctx *DictMakerContext) any

	// Visit a parse tree produced by ExprParser#setMaker.
	VisitSetMaker(ctx *SetMakerContext) any

	// Visit a parse tree produced by ExprParser#kvDictItem.
	VisitKvDictItem(ctx *KvDictItemContext) any

	// Visit a parse tree produced by ExprParser#starsDictItem.
	VisitStarsDictItem(ctx *StarsDictItemContext) any
}
