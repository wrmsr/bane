package types

//

type Cmp int8

const (
	CmpLess    = -1
	CmpEqual   = 0
	CmpGreater = 1
)

//

type Comparer[T any] interface {
	Compare(o T) Cmp
}

//

type CmpImpl[T any] func(o T) Cmp
