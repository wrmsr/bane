package types

type Kv[K comparable, V any] struct {
	K K
	V V
}
