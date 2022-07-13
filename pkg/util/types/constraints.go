package types

import "golang.org/x/exp/constraints"

type Rational interface {
	constraints.Integer | constraints.Float
}

type Numeric interface {
	Rational | constraints.Complex
}
