// Code generated from C.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // C

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

// A complete Visitor for a parse tree produced by CParser.
type CVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CParser#compilationUnit.
	VisitCompilationUnit(ctx *CompilationUnitContext) any

	// Visit a parse tree produced by CParser#translationUnit.
	VisitTranslationUnit(ctx *TranslationUnitContext) any

	// Visit a parse tree produced by CParser#externalDeclaration.
	VisitExternalDeclaration(ctx *ExternalDeclarationContext) any

	// Visit a parse tree produced by CParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) any

	// Visit a parse tree produced by CParser#declarationList.
	VisitDeclarationList(ctx *DeclarationListContext) any

	// Visit a parse tree produced by CParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) any

	// Visit a parse tree produced by CParser#declarationSpecifiers.
	VisitDeclarationSpecifiers(ctx *DeclarationSpecifiersContext) any

	// Visit a parse tree produced by CParser#declarationSpecifiers2.
	VisitDeclarationSpecifiers2(ctx *DeclarationSpecifiers2Context) any

	// Visit a parse tree produced by CParser#declarationSpecifier.
	VisitDeclarationSpecifier(ctx *DeclarationSpecifierContext) any

	// Visit a parse tree produced by CParser#initDeclaratorList.
	VisitInitDeclaratorList(ctx *InitDeclaratorListContext) any

	// Visit a parse tree produced by CParser#initDeclarator.
	VisitInitDeclarator(ctx *InitDeclaratorContext) any

	// Visit a parse tree produced by CParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) any

	// Visit a parse tree produced by CParser#genericSelection.
	VisitGenericSelection(ctx *GenericSelectionContext) any

	// Visit a parse tree produced by CParser#genericAssocList.
	VisitGenericAssocList(ctx *GenericAssocListContext) any

	// Visit a parse tree produced by CParser#genericAssociation.
	VisitGenericAssociation(ctx *GenericAssociationContext) any

	// Visit a parse tree produced by CParser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) any

	// Visit a parse tree produced by CParser#postfixExpression2.
	VisitPostfixExpression2(ctx *PostfixExpression2Context) any

	// Visit a parse tree produced by CParser#argumentExpressionList.
	VisitArgumentExpressionList(ctx *ArgumentExpressionListContext) any

	// Visit a parse tree produced by CParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) any

	// Visit a parse tree produced by CParser#unaryOperator.
	VisitUnaryOperator(ctx *UnaryOperatorContext) any

	// Visit a parse tree produced by CParser#castExpression.
	VisitCastExpression(ctx *CastExpressionContext) any

	// Visit a parse tree produced by CParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) any

	// Visit a parse tree produced by CParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) any

	// Visit a parse tree produced by CParser#shiftExpression.
	VisitShiftExpression(ctx *ShiftExpressionContext) any

	// Visit a parse tree produced by CParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) any

	// Visit a parse tree produced by CParser#equalityExpression.
	VisitEqualityExpression(ctx *EqualityExpressionContext) any

	// Visit a parse tree produced by CParser#andExpression.
	VisitAndExpression(ctx *AndExpressionContext) any

	// Visit a parse tree produced by CParser#exclusiveOrExpression.
	VisitExclusiveOrExpression(ctx *ExclusiveOrExpressionContext) any

	// Visit a parse tree produced by CParser#inclusiveOrExpression.
	VisitInclusiveOrExpression(ctx *InclusiveOrExpressionContext) any

	// Visit a parse tree produced by CParser#logicalAndExpression.
	VisitLogicalAndExpression(ctx *LogicalAndExpressionContext) any

	// Visit a parse tree produced by CParser#logicalOrExpression.
	VisitLogicalOrExpression(ctx *LogicalOrExpressionContext) any

	// Visit a parse tree produced by CParser#conditionalExpression.
	VisitConditionalExpression(ctx *ConditionalExpressionContext) any

	// Visit a parse tree produced by CParser#assignmentExpression.
	VisitAssignmentExpression(ctx *AssignmentExpressionContext) any

	// Visit a parse tree produced by CParser#assignmentOperator.
	VisitAssignmentOperator(ctx *AssignmentOperatorContext) any

	// Visit a parse tree produced by CParser#expression.
	VisitExpression(ctx *ExpressionContext) any

	// Visit a parse tree produced by CParser#constantExpression.
	VisitConstantExpression(ctx *ConstantExpressionContext) any

	// Visit a parse tree produced by CParser#storageClassSpecifier.
	VisitStorageClassSpecifier(ctx *StorageClassSpecifierContext) any

	// Visit a parse tree produced by CParser#typeSpecifier.
	VisitTypeSpecifier(ctx *TypeSpecifierContext) any

	// Visit a parse tree produced by CParser#structOrUnionSpecifier.
	VisitStructOrUnionSpecifier(ctx *StructOrUnionSpecifierContext) any

	// Visit a parse tree produced by CParser#structOrUnion.
	VisitStructOrUnion(ctx *StructOrUnionContext) any

	// Visit a parse tree produced by CParser#structDeclarationList.
	VisitStructDeclarationList(ctx *StructDeclarationListContext) any

	// Visit a parse tree produced by CParser#structDeclaration.
	VisitStructDeclaration(ctx *StructDeclarationContext) any

	// Visit a parse tree produced by CParser#specifierQualifierList.
	VisitSpecifierQualifierList(ctx *SpecifierQualifierListContext) any

	// Visit a parse tree produced by CParser#structDeclaratorList.
	VisitStructDeclaratorList(ctx *StructDeclaratorListContext) any

	// Visit a parse tree produced by CParser#structDeclarator.
	VisitStructDeclarator(ctx *StructDeclaratorContext) any

	// Visit a parse tree produced by CParser#enumSpecifier.
	VisitEnumSpecifier(ctx *EnumSpecifierContext) any

	// Visit a parse tree produced by CParser#enumeratorList.
	VisitEnumeratorList(ctx *EnumeratorListContext) any

	// Visit a parse tree produced by CParser#enumerator.
	VisitEnumerator(ctx *EnumeratorContext) any

	// Visit a parse tree produced by CParser#enumerationConstant.
	VisitEnumerationConstant(ctx *EnumerationConstantContext) any

	// Visit a parse tree produced by CParser#atomicTypeSpecifier.
	VisitAtomicTypeSpecifier(ctx *AtomicTypeSpecifierContext) any

	// Visit a parse tree produced by CParser#typeQualifier.
	VisitTypeQualifier(ctx *TypeQualifierContext) any

	// Visit a parse tree produced by CParser#functionSpecifier.
	VisitFunctionSpecifier(ctx *FunctionSpecifierContext) any

	// Visit a parse tree produced by CParser#alignmentSpecifier.
	VisitAlignmentSpecifier(ctx *AlignmentSpecifierContext) any

	// Visit a parse tree produced by CParser#declarator.
	VisitDeclarator(ctx *DeclaratorContext) any

	// Visit a parse tree produced by CParser#staticBracket2DirectDeclarator.
	VisitStaticBracket2DirectDeclarator(ctx *StaticBracket2DirectDeclaratorContext) any

	// Visit a parse tree produced by CParser#bracketStarDirectDeclarator.
	VisitBracketStarDirectDeclarator(ctx *BracketStarDirectDeclaratorContext) any

	// Visit a parse tree produced by CParser#bitFieldDirectDeclarator.
	VisitBitFieldDirectDeclarator(ctx *BitFieldDirectDeclaratorContext) any

	// Visit a parse tree produced by CParser#vcSpecificDirectDeclarator.
	VisitVcSpecificDirectDeclarator(ctx *VcSpecificDirectDeclaratorContext) any

	// Visit a parse tree produced by CParser#parenDirectDeclarator.
	VisitParenDirectDeclarator(ctx *ParenDirectDeclaratorContext) any

	// Visit a parse tree produced by CParser#paramParenDirectDeclarator.
	VisitParamParenDirectDeclarator(ctx *ParamParenDirectDeclaratorContext) any

	// Visit a parse tree produced by CParser#bracketDirectDeclarator.
	VisitBracketDirectDeclarator(ctx *BracketDirectDeclaratorContext) any

	// Visit a parse tree produced by CParser#staticBracketDirectDeclarator.
	VisitStaticBracketDirectDeclarator(ctx *StaticBracketDirectDeclaratorContext) any

	// Visit a parse tree produced by CParser#vcSpecific2DirectDeclarator.
	VisitVcSpecific2DirectDeclarator(ctx *VcSpecific2DirectDeclaratorContext) any

	// Visit a parse tree produced by CParser#identifierParenDirectDeclarator.
	VisitIdentifierParenDirectDeclarator(ctx *IdentifierParenDirectDeclaratorContext) any

	// Visit a parse tree produced by CParser#identifierDirectDeclarator.
	VisitIdentifierDirectDeclarator(ctx *IdentifierDirectDeclaratorContext) any

	// Visit a parse tree produced by CParser#vcSpecificModifier.
	VisitVcSpecificModifier(ctx *VcSpecificModifierContext) any

	// Visit a parse tree produced by CParser#gccDeclaratorExtension.
	VisitGccDeclaratorExtension(ctx *GccDeclaratorExtensionContext) any

	// Visit a parse tree produced by CParser#gccAttributeSpecifier.
	VisitGccAttributeSpecifier(ctx *GccAttributeSpecifierContext) any

	// Visit a parse tree produced by CParser#gccAttributeList.
	VisitGccAttributeList(ctx *GccAttributeListContext) any

	// Visit a parse tree produced by CParser#gccAttribute.
	VisitGccAttribute(ctx *GccAttributeContext) any

	// Visit a parse tree produced by CParser#nestedParenthesesBlock.
	VisitNestedParenthesesBlock(ctx *NestedParenthesesBlockContext) any

	// Visit a parse tree produced by CParser#pointer.
	VisitPointer(ctx *PointerContext) any

	// Visit a parse tree produced by CParser#typeQualifierList.
	VisitTypeQualifierList(ctx *TypeQualifierListContext) any

	// Visit a parse tree produced by CParser#parameterTypeList.
	VisitParameterTypeList(ctx *ParameterTypeListContext) any

	// Visit a parse tree produced by CParser#parameterList.
	VisitParameterList(ctx *ParameterListContext) any

	// Visit a parse tree produced by CParser#parameterDeclaration.
	VisitParameterDeclaration(ctx *ParameterDeclarationContext) any

	// Visit a parse tree produced by CParser#identifierList.
	VisitIdentifierList(ctx *IdentifierListContext) any

	// Visit a parse tree produced by CParser#typeName.
	VisitTypeName(ctx *TypeNameContext) any

	// Visit a parse tree produced by CParser#abstractDeclarator.
	VisitAbstractDeclarator(ctx *AbstractDeclaratorContext) any

	// Visit a parse tree produced by CParser#parameterTypeListDirectAbstractDeclarator.
	VisitParameterTypeListDirectAbstractDeclarator(ctx *ParameterTypeListDirectAbstractDeclaratorContext) any

	// Visit a parse tree produced by CParser#parenDirectAbstractDeclarator.
	VisitParenDirectAbstractDeclarator(ctx *ParenDirectAbstractDeclaratorContext) any

	// Visit a parse tree produced by CParser#bracketStarDirectAbstractDeclarator.
	VisitBracketStarDirectAbstractDeclarator(ctx *BracketStarDirectAbstractDeclaratorContext) any

	// Visit a parse tree produced by CParser#recTypeQualifierListStaticDirectAbstractDeclarator.
	VisitRecTypeQualifierListStaticDirectAbstractDeclarator(ctx *RecTypeQualifierListStaticDirectAbstractDeclaratorContext) any

	// Visit a parse tree produced by CParser#recBracketStarDirectAbstractDeclarator.
	VisitRecBracketStarDirectAbstractDeclarator(ctx *RecBracketStarDirectAbstractDeclaratorContext) any

	// Visit a parse tree produced by CParser#typeQualifierListDirectAbstractDeclarator.
	VisitTypeQualifierListDirectAbstractDeclarator(ctx *TypeQualifierListDirectAbstractDeclaratorContext) any

	// Visit a parse tree produced by CParser#typeQualifierListStaticDirectAbstractDeclarator.
	VisitTypeQualifierListStaticDirectAbstractDeclarator(ctx *TypeQualifierListStaticDirectAbstractDeclaratorContext) any

	// Visit a parse tree produced by CParser#recStaticTypeQualifierListDirectAbstractDeclarator.
	VisitRecStaticTypeQualifierListDirectAbstractDeclarator(ctx *RecStaticTypeQualifierListDirectAbstractDeclaratorContext) any

	// Visit a parse tree produced by CParser#recParameterTypeListDirectAbstractDeclarator.
	VisitRecParameterTypeListDirectAbstractDeclarator(ctx *RecParameterTypeListDirectAbstractDeclaratorContext) any

	// Visit a parse tree produced by CParser#recTypeQualifierListDirectAbstractDeclarator.
	VisitRecTypeQualifierListDirectAbstractDeclarator(ctx *RecTypeQualifierListDirectAbstractDeclaratorContext) any

	// Visit a parse tree produced by CParser#staticTypeQualifierListDirectAbstractDeclarator.
	VisitStaticTypeQualifierListDirectAbstractDeclarator(ctx *StaticTypeQualifierListDirectAbstractDeclaratorContext) any

	// Visit a parse tree produced by CParser#typedefName.
	VisitTypedefName(ctx *TypedefNameContext) any

	// Visit a parse tree produced by CParser#initializer.
	VisitInitializer(ctx *InitializerContext) any

	// Visit a parse tree produced by CParser#initializerList.
	VisitInitializerList(ctx *InitializerListContext) any

	// Visit a parse tree produced by CParser#designation.
	VisitDesignation(ctx *DesignationContext) any

	// Visit a parse tree produced by CParser#designatorList.
	VisitDesignatorList(ctx *DesignatorListContext) any

	// Visit a parse tree produced by CParser#designator.
	VisitDesignator(ctx *DesignatorContext) any

	// Visit a parse tree produced by CParser#staticAssertDeclaration.
	VisitStaticAssertDeclaration(ctx *StaticAssertDeclarationContext) any

	// Visit a parse tree produced by CParser#statement.
	VisitStatement(ctx *StatementContext) any

	// Visit a parse tree produced by CParser#labeledStatement.
	VisitLabeledStatement(ctx *LabeledStatementContext) any

	// Visit a parse tree produced by CParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) any

	// Visit a parse tree produced by CParser#blockItemList.
	VisitBlockItemList(ctx *BlockItemListContext) any

	// Visit a parse tree produced by CParser#blockItem.
	VisitBlockItem(ctx *BlockItemContext) any

	// Visit a parse tree produced by CParser#expressionStatement.
	VisitExpressionStatement(ctx *ExpressionStatementContext) any

	// Visit a parse tree produced by CParser#selectionStatement.
	VisitSelectionStatement(ctx *SelectionStatementContext) any

	// Visit a parse tree produced by CParser#iterationStatement.
	VisitIterationStatement(ctx *IterationStatementContext) any

	// Visit a parse tree produced by CParser#forCondition.
	VisitForCondition(ctx *ForConditionContext) any

	// Visit a parse tree produced by CParser#forDeclaration.
	VisitForDeclaration(ctx *ForDeclarationContext) any

	// Visit a parse tree produced by CParser#forExpression.
	VisitForExpression(ctx *ForExpressionContext) any

	// Visit a parse tree produced by CParser#jumpStatement.
	VisitJumpStatement(ctx *JumpStatementContext) any
}
