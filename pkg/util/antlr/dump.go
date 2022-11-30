package antlr

import (
	"fmt"
	"reflect"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

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
