package slices

import "fmt"

//

type SliceLengthError struct {
	Actual, Expected int
}

func (e SliceLengthError) Error() string {
	return fmt.Sprintf("slice length error: actual=%d expected=%d", e.Actual, e.Expected)
}

func checkLength[T any](s []T, e int) []T {
	if len(s) != e {
		panic(SliceLengthError{Actual: len(s), Expected: e})
	}
	return s
}

//

func Unpack1[T any](s []T) T {
	checkLength(s, 1)
	return s[0]
}

func Unpack2[T any](s []T) (T, T) {
	checkLength(s, 2)
	return s[0], s[1]
}

func Unpack3[T any](s []T) (T, T, T) {
	checkLength(s, 3)
	return s[0], s[1], s[2]
}

func Unpack4[T any](s []T) (T, T, T, T) {
	checkLength(s, 4)
	return s[0], s[1], s[2], s[3]
}

//

func PtrUnpack1[T any](s []T) *T {
	checkLength(s, 1)
	return &s[0]
}

func PtrUnpack2[T any](s []T) (*T, *T) {
	checkLength(s, 2)
	return &s[0], &s[1]
}

func PtrUnpack3[T any](s []T) (*T, *T, *T) {
	checkLength(s, 3)
	return &s[0], &s[1], &s[2]
}

func PtrUnpack4[T any](s []T) (*T, *T, *T, *T) {
	checkLength(s, 4)
	return &s[0], &s[1], &s[2], &s[3]
}

//

func Dup2[T any](v T) (T, T) {
	return v, v
}

func Dup3[T any](v T) (T, T, T) {
	return v, v, v
}

func Dup4[T any](v T) (T, T, T, T) {
	return v, v, v, v
}
