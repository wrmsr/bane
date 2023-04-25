package instrs

//

type Class int8

const (
	Control Class = iota + 1
	Memory
	Numeric
	Parametric
	Reference
	Table
	Variable
	Vector
)

func (c Class) String() string {
	switch c {
	case 0:
		return "?"
	case Control:
		return "control"
	case Memory:
		return "memory"
	case Numeric:
		return "numeric"
	case Parametric:
		return "parametric"
	case Reference:
		return "reference"
	case Table:
		return "table"
	case Variable:
		return "variable"
	case Vector:
		return "vector"
	}
	panic(c)
}

//

type Access int8

const (
	Load Access = iota + 1
	Store
)

func (a Access) String() string {
	switch a {
	case 0:
		return "-"
	case Load:
		return "load"
	case Store:
		return "store"
	}
	panic(a)
}
