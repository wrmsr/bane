package containers

//

type PerMap[K comparable, V any] interface {
	Map[K, V]

	Put(k K, v V) PerMap[K, V]
	Delete(k K) PerMap[K, V]
	Default(k K, v V) PerMap[K, V]
}
