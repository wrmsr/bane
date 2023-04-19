// Code generated from Minml.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Minml

import antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

type BaseMinmlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMinmlVisitor) VisitRoot(ctx *RootContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitValue(ctx *ValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitObj(ctx *ObjContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitPair(ctx *PairContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitArray(ctx *ArrayContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitIdentifier(ctx *IdentifierContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitStringValue(ctx *StringValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitNumber(ctx *NumberContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitTrue(ctx *TrueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitFalse(ctx *FalseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMinmlVisitor) VisitNull(ctx *NullContext) any {
	return v.VisitChildren(ctx)
}
