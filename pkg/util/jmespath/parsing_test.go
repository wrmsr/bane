package jmespath

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	"github.com/wrmsr/bane/pkg/util/jmespath/parser"
	ju "github.com/wrmsr/bane/pkg/util/json"
	msh "github.com/wrmsr/bane/pkg/util/marshal"
	bt "github.com/wrmsr/bane/pkg/util/types"
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

type TestError int8

const (
	NoError TestError = iota
	InvalidArityError
	InvalidTypeError
	InvalidValueError
	SyntaxError
	UnknownFunctionError
)

func ParseTestError(s string) TestError {
	switch s {
	case "invalid-arity":
		return InvalidArityError
	case "invalid-type":
		return InvalidTypeError
	case "invalid-value":
		return InvalidValueError
	case "syntax":
		return SyntaxError
	case "unknown-function":
		return UnknownFunctionError
	}
	panic(s)
}

func (e *TestError) UnmarshalText(b []byte) error {
	*e = ParseTestError(string(b))
	return nil
}

type TestBench int8

const (
	NoBench TestBench = iota
	ParseBench
	InterpretBench
	FullBench
)

func ParseTestBench(s string) TestBench {
	switch s {
	case "parse":
		return ParseBench
	case "interpret":
		return InterpretBench
	case "full":
		return FullBench
	}
	panic(s)
}

func (e *TestBench) UnmarshalText(b []byte) error {
	*e = ParseTestBench(string(b))
	return nil
}

type TestCase struct {
	Expression string           `json:"expression"`
	Result     bt.Optional[any] `json:"result"`
	Error      TestError        `json:"error"`
	Bench      TestBench        `json:"bench"`
	Comment    string           `json:"comment"`
}

type TestItem struct {
	Given   any        `json:"given"`
	Comment string     `json:"comment"`
	Cases   []TestCase `json:"cases"`
}

func TestParsing2(t *testing.T) {
	dir := filepath.Join(paths.FindProjectRoot(), "pkg/util/jmespath/dev/compliance/tests")
	check.Must(filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		buf := check.Must1(os.ReadFile(path))
		tis := check.Must1(ju.UnmarshalAs[[]TestItem](buf))
		for _, ti := range tis {
			for _, tc := range ti.Cases {
				if tc.Error == NoError {
					fmt.Println(tc.Expression)

					root := testParse(tc.Expression)
					fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&root)))))

					rt := SimpleRuntime{}
					rt.AddFunctions(Functions...)

					o := Evaluator[any]{rt: rt}.Eval(root, ti.Given)
					fmt.Println(o)

					if tc.Result.Present() {
						r := tc.Result.Value()
						if !reflect.DeepEqual(o, r) {
							Evaluator[any]{rt: rt}.Eval(root, ti.Given)
							tu.AssertDeepEqual(t, o, r)
						}
					}
				}
			}
		}

		return nil
	}))
}
