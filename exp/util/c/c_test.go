package c

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/exp/util/c/parser"
	antlru "github.com/wrmsr/bane/pkg/util/antlr"
)

func TestC(t *testing.T) {
	cu := func(p *parser.CParser) antlr.ParseTree { return p.CompilationUnit() }
	ex := func(p *parser.CParser) antlr.ParseTree { return p.PrimaryExpression() }
	ts := func(p *parser.CParser) antlr.ParseTree { return p.TypeSpecifier() }

	_ = cu
	_ = ex
	_ = ts

	for _, src := range []struct {
		s   string
		pfn func(p *parser.CParser) antlr.ParseTree
	}{
		//{`1`, ex},
		//{`int`, ts},
		//{`int foo(int x) {
		//	return x;
		//}`, cu},
		//{`int main(int argc, const char * const *argv) {
		//	printf("hi\n");
		//	return 0;
		//}`, cu},
		{`int main(int argc, const char * const *argv) {
			int c = argc + 2;
			printf("hi %d\n", c);
			return 0;
		}`, cu},
	} {
		fmt.Println(src.s)

		is := antlr.NewInputStream(src.s)
		lexer := parser.NewCLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewCParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.BuildParseTrees = true

		tree := src.pfn(p)

		fmt.Println(tree)

		antlru.Dump(tree, "")

		//tu.AssertNoErr(t, dot.Open(context.Background(), dot.RenderString(antlru.Dot(tree))))

		v := &parseVisitor{}
		//v.BaseParseTreeVisitor = v

		n := tree.Accept(v)
		fmt.Printf("%#v\n", n)

		//fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(n)))))

		fmt.Println()
	}
}
