// Code generated from Protobuf3.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Protobuf3

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

// A complete Visitor for a parse tree produced by Protobuf3Parser.
type Protobuf3Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Protobuf3Parser#proto.
	VisitProto(ctx *ProtoContext) any

	// Visit a parse tree produced by Protobuf3Parser#syntax.
	VisitSyntax(ctx *SyntaxContext) any

	// Visit a parse tree produced by Protobuf3Parser#syntaxExtra.
	VisitSyntaxExtra(ctx *SyntaxExtraContext) any

	// Visit a parse tree produced by Protobuf3Parser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) any

	// Visit a parse tree produced by Protobuf3Parser#packageStatement.
	VisitPackageStatement(ctx *PackageStatementContext) any

	// Visit a parse tree produced by Protobuf3Parser#option.
	VisitOption(ctx *OptionContext) any

	// Visit a parse tree produced by Protobuf3Parser#optionName.
	VisitOptionName(ctx *OptionNameContext) any

	// Visit a parse tree produced by Protobuf3Parser#optionBody.
	VisitOptionBody(ctx *OptionBodyContext) any

	// Visit a parse tree produced by Protobuf3Parser#optionBodyVariable.
	VisitOptionBodyVariable(ctx *OptionBodyVariableContext) any

	// Visit a parse tree produced by Protobuf3Parser#topLevelDef.
	VisitTopLevelDef(ctx *TopLevelDefContext) any

	// Visit a parse tree produced by Protobuf3Parser#message.
	VisitMessage(ctx *MessageContext) any

	// Visit a parse tree produced by Protobuf3Parser#messageBody.
	VisitMessageBody(ctx *MessageBodyContext) any

	// Visit a parse tree produced by Protobuf3Parser#messageBodyContent.
	VisitMessageBodyContent(ctx *MessageBodyContentContext) any

	// Visit a parse tree produced by Protobuf3Parser#enumDefinition.
	VisitEnumDefinition(ctx *EnumDefinitionContext) any

	// Visit a parse tree produced by Protobuf3Parser#enumBody.
	VisitEnumBody(ctx *EnumBodyContext) any

	// Visit a parse tree produced by Protobuf3Parser#enumField.
	VisitEnumField(ctx *EnumFieldContext) any

	// Visit a parse tree produced by Protobuf3Parser#enumValueOption.
	VisitEnumValueOption(ctx *EnumValueOptionContext) any

	// Visit a parse tree produced by Protobuf3Parser#extend.
	VisitExtend(ctx *ExtendContext) any

	// Visit a parse tree produced by Protobuf3Parser#service.
	VisitService(ctx *ServiceContext) any

	// Visit a parse tree produced by Protobuf3Parser#rpc.
	VisitRpc(ctx *RpcContext) any

	// Visit a parse tree produced by Protobuf3Parser#reserved.
	VisitReserved(ctx *ReservedContext) any

	// Visit a parse tree produced by Protobuf3Parser#ranges.
	VisitRanges(ctx *RangesContext) any

	// Visit a parse tree produced by Protobuf3Parser#rangeRule.
	VisitRangeRule(ctx *RangeRuleContext) any

	// Visit a parse tree produced by Protobuf3Parser#fieldNames.
	VisitFieldNames(ctx *FieldNamesContext) any

	// Visit a parse tree produced by Protobuf3Parser#typeRule.
	VisitTypeRule(ctx *TypeRuleContext) any

	// Visit a parse tree produced by Protobuf3Parser#simpleType.
	VisitSimpleType(ctx *SimpleTypeContext) any

	// Visit a parse tree produced by Protobuf3Parser#fieldNumber.
	VisitFieldNumber(ctx *FieldNumberContext) any

	// Visit a parse tree produced by Protobuf3Parser#field.
	VisitField(ctx *FieldContext) any

	// Visit a parse tree produced by Protobuf3Parser#fieldOptions.
	VisitFieldOptions(ctx *FieldOptionsContext) any

	// Visit a parse tree produced by Protobuf3Parser#fieldOption.
	VisitFieldOption(ctx *FieldOptionContext) any

	// Visit a parse tree produced by Protobuf3Parser#oneof.
	VisitOneof(ctx *OneofContext) any

	// Visit a parse tree produced by Protobuf3Parser#oneofField.
	VisitOneofField(ctx *OneofFieldContext) any

	// Visit a parse tree produced by Protobuf3Parser#mapField.
	VisitMapField(ctx *MapFieldContext) any

	// Visit a parse tree produced by Protobuf3Parser#keyType.
	VisitKeyType(ctx *KeyTypeContext) any

	// Visit a parse tree produced by Protobuf3Parser#reservedWord.
	VisitReservedWord(ctx *ReservedWordContext) any

	// Visit a parse tree produced by Protobuf3Parser#fullIdent.
	VisitFullIdent(ctx *FullIdentContext) any

	// Visit a parse tree produced by Protobuf3Parser#messageName.
	VisitMessageName(ctx *MessageNameContext) any

	// Visit a parse tree produced by Protobuf3Parser#enumName.
	VisitEnumName(ctx *EnumNameContext) any

	// Visit a parse tree produced by Protobuf3Parser#messageOrEnumName.
	VisitMessageOrEnumName(ctx *MessageOrEnumNameContext) any

	// Visit a parse tree produced by Protobuf3Parser#fieldName.
	VisitFieldName(ctx *FieldNameContext) any

	// Visit a parse tree produced by Protobuf3Parser#oneofName.
	VisitOneofName(ctx *OneofNameContext) any

	// Visit a parse tree produced by Protobuf3Parser#mapName.
	VisitMapName(ctx *MapNameContext) any

	// Visit a parse tree produced by Protobuf3Parser#serviceName.
	VisitServiceName(ctx *ServiceNameContext) any

	// Visit a parse tree produced by Protobuf3Parser#rpcName.
	VisitRpcName(ctx *RpcNameContext) any

	// Visit a parse tree produced by Protobuf3Parser#messageType.
	VisitMessageType(ctx *MessageTypeContext) any

	// Visit a parse tree produced by Protobuf3Parser#messageOrEnumType.
	VisitMessageOrEnumType(ctx *MessageOrEnumTypeContext) any

	// Visit a parse tree produced by Protobuf3Parser#emptyStatement.
	VisitEmptyStatement(ctx *EmptyStatementContext) any

	// Visit a parse tree produced by Protobuf3Parser#constant.
	VisitConstant(ctx *ConstantContext) any
}
