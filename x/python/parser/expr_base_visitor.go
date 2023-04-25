// Code generated from Expr.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Expr

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

type BaseExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseExprVisitor) VisitSingleExpr(ctx *SingleExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitExpr(ctx *ExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitOrTest(ctx *OrTestContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAndTest(ctx *AndTestContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitNotTest(ctx *NotTestContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitComparison(ctx *ComparisonContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitCompOp(ctx *CompOpContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitExprMain(ctx *ExprMainContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitExprCont(ctx *ExprContContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitXorExpr(ctx *XorExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitXorExprCont(ctx *XorExprContContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAndExpr(ctx *AndExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAndExprCont(ctx *AndExprContContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitShiftExpr(ctx *ShiftExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitShiftExprCont(ctx *ShiftExprContContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitArithExpr(ctx *ArithExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitArithExprCont(ctx *ArithExprContContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitTerm(ctx *TermContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitTermCont(ctx *TermContContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitFactor(ctx *FactorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitPower(ctx *PowerContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAtomExpr(ctx *AtomExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitParenAtom(ctx *ParenAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitBracketAtom(ctx *BracketAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitDictOrSetAtom(ctx *DictOrSetAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitConstAtom(ctx *ConstAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitConstVal(ctx *ConstValContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitTestListComp(ctx *TestListCompContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitExprOrStarExpr(ctx *ExprOrStarExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitStarExpr(ctx *StarExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitTrailer(ctx *TrailerContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitSubscriptList(ctx *SubscriptListContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitExprSubscript(ctx *ExprSubscriptContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitSliceSubscript(ctx *SliceSubscriptContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitSliceOp(ctx *SliceOpContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitDictMaker(ctx *DictMakerContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitSetMaker(ctx *SetMakerContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitKvDictItem(ctx *KvDictItemContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitStarsDictItem(ctx *StarsDictItemContext) any {
	return v.VisitChildren(ctx)
}
