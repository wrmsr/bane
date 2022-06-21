// Code generated from Json.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Json

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseJsonVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJsonVisitor) VisitJson(ctx *JsonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonVisitor) VisitObj(ctx *ObjContext) interface{} {
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
