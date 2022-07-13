package jmespath

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/jmespath/parser"
	ju "github.com/wrmsr/bane/pkg/util/json"
	msh "github.com/wrmsr/bane/pkg/util/marshal"
)

func TestParsing(t *testing.T) {
	is := antlr.NewInputStream(testExpr)
	lexer := parser.NewJmespathLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewJmespathParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.SingleExpression()

	v := &parseVisitor{}
	root := tree.Accept(v).(Node)
	fmt.Printf("%+v\n", &root)

	mv := check.Must1(msh.Marshal(&root))
	fmt.Println(check.Must1(ju.MarshalPretty(mv)))

	var v2 Node
	check.Must(msh.Unmarshal(mv, &v2))
	fmt.Println(v2)
}
