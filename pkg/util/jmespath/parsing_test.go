package jmespath

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	"github.com/wrmsr/bane/pkg/util/jmespath/parser"
	ju "github.com/wrmsr/bane/pkg/util/json"
	msh "github.com/wrmsr/bane/pkg/util/marshal"
)

func testParse(s string) Node {
	is := antlr.NewInputStream(s)
	lexer := parser.NewJmespathLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewJmespathParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.SingleExpression()

	v := &parseVisitor{}
	return tree.Accept(v).(Node)
}

func TestParsing(t *testing.T) {
	root := testParse(testExpr)
	fmt.Printf("%+v\n", &root)

	mv := check.Must1(msh.Marshal(&root))
	fmt.Println(check.Must1(ju.MarshalPretty(mv)))

	var v2 Node
	check.Must(msh.Unmarshal(mv, &v2))

	fmt.Println(root)
	fmt.Println(v2)
	tu.AssertDeepEqual(t, root, v2)
}

func TestParsing2(t *testing.T) {

}
