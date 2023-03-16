// Code generated from Json.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Json

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

type BaseJsonVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJsonVisitor) VisitJson(ctx *JsonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonVisitor) VisitObject(ctx *ObjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}
