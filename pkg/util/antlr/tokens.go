package antlr

/*
class TokenAnalysis:

    def __init__(self, root: antlr4.ParserRuleContext) -> None:
        super().__init__()

        self._root = check.isinstance(root, antlr4.ParserRuleContext)

    @properties.cached
    @property
    def toks(self) -> ocol.IndexedSeq[antlr4.Token]:
        return ocol.IndexedSeq(self._root.parser.getInputStream().tokens, identity=True)

    @properties.cached
    @property
    def ctxs(self) -> ocol.SortedMapping[int, ta.AbstractSet[antlr4.ParserRuleContext]]:
        def rec(ctx):
            if isinstance(ctx, antlr4.TerminalNode):
                tok = ctx.symbol  # noqa
            else:
                tok = ctx.start
            check.isinstance(tok, antlr4.Token)
            dct.setdefault(self.toks.idx(tok), ocol.IdentitySet()).add(ctx)

            if not isinstance(ctx, antlr4.TerminalNode):
                for cctx in ctx.children or []:
                    rec(cctx)

        dct = ocol.SkipListDict()
        rec(self._root)
        return dct

    @properties.cached
    @property
    def rctxs(self) -> ta.Mapping[antlr4.ParserRuleContext, int]:
        return ocol.IdentityKeyDict((c, i) for i, cs in self.ctxs.items() for c in cs)
*/

/*
class TokenAnalysis:

    def __init__(self, root: no.Node) -> None:
        super().__init__()

        self._root = check.isinstance(root, no.Node)

        self._basic = ana.basic(self._root)

        self._tok_ana = tokens.TokenAnalysis(self._root.meeta[antlr4.ParserRuleContext])

    @properties.cached
    @property
    def rnodes(self) -> ta.Mapping[no.Node, int]:
        return ocol.IdentityKeyDict(
            (n, self._tok_ana.rctxs[n.meta[antlr4.ParserRuleContext]])
            for n in self._basic.nodes
        )

    def get_node_toks(self, node: no.Node) -> ta.Sequence[antlr4.Token]:
        l = self.rnodes[node]
        try:
            r = next(self._tok_ana.ctxs.itemsfrom(l + 1))[0]
        except StopIteration:
            r = -1
        return self._tok_ana.toks[l:r]

    def get_node_comments(self, node: no.Node) -> ta.Sequence[str]:
        return [t.text for t in self.get_node_toks(node) if t.channel == 1]
*/
