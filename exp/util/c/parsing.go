package c

import (
	"fmt"

	antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

	"github.com/wrmsr/bane/exp/util/c/parser"
	"github.com/wrmsr/bane/pkg/util/check"
)

type parseVisitor struct{}

func maybeNode(o any) Node {
	if o == nil {
		return nil
	}
	return o.(Node)
}

func (v *parseVisitor) defaultResult() any {
	return nil
}

func (v *parseVisitor) aggregateResult(aggregate, nextResult any) any {
	if aggregate != nil {
		panic("should be nil")
	}
	if nextResult == nil {
		panic("must not be nil")
	}
	return nextResult
}

func (v *parseVisitor) error(e error) any {
	panic(e)
}

func (v *parseVisitor) unimplemented(ctx antlr.ParserRuleContext) any {
	return v.error(fmt.Errorf("intentionally unimplemented: %s", ctx))
}

//

var _ parser.CVisitor = &parseVisitor{}

func (v *parseVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *parseVisitor) VisitChildren(node antlr.RuleNode) any {
	var result = v.defaultResult()
	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		if pt, ok := child.(antlr.ParseTree); ok {
			childResult := pt.Accept(v)
			result = v.aggregateResult(result, childResult)
		}
	}
	return result
}

func (v *parseVisitor) VisitTerminal(node antlr.TerminalNode) any {
	return v.defaultResult()
}

func (v *parseVisitor) VisitErrorNode(node antlr.ErrorNode) any {
	return v.defaultResult()
}

func (v *parseVisitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) any {
	return v.Visit(ctx.TranslationUnit())
}

func (v *parseVisitor) VisitTranslationUnit(ctx *parser.TranslationUnitContext) any {
	eds := ctx.AllExternalDeclaration()
	ds := make([]Declaration, len(eds))
	for i, ed := range eds {
		e := v.Visit(ed)
		ds[i] = e.(Declaration)
	}
	return TranslationUnit{Ds: ds}
}

