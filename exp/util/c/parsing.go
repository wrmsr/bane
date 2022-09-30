package c

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/exp/util/c/parser"
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

//

var _ parser.CVisitor = &parseVisitor{}

func (v *parseVisitor) Visit(tree antlr.ParseTree) any {
	panic("implement me")
}

func (v *parseVisitor) VisitChildren(node antlr.RuleNode) any {
	panic("implement me")
}

func (v *parseVisitor) VisitTerminal(node antlr.TerminalNode) any {
	panic("implement me")
}

func (v *parseVisitor) VisitErrorNode(node antlr.ErrorNode) any {
	panic("implement me")
}

func (v *parseVisitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitTranslationUnit(ctx *parser.TranslationUnitContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitExternalDeclaration(ctx *parser.ExternalDeclarationContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitFunctionDefinition(ctx *parser.FunctionDefinitionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDeclarationList(ctx *parser.DeclarationListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDeclaration(ctx *parser.DeclarationContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDeclarationSpecifiers(ctx *parser.DeclarationSpecifiersContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDeclarationSpecifiers2(ctx *parser.DeclarationSpecifiers2Context) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDeclarationSpecifier(ctx *parser.DeclarationSpecifierContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitInitDeclaratorList(ctx *parser.InitDeclaratorListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitInitDeclarator(ctx *parser.InitDeclaratorContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitGenericSelection(ctx *parser.GenericSelectionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitGenericAssocList(ctx *parser.GenericAssocListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitGenericAssociation(ctx *parser.GenericAssociationContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitArgumentExpressionList(ctx *parser.ArgumentExpressionListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitUnaryOperator(ctx *parser.UnaryOperatorContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitCastExpression(ctx *parser.CastExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitShiftExpression(ctx *parser.ShiftExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitAndExpression(ctx *parser.AndExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitExclusiveOrExpression(ctx *parser.ExclusiveOrExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitInclusiveOrExpression(ctx *parser.InclusiveOrExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitConditionalExpression(ctx *parser.ConditionalExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitAssignmentExpression(ctx *parser.AssignmentExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitAssignmentOperator(ctx *parser.AssignmentOperatorContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitExpression(ctx *parser.ExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitConstantExpression(ctx *parser.ConstantExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitStorageClassSpecifier(ctx *parser.StorageClassSpecifierContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitTypeSpecifier(ctx *parser.TypeSpecifierContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitStructOrUnionSpecifier(ctx *parser.StructOrUnionSpecifierContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitStructOrUnion(ctx *parser.StructOrUnionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitStructDeclarationList(ctx *parser.StructDeclarationListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitStructDeclaration(ctx *parser.StructDeclarationContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitSpecifierQualifierList(ctx *parser.SpecifierQualifierListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitStructDeclaratorList(ctx *parser.StructDeclaratorListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitStructDeclarator(ctx *parser.StructDeclaratorContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitEnumSpecifier(ctx *parser.EnumSpecifierContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitEnumeratorList(ctx *parser.EnumeratorListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitEnumerator(ctx *parser.EnumeratorContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitEnumerationConstant(ctx *parser.EnumerationConstantContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitAtomicTypeSpecifier(ctx *parser.AtomicTypeSpecifierContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitTypeQualifier(ctx *parser.TypeQualifierContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitFunctionSpecifier(ctx *parser.FunctionSpecifierContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitAlignmentSpecifier(ctx *parser.AlignmentSpecifierContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDeclarator(ctx *parser.DeclaratorContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDirectDeclarator(ctx *parser.DirectDeclaratorContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitVcSpecificModifer(ctx *parser.VcSpecificModiferContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitGccDeclaratorExtension(ctx *parser.GccDeclaratorExtensionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitGccAttributeSpecifier(ctx *parser.GccAttributeSpecifierContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitGccAttributeList(ctx *parser.GccAttributeListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitGccAttribute(ctx *parser.GccAttributeContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitNestedParenthesesBlock(ctx *parser.NestedParenthesesBlockContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitPointer(ctx *parser.PointerContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitTypeQualifierList(ctx *parser.TypeQualifierListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitParameterTypeList(ctx *parser.ParameterTypeListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitParameterList(ctx *parser.ParameterListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitParameterDeclaration(ctx *parser.ParameterDeclarationContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitIdentifierList(ctx *parser.IdentifierListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitTypeName(ctx *parser.TypeNameContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitAbstractDeclarator(ctx *parser.AbstractDeclaratorContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDirectAbstractDeclarator(ctx *parser.DirectAbstractDeclaratorContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitTypedefName(ctx *parser.TypedefNameContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitInitializer(ctx *parser.InitializerContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitInitializerList(ctx *parser.InitializerListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDesignation(ctx *parser.DesignationContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDesignatorList(ctx *parser.DesignatorListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitDesignator(ctx *parser.DesignatorContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitStaticAssertDeclaration(ctx *parser.StaticAssertDeclarationContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitStatement(ctx *parser.StatementContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitLabeledStatement(ctx *parser.LabeledStatementContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitCompoundStatement(ctx *parser.CompoundStatementContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitBlockItemList(ctx *parser.BlockItemListContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitBlockItem(ctx *parser.BlockItemContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitSelectionStatement(ctx *parser.SelectionStatementContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitIterationStatement(ctx *parser.IterationStatementContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitForCondition(ctx *parser.ForConditionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitForDeclaration(ctx *parser.ForDeclarationContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitForExpression(ctx *parser.ForExpressionContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitJumpStatement(ctx *parser.JumpStatementContext) any {
	panic("implement me")
}