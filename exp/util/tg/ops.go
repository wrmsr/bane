package tg

import (
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
)

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

//

type RealizedOp struct {
	data *Buffer
	srcs []*Buffer
	ot   OpType
}

func realizeUnary(data *LazyBuffer) RealizedOp {
	panic("nyi")
}

func realizeBinary(data *LazyBuffer) RealizedOp {
	obs := data.op.GetBuffers()

	realSrcs := make(map[*LazyBuffer]*Buffer)
	rds := make([]*Buffer, len(obs))
	for i, x := range obs {
		rd := x.Realize()
		rds[i] = rd
		realSrcs[x] = rd
	}

	var astEval func(Lazy) *Buffer
	astEval = func(x Lazy) *Buffer {
		if x, ok := x.(*LazyBuffer); ok {
			return realSrcs[x]
		}
		lo := x.(*LazyOp)
		switch lo.op.Type() {
		case UnaryOpType:
			return astEval(lo.srcs[0]).UnaryOp(lo.op)
		case BinaryOpType:
			return astEval(lo.srcs[0]).BinaryOp(lo.op, astEval(lo.srcs[1]))
		}
		panic("unhandled")
	}

	ret := astEval(data.op)

	return RealizedOp{
		data: ret,
		srcs: rds,
		ot:   BinaryOpType,
	}
}

func realizeMovement(data *LazyBuffer) RealizedOp {
	realSrc := data.op.GetBuffers()[0].Realize()
	ops := slices.Reversed(data.op.GetOps())
	x := realSrc
	for i, o := range ops {
		x = x.MovementOp(o.op, o.arg)
		_ = i
	}
	return RealizedOp{
		data: x,
		srcs: []*Buffer{realSrc},
		ot:   MovementOpType,
	}
}

func realizeReduce(data *LazyBuffer) RealizedOp {
	realSrc := data.op.srcs[0].(*LazyBuffer).Realize()
	return RealizedOp{
		data: realSrc.ReduceOp(data.op.op, data.op.arg.(Shape)),
		srcs: []*Buffer{realSrc},
		ot:   ReduceOpType,
	}
}

func realizeProcessing(data *LazyBuffer) RealizedOp {
	rsx := data.op.srcs[0].(*LazyBuffer).Realize()
	rsw := data.op.srcs[1].(*LazyBuffer).Realize()
	ret := rsx.ProcessingOp(data.op.op, rsw, data.op.arg)
	return RealizedOp{
		data: ret,
		srcs: []*Buffer{rsx, rsw},
		ot:   ProcessingOpType,
	}
}

func realizeLoad(data *LazyBuffer) RealizedOp {
	check.Condition(data.op.op == FromCpuOp)
	return RealizedOp{
		data: data.op.arg.(*Buffer),
		ot:   LoadOpType,
	}
}

var realize = make(map[OpType]func(*LazyBuffer) RealizedOp)

func init() {
	realize[UnaryOpType] = realizeUnary
	realize[BinaryOpType] = realizeBinary
	realize[MovementOpType] = realizeMovement
	realize[ReduceOpType] = realizeReduce
	realize[ProcessingOpType] = realizeProcessing
	realize[LoadOpType] = realizeLoad
}
