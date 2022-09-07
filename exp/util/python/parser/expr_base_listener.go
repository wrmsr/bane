// Code generated from Expr.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExprListener is a complete listener for a parse tree produced by ExprParser.
type BaseExprListener struct{}

var _ ExprListener = &BaseExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSingleExpr is called when production singleExpr is entered.
func (s *BaseExprListener) EnterSingleExpr(ctx *SingleExprContext) {}

// ExitSingleExpr is called when production singleExpr is exited.
func (s *BaseExprListener) ExitSingleExpr(ctx *SingleExprContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseExprListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseExprListener) ExitExpr(ctx *ExprContext) {}

// EnterOrTest is called when production orTest is entered.
func (s *BaseExprListener) EnterOrTest(ctx *OrTestContext) {}

// ExitOrTest is called when production orTest is exited.
func (s *BaseExprListener) ExitOrTest(ctx *OrTestContext) {}

// EnterAndTest is called when production andTest is entered.
func (s *BaseExprListener) EnterAndTest(ctx *AndTestContext) {}

// ExitAndTest is called when production andTest is exited.
func (s *BaseExprListener) ExitAndTest(ctx *AndTestContext) {}

// EnterNotTest is called when production notTest is entered.
func (s *BaseExprListener) EnterNotTest(ctx *NotTestContext) {}

// ExitNotTest is called when production notTest is exited.
func (s *BaseExprListener) ExitNotTest(ctx *NotTestContext) {}

// EnterComparison is called when production comparison is entered.
func (s *BaseExprListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseExprListener) ExitComparison(ctx *ComparisonContext) {}

// EnterCompOp is called when production compOp is entered.
func (s *BaseExprListener) EnterCompOp(ctx *CompOpContext) {}

// ExitCompOp is called when production compOp is exited.
func (s *BaseExprListener) ExitCompOp(ctx *CompOpContext) {}

// EnterExprMain is called when production exprMain is entered.
func (s *BaseExprListener) EnterExprMain(ctx *ExprMainContext) {}

// ExitExprMain is called when production exprMain is exited.
func (s *BaseExprListener) ExitExprMain(ctx *ExprMainContext) {}

// EnterExprCont is called when production exprCont is entered.
func (s *BaseExprListener) EnterExprCont(ctx *ExprContContext) {}

// ExitExprCont is called when production exprCont is exited.
func (s *BaseExprListener) ExitExprCont(ctx *ExprContContext) {}

// EnterXorExpr is called when production xorExpr is entered.
func (s *BaseExprListener) EnterXorExpr(ctx *XorExprContext) {}

// ExitXorExpr is called when production xorExpr is exited.
func (s *BaseExprListener) ExitXorExpr(ctx *XorExprContext) {}

// EnterXorExprCont is called when production xorExprCont is entered.
func (s *BaseExprListener) EnterXorExprCont(ctx *XorExprContContext) {}

