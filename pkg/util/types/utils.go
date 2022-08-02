package types

func Is[I any](v any) bool {
	_, ok := v.(I)
	return ok
}

func CanAssign[F, T any]() bool {
	var z F
	var a any
	a = z
	if _, ok := a.(T); ok {
		return true
	}
	return false
}

func As[F, T any](f F) T {
	var a any
	a = f
	return a.(T)
}

func Zero[T any]() T {
	var z T
	return z
}

func PtrTo[T any](v T) *T {
	return &v
}

func BoolNum[T Rational](b bool) T {
	if b {
		return 1
	}
	return 0
}

func Choose[T any](b bool, t T, f T) T {
	if b {
		return t
	}
	return f
}

func Coalesce[T comparable](vs ...T) T {
	var z T
	for _, v := range vs {
		if v != z {
			return v
		}
	}
	return z
}

func CoalesceR[T comparable](vs ...T) T {
	var z T
	for i := len(vs) - 1; i >= 0; i-- {
		v := vs[i]
		if v != z {
			return v
		}
	}
	return z
}

func Identity[T any]() func(T) T {
	return func(v T) T { return v }
}

func Const[T any](v T) func() T {
	return func() T { return v }
}

func CloneSlice[T any](s []T) []T {
	r := make([]T, len(s))
	copy(r, s)
	return r
}
