package c

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/exp/util/c/parser"
	antlru "github.com/wrmsr/bane/pkg/util/antlr"
)

func TestC(t *testing.T) {
	src := `
int main(int argc, const char * const *argv) {
	printf("hi\n");
	return 0;
}
`

	is := antlr.NewInputStream(src)
	lexer := parser.NewCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.CompilationUnit()

	fmt.Println(tree)

	antlru.Dump(tree, "")
}
