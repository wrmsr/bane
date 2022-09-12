//go:build !nodev

package test

import (
	"fmt"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	antlru "github.com/wrmsr/bane/pkg/util/antlr"
	"github.com/wrmsr/bane/pkg/util/antlr/dev/test/parser"
)

//

type SuperChatVisitor struct {
	Super parser.ChatVisitor
}

func (v SuperChatVisitor) SuperVisitChildren(node antlr.RuleNode) any {
	for i := 0; i < node.GetChildCount(); i++ {
		child := node.GetChild(i)
		if pt, ok := child.(antlr.ParseTree); ok {
			pt.Accept(v.Super)
		}
	}
	return nil
}

var _ parser.ChatVisitor = &SuperChatVisitor{}

func (v SuperChatVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(v)
}

func (v SuperChatVisitor) VisitChildren(node antlr.RuleNode) any {
	return v.SuperVisitChildren(node)
}

func (v SuperChatVisitor) VisitTerminal(node antlr.TerminalNode) any {
	return nil
}

func (v SuperChatVisitor) VisitErrorNode(node antlr.ErrorNode) any {
	return nil
}

func (v SuperChatVisitor) VisitChat(ctx *parser.ChatContext) any         { return v.VisitChildren(ctx) }
func (v SuperChatVisitor) VisitLine(ctx *parser.LineContext) any         { return v.VisitChildren(ctx) }
func (v SuperChatVisitor) VisitName(ctx *parser.NameContext) any         { return v.VisitChildren(ctx) }
func (v SuperChatVisitor) VisitCommand(ctx *parser.CommandContext) any   { return v.VisitChildren(ctx) }
func (v SuperChatVisitor) VisitMessage(ctx *parser.MessageContext) any   { return v.VisitChildren(ctx) }
func (v SuperChatVisitor) VisitEmoticon(ctx *parser.EmoticonContext) any { return v.VisitChildren(ctx) }
func (v SuperChatVisitor) VisitLink(ctx *parser.LinkContext) any         { return v.VisitChildren(ctx) }
func (v SuperChatVisitor) VisitColor(ctx *parser.ColorContext) any       { return v.VisitChildren(ctx) }
func (v SuperChatVisitor) VisitMention(ctx *parser.MentionContext) any   { return v.VisitChildren(ctx) }

//

type MyChatVisitor struct {
	SuperChatVisitor
}

func (v *MyChatVisitor) VisitName(ctx *parser.NameContext) any {
	fmt.Printf("got name: %v\n", ctx.GetText())
	return v.SuperChatVisitor.VisitName(ctx)
}

//

func TestChat(t *testing.T) {
	testChat := `
barf says: hi // comment0
// comment1
xarf says: /* comment2 */ xi
`

	is := antlr.NewInputStream(testChat)
	lexer := parser.NewChatLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewChatParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Chat()
	toks := stream.GetAllTokens()

	ta := antlru.NewTokenAnalysis(tree, stream.GetAllTokens())

	for _, c := range tree.GetChildren() {
		if lc, ok := c.(*parser.LineContext); ok {
			fmt.Println(lc.GetText())
			l, r := lc.GetStart().GetTokenIndex(), lc.GetStop().GetTokenIndex()
			for i := l; i < r; i++ {
				fmt.Println(toks[i])
			}
			tok := antlru.GetTreeToken(lc)
			_ = ta.TreeIdxsByStart()[tok.GetStart()]
		}
	}

	v := MyChatVisitor{}
	v.Super = &v
	tree.Accept(&v)
}
