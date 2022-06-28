package container

//

type PerMap[K, V any] interface {
	Map[K, V]
	Persistent

	With(k K, v V) PerMap[K, V]
	Without(k K) PerMap[K, V]
	Default(k K, v V) PerMap[K, V]
}
