// Code generated from Expr.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseExprVisitor) VisitSingleExpr(ctx *SingleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitOrTest(ctx *OrTestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAndTest(ctx *AndTestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitNotTest(ctx *NotTestContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitComparison(ctx *ComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitCompOp(ctx *CompOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitExprMain(ctx *ExprMainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitExprCont(ctx *ExprContContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitXorExpr(ctx *XorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitXorExprCont(ctx *XorExprContContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAndExprCont(ctx *AndExprContContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitShiftExpr(ctx *ShiftExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitShiftExprCont(ctx *ShiftExprContContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitArithExpr(ctx *ArithExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitArithExprCont(ctx *ArithExprContContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitTermCont(ctx *TermContContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitPower(ctx *PowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAtomExpr(ctx *AtomExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitParenAtom(ctx *ParenAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitBracketAtom(ctx *BracketAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitDictOrSetAtom(ctx *DictOrSetAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitConstAtom(ctx *ConstAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitConst(ctx *ConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitTestListComp(ctx *TestListCompContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitExprOrStarExpr(ctx *ExprOrStarExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitStarExpr(ctx *StarExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitTrailer(ctx *TrailerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitSubscriptList(ctx *SubscriptListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitExprSubscript(ctx *ExprSubscriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitSliceSubscript(ctx *SliceSubscriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitSliceOp(ctx *SliceOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitDictMaker(ctx *DictMakerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitSetMaker(ctx *SetMakerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitKvDictItem(ctx *KvDictItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitStarsDictItem(ctx *StarsDictItemContext) interface{} {
	return v.VisitChildren(ctx)
}
