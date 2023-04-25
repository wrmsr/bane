package types

type Prim int8

const (
	None Prim = iota

	I32
	I64
	F32
	F64

	V128
)

func (p Prim) String() string {
	switch p {
	case None:
		return "none"
	case I32:
		return "i32"
	case I64:
		return "i64"
	case F32:
		return "f32"
	case F64:
		return "f64"
	case V128:
		return "v128"
	}
	panic(p)
}

func ParsePrim(s string) Prim {
	switch s {
	case "none":
		return None
	case "i32":
		return I32
	case "i64":
		return I64
	case "f32":
		return F32
	case "f64":
		return F64
	case "v128":
		return V128
	}
	panic(s)
}