func (v *parseVisitor) VisitExternalDeclaration(ctx *parser.ExternalDeclarationContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitFunctionDefinition(ctx *parser.FunctionDefinitionContext) any {
	var fd FunctionDeclaration
	if ds := ctx.DeclarationSpecifiers(); ds != nil {
		fd.Ds = v.Visit(ds).(DeclarationSpecifiers)
	}
	fd.D = v.Visit(ctx.Declarator()).(Declarator)
	if dl := ctx.DeclarationList(); dl != nil {
		// FIXME:
		_ = v.Visit(dl)
	}
	fd.B = v.Visit(ctx.CompoundStatement()).(CompoundStatement)
	return fd
}

func (v *parseVisitor) VisitDeclarationList(ctx *parser.DeclarationListContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitDeclaration(ctx *parser.DeclarationContext) any {
	if ds := ctx.DeclarationSpecifiers(); ds != nil {
		dss := v.Visit(ds).(DeclarationSpecifiers)
		_ = dss
		var s []Declarator
		if idl := ctx.InitDeclaratorList(); idl != nil {
			ids := idl.(*parser.InitDeclaratorListContext).AllInitDeclarator()
			s = make([]Declarator, len(ids))
			for i, id := range ids {
				dcl := v.Visit(id)
				s[i] = dcl.(Declarator)
			}
		}
		return DeclaratorsDeclaration{
			S:  dss,
			Ds: s,
		}
	}
	panic("unimplemented")
}

func (v *parseVisitor) VisitDeclarationSpecifiers(ctx *parser.DeclarationSpecifiersContext) any {
	ds := ctx.AllDeclarationSpecifier()
	s := make([]DeclarationSpecifier, len(ds))
	for i, d := range ds {
		s[i] = v.Visit(d).(DeclarationSpecifier)
	}
	return DeclarationSpecifiers{S: s}
}

func (v *parseVisitor) VisitDeclarationSpecifiers2(ctx *parser.DeclarationSpecifiers2Context) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitDeclarationSpecifier(ctx *parser.DeclarationSpecifierContext) any {
	var s any
	if c := ctx.StorageClassSpecifier(); c != nil {
		s = v.Visit(c)
	}
	if c := ctx.TypeSpecifier(); c != nil {
		s = v.Visit(c)
	}
	if c := ctx.TypeQualifier(); c != nil {
		s = v.Visit(c)
	}
	if c := ctx.FunctionSpecifier(); c != nil {
		panic(c)
	}
	if c := ctx.AlignmentSpecifier(); c != nil {
		panic(c)
	}
	if s == nil {
		panic("unhandled")
	}
	return s
}

func (v *parseVisitor) VisitInitDeclaratorList(ctx *parser.InitDeclaratorListContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitInitDeclarator(ctx *parser.InitDeclaratorContext) any {
	dcl := v.Visit(ctx.Declarator()).(Declarator)
	if i := ctx.Initializer(); i != nil {
		dcl = InitDeclarator{
			D: dcl,
			I: v.Visit(i).(Expression),
		}
	}
	return dcl
}

func (v *parseVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) any {
	if c := ctx.Identifier(); c != nil {
		return Identifier{
			S: c.GetText(),
		}
	}
	if c := ctx.Constant(); c != nil {
		return Constant{
			S: c.GetText(),
		}
	}
	if cs := ctx.AllStringLiteral(); len(cs) > 0 {
		s := make([]string, len(cs))
		for i, c := range cs {
			s[i] = c.GetText()
		}
		return StringLiteral{
			S: s,
		}
	}
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitGenericSelection(ctx *parser.GenericSelectionContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitGenericAssocList(ctx *parser.GenericAssocListContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitGenericAssociation(ctx *parser.GenericAssociationContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) any {
	var ret Expression
	if pe := ctx.PrimaryExpression(); pe != nil {
		check.Nil(ctx.TypeName())
		check.Nil(ctx.InitializerList())
		ret = v.Visit(pe).(Expression)
	}

	check.NotNil(ret)
	for _, pe2 := range ctx.AllPostfixExpression2() {
		rhs := v.Visit(pe2).(Expression)
		switch rhs := rhs.(type) {

		case Call:
			check.Nil(rhs.Fn)
			rhs.Fn = ret
			ret = rhs

		default:
			panic(rhs)
		}
	}

	return ret
}

func (v *parseVisitor) VisitPostfixExpression2(ctx *parser.PostfixExpression2Context) any {
	switch s := ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetSymbol(); s.GetTokenType() {

	case parser.CParserLeftParen:
		var args []Expression
		if ael := ctx.ArgumentExpressionList(); ael != nil {
			aes := ael.(*parser.ArgumentExpressionListContext).AllAssignmentExpression()
			args = make([]Expression, len(aes))
			for i, ae := range aes {
				args[i] = v.Visit(ae).(Expression)
			}
		}
		return Call{
			Args: args,
		}

	}
	panic("nyi")
}

func (v *parseVisitor) VisitArgumentExpressionList(ctx *parser.ArgumentExpressionListContext) any {
	return v.unimplemented(ctx)
}

func (v *parseVisitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitUnaryOperator(ctx *parser.UnaryOperatorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitCastExpression(ctx *parser.CastExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) any {
	return v.VisitChildren(ctx)
}

func parseBinaryOp(i int) BinaryOp {
	switch i {

	case parser.CParserPlus:
		return AddOp

	}
	panic(i)
}

func (v *parseVisitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) any {
	cs := ctx.GetChildren()
	ret := v.Visit(cs[0].(antlr.ParseTree)).(Expression)
	for i := 1; i < len(cs); i += 2 {
		ot := cs[i].(*antlr.TerminalNodeImpl).GetSymbol().GetTokenType()
		op := parseBinaryOp(ot)
		r := v.Visit(cs[i+1].(antlr.ParseTree)).(Expression)
		ret = Binary{
			Op: op,
			L:  ret,
			R:  r,
		}
	}
	return ret
}

func (v *parseVisitor) VisitShiftExpression(ctx *parser.ShiftExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitAndExpression(ctx *parser.AndExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitExclusiveOrExpression(ctx *parser.ExclusiveOrExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitInclusiveOrExpression(ctx *parser.InclusiveOrExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitConditionalExpression(ctx *parser.ConditionalExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitAssignmentExpression(ctx *parser.AssignmentExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitAssignmentOperator(ctx *parser.AssignmentOperatorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitExpression(ctx *parser.ExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitConstantExpression(ctx *parser.ConstantExpressionContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitStorageClassSpecifier(ctx *parser.StorageClassSpecifierContext) any {
	return ParseStorageClassSpecifier(ctx.GetText())
}

func (v *parseVisitor) VisitTypeSpecifier(ctx *parser.TypeSpecifierContext) any {
	if c := ctx.AtomicTypeSpecifier(); c != nil {
		panic("unimplemented")
	}
	if c := ctx.StructOrUnionSpecifier(); c != nil {
		return v.Visit(c)
	}
	if c := ctx.EnumSpecifier(); c != nil {
		panic("unimplemented")
	}
	if c := ctx.TypedefName(); c != nil {
		return v.Visit(c)
	}
	if c := ctx.ConstantExpression(); c != nil {
		panic("unimplemented")
	}
	return ParseBuiltinType(ctx.GetText())
}

func (v *parseVisitor) VisitStructOrUnionSpecifier(ctx *parser.StructOrUnionSpecifierContext) any {
	var ret StructOrUnionSpecifier
	switch sol := ctx.StructOrUnion().GetText(); sol {
	case "struct":
		ret.U = false // nop
	case "union":
		ret.U = true
	default:
		panic(sol)
	}
	if id := ctx.Identifier(); id != nil {
		ret.I = id.GetText()
	}
	if sdl := ctx.StructDeclarationList(); sdl != nil {
		ret.Ds = v.Visit(ctx.StructDeclarationList()).([]StructDeclaration)
	}
	return ret
}

func (v *parseVisitor) VisitStructOrUnion(ctx *parser.StructOrUnionContext) any {
	return v.unimplemented(ctx)
}

func (v *parseVisitor) VisitStructDeclarationList(ctx *parser.StructDeclarationListContext) any {
	cs := ctx.AllStructDeclaration()
	ds := make([]StructDeclaration, len(cs))
	for i, c := range cs {
		ds[i] = v.Visit(c).(StructDeclaration)
	}
	return ds
}

func (v *parseVisitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) any {
	if sql := ctx.SpecifierQualifierList(); sql != nil {
		ret := StructDeclaration{
			Sqs: v.Visit(sql).([]SpecifierQualifier),
		}
		if sdl := ctx.StructDeclaratorList(); sdl != nil {
			ret.Sds = v.Visit(sdl).([]StructDeclarator)
		}
		return ret
	}
	if sad := ctx.StaticAssertDeclaration(); sad != nil {
		panic(sad)
	}
	panic("unimplemented")
}

func (v *parseVisitor) VisitSpecifierQualifierList(ctx *parser.SpecifierQualifierListContext) any {
	var ret []SpecifierQualifier
	cur := ctx
	for {
		if ts := cur.TypeSpecifier(); ts != nil {
			ret = append(ret, SpecifierQualifier{Ts: v.Visit(ts).(TypeSpecifier)})
		} else if tq := cur.TypeQualifier(); tq != nil {
			ret = append(ret, SpecifierQualifier{Tq: v.Visit(tq).(TypeQualifier)})
		} else {
			panic(cur)
		}
		nxt := cur.SpecifierQualifierList()
		if nxt == nil {
			break
		}
		cur = nxt.(*parser.SpecifierQualifierListContext)
	}
	return ret
}

func (v *parseVisitor) VisitStructDeclaratorList(ctx *parser.StructDeclaratorListContext) any {
	cs := ctx.AllStructDeclarator()
	ret := make([]StructDeclarator, len(cs))
	for i, c := range cs {
		ret[i] = v.Visit(c).(StructDeclarator)
	}
	return ret
}

func (v *parseVisitor) VisitStructDeclarator(ctx *parser.StructDeclaratorContext) any {
	if c := ctx.ConstantExpression(); c != nil {
		panic(c)
	}
	return StructDeclarator{
		D: v.Visit(ctx.Declarator()).(Declarator),
	}
}

func (v *parseVisitor) VisitEnumSpecifier(ctx *parser.EnumSpecifierContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitEnumeratorList(ctx *parser.EnumeratorListContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitEnumerator(ctx *parser.EnumeratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitEnumerationConstant(ctx *parser.EnumerationConstantContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitAtomicTypeSpecifier(ctx *parser.AtomicTypeSpecifierContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitTypeQualifier(ctx *parser.TypeQualifierContext) any {
	return ParseTypeQualifier(ctx.GetText())
}

func (v *parseVisitor) VisitFunctionSpecifier(ctx *parser.FunctionSpecifierContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitAlignmentSpecifier(ctx *parser.AlignmentSpecifierContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitDeclarator(ctx *parser.DeclaratorContext) any {
	if gces := ctx.AllGccDeclaratorExtension(); len(gces) > 0 {
		panic(gces)
	}
	d := v.Visit(ctx.DirectDeclarator()).(Declarator)
	if p := ctx.Pointer(); p != nil {
		pn := v.Visit(p).(Pointer)
		d = PointerDeclarator{P: pn, D: d}
	}
	return d
}

func (v *parseVisitor) VisitDirectDeclarator(ctx *parser.DirectDeclaratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitVcSpecificModifer(ctx *parser.VcSpecificModiferContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitGccDeclaratorExtension(ctx *parser.GccDeclaratorExtensionContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitGccAttributeSpecifier(ctx *parser.GccAttributeSpecifierContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitGccAttributeList(ctx *parser.GccAttributeListContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitGccAttribute(ctx *parser.GccAttributeContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitNestedParenthesesBlock(ctx *parser.NestedParenthesesBlockContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitPointer(ctx *parser.PointerContext) any {
	var s []PointerLevel
	for _, c := range ctx.GetChildren() {
		switch c := c.(type) {
		case *antlr.TerminalNodeImpl:
			s = append(s, PointerLevel{Ca: c.GetSymbol().GetTokenType() == parser.CParserCaret})
		case *parser.TypeQualifierListContext:
			l := &s[len(s)-1]
			if l.Q != nil {
				panic(l.Q)
			}
			q := v.Visit(c).(TypeQualifierList)
			l.Q = &q
		default:
			panic(c)
		}
	}
	return Pointer{S: s}
}

func (v *parseVisitor) VisitTypeQualifierList(ctx *parser.TypeQualifierListContext) any {
	tqs := ctx.AllTypeQualifier()
	s := make([]TypeQualifier, len(tqs))
	for i, tq := range tqs {
		s[i] = v.Visit(tq).(TypeQualifier)
	}
	return TypeQualifierList{S: s}
}

func (v *parseVisitor) VisitParameterTypeList(ctx *parser.ParameterTypeListContext) any {
	pl := v.Visit(ctx.ParameterList()).(ParameterList)
	// FIXME: '...'
	return pl
}

func (v *parseVisitor) VisitParameterList(ctx *parser.ParameterListContext) any {
	pds := ctx.AllParameterDeclaration()
	ps := make([]ParameterDeclaration, len(pds))
	for i, pd := range pds {
		ps[i] = v.Visit(pd).(ParameterDeclaration)
	}
	return ParameterList{Ps: ps}
}

func (v *parseVisitor) VisitParameterDeclaration(ctx *parser.ParameterDeclarationContext) any {
	ds := v.Visit(ctx.DeclarationSpecifiers()).(DeclarationSpecifiers)
	d := v.Visit(ctx.Declarator()).(Declarator)
	return ParameterDeclaration{S: ds, D: d}
}

func (v *parseVisitor) VisitIdentifierList(ctx *parser.IdentifierListContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitTypeName(ctx *parser.TypeNameContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitAbstractDeclarator(ctx *parser.AbstractDeclaratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitDirectAbstractDeclarator(ctx *parser.DirectAbstractDeclaratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitTypedefName(ctx *parser.TypedefNameContext) any {
	return TypedefTypeSpecifier{I: Identifier{S: ctx.Identifier().GetText()}}
}

func (v *parseVisitor) VisitInitializer(ctx *parser.InitializerContext) any {
	if ae := ctx.AssignmentExpression(); ae != nil {
		return v.Visit(ae)
	}
	panic("unimplemented")
}

func (v *parseVisitor) VisitInitializerList(ctx *parser.InitializerListContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitDesignation(ctx *parser.DesignationContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitDesignatorList(ctx *parser.DesignatorListContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitDesignator(ctx *parser.DesignatorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitStaticAssertDeclaration(ctx *parser.StaticAssertDeclarationContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitStatement(ctx *parser.StatementContext) any {
	if e := ctx.LabeledStatement(); e != nil {
		panic(e)
	}
	if e := ctx.CompoundStatement(); e != nil {
		panic(e)
	}
	if e := ctx.ExpressionStatement(); e != nil {
		return v.Visit(e)
	}
	if e := ctx.SelectionStatement(); e != nil {
		panic(e)
	}
	if e := ctx.IterationStatement(); e != nil {
		panic(e)
	}
	if e := ctx.JumpStatement(); e != nil {
		return v.Visit(e)
	}
	panic("unhandled")
}

func (v *parseVisitor) VisitLabeledStatement(ctx *parser.LabeledStatementContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitCompoundStatement(ctx *parser.CompoundStatementContext) any {
	var s []BlockItem
	if bil := ctx.BlockItemList(); bil != nil {
		bis := bil.(*parser.BlockItemListContext).AllBlockItem()
		s = make([]BlockItem, len(bis))
		for i, bi := range bis {
			e := v.Visit(bi)
			s[i] = e.(BlockItem)
		}
	}
	return CompoundStatement{S: s}
}

func (v *parseVisitor) VisitBlockItemList(ctx *parser.BlockItemListContext) any {
	return v.unimplemented(ctx)
}

func (v *parseVisitor) VisitBlockItem(ctx *parser.BlockItemContext) any {
	return v.VisitChildren(ctx)
}

func (v *parseVisitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) any {
	e := v.Visit(ctx.Expression()).(Expression)
	return ExpressionStatement{Expression: e}
}

func (v *parseVisitor) VisitSelectionStatement(ctx *parser.SelectionStatementContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitIterationStatement(ctx *parser.IterationStatementContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitForCondition(ctx *parser.ForConditionContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitForDeclaration(ctx *parser.ForDeclarationContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitForExpression(ctx *parser.ForExpressionContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitJumpStatement(ctx *parser.JumpStatementContext) any {
	switch s := ctx.GetChild(0).(*antlr.TerminalNodeImpl).GetSymbol(); s.GetTokenType() {
	case parser.CParserReturn:
		var e Expression
		if ec := ctx.Expression(); ec != nil {
			e = v.Visit(ec).(Expression)
		}
		return ReturnStatement{V: e}
	default:
		panic(s)
	}
}

func (v *parseVisitor) VisitStaticBracket2DirectDeclarator(ctx *parser.StaticBracket2DirectDeclaratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitBracketStarDirectDeclarator(ctx *parser.BracketStarDirectDeclaratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitBitFieldDirectDeclarator(ctx *parser.BitFieldDirectDeclaratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitVcSpecificDirectDeclarator(ctx *parser.VcSpecificDirectDeclaratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitParenDirectDeclarator(ctx *parser.ParenDirectDeclaratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitParamParenDirectDeclarator(ctx *parser.ParamParenDirectDeclaratorContext) any {
	d := v.Visit(ctx.DirectDeclarator()).(Declarator)
	ps := v.Visit(ctx.ParameterTypeList()).(ParameterList)
	return ParameterListDeclarator{D: d, Ps: ps}
}

func (v *parseVisitor) VisitBracketDirectDeclarator(ctx *parser.BracketDirectDeclaratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitStaticBracketDirectDeclarator(ctx *parser.StaticBracketDirectDeclaratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitVcSpecific2DirectDeclarator(ctx *parser.VcSpecific2DirectDeclaratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitIdentifierParenDirectDeclarator(ctx *parser.IdentifierParenDirectDeclaratorContext) any {
	panic("unimplemented")
}

func (v *parseVisitor) VisitIdentifierDirectDeclarator(ctx *parser.IdentifierDirectDeclaratorContext) any {
	return IdentifierDeclarator{I: Identifier{S: ctx.Identifier().GetText()}}
}
