package lisp

type Value interface {
	String() string

	isValue()
}

type Cons struct {
	Car, Cdr Value
}
