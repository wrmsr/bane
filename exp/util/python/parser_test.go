package python

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/exp/util/python/parser"
)

func TestParser(t *testing.T) {
	for _, testExpr := range []string{
		`420`,
		`None`,
		`"abc"`,
		`'def'`,
		`1+2`,
		//`{'descr': '<f4', 'fortran_order': False, 'shape': (3, 3), }`,
	} {
		is := antlr.NewInputStream(testExpr)
		lexer := parser.NewExprLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewExprParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.BuildParseTrees = true
		tree := p.Expr()

		var dump func(any, string)
		dump = func(o any, p string) {
			if o, ok := o.(antlr.ParserRuleContext); ok {
				fmt.Printf("%s%s\n", p, reflect.TypeOf(o).String())
				for _, c := range o.GetChildren() {
					dump(c, p+"  ")
				}
				return
			}

			if o, ok := o.(antlr.TerminalNode); ok {
				fmt.Printf("%s%s (%s)\n", p, reflect.TypeOf(o).String(), o.GetText())
				return
			}

			panic(o)
		}
		dump(tree, "")

		v := parseVisitor{}
		n := tree.Accept(&v).(Node)
		fmt.Printf("%#v\n", n)
	}
}
