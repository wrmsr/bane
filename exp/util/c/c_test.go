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

	for _, src := range []struct {
		s   string
		pfn func(p *parser.CParser) antlr.ParseTree
	}{
		{`1`, ex},
		{`int main(int argc, const char * const *argv) {
			printf("hi\n");
			return 0;
		}`, cu},
	} {
		is := antlr.NewInputStream(src.s)
		lexer := parser.NewCLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewCParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.BuildParseTrees = true

		tree := src.pfn(p)

		fmt.Println(tree)

		antlru.Dump(tree, "")

		v := &parseVisitor{}
		//v.BaseParseTreeVisitor = v

		n := tree.Accept(v)
		fmt.Printf("%#v\n", n)
	}
}
