package json

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/pkg/util/encoding/json/parser"
)

type parseVisitor struct{}

var _ parser.JsonVisitor = &parseVisitor{}

func (v *parseVisitor) defaultResult() any {
	return nil
}

func (v *parseVisitor) aggregateResult(aggregate, nextResult any) any {
	if aggregate != nil {
		panic("should be nil")
	}
	if nextResult == nil {
		panic("must not be nil")
	}
	return nextResult
}

func (v *parseVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v *parseVisitor) VisitChildren(node antlr.RuleNode) any {
	var result = v.defaultResult()
	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		if pt, ok := child.(antlr.ParseTree); ok {
			childResult := pt.Accept(v)
			result = v.aggregateResult(result, childResult)
		}
	}
	return result
}

func (v *parseVisitor) VisitTerminal(node antlr.TerminalNode) any {
	panic("implement me")
}

func (v *parseVisitor) VisitErrorNode(node antlr.ErrorNode) any {
	panic("implement me")
}

func (v *parseVisitor) VisitJson(ctx *parser.JsonContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitObj(ctx *parser.ObjContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitPair(ctx *parser.PairContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitArray(ctx *parser.ArrayContext) any {
	panic("implement me")
}

func (v *parseVisitor) VisitValue(ctx *parser.ValueContext) any {
	panic("implement me")
}
