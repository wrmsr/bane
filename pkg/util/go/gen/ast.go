package gen

import "go/token"

func InfixOpFromToken(tok token.Token) InfixOp {
	switch tok {

	case token.ADD:
		return AddOp
	case token.SUB:
		return SubOp
	case token.MUL:
		return MulOp
	case token.QUO:
		return DivOp
	case token.REM:
		return ModOp

	case token.LSS:
		return LtOp
	case token.LEQ:
		return LteOp
	case token.EQL:
		return EqOp
	case token.NEQ:
		return NeOp
	case token.GTR:
		return GtOp
	case token.GEQ:
		return GteOp

	case token.LAND:
		return AndOp
	case token.LOR:
		return OrOp

	}
	panic(tok)
}

func UnaryOpFromToken(tok token.Token) UnaryOp {
	switch tok {

	case token.NOT:
		return NotOp

	case token.SUB:
		return NegOp

	}
	panic(tok)
}