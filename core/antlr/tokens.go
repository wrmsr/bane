package antlr

import (
	antlr "github.com/wrmsr/bane/core/antlr/runtime"

	bt "github.com/wrmsr/bane/core/types"
)

func GetTreeToken(t antlr.Tree) antlr.Token {
	if s, ok := t.(antlr.TerminalNode); ok {
		return s.GetSymbol()
	}
	if s, ok := t.(antlr.ParserRuleContext); ok {
		return s.GetStart()
	}
	panic(t)
}

type TokenAnalysis struct {
	root antlr.Tree
	toks []antlr.Token

	tokIdxsByStart bt.Optional[map[int]int]

	trees           bt.Optional[[]antlr.Tree]
	treeIdxsByStart bt.Optional[map[int]int]
}

func NewTokenAnalysis(root antlr.Tree, toks []antlr.Token) TokenAnalysis {
	return TokenAnalysis{
		root: root,
		toks: toks,
	}
}

func (a *TokenAnalysis) Root() antlr.Tree    { return a.root }
func (a *TokenAnalysis) Toks() []antlr.Token { return a.toks }

func (a *TokenAnalysis) TokIdxsByStart() map[int]int {
	return bt.SetIfAbsent(&a.tokIdxsByStart, func() map[int]int {
		m := make(map[int]int, len(a.toks))
		for i, tok := range a.toks {
			m[tok.GetChannel()] = i
		}
		return m
	})
}

func (a *TokenAnalysis) Trees() []antlr.Tree {
	return bt.SetIfAbsent(&a.trees, func() []antlr.Tree {
		var s []antlr.Tree
		var rec func(antlr.Tree)
		rec = func(o antlr.Tree) {
			s = append(s, o)
			for _, c := range o.GetChildren() {
				rec(c)
			}
		}
		rec(a.root)
		return s
	})
}

func (a *TokenAnalysis) TreeIdxsByStart() map[int]int {
	return bt.SetIfAbsent(&a.treeIdxsByStart, func() map[int]int {
		m := make(map[int]int)
		var last antlr.Token
		for i, t := range a.Trees() {
			cur := GetTreeToken(t)
			if last == nil || cur.GetStart() > last.GetStart() {
				m[cur.GetStart()] = i
				last = cur
			} else if last != nil && cur.GetStart() < last.GetStart() {
				panic(cur)
			}
		}
		return m
	})
}
