// Code generated from Protobuf3.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Protobuf3

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

type BaseProtobuf3Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseProtobuf3Visitor) VisitProto(ctx *ProtoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitSyntax(ctx *SyntaxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitSyntaxExtra(ctx *SyntaxExtraContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitPackageStatement(ctx *PackageStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOption(ctx *OptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOptionName(ctx *OptionNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOptionBody(ctx *OptionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOptionBodyVariable(ctx *OptionBodyVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitTopLevelDef(ctx *TopLevelDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessage(ctx *MessageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessageBody(ctx *MessageBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessageBodyContent(ctx *MessageBodyContentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitEnumDefinition(ctx *EnumDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitEnumBody(ctx *EnumBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitEnumField(ctx *EnumFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitEnumValueOption(ctx *EnumValueOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitExtend(ctx *ExtendContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitService(ctx *ServiceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitRpc(ctx *RpcContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitReserved(ctx *ReservedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitRanges(ctx *RangesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitRangeRule(ctx *RangeRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitFieldNames(ctx *FieldNamesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitTypeRule(ctx *TypeRuleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitSimpleType(ctx *SimpleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitFieldNumber(ctx *FieldNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitField(ctx *FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitFieldOptions(ctx *FieldOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitFieldOption(ctx *FieldOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOneof(ctx *OneofContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOneofField(ctx *OneofFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMapField(ctx *MapFieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitKeyType(ctx *KeyTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitReservedWord(ctx *ReservedWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitFullIdent(ctx *FullIdentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessageName(ctx *MessageNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitEnumName(ctx *EnumNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessageOrEnumName(ctx *MessageOrEnumNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitFieldName(ctx *FieldNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOneofName(ctx *OneofNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMapName(ctx *MapNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitServiceName(ctx *ServiceNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitRpcName(ctx *RpcNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessageType(ctx *MessageTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessageOrEnumType(ctx *MessageOrEnumTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitEmptyStatement(ctx *EmptyStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}
