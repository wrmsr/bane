package antlr

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/wrmsr/bane/pkg/util/slices"
	stru "github.com/wrmsr/bane/pkg/util/strings"
)

//

func Dump(obj any, prefix string) {
	if o, ok := obj.(antlr.ParserRuleContext); ok {
		fmt.Printf("%s%s\n", prefix, reflect.TypeOf(o).String())
		for _, c := range o.GetChildren() {
			Dump(c, prefix+"  ")
		}
		return
	}

	if o, ok := obj.(antlr.TerminalNode); ok {
		fmt.Printf("%s%s (%s)\n", prefix, reflect.TypeOf(o).String(), o.GetText())
		return
	}

	panic(obj)
}

func FindTreeChildren[T antlr.ParserRuleContext](parent antlr.Tree) []T {
	var cs []T
	for _, ctx := range parent.GetChildren() {
		if c, ok := ctx.(T); ok {
			cs = append(cs, c)
			break
		}
	}
	return cs
}

//

type GenericError struct {
	Args []any
}

func (e GenericError) Error() string {
	return fmt.Sprintf("generic antlr error: %s", strings.Join(slices.Map[any, string](stru.ToString, e.Args), ", "))
}

//

type PanicErrorListener struct{}

var _ antlr.ErrorListener = PanicErrorListener{}

func (l PanicErrorListener) error(args ...any) {
	panic(GenericError{Args: args})
}

func (l PanicErrorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol any,
	line,
	column int,
	msg string,
	e antlr.RecognitionException,
) {
	l.error(
		offendingSymbol,
		line,
		column,
		msg,
		e,
	)
}

func (l PanicErrorListener) ReportAmbiguity(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex,
	stopIndex int,
	exact bool,
	ambigAlts *antlr.BitSet,
	configs antlr.ATNConfigSet,
) {
	l.error(
		dfa,
		startIndex,
		stopIndex,
		exact,
		ambigAlts,
		configs,
	)
}

func (l PanicErrorListener) ReportAttemptingFullContext(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex,
	stopIndex int,
	conflictingAlts *antlr.BitSet,
	configs antlr.ATNConfigSet,
) {
	l.error(
		dfa,
		startIndex,
		stopIndex,
		conflictingAlts,
		configs,
	)
}

func (l PanicErrorListener) ReportContextSensitivity(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex,
	stopIndex,
	prediction int,
	configs antlr.ATNConfigSet,
) {
	l.error(
		dfa,
		startIndex,
		stopIndex,
		prediction,
		configs,
	)
}
