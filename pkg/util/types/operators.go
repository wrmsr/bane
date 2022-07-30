package types

import "golang.org/x/exp/constraints"

func Add[T Rational](l, r T) T {
	return l + r
}

func Sub[T Rational](l, r T) T {
	return l - r
}

func Mul[T Rational](l, r T) T {
	return l * r
}

func Div[T Rational](l, r T) T {
	return l / r
}

func Mod[T constraints.Integer](l, r T) T {
	return l % r
}