// ExitXorExprCont is called when production xorExprCont is exited.
func (s *BaseExprListener) ExitXorExprCont(ctx *XorExprContContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseExprListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseExprListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterAndExprCont is called when production andExprCont is entered.
func (s *BaseExprListener) EnterAndExprCont(ctx *AndExprContContext) {}

// ExitAndExprCont is called when production andExprCont is exited.
func (s *BaseExprListener) ExitAndExprCont(ctx *AndExprContContext) {}

// EnterShiftExpr is called when production shiftExpr is entered.
func (s *BaseExprListener) EnterShiftExpr(ctx *ShiftExprContext) {}

// ExitShiftExpr is called when production shiftExpr is exited.
func (s *BaseExprListener) ExitShiftExpr(ctx *ShiftExprContext) {}

// EnterShiftExprCont is called when production shiftExprCont is entered.
func (s *BaseExprListener) EnterShiftExprCont(ctx *ShiftExprContContext) {}

// ExitShiftExprCont is called when production shiftExprCont is exited.
func (s *BaseExprListener) ExitShiftExprCont(ctx *ShiftExprContContext) {}

// EnterArithExpr is called when production arithExpr is entered.
func (s *BaseExprListener) EnterArithExpr(ctx *ArithExprContext) {}

// ExitArithExpr is called when production arithExpr is exited.
func (s *BaseExprListener) ExitArithExpr(ctx *ArithExprContext) {}

// EnterArithExprCont is called when production arithExprCont is entered.
func (s *BaseExprListener) EnterArithExprCont(ctx *ArithExprContContext) {}

// ExitArithExprCont is called when production arithExprCont is exited.
func (s *BaseExprListener) ExitArithExprCont(ctx *ArithExprContContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseExprListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseExprListener) ExitTerm(ctx *TermContext) {}

// EnterTermCont is called when production termCont is entered.
func (s *BaseExprListener) EnterTermCont(ctx *TermContContext) {}

// ExitTermCont is called when production termCont is exited.
func (s *BaseExprListener) ExitTermCont(ctx *TermContContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseExprListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseExprListener) ExitFactor(ctx *FactorContext) {}

// EnterPower is called when production power is entered.
func (s *BaseExprListener) EnterPower(ctx *PowerContext) {}

// ExitPower is called when production power is exited.
func (s *BaseExprListener) ExitPower(ctx *PowerContext) {}

// EnterAtomExpr is called when production atomExpr is entered.
func (s *BaseExprListener) EnterAtomExpr(ctx *AtomExprContext) {}

// ExitAtomExpr is called when production atomExpr is exited.
func (s *BaseExprListener) ExitAtomExpr(ctx *AtomExprContext) {}

// EnterParenAtom is called when production parenAtom is entered.
func (s *BaseExprListener) EnterParenAtom(ctx *ParenAtomContext) {}

// ExitParenAtom is called when production parenAtom is exited.
func (s *BaseExprListener) ExitParenAtom(ctx *ParenAtomContext) {}

// EnterBracketAtom is called when production bracketAtom is entered.
func (s *BaseExprListener) EnterBracketAtom(ctx *BracketAtomContext) {}

// ExitBracketAtom is called when production bracketAtom is exited.
func (s *BaseExprListener) ExitBracketAtom(ctx *BracketAtomContext) {}

// EnterDictOrSetAtom is called when production dictOrSetAtom is entered.
func (s *BaseExprListener) EnterDictOrSetAtom(ctx *DictOrSetAtomContext) {}

// ExitDictOrSetAtom is called when production dictOrSetAtom is exited.
func (s *BaseExprListener) ExitDictOrSetAtom(ctx *DictOrSetAtomContext) {}

// EnterConstAtom is called when production constAtom is entered.
func (s *BaseExprListener) EnterConstAtom(ctx *ConstAtomContext) {}

// ExitConstAtom is called when production constAtom is exited.
func (s *BaseExprListener) ExitConstAtom(ctx *ConstAtomContext) {}

// EnterConst is called when production const is entered.
func (s *BaseExprListener) EnterConst(ctx *ConstContext) {}

// ExitConst is called when production const is exited.
func (s *BaseExprListener) ExitConst(ctx *ConstContext) {}

// EnterTestListComp is called when production testListComp is entered.
func (s *BaseExprListener) EnterTestListComp(ctx *TestListCompContext) {}

// ExitTestListComp is called when production testListComp is exited.
func (s *BaseExprListener) ExitTestListComp(ctx *TestListCompContext) {}

// EnterExprOrStarExpr is called when production exprOrStarExpr is entered.
func (s *BaseExprListener) EnterExprOrStarExpr(ctx *ExprOrStarExprContext) {}

// ExitExprOrStarExpr is called when production exprOrStarExpr is exited.
func (s *BaseExprListener) ExitExprOrStarExpr(ctx *ExprOrStarExprContext) {}

// EnterStarExpr is called when production starExpr is entered.
func (s *BaseExprListener) EnterStarExpr(ctx *StarExprContext) {}

// ExitStarExpr is called when production starExpr is exited.
func (s *BaseExprListener) ExitStarExpr(ctx *StarExprContext) {}

// EnterTrailer is called when production trailer is entered.
func (s *BaseExprListener) EnterTrailer(ctx *TrailerContext) {}

// ExitTrailer is called when production trailer is exited.
func (s *BaseExprListener) ExitTrailer(ctx *TrailerContext) {}

// EnterSubscriptList is called when production subscriptList is entered.
func (s *BaseExprListener) EnterSubscriptList(ctx *SubscriptListContext) {}

// ExitSubscriptList is called when production subscriptList is exited.
func (s *BaseExprListener) ExitSubscriptList(ctx *SubscriptListContext) {}

// EnterExprSubscript is called when production exprSubscript is entered.
func (s *BaseExprListener) EnterExprSubscript(ctx *ExprSubscriptContext) {}

// ExitExprSubscript is called when production exprSubscript is exited.
func (s *BaseExprListener) ExitExprSubscript(ctx *ExprSubscriptContext) {}

// EnterSliceSubscript is called when production sliceSubscript is entered.
func (s *BaseExprListener) EnterSliceSubscript(ctx *SliceSubscriptContext) {}

// ExitSliceSubscript is called when production sliceSubscript is exited.
func (s *BaseExprListener) ExitSliceSubscript(ctx *SliceSubscriptContext) {}

// EnterSliceOp is called when production sliceOp is entered.
func (s *BaseExprListener) EnterSliceOp(ctx *SliceOpContext) {}

// ExitSliceOp is called when production sliceOp is exited.
func (s *BaseExprListener) ExitSliceOp(ctx *SliceOpContext) {}

// EnterDictMaker is called when production dictMaker is entered.
func (s *BaseExprListener) EnterDictMaker(ctx *DictMakerContext) {}

// ExitDictMaker is called when production dictMaker is exited.
func (s *BaseExprListener) ExitDictMaker(ctx *DictMakerContext) {}

// EnterSetMaker is called when production setMaker is entered.
func (s *BaseExprListener) EnterSetMaker(ctx *SetMakerContext) {}

// ExitSetMaker is called when production setMaker is exited.
func (s *BaseExprListener) ExitSetMaker(ctx *SetMakerContext) {}

// EnterKvDictItem is called when production kvDictItem is entered.
func (s *BaseExprListener) EnterKvDictItem(ctx *KvDictItemContext) {}

// ExitKvDictItem is called when production kvDictItem is exited.
func (s *BaseExprListener) ExitKvDictItem(ctx *KvDictItemContext) {}

// EnterStarsDictItem is called when production starsDictItem is entered.
func (s *BaseExprListener) EnterStarsDictItem(ctx *StarsDictItemContext) {}

// ExitStarsDictItem is called when production starsDictItem is exited.
func (s *BaseExprListener) ExitStarsDictItem(ctx *StarsDictItemContext) {}
