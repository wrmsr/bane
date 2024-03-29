// Code generated from Expr.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Expr

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

// ExprListener is a complete listener for a parse tree produced by ExprParser.
type ExprListener interface {
	antlr.ParseTreeListener

	// EnterSingleExpr is called when entering the singleExpr production.
	EnterSingleExpr(c *SingleExprContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterOrTest is called when entering the orTest production.
	EnterOrTest(c *OrTestContext)

	// EnterAndTest is called when entering the andTest production.
	EnterAndTest(c *AndTestContext)

	// EnterNotTest is called when entering the notTest production.
	EnterNotTest(c *NotTestContext)

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterCompOp is called when entering the compOp production.
	EnterCompOp(c *CompOpContext)

	// EnterExprMain is called when entering the exprMain production.
	EnterExprMain(c *ExprMainContext)

	// EnterExprCont is called when entering the exprCont production.
	EnterExprCont(c *ExprContContext)

	// EnterXorExpr is called when entering the xorExpr production.
	EnterXorExpr(c *XorExprContext)

	// EnterXorExprCont is called when entering the xorExprCont production.
	EnterXorExprCont(c *XorExprContContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterAndExprCont is called when entering the andExprCont production.
	EnterAndExprCont(c *AndExprContContext)

	// EnterShiftExpr is called when entering the shiftExpr production.
	EnterShiftExpr(c *ShiftExprContext)

	// EnterShiftExprCont is called when entering the shiftExprCont production.
	EnterShiftExprCont(c *ShiftExprContContext)

	// EnterArithExpr is called when entering the arithExpr production.
	EnterArithExpr(c *ArithExprContext)

	// EnterArithExprCont is called when entering the arithExprCont production.
	EnterArithExprCont(c *ArithExprContContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterTermCont is called when entering the termCont production.
	EnterTermCont(c *TermContContext)

	// EnterFactor is called when entering the factor production.
	EnterFactor(c *FactorContext)

	// EnterPower is called when entering the power production.
	EnterPower(c *PowerContext)

	// EnterAtomExpr is called when entering the atomExpr production.
	EnterAtomExpr(c *AtomExprContext)

	// EnterParenAtom is called when entering the parenAtom production.
	EnterParenAtom(c *ParenAtomContext)

	// EnterBracketAtom is called when entering the bracketAtom production.
	EnterBracketAtom(c *BracketAtomContext)

	// EnterDictOrSetAtom is called when entering the dictOrSetAtom production.
	EnterDictOrSetAtom(c *DictOrSetAtomContext)

	// EnterConstAtom is called when entering the constAtom production.
	EnterConstAtom(c *ConstAtomContext)

	// EnterConstVal is called when entering the constVal production.
	EnterConstVal(c *ConstValContext)

	// EnterTestListComp is called when entering the testListComp production.
	EnterTestListComp(c *TestListCompContext)

	// EnterExprOrStarExpr is called when entering the exprOrStarExpr production.
	EnterExprOrStarExpr(c *ExprOrStarExprContext)

	// EnterStarExpr is called when entering the starExpr production.
	EnterStarExpr(c *StarExprContext)

	// EnterTrailer is called when entering the trailer production.
	EnterTrailer(c *TrailerContext)

	// EnterSubscriptList is called when entering the subscriptList production.
	EnterSubscriptList(c *SubscriptListContext)

	// EnterExprSubscript is called when entering the exprSubscript production.
	EnterExprSubscript(c *ExprSubscriptContext)

	// EnterSliceSubscript is called when entering the sliceSubscript production.
	EnterSliceSubscript(c *SliceSubscriptContext)

	// EnterSliceOp is called when entering the sliceOp production.
	EnterSliceOp(c *SliceOpContext)

	// EnterDictMaker is called when entering the dictMaker production.
	EnterDictMaker(c *DictMakerContext)

	// EnterSetMaker is called when entering the setMaker production.
	EnterSetMaker(c *SetMakerContext)

	// EnterKvDictItem is called when entering the kvDictItem production.
	EnterKvDictItem(c *KvDictItemContext)

	// EnterStarsDictItem is called when entering the starsDictItem production.
	EnterStarsDictItem(c *StarsDictItemContext)

	// ExitSingleExpr is called when exiting the singleExpr production.
	ExitSingleExpr(c *SingleExprContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitOrTest is called when exiting the orTest production.
	ExitOrTest(c *OrTestContext)

	// ExitAndTest is called when exiting the andTest production.
	ExitAndTest(c *AndTestContext)

	// ExitNotTest is called when exiting the notTest production.
	ExitNotTest(c *NotTestContext)

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitCompOp is called when exiting the compOp production.
	ExitCompOp(c *CompOpContext)

	// ExitExprMain is called when exiting the exprMain production.
	ExitExprMain(c *ExprMainContext)

	// ExitExprCont is called when exiting the exprCont production.
	ExitExprCont(c *ExprContContext)

	// ExitXorExpr is called when exiting the xorExpr production.
	ExitXorExpr(c *XorExprContext)

	// ExitXorExprCont is called when exiting the xorExprCont production.
	ExitXorExprCont(c *XorExprContContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitAndExprCont is called when exiting the andExprCont production.
	ExitAndExprCont(c *AndExprContContext)

	// ExitShiftExpr is called when exiting the shiftExpr production.
	ExitShiftExpr(c *ShiftExprContext)

	// ExitShiftExprCont is called when exiting the shiftExprCont production.
	ExitShiftExprCont(c *ShiftExprContContext)

	// ExitArithExpr is called when exiting the arithExpr production.
	ExitArithExpr(c *ArithExprContext)

	// ExitArithExprCont is called when exiting the arithExprCont production.
	ExitArithExprCont(c *ArithExprContContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitTermCont is called when exiting the termCont production.
	ExitTermCont(c *TermContContext)

	// ExitFactor is called when exiting the factor production.
	ExitFactor(c *FactorContext)

	// ExitPower is called when exiting the power production.
	ExitPower(c *PowerContext)

	// ExitAtomExpr is called when exiting the atomExpr production.
	ExitAtomExpr(c *AtomExprContext)

	// ExitParenAtom is called when exiting the parenAtom production.
	ExitParenAtom(c *ParenAtomContext)

	// ExitBracketAtom is called when exiting the bracketAtom production.
	ExitBracketAtom(c *BracketAtomContext)

	// ExitDictOrSetAtom is called when exiting the dictOrSetAtom production.
	ExitDictOrSetAtom(c *DictOrSetAtomContext)

	// ExitConstAtom is called when exiting the constAtom production.
	ExitConstAtom(c *ConstAtomContext)

	// ExitConstVal is called when exiting the constVal production.
	ExitConstVal(c *ConstValContext)

	// ExitTestListComp is called when exiting the testListComp production.
	ExitTestListComp(c *TestListCompContext)

	// ExitExprOrStarExpr is called when exiting the exprOrStarExpr production.
	ExitExprOrStarExpr(c *ExprOrStarExprContext)

	// ExitStarExpr is called when exiting the starExpr production.
	ExitStarExpr(c *StarExprContext)

	// ExitTrailer is called when exiting the trailer production.
	ExitTrailer(c *TrailerContext)

	// ExitSubscriptList is called when exiting the subscriptList production.
	ExitSubscriptList(c *SubscriptListContext)

	// ExitExprSubscript is called when exiting the exprSubscript production.
	ExitExprSubscript(c *ExprSubscriptContext)

	// ExitSliceSubscript is called when exiting the sliceSubscript production.
	ExitSliceSubscript(c *SliceSubscriptContext)

	// ExitSliceOp is called when exiting the sliceOp production.
	ExitSliceOp(c *SliceOpContext)

	// ExitDictMaker is called when exiting the dictMaker production.
	ExitDictMaker(c *DictMakerContext)

	// ExitSetMaker is called when exiting the setMaker production.
	ExitSetMaker(c *SetMakerContext)

	// ExitKvDictItem is called when exiting the kvDictItem production.
	ExitKvDictItem(c *KvDictItemContext)

	// ExitStarsDictItem is called when exiting the starsDictItem production.
	ExitStarsDictItem(c *StarsDictItemContext)
}
