// Code generated from Hocon.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Hocon

import antlr "github.com/wrmsr/bane/core/antlr/runtime"

type BaseHoconVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseHoconVisitor) VisitHocon(ctx *HoconContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitProp(ctx *PropContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitObj(ctx *ObjContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitObjectBegin(ctx *ObjectBeginContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitObjectEnd(ctx *ObjectEndContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitObjectData(ctx *ObjectDataContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayData(ctx *ArrayDataContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitStringData(ctx *StringDataContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitReferenceData(ctx *ReferenceDataContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitNumberData(ctx *NumberDataContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitKey(ctx *KeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitPath(ctx *PathContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayBegin(ctx *ArrayBeginContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayEnd(ctx *ArrayEndContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArray(ctx *ArrayContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayValue(ctx *ArrayValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayString(ctx *ArrayStringContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayReference(ctx *ArrayReferenceContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayNumber(ctx *ArrayNumberContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayObj(ctx *ArrayObjContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitArrayArray(ctx *ArrayArrayContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitV_string(ctx *V_stringContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitV_rawstring(ctx *V_rawstringContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitV_reference(ctx *V_referenceContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseHoconVisitor) VisitRawstring(ctx *RawstringContext) any {
	return v.VisitChildren(ctx)
}
