// Code generated from C.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // C

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

// CListener is a complete listener for a parse tree produced by CParser.
type CListener interface {
	antlr.ParseTreeListener

	// EnterCompilationUnit is called when entering the compilationUnit production.
	EnterCompilationUnit(c *CompilationUnitContext)

	// EnterTranslationUnit is called when entering the translationUnit production.
	EnterTranslationUnit(c *TranslationUnitContext)

	// EnterExternalDeclaration is called when entering the externalDeclaration production.
	EnterExternalDeclaration(c *ExternalDeclarationContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterDeclarationList is called when entering the declarationList production.
	EnterDeclarationList(c *DeclarationListContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterDeclarationSpecifiers is called when entering the declarationSpecifiers production.
	EnterDeclarationSpecifiers(c *DeclarationSpecifiersContext)

	// EnterDeclarationSpecifiers2 is called when entering the declarationSpecifiers2 production.
	EnterDeclarationSpecifiers2(c *DeclarationSpecifiers2Context)

	// EnterDeclarationSpecifier is called when entering the declarationSpecifier production.
	EnterDeclarationSpecifier(c *DeclarationSpecifierContext)

	// EnterInitDeclaratorList is called when entering the initDeclaratorList production.
	EnterInitDeclaratorList(c *InitDeclaratorListContext)

	// EnterInitDeclarator is called when entering the initDeclarator production.
	EnterInitDeclarator(c *InitDeclaratorContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterGenericSelection is called when entering the genericSelection production.
	EnterGenericSelection(c *GenericSelectionContext)

	// EnterGenericAssocList is called when entering the genericAssocList production.
	EnterGenericAssocList(c *GenericAssocListContext)

	// EnterGenericAssociation is called when entering the genericAssociation production.
	EnterGenericAssociation(c *GenericAssociationContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterPostfixExpression2 is called when entering the postfixExpression2 production.
	EnterPostfixExpression2(c *PostfixExpression2Context)

	// EnterArgumentExpressionList is called when entering the argumentExpressionList production.
	EnterArgumentExpressionList(c *ArgumentExpressionListContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterCastExpression is called when entering the castExpression production.
	EnterCastExpression(c *CastExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterShiftExpression is called when entering the shiftExpression production.
	EnterShiftExpression(c *ShiftExpressionContext)

	// EnterRelationalExpression is called when entering the relationalExpression production.
	EnterRelationalExpression(c *RelationalExpressionContext)

	// EnterEqualityExpression is called when entering the equalityExpression production.
	EnterEqualityExpression(c *EqualityExpressionContext)

	// EnterAndExpression is called when entering the andExpression production.
	EnterAndExpression(c *AndExpressionContext)

	// EnterExclusiveOrExpression is called when entering the exclusiveOrExpression production.
	EnterExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// EnterInclusiveOrExpression is called when entering the inclusiveOrExpression production.
	EnterInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// EnterLogicalAndExpression is called when entering the logicalAndExpression production.
	EnterLogicalAndExpression(c *LogicalAndExpressionContext)

	// EnterLogicalOrExpression is called when entering the logicalOrExpression production.
	EnterLogicalOrExpression(c *LogicalOrExpressionContext)

	// EnterConditionalExpression is called when entering the conditionalExpression production.
	EnterConditionalExpression(c *ConditionalExpressionContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterAssignmentOperator is called when entering the assignmentOperator production.
	EnterAssignmentOperator(c *AssignmentOperatorContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterConstantExpression is called when entering the constantExpression production.
	EnterConstantExpression(c *ConstantExpressionContext)

	// EnterStorageClassSpecifier is called when entering the storageClassSpecifier production.
	EnterStorageClassSpecifier(c *StorageClassSpecifierContext)

	// EnterTypeSpecifier is called when entering the typeSpecifier production.
	EnterTypeSpecifier(c *TypeSpecifierContext)

	// EnterStructOrUnionSpecifier is called when entering the structOrUnionSpecifier production.
	EnterStructOrUnionSpecifier(c *StructOrUnionSpecifierContext)

	// EnterStructOrUnion is called when entering the structOrUnion production.
	EnterStructOrUnion(c *StructOrUnionContext)

	// EnterStructDeclarationList is called when entering the structDeclarationList production.
	EnterStructDeclarationList(c *StructDeclarationListContext)

	// EnterStructDeclaration is called when entering the structDeclaration production.
	EnterStructDeclaration(c *StructDeclarationContext)

	// EnterSpecifierQualifierList is called when entering the specifierQualifierList production.
	EnterSpecifierQualifierList(c *SpecifierQualifierListContext)

	// EnterStructDeclaratorList is called when entering the structDeclaratorList production.
	EnterStructDeclaratorList(c *StructDeclaratorListContext)

	// EnterStructDeclarator is called when entering the structDeclarator production.
	EnterStructDeclarator(c *StructDeclaratorContext)

	// EnterEnumSpecifier is called when entering the enumSpecifier production.
	EnterEnumSpecifier(c *EnumSpecifierContext)

	// EnterEnumeratorList is called when entering the enumeratorList production.
	EnterEnumeratorList(c *EnumeratorListContext)

	// EnterEnumerator is called when entering the enumerator production.
	EnterEnumerator(c *EnumeratorContext)

	// EnterEnumerationConstant is called when entering the enumerationConstant production.
	EnterEnumerationConstant(c *EnumerationConstantContext)

	// EnterAtomicTypeSpecifier is called when entering the atomicTypeSpecifier production.
	EnterAtomicTypeSpecifier(c *AtomicTypeSpecifierContext)

	// EnterTypeQualifier is called when entering the typeQualifier production.
	EnterTypeQualifier(c *TypeQualifierContext)

	// EnterFunctionSpecifier is called when entering the functionSpecifier production.
	EnterFunctionSpecifier(c *FunctionSpecifierContext)

	// EnterAlignmentSpecifier is called when entering the alignmentSpecifier production.
	EnterAlignmentSpecifier(c *AlignmentSpecifierContext)

	// EnterDeclarator is called when entering the declarator production.
	EnterDeclarator(c *DeclaratorContext)

	// EnterStaticBracket2DirectDeclarator is called when entering the staticBracket2DirectDeclarator production.
	EnterStaticBracket2DirectDeclarator(c *StaticBracket2DirectDeclaratorContext)

	// EnterBracketStarDirectDeclarator is called when entering the bracketStarDirectDeclarator production.
	EnterBracketStarDirectDeclarator(c *BracketStarDirectDeclaratorContext)

	// EnterBitFieldDirectDeclarator is called when entering the bitFieldDirectDeclarator production.
	EnterBitFieldDirectDeclarator(c *BitFieldDirectDeclaratorContext)

	// EnterVcSpecificDirectDeclarator is called when entering the vcSpecificDirectDeclarator production.
	EnterVcSpecificDirectDeclarator(c *VcSpecificDirectDeclaratorContext)

	// EnterParenDirectDeclarator is called when entering the parenDirectDeclarator production.
	EnterParenDirectDeclarator(c *ParenDirectDeclaratorContext)

	// EnterParamParenDirectDeclarator is called when entering the paramParenDirectDeclarator production.
	EnterParamParenDirectDeclarator(c *ParamParenDirectDeclaratorContext)

	// EnterBracketDirectDeclarator is called when entering the bracketDirectDeclarator production.
	EnterBracketDirectDeclarator(c *BracketDirectDeclaratorContext)

	// EnterStaticBracketDirectDeclarator is called when entering the staticBracketDirectDeclarator production.
	EnterStaticBracketDirectDeclarator(c *StaticBracketDirectDeclaratorContext)

	// EnterVcSpecific2DirectDeclarator is called when entering the vcSpecific2DirectDeclarator production.
	EnterVcSpecific2DirectDeclarator(c *VcSpecific2DirectDeclaratorContext)

	// EnterIdentifierParenDirectDeclarator is called when entering the identifierParenDirectDeclarator production.
	EnterIdentifierParenDirectDeclarator(c *IdentifierParenDirectDeclaratorContext)

	// EnterIdentifierDirectDeclarator is called when entering the identifierDirectDeclarator production.
	EnterIdentifierDirectDeclarator(c *IdentifierDirectDeclaratorContext)

	// EnterVcSpecificModifier is called when entering the vcSpecificModifier production.
	EnterVcSpecificModifier(c *VcSpecificModifierContext)

	// EnterGccDeclaratorExtension is called when entering the gccDeclaratorExtension production.
	EnterGccDeclaratorExtension(c *GccDeclaratorExtensionContext)

	// EnterGccAttributeSpecifier is called when entering the gccAttributeSpecifier production.
	EnterGccAttributeSpecifier(c *GccAttributeSpecifierContext)

	// EnterGccAttributeList is called when entering the gccAttributeList production.
	EnterGccAttributeList(c *GccAttributeListContext)

	// EnterGccAttribute is called when entering the gccAttribute production.
	EnterGccAttribute(c *GccAttributeContext)

	// EnterNestedParenthesesBlock is called when entering the nestedParenthesesBlock production.
	EnterNestedParenthesesBlock(c *NestedParenthesesBlockContext)

	// EnterPointer is called when entering the pointer production.
	EnterPointer(c *PointerContext)

	// EnterTypeQualifierList is called when entering the typeQualifierList production.
	EnterTypeQualifierList(c *TypeQualifierListContext)

	// EnterParameterTypeList is called when entering the parameterTypeList production.
	EnterParameterTypeList(c *ParameterTypeListContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameterDeclaration is called when entering the parameterDeclaration production.
	EnterParameterDeclaration(c *ParameterDeclarationContext)

	// EnterIdentifierList is called when entering the identifierList production.
	EnterIdentifierList(c *IdentifierListContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// EnterAbstractDeclarator is called when entering the abstractDeclarator production.
	EnterAbstractDeclarator(c *AbstractDeclaratorContext)

	// EnterParameterTypeListDirectAbstractDeclarator is called when entering the parameterTypeListDirectAbstractDeclarator production.
	EnterParameterTypeListDirectAbstractDeclarator(c *ParameterTypeListDirectAbstractDeclaratorContext)

	// EnterParenDirectAbstractDeclarator is called when entering the parenDirectAbstractDeclarator production.
	EnterParenDirectAbstractDeclarator(c *ParenDirectAbstractDeclaratorContext)

	// EnterBracketStarDirectAbstractDeclarator is called when entering the bracketStarDirectAbstractDeclarator production.
	EnterBracketStarDirectAbstractDeclarator(c *BracketStarDirectAbstractDeclaratorContext)

	// EnterRecTypeQualifierListStaticDirectAbstractDeclarator is called when entering the recTypeQualifierListStaticDirectAbstractDeclarator production.
	EnterRecTypeQualifierListStaticDirectAbstractDeclarator(c *RecTypeQualifierListStaticDirectAbstractDeclaratorContext)

	// EnterRecBracketStarDirectAbstractDeclarator is called when entering the recBracketStarDirectAbstractDeclarator production.
	EnterRecBracketStarDirectAbstractDeclarator(c *RecBracketStarDirectAbstractDeclaratorContext)

	// EnterTypeQualifierListDirectAbstractDeclarator is called when entering the typeQualifierListDirectAbstractDeclarator production.
	EnterTypeQualifierListDirectAbstractDeclarator(c *TypeQualifierListDirectAbstractDeclaratorContext)

	// EnterTypeQualifierListStaticDirectAbstractDeclarator is called when entering the typeQualifierListStaticDirectAbstractDeclarator production.
	EnterTypeQualifierListStaticDirectAbstractDeclarator(c *TypeQualifierListStaticDirectAbstractDeclaratorContext)

	// EnterRecStaticTypeQualifierListDirectAbstractDeclarator is called when entering the recStaticTypeQualifierListDirectAbstractDeclarator production.
	EnterRecStaticTypeQualifierListDirectAbstractDeclarator(c *RecStaticTypeQualifierListDirectAbstractDeclaratorContext)

	// EnterRecParameterTypeListDirectAbstractDeclarator is called when entering the recParameterTypeListDirectAbstractDeclarator production.
	EnterRecParameterTypeListDirectAbstractDeclarator(c *RecParameterTypeListDirectAbstractDeclaratorContext)

	// EnterRecTypeQualifierListDirectAbstractDeclarator is called when entering the recTypeQualifierListDirectAbstractDeclarator production.
	EnterRecTypeQualifierListDirectAbstractDeclarator(c *RecTypeQualifierListDirectAbstractDeclaratorContext)

	// EnterStaticTypeQualifierListDirectAbstractDeclarator is called when entering the staticTypeQualifierListDirectAbstractDeclarator production.
	EnterStaticTypeQualifierListDirectAbstractDeclarator(c *StaticTypeQualifierListDirectAbstractDeclaratorContext)

	// EnterTypedefName is called when entering the typedefName production.
	EnterTypedefName(c *TypedefNameContext)

	// EnterInitializer is called when entering the initializer production.
	EnterInitializer(c *InitializerContext)

	// EnterInitializerList is called when entering the initializerList production.
	EnterInitializerList(c *InitializerListContext)

	// EnterDesignation is called when entering the designation production.
	EnterDesignation(c *DesignationContext)

	// EnterDesignatorList is called when entering the designatorList production.
	EnterDesignatorList(c *DesignatorListContext)

	// EnterDesignator is called when entering the designator production.
	EnterDesignator(c *DesignatorContext)

	// EnterStaticAssertDeclaration is called when entering the staticAssertDeclaration production.
	EnterStaticAssertDeclaration(c *StaticAssertDeclarationContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLabeledStatement is called when entering the labeledStatement production.
	EnterLabeledStatement(c *LabeledStatementContext)

	// EnterCompoundStatement is called when entering the compoundStatement production.
	EnterCompoundStatement(c *CompoundStatementContext)

	// EnterBlockItemList is called when entering the blockItemList production.
	EnterBlockItemList(c *BlockItemListContext)

	// EnterBlockItem is called when entering the blockItem production.
	EnterBlockItem(c *BlockItemContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterSelectionStatement is called when entering the selectionStatement production.
	EnterSelectionStatement(c *SelectionStatementContext)

	// EnterIterationStatement is called when entering the iterationStatement production.
	EnterIterationStatement(c *IterationStatementContext)

	// EnterForCondition is called when entering the forCondition production.
	EnterForCondition(c *ForConditionContext)

	// EnterForDeclaration is called when entering the forDeclaration production.
	EnterForDeclaration(c *ForDeclarationContext)

	// EnterForExpression is called when entering the forExpression production.
	EnterForExpression(c *ForExpressionContext)

	// EnterJumpStatement is called when entering the jumpStatement production.
	EnterJumpStatement(c *JumpStatementContext)

	// ExitCompilationUnit is called when exiting the compilationUnit production.
	ExitCompilationUnit(c *CompilationUnitContext)

	// ExitTranslationUnit is called when exiting the translationUnit production.
	ExitTranslationUnit(c *TranslationUnitContext)

	// ExitExternalDeclaration is called when exiting the externalDeclaration production.
	ExitExternalDeclaration(c *ExternalDeclarationContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitDeclarationList is called when exiting the declarationList production.
	ExitDeclarationList(c *DeclarationListContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitDeclarationSpecifiers is called when exiting the declarationSpecifiers production.
	ExitDeclarationSpecifiers(c *DeclarationSpecifiersContext)

	// ExitDeclarationSpecifiers2 is called when exiting the declarationSpecifiers2 production.
	ExitDeclarationSpecifiers2(c *DeclarationSpecifiers2Context)

	// ExitDeclarationSpecifier is called when exiting the declarationSpecifier production.
	ExitDeclarationSpecifier(c *DeclarationSpecifierContext)

	// ExitInitDeclaratorList is called when exiting the initDeclaratorList production.
	ExitInitDeclaratorList(c *InitDeclaratorListContext)

	// ExitInitDeclarator is called when exiting the initDeclarator production.
	ExitInitDeclarator(c *InitDeclaratorContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitGenericSelection is called when exiting the genericSelection production.
	ExitGenericSelection(c *GenericSelectionContext)

	// ExitGenericAssocList is called when exiting the genericAssocList production.
	ExitGenericAssocList(c *GenericAssocListContext)

	// ExitGenericAssociation is called when exiting the genericAssociation production.
	ExitGenericAssociation(c *GenericAssociationContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitPostfixExpression2 is called when exiting the postfixExpression2 production.
	ExitPostfixExpression2(c *PostfixExpression2Context)

	// ExitArgumentExpressionList is called when exiting the argumentExpressionList production.
	ExitArgumentExpressionList(c *ArgumentExpressionListContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitCastExpression is called when exiting the castExpression production.
	ExitCastExpression(c *CastExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitShiftExpression is called when exiting the shiftExpression production.
	ExitShiftExpression(c *ShiftExpressionContext)

	// ExitRelationalExpression is called when exiting the relationalExpression production.
	ExitRelationalExpression(c *RelationalExpressionContext)

	// ExitEqualityExpression is called when exiting the equalityExpression production.
	ExitEqualityExpression(c *EqualityExpressionContext)

	// ExitAndExpression is called when exiting the andExpression production.
	ExitAndExpression(c *AndExpressionContext)

	// ExitExclusiveOrExpression is called when exiting the exclusiveOrExpression production.
	ExitExclusiveOrExpression(c *ExclusiveOrExpressionContext)

	// ExitInclusiveOrExpression is called when exiting the inclusiveOrExpression production.
	ExitInclusiveOrExpression(c *InclusiveOrExpressionContext)

	// ExitLogicalAndExpression is called when exiting the logicalAndExpression production.
	ExitLogicalAndExpression(c *LogicalAndExpressionContext)

	// ExitLogicalOrExpression is called when exiting the logicalOrExpression production.
	ExitLogicalOrExpression(c *LogicalOrExpressionContext)

	// ExitConditionalExpression is called when exiting the conditionalExpression production.
	ExitConditionalExpression(c *ConditionalExpressionContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitAssignmentOperator is called when exiting the assignmentOperator production.
	ExitAssignmentOperator(c *AssignmentOperatorContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitConstantExpression is called when exiting the constantExpression production.
	ExitConstantExpression(c *ConstantExpressionContext)

	// ExitStorageClassSpecifier is called when exiting the storageClassSpecifier production.
	ExitStorageClassSpecifier(c *StorageClassSpecifierContext)

	// ExitTypeSpecifier is called when exiting the typeSpecifier production.
	ExitTypeSpecifier(c *TypeSpecifierContext)

	// ExitStructOrUnionSpecifier is called when exiting the structOrUnionSpecifier production.
	ExitStructOrUnionSpecifier(c *StructOrUnionSpecifierContext)

	// ExitStructOrUnion is called when exiting the structOrUnion production.
	ExitStructOrUnion(c *StructOrUnionContext)

	// ExitStructDeclarationList is called when exiting the structDeclarationList production.
	ExitStructDeclarationList(c *StructDeclarationListContext)

	// ExitStructDeclaration is called when exiting the structDeclaration production.
	ExitStructDeclaration(c *StructDeclarationContext)

	// ExitSpecifierQualifierList is called when exiting the specifierQualifierList production.
	ExitSpecifierQualifierList(c *SpecifierQualifierListContext)

	// ExitStructDeclaratorList is called when exiting the structDeclaratorList production.
	ExitStructDeclaratorList(c *StructDeclaratorListContext)

	// ExitStructDeclarator is called when exiting the structDeclarator production.
	ExitStructDeclarator(c *StructDeclaratorContext)

	// ExitEnumSpecifier is called when exiting the enumSpecifier production.
	ExitEnumSpecifier(c *EnumSpecifierContext)

	// ExitEnumeratorList is called when exiting the enumeratorList production.
	ExitEnumeratorList(c *EnumeratorListContext)

	// ExitEnumerator is called when exiting the enumerator production.
	ExitEnumerator(c *EnumeratorContext)

	// ExitEnumerationConstant is called when exiting the enumerationConstant production.
	ExitEnumerationConstant(c *EnumerationConstantContext)

	// ExitAtomicTypeSpecifier is called when exiting the atomicTypeSpecifier production.
	ExitAtomicTypeSpecifier(c *AtomicTypeSpecifierContext)

	// ExitTypeQualifier is called when exiting the typeQualifier production.
	ExitTypeQualifier(c *TypeQualifierContext)

	// ExitFunctionSpecifier is called when exiting the functionSpecifier production.
	ExitFunctionSpecifier(c *FunctionSpecifierContext)

	// ExitAlignmentSpecifier is called when exiting the alignmentSpecifier production.
	ExitAlignmentSpecifier(c *AlignmentSpecifierContext)

	// ExitDeclarator is called when exiting the declarator production.
	ExitDeclarator(c *DeclaratorContext)

	// ExitStaticBracket2DirectDeclarator is called when exiting the staticBracket2DirectDeclarator production.
	ExitStaticBracket2DirectDeclarator(c *StaticBracket2DirectDeclaratorContext)

	// ExitBracketStarDirectDeclarator is called when exiting the bracketStarDirectDeclarator production.
	ExitBracketStarDirectDeclarator(c *BracketStarDirectDeclaratorContext)

	// ExitBitFieldDirectDeclarator is called when exiting the bitFieldDirectDeclarator production.
	ExitBitFieldDirectDeclarator(c *BitFieldDirectDeclaratorContext)

	// ExitVcSpecificDirectDeclarator is called when exiting the vcSpecificDirectDeclarator production.
	ExitVcSpecificDirectDeclarator(c *VcSpecificDirectDeclaratorContext)

	// ExitParenDirectDeclarator is called when exiting the parenDirectDeclarator production.
	ExitParenDirectDeclarator(c *ParenDirectDeclaratorContext)

	// ExitParamParenDirectDeclarator is called when exiting the paramParenDirectDeclarator production.
	ExitParamParenDirectDeclarator(c *ParamParenDirectDeclaratorContext)

	// ExitBracketDirectDeclarator is called when exiting the bracketDirectDeclarator production.
	ExitBracketDirectDeclarator(c *BracketDirectDeclaratorContext)

	// ExitStaticBracketDirectDeclarator is called when exiting the staticBracketDirectDeclarator production.
	ExitStaticBracketDirectDeclarator(c *StaticBracketDirectDeclaratorContext)

	// ExitVcSpecific2DirectDeclarator is called when exiting the vcSpecific2DirectDeclarator production.
	ExitVcSpecific2DirectDeclarator(c *VcSpecific2DirectDeclaratorContext)

	// ExitIdentifierParenDirectDeclarator is called when exiting the identifierParenDirectDeclarator production.
	ExitIdentifierParenDirectDeclarator(c *IdentifierParenDirectDeclaratorContext)

	// ExitIdentifierDirectDeclarator is called when exiting the identifierDirectDeclarator production.
	ExitIdentifierDirectDeclarator(c *IdentifierDirectDeclaratorContext)

	// ExitVcSpecificModifier is called when exiting the vcSpecificModifier production.
	ExitVcSpecificModifier(c *VcSpecificModifierContext)

	// ExitGccDeclaratorExtension is called when exiting the gccDeclaratorExtension production.
	ExitGccDeclaratorExtension(c *GccDeclaratorExtensionContext)

	// ExitGccAttributeSpecifier is called when exiting the gccAttributeSpecifier production.
	ExitGccAttributeSpecifier(c *GccAttributeSpecifierContext)

	// ExitGccAttributeList is called when exiting the gccAttributeList production.
	ExitGccAttributeList(c *GccAttributeListContext)

	// ExitGccAttribute is called when exiting the gccAttribute production.
	ExitGccAttribute(c *GccAttributeContext)

	// ExitNestedParenthesesBlock is called when exiting the nestedParenthesesBlock production.
	ExitNestedParenthesesBlock(c *NestedParenthesesBlockContext)

	// ExitPointer is called when exiting the pointer production.
	ExitPointer(c *PointerContext)

	// ExitTypeQualifierList is called when exiting the typeQualifierList production.
	ExitTypeQualifierList(c *TypeQualifierListContext)

	// ExitParameterTypeList is called when exiting the parameterTypeList production.
	ExitParameterTypeList(c *ParameterTypeListContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameterDeclaration is called when exiting the parameterDeclaration production.
	ExitParameterDeclaration(c *ParameterDeclarationContext)

	// ExitIdentifierList is called when exiting the identifierList production.
	ExitIdentifierList(c *IdentifierListContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)

	// ExitAbstractDeclarator is called when exiting the abstractDeclarator production.
	ExitAbstractDeclarator(c *AbstractDeclaratorContext)

	// ExitParameterTypeListDirectAbstractDeclarator is called when exiting the parameterTypeListDirectAbstractDeclarator production.
	ExitParameterTypeListDirectAbstractDeclarator(c *ParameterTypeListDirectAbstractDeclaratorContext)

	// ExitParenDirectAbstractDeclarator is called when exiting the parenDirectAbstractDeclarator production.
	ExitParenDirectAbstractDeclarator(c *ParenDirectAbstractDeclaratorContext)

	// ExitBracketStarDirectAbstractDeclarator is called when exiting the bracketStarDirectAbstractDeclarator production.
	ExitBracketStarDirectAbstractDeclarator(c *BracketStarDirectAbstractDeclaratorContext)

	// ExitRecTypeQualifierListStaticDirectAbstractDeclarator is called when exiting the recTypeQualifierListStaticDirectAbstractDeclarator production.
	ExitRecTypeQualifierListStaticDirectAbstractDeclarator(c *RecTypeQualifierListStaticDirectAbstractDeclaratorContext)

	// ExitRecBracketStarDirectAbstractDeclarator is called when exiting the recBracketStarDirectAbstractDeclarator production.
	ExitRecBracketStarDirectAbstractDeclarator(c *RecBracketStarDirectAbstractDeclaratorContext)

	// ExitTypeQualifierListDirectAbstractDeclarator is called when exiting the typeQualifierListDirectAbstractDeclarator production.
	ExitTypeQualifierListDirectAbstractDeclarator(c *TypeQualifierListDirectAbstractDeclaratorContext)

	// ExitTypeQualifierListStaticDirectAbstractDeclarator is called when exiting the typeQualifierListStaticDirectAbstractDeclarator production.
	ExitTypeQualifierListStaticDirectAbstractDeclarator(c *TypeQualifierListStaticDirectAbstractDeclaratorContext)

	// ExitRecStaticTypeQualifierListDirectAbstractDeclarator is called when exiting the recStaticTypeQualifierListDirectAbstractDeclarator production.
	ExitRecStaticTypeQualifierListDirectAbstractDeclarator(c *RecStaticTypeQualifierListDirectAbstractDeclaratorContext)

	// ExitRecParameterTypeListDirectAbstractDeclarator is called when exiting the recParameterTypeListDirectAbstractDeclarator production.
	ExitRecParameterTypeListDirectAbstractDeclarator(c *RecParameterTypeListDirectAbstractDeclaratorContext)

	// ExitRecTypeQualifierListDirectAbstractDeclarator is called when exiting the recTypeQualifierListDirectAbstractDeclarator production.
	ExitRecTypeQualifierListDirectAbstractDeclarator(c *RecTypeQualifierListDirectAbstractDeclaratorContext)

	// ExitStaticTypeQualifierListDirectAbstractDeclarator is called when exiting the staticTypeQualifierListDirectAbstractDeclarator production.
	ExitStaticTypeQualifierListDirectAbstractDeclarator(c *StaticTypeQualifierListDirectAbstractDeclaratorContext)

	// ExitTypedefName is called when exiting the typedefName production.
	ExitTypedefName(c *TypedefNameContext)

	// ExitInitializer is called when exiting the initializer production.
	ExitInitializer(c *InitializerContext)

	// ExitInitializerList is called when exiting the initializerList production.
	ExitInitializerList(c *InitializerListContext)

	// ExitDesignation is called when exiting the designation production.
	ExitDesignation(c *DesignationContext)

	// ExitDesignatorList is called when exiting the designatorList production.
	ExitDesignatorList(c *DesignatorListContext)

	// ExitDesignator is called when exiting the designator production.
	ExitDesignator(c *DesignatorContext)

	// ExitStaticAssertDeclaration is called when exiting the staticAssertDeclaration production.
	ExitStaticAssertDeclaration(c *StaticAssertDeclarationContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLabeledStatement is called when exiting the labeledStatement production.
	ExitLabeledStatement(c *LabeledStatementContext)

	// ExitCompoundStatement is called when exiting the compoundStatement production.
	ExitCompoundStatement(c *CompoundStatementContext)

	// ExitBlockItemList is called when exiting the blockItemList production.
	ExitBlockItemList(c *BlockItemListContext)

	// ExitBlockItem is called when exiting the blockItem production.
	ExitBlockItem(c *BlockItemContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitSelectionStatement is called when exiting the selectionStatement production.
	ExitSelectionStatement(c *SelectionStatementContext)

	// ExitIterationStatement is called when exiting the iterationStatement production.
	ExitIterationStatement(c *IterationStatementContext)

	// ExitForCondition is called when exiting the forCondition production.
	ExitForCondition(c *ForConditionContext)

	// ExitForDeclaration is called when exiting the forDeclaration production.
	ExitForDeclaration(c *ForDeclarationContext)

	// ExitForExpression is called when exiting the forExpression production.
	ExitForExpression(c *ForExpressionContext)

	// ExitJumpStatement is called when exiting the jumpStatement production.
	ExitJumpStatement(c *JumpStatementContext)
}
