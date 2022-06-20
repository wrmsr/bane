package containers

//

type PerMap[K comparable, V any] interface {
	Map[K, V]
	Persistent

	With(k K, v V) PerMap[K, V]
	Without(k K) PerMap[K, V]
	Default(k K, v V) PerMap[K, V]
}
