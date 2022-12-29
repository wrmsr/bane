package antlr

import (
	antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"
)

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

func WalkChildren(obj any, fn func(any) bool) bool {
	if o, ok := obj.(antlr.ParserRuleContext); ok {
		if !fn(o) {
			return false
		}
		for _, c := range o.GetChildren() {
			if !WalkChildren(c, fn) {
				return false
			}
		}
		return true
	}

	if _, ok := obj.(antlr.TerminalNode); ok {
		return fn(obj)
	}

	panic(obj)
}

func BuildIndex(obj any) map[any]int {
	m := make(map[any]int)
	WalkChildren(obj, func(obj any) bool {
		m[obj] = len(m)
		return true
	})
	return m
}
