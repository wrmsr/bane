package container

import bt "github.com/wrmsr/bane/pkg/util/types"

//

type PersistentMap[K, V any] interface {
	Map[K, V]
	Persistent[bt.Kv[K, V]]

	With(k K, v V) PersistentMap[K, V]
	Without(k K) PersistentMap[K, V]
	Default(k K, v V) PersistentMap[K, V]
}
