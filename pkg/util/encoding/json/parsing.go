package json

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/pkg/util/encoding/json/parser"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type parseVisitor struct{}

var _ parser.JsonVisitor = &parseVisitor{}

//

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

//

func (v *parseVisitor) VisitTerminal(node antlr.TerminalNode) any {
	panic(node)
}

func (v *parseVisitor) VisitErrorNode(node antlr.ErrorNode) any {
	panic(node)
}

func (v *parseVisitor) VisitJson(ctx *parser.JsonContext) any {
	return v.Visit(ctx.Value())
}

func (v *parseVisitor) VisitObject(ctx *parser.ObjectContext) any {
	var ps []Pair
	cc := ctx.GetChildCount()
	for i := 0; i < cc; i++ {
		if c, ok := ctx.GetChild(i).(antlr.ParseTree); ok && !bt.Is[antlr.TerminalNode](c) {
			ps = append(ps, v.Visit(c).(Pair))
		}
	}
	return Object{Pairs: ps}
}

func chompQuotes(s string) string {
	return s[1 : len(s)-1]
}

func (v *parseVisitor) VisitPair(ctx *parser.PairContext) any {
	return bt.KvOf(chompQuotes(ctx.STRING().GetText()), v.Visit(ctx.Value()).(Node))
}

func (v *parseVisitor) VisitArray(ctx *parser.ArrayContext) any {
	var vs []Node
	cc := ctx.GetChildCount()
	for i := 0; i < cc; i++ {
		if c, ok := ctx.GetChild(i).(antlr.ParseTree); ok && !bt.Is[antlr.TerminalNode](c) {
			vs = append(vs, v.Visit(c).(Node))
		}
	}
	return Array{Values: vs}
}

func (v *parseVisitor) VisitValue(ctx *parser.ValueContext) any {
	if s := ctx.STRING(); s != nil {
		return String{S: chompQuotes(s.GetText())}
	}
	if n := ctx.NUMBER(); n != nil {
		return Number{S: chompQuotes(n.GetText())}
	}
	if o := ctx.Object(); o != nil {
		return v.Visit(ctx.Object())
	}
	if a := ctx.Array(); a != nil {
		return v.Visit(ctx.Array())
	}
	txt := ctx.GetText()
	switch txt {
	case "true":
		return True{}
	case "false":
		return False{}
	case "null":
		return Null{}
	}
	panic(fmt.Errorf("invalid raw value: %s", txt))
}
