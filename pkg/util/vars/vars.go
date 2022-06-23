package vars

type Var[T any] interface {
	Get() T
	Set(v T)
}
