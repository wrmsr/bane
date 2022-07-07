package container

import "fmt"

func DecayList[T any](l MutList[T]) List[T] {
	switch l := l.(type) {
	case *mutListImpl[T]:
		return l.Decay()
	}
	panic(fmt.Errorf("unhandled type: %T", l))
}

func DecaySet[T comparable](s MutSet[T]) Set[T] {
	switch s := s.(type) {
	case *mutSetImpl[T]:
		return s.Decay()
	}
	panic(fmt.Errorf("unhandled type: %T", s))
}

func DecayMap[K comparable, V any](m MutMap[K, V]) Map[K, V] {
	switch m := m.(type) {
	case *mutMapImpl[K, V]:
		return m.Decay()
	}
	panic(fmt.Errorf("unhandled type: %T", m))
}
