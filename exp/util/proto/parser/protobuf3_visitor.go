// Code generated from Protobuf3.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Protobuf3

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

// A complete Visitor for a parse tree produced by Protobuf3Parser.
type Protobuf3Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Protobuf3Parser#proto.
	VisitProto(ctx *ProtoContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#syntax.
	VisitSyntax(ctx *SyntaxContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#syntaxExtra.
	VisitSyntaxExtra(ctx *SyntaxExtraContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#packageStatement.
	VisitPackageStatement(ctx *PackageStatementContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#option.
	VisitOption(ctx *OptionContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#optionName.
	VisitOptionName(ctx *OptionNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#optionBody.
	VisitOptionBody(ctx *OptionBodyContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#optionBodyVariable.
	VisitOptionBodyVariable(ctx *OptionBodyVariableContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#topLevelDef.
	VisitTopLevelDef(ctx *TopLevelDefContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#message.
	VisitMessage(ctx *MessageContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#messageBody.
	VisitMessageBody(ctx *MessageBodyContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#messageBodyContent.
	VisitMessageBodyContent(ctx *MessageBodyContentContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumDefinition.
	VisitEnumDefinition(ctx *EnumDefinitionContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumBody.
	VisitEnumBody(ctx *EnumBodyContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumField.
	VisitEnumField(ctx *EnumFieldContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumValueOption.
	VisitEnumValueOption(ctx *EnumValueOptionContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#extend.
	VisitExtend(ctx *ExtendContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#service.
	VisitService(ctx *ServiceContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#rpc.
	VisitRpc(ctx *RpcContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#reserved.
	VisitReserved(ctx *ReservedContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#ranges.
	VisitRanges(ctx *RangesContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#rangeRule.
	VisitRangeRule(ctx *RangeRuleContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#fieldNames.
	VisitFieldNames(ctx *FieldNamesContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#typeRule.
	VisitTypeRule(ctx *TypeRuleContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#simpleType.
	VisitSimpleType(ctx *SimpleTypeContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#fieldNumber.
	VisitFieldNumber(ctx *FieldNumberContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#field.
	VisitField(ctx *FieldContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#fieldOptions.
	VisitFieldOptions(ctx *FieldOptionsContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#fieldOption.
	VisitFieldOption(ctx *FieldOptionContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#oneof.
	VisitOneof(ctx *OneofContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#oneofField.
	VisitOneofField(ctx *OneofFieldContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#mapField.
	VisitMapField(ctx *MapFieldContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#keyType.
	VisitKeyType(ctx *KeyTypeContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#reservedWord.
	VisitReservedWord(ctx *ReservedWordContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#fullIdent.
	VisitFullIdent(ctx *FullIdentContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#messageName.
	VisitMessageName(ctx *MessageNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#enumName.
	VisitEnumName(ctx *EnumNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#messageOrEnumName.
	VisitMessageOrEnumName(ctx *MessageOrEnumNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#fieldName.
	VisitFieldName(ctx *FieldNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#oneofName.
	VisitOneofName(ctx *OneofNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#mapName.
	VisitMapName(ctx *MapNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#serviceName.
	VisitServiceName(ctx *ServiceNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#rpcName.
	VisitRpcName(ctx *RpcNameContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#messageType.
	VisitMessageType(ctx *MessageTypeContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#messageOrEnumType.
	VisitMessageOrEnumType(ctx *MessageOrEnumTypeContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#emptyStatement.
	VisitEmptyStatement(ctx *EmptyStatementContext) interface{}

	// Visit a parse tree produced by Protobuf3Parser#constant.
	VisitConstant(ctx *ConstantContext) interface{}
}
