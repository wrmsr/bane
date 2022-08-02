package types

import "golang.org/x/exp/constraints"

func Eq[T Equality](l, r T) bool {
	return l == r
}

func Ne[T Equality](l, r T) bool {
	return l != r
}

func Lt[T constraints.Ordered](l, r T) bool {
	return l < r
}

func Le[T constraints.Ordered](l, r T) bool {
	return l <= r
}

func Gt[T constraints.Ordered](l, r T) bool {
	return l > r
}

func Ge[T constraints.Ordered](l, r T) bool {
	return l >= r
}

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
