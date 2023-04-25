package antlr

import (
	"fmt"
	"strings"

	antlr "github.com/wrmsr/bane/core/antlr/runtime"

	"github.com/wrmsr/bane/core/slices"
	stru "github.com/wrmsr/bane/core/strings"
)

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
