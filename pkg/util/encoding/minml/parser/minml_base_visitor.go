// Code generated from Minml.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Minml

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseMinmlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMinmlVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitObj(ctx *ObjContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitPair(ctx *PairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitArray(ctx *ArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitStringValue(ctx *StringValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitTrue(ctx *TrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitFalse(ctx *FalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}
