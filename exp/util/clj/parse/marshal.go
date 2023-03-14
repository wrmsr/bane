package parse

import msh "github.com/wrmsr/bane/pkg/util/marshal"

var _ = msh.RegisterTo[Node](
	msh.SetImplOf[BoolNode](),
	msh.SetImplOf[CharacterNode](),
	msh.SetImplOf[CommentNode](),
	msh.SetImplOf[DerefNode](),
	msh.SetImplOf[FnLiteralNode](),
	msh.SetImplOf[KeywordNode](),
	msh.SetImplOf[ListNode](),
	msh.SetImplOf[MapNode](),
	msh.SetImplOf[MetadataNode](),
	msh.SetImplOf[NewlineNode](),
	msh.SetImplOf[NilNode](),
	msh.SetImplOf[NumberNode](),
	msh.SetImplOf[QuoteNode](),
	msh.SetImplOf[ReaderCondNode](),
	msh.SetImplOf[ReaderCondSpliceNode](),
	msh.SetImplOf[ReaderDiscardNode](),
	msh.SetImplOf[ReaderEvalNode](),
	msh.SetImplOf[RegexNode](),
	msh.SetImplOf[SetNode](),
	msh.SetImplOf[StringNode](),
	msh.SetImplOf[SymbolNode](),
	msh.SetImplOf[SyntaxQuoteNode](),
	msh.SetImplOf[TagNode](),
	msh.SetImplOf[UnquoteNode](),
	msh.SetImplOf[UnquoteSpliceNode](),
	msh.SetImplOf[VarQuoteNode](),
	msh.SetImplOf[VectorNode](),
)
