package tg

//

type OpType uint8

const (
	InvalidOpType OpType = iota
	UnaryOpType
	BinaryOpType
	ReduceOpType
	MovementOpType
	ProcessingOpType
	LoadOpType
)

func (t OpType) String() string {
	switch t {
	case UnaryOpType:
		return "unary"
	case BinaryOpType:
		return "binary"
	case ReduceOpType:
		return "reduce"
	case MovementOpType:
		return "movement"
	case ProcessingOpType:
		return "processing"
	case LoadOpType:
		return "load"
	}
	panic(t)
}

//

type Op uint8

const (
	InvalidOp Op = iota

	_FirstUnaryOp
	NoOp
	NegOp
	ReluOp
	ExpOp
	LogOp
	SignOp

	_FirstBinaryOp
	AddOp
	SubOp
	MulOp
	DivOp
	PowOp
	CmpEqOp

	_FirstReduceOp
	SumOp
	MaxOp

	_FirstMovementOp
	ReshapeOp
	PermuteOp
	SliceOp
	ExpandOp
	FlipOp

	_FirstProcessingOp
	ConvOp

	_FirstLoadOp
	FromCpuOp

	_LastOp
)

func (o Op) String() string {
	switch o {

	case NoOp:
		return "nop"
	case NegOp:
		return "neg"
	case ReluOp:
		return "relu"
	case ExpOp:
		return "exp"
	case LogOp:
		return "log"
	case SignOp:
		return "sign"

	case AddOp:
		return "add"
	case SubOp:
		return "sub"
	case MulOp:
		return "mul"
	case DivOp:
		return "div"
	case PowOp:
		return "pow"
	case CmpEqOp:
		return "cmp_eq"

	case SumOp:
		return "sum"
	case MaxOp:
		return "max"

	case ReshapeOp:
		return "reshape"
	case PermuteOp:
		return "permute"
	case SliceOp:
		return "slice"
	case ExpandOp:
		return "expand"
	case FlipOp:
		return "flip"

	case ConvOp:
		return "conv"

	case FromCpuOp:
		return "from_cpu"

	}
	panic(o)
}

func (o Op) Type() OpType {
	switch {
	case o < _FirstUnaryOp:
		break
	case o < _FirstBinaryOp:
		return UnaryOpType
	case o < _FirstReduceOp:
		return BinaryOpType
	case o < _FirstMovementOp:
		return ReduceOpType
	case o < _FirstProcessingOp:
		return MovementOpType
	case o < _FirstLoadOp:
		return ProcessingOpType
	case o < _LastOp:
		return LoadOpType
	}
	panic(o)
}
