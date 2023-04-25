// Code generated from Hocon.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Hocon

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

// BaseHoconListener is a complete listener for a parse tree produced by HoconParser.
type BaseHoconListener struct{}

var _ HoconListener = &BaseHoconListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseHoconListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseHoconListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseHoconListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseHoconListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterHocon is called when production hocon is entered.
func (s *BaseHoconListener) EnterHocon(ctx *HoconContext) {}

// ExitHocon is called when production hocon is exited.
func (s *BaseHoconListener) ExitHocon(ctx *HoconContext) {}

// EnterProp is called when production prop is entered.
func (s *BaseHoconListener) EnterProp(ctx *PropContext) {}

// ExitProp is called when production prop is exited.
func (s *BaseHoconListener) ExitProp(ctx *PropContext) {}

// EnterObj is called when production obj is entered.
func (s *BaseHoconListener) EnterObj(ctx *ObjContext) {}

// ExitObj is called when production obj is exited.
func (s *BaseHoconListener) ExitObj(ctx *ObjContext) {}

// EnterObjectBegin is called when production objectBegin is entered.
func (s *BaseHoconListener) EnterObjectBegin(ctx *ObjectBeginContext) {}

// ExitObjectBegin is called when production objectBegin is exited.
func (s *BaseHoconListener) ExitObjectBegin(ctx *ObjectBeginContext) {}

// EnterObjectEnd is called when production objectEnd is entered.
func (s *BaseHoconListener) EnterObjectEnd(ctx *ObjectEndContext) {}

// ExitObjectEnd is called when production objectEnd is exited.
func (s *BaseHoconListener) ExitObjectEnd(ctx *ObjectEndContext) {}

// EnterObjectData is called when production objectData is entered.
func (s *BaseHoconListener) EnterObjectData(ctx *ObjectDataContext) {}

// ExitObjectData is called when production objectData is exited.
func (s *BaseHoconListener) ExitObjectData(ctx *ObjectDataContext) {}

// EnterArrayData is called when production arrayData is entered.
func (s *BaseHoconListener) EnterArrayData(ctx *ArrayDataContext) {}

// ExitArrayData is called when production arrayData is exited.
func (s *BaseHoconListener) ExitArrayData(ctx *ArrayDataContext) {}

// EnterStringData is called when production stringData is entered.
func (s *BaseHoconListener) EnterStringData(ctx *StringDataContext) {}

// ExitStringData is called when production stringData is exited.
func (s *BaseHoconListener) ExitStringData(ctx *StringDataContext) {}

// EnterReferenceData is called when production referenceData is entered.
func (s *BaseHoconListener) EnterReferenceData(ctx *ReferenceDataContext) {}

// ExitReferenceData is called when production referenceData is exited.
func (s *BaseHoconListener) ExitReferenceData(ctx *ReferenceDataContext) {}

// EnterNumberData is called when production numberData is entered.
func (s *BaseHoconListener) EnterNumberData(ctx *NumberDataContext) {}

// ExitNumberData is called when production numberData is exited.
func (s *BaseHoconListener) ExitNumberData(ctx *NumberDataContext) {}

// EnterKey is called when production key is entered.
func (s *BaseHoconListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseHoconListener) ExitKey(ctx *KeyContext) {}

// EnterPath is called when production path is entered.
func (s *BaseHoconListener) EnterPath(ctx *PathContext) {}

// ExitPath is called when production path is exited.
func (s *BaseHoconListener) ExitPath(ctx *PathContext) {}

// EnterArrayBegin is called when production arrayBegin is entered.
func (s *BaseHoconListener) EnterArrayBegin(ctx *ArrayBeginContext) {}

// ExitArrayBegin is called when production arrayBegin is exited.
func (s *BaseHoconListener) ExitArrayBegin(ctx *ArrayBeginContext) {}

// EnterArrayEnd is called when production arrayEnd is entered.
func (s *BaseHoconListener) EnterArrayEnd(ctx *ArrayEndContext) {}

// ExitArrayEnd is called when production arrayEnd is exited.
func (s *BaseHoconListener) ExitArrayEnd(ctx *ArrayEndContext) {}

// EnterArray is called when production array is entered.
func (s *BaseHoconListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseHoconListener) ExitArray(ctx *ArrayContext) {}

// EnterArrayValue is called when production arrayValue is entered.
func (s *BaseHoconListener) EnterArrayValue(ctx *ArrayValueContext) {}

// ExitArrayValue is called when production arrayValue is exited.
func (s *BaseHoconListener) ExitArrayValue(ctx *ArrayValueContext) {}

// EnterArrayString is called when production arrayString is entered.
func (s *BaseHoconListener) EnterArrayString(ctx *ArrayStringContext) {}

// ExitArrayString is called when production arrayString is exited.
func (s *BaseHoconListener) ExitArrayString(ctx *ArrayStringContext) {}

// EnterArrayReference is called when production arrayReference is entered.
func (s *BaseHoconListener) EnterArrayReference(ctx *ArrayReferenceContext) {}

// ExitArrayReference is called when production arrayReference is exited.
func (s *BaseHoconListener) ExitArrayReference(ctx *ArrayReferenceContext) {}

// EnterArrayNumber is called when production arrayNumber is entered.
func (s *BaseHoconListener) EnterArrayNumber(ctx *ArrayNumberContext) {}

// ExitArrayNumber is called when production arrayNumber is exited.
func (s *BaseHoconListener) ExitArrayNumber(ctx *ArrayNumberContext) {}

// EnterArrayObj is called when production arrayObj is entered.
func (s *BaseHoconListener) EnterArrayObj(ctx *ArrayObjContext) {}

// ExitArrayObj is called when production arrayObj is exited.
func (s *BaseHoconListener) ExitArrayObj(ctx *ArrayObjContext) {}

// EnterArrayArray is called when production arrayArray is entered.
func (s *BaseHoconListener) EnterArrayArray(ctx *ArrayArrayContext) {}

// ExitArrayArray is called when production arrayArray is exited.
func (s *BaseHoconListener) ExitArrayArray(ctx *ArrayArrayContext) {}

// EnterV_string is called when production v_string is entered.
func (s *BaseHoconListener) EnterV_string(ctx *V_stringContext) {}

// ExitV_string is called when production v_string is exited.
func (s *BaseHoconListener) ExitV_string(ctx *V_stringContext) {}

// EnterV_rawstring is called when production v_rawstring is entered.
func (s *BaseHoconListener) EnterV_rawstring(ctx *V_rawstringContext) {}

// ExitV_rawstring is called when production v_rawstring is exited.
func (s *BaseHoconListener) ExitV_rawstring(ctx *V_rawstringContext) {}

// EnterV_reference is called when production v_reference is entered.
func (s *BaseHoconListener) EnterV_reference(ctx *V_referenceContext) {}

// ExitV_reference is called when production v_reference is exited.
func (s *BaseHoconListener) ExitV_reference(ctx *V_referenceContext) {}

// EnterRawstring is called when production rawstring is entered.
func (s *BaseHoconListener) EnterRawstring(ctx *RawstringContext) {}

// ExitRawstring is called when production rawstring is exited.
func (s *BaseHoconListener) ExitRawstring(ctx *RawstringContext) {}
