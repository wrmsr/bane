// Code generated from Protobuf3.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Protobuf3

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

type BaseProtobuf3Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseProtobuf3Visitor) VisitProto(ctx *ProtoContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitSyntax(ctx *SyntaxContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitSyntaxExtra(ctx *SyntaxExtraContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitImportStatement(ctx *ImportStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitPackageStatement(ctx *PackageStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOption(ctx *OptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOptionName(ctx *OptionNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOptionBody(ctx *OptionBodyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOptionBodyVariable(ctx *OptionBodyVariableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitTopLevelDef(ctx *TopLevelDefContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessage(ctx *MessageContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessageBody(ctx *MessageBodyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessageBodyContent(ctx *MessageBodyContentContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitEnumDefinition(ctx *EnumDefinitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitEnumBody(ctx *EnumBodyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitEnumField(ctx *EnumFieldContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitEnumValueOption(ctx *EnumValueOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitExtend(ctx *ExtendContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitService(ctx *ServiceContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitRpc(ctx *RpcContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitReserved(ctx *ReservedContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitRanges(ctx *RangesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitRangeRule(ctx *RangeRuleContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitFieldNames(ctx *FieldNamesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitTypeRule(ctx *TypeRuleContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitSimpleType(ctx *SimpleTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitFieldNumber(ctx *FieldNumberContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitField(ctx *FieldContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitFieldOptions(ctx *FieldOptionsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitFieldOption(ctx *FieldOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOneof(ctx *OneofContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOneofField(ctx *OneofFieldContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMapField(ctx *MapFieldContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitKeyType(ctx *KeyTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitReservedWord(ctx *ReservedWordContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitFullIdent(ctx *FullIdentContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessageName(ctx *MessageNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitEnumName(ctx *EnumNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessageOrEnumName(ctx *MessageOrEnumNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitFieldName(ctx *FieldNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitOneofName(ctx *OneofNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMapName(ctx *MapNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitServiceName(ctx *ServiceNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitRpcName(ctx *RpcNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessageType(ctx *MessageTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitMessageOrEnumType(ctx *MessageOrEnumTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitEmptyStatement(ctx *EmptyStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseProtobuf3Visitor) VisitConstant(ctx *ConstantContext) any {
	return v.VisitChildren(ctx)
}
