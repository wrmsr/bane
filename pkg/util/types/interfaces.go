package types

type ErrFn = func() error

type Stringer interface {
	String() string
}
