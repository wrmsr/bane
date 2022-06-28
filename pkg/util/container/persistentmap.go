package container

//

type PersistentMap[K, V any] interface {
	Map[K, V]
	Persistent

	With(k K, v V) PersistentMap[K, V]
	Without(k K) PersistentMap[K, V]
	Default(k K, v V) PersistentMap[K, V]
}
