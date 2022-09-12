// Code generated from Hocon.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Hocon

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseHoconVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseHoconVisitor) VisitHocon(ctx *HoconContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitProp(ctx *PropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitObjectBegin(ctx *ObjectBeginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitObjectEnd(ctx *ObjectEndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitObjectData(ctx *ObjectDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayData(ctx *ArrayDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitStringData(ctx *StringDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitReferenceData(ctx *ReferenceDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitNumberData(ctx *NumberDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitKey(ctx *KeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitPath(ctx *PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayBegin(ctx *ArrayBeginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayEnd(ctx *ArrayEndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayValue(ctx *ArrayValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayString(ctx *ArrayStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayReference(ctx *ArrayReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayNumber(ctx *ArrayNumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayObj(ctx *ArrayObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayArray(ctx *ArrayArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitV_string(ctx *V_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitV_rawstring(ctx *V_rawstringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitV_reference(ctx *V_referenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitRawstring(ctx *RawstringContext) interface{} {
	return v.VisitChildren(ctx)
}
