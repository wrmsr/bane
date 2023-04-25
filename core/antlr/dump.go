package antlr

import (
	"fmt"
	"reflect"

	antlr "github.com/wrmsr/bane/core/antlr/runtime"
)

func Dump(obj any, prefix string) {
	if o, ok := obj.(antlr.ParserRuleContext); ok {
		fmt.Printf("%s%s (%d-%d)\n", prefix, reflect.TypeOf(o).String(), o.GetStart().GetStart(), o.GetStop().GetStop())
		for _, c := range o.GetChildren() {
			Dump(c, prefix+"  ")
		}
		return
	}

	if o, ok := obj.(antlr.TerminalNode); ok {
		si := o.GetSourceInterval()
		fmt.Printf("%s%s (%d-%d) (%s)\n", prefix, reflect.TypeOf(o).String(), si.Start, si.Stop, o.GetText())
		return
	}

	panic(obj)
}
