package json

import (
	"fmt"
	"testing"

	antlr "github.com/wrmsr/bane/core/antlr/runtime"

	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/encoding/json/parser"
	ju "github.com/wrmsr/bane/core/json"
	msh "github.com/wrmsr/bane/core/marshal"
)

func TestParsing(t *testing.T) {
	s := `{"foo": ["bar", 420, true, null]}`

	is := antlr.NewInputStream(s)
	lexer := parser.NewJsonLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewJsonParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Json()

	v := &parseVisitor{}
	root := tree.Accept(v).(Node)
	fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&root)))))
}
