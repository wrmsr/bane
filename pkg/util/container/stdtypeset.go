package container

import (
	"reflect"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type StdTypeSet struct {
	m map[reflect.Type]struct{}
}

func NewStdTypeSet(it bt.Iterable[reflect.Type]) StdTypeSet {
	s := StdTypeSet{}
	if it != nil {
		m := make(map[reflect.Type]struct{})
		for it := it.Iterate(); it.HasNext(); {
			s.m[it.Next()] = struct{}{}
		}
		s.m = m
	}
	return s
}

func NewStdTypeSetOf(vs ...reflect.Type) StdTypeSet {
	return NewStdTypeSet(its.Of(vs...))
}

var _ Set[reflect.Type] = StdTypeSet{}

func (s StdTypeSet) Len() int {
	if s.m == nil {
		return 0
	}
	return len(s.m)
}

func (s StdTypeSet) Contains(t reflect.Type) bool {
	if s.m == nil {
		return false
	}
	_, ok := s.m[t]
	return ok
}

func (s StdTypeSet) Iterate() bt.Iterator[reflect.Type] {
	if s.m == nil {
		return its.Empty[reflect.Type]().Iterate()
	}
	return its.Map(its.OfTypeMap(s.m), func(kv bt.Kv[reflect.Type, struct{}]) reflect.Type { return kv.K }).Iterate()
}

func (s StdTypeSet) ForEach(fn func(reflect.Type) bool) bool {
	if s.m == nil {
		return true
	}
	for v := range s.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

//

type MutStdTypeSet struct {
	s StdTypeSet
}

func NewMutStdTypeSet(it bt.Iterable[reflect.Type]) *MutStdTypeSet {
	return &MutStdTypeSet{s: NewStdTypeSet(it)}
}

func NewMutStdTypeSetOf(vs ...reflect.Type) *MutStdTypeSet {
	return NewMutStdTypeSet(its.Of(vs...))
}

func WrapTypeSet(s map[reflect.Type]struct{}) *MutStdTypeSet {
	return &MutStdTypeSet{StdTypeSet{s}}
}

var _ MutSet[reflect.Type] = &MutStdTypeSet{}

func (s *MutStdTypeSet) isMut() {}

func (s *MutStdTypeSet) Len() int                                  { return s.s.Len() }
func (s *MutStdTypeSet) Contains(v reflect.Type) bool              { return s.s.Contains(v) }
func (s *MutStdTypeSet) Iterate() bt.Iterator[reflect.Type]        { return s.s.Iterate() }
func (s *MutStdTypeSet) ForEach(fn func(v reflect.Type) bool) bool { return s.s.ForEach(fn) }

func (s *MutStdTypeSet) lazyInit() {
	if s.s.m == nil {
		s.s.m = make(map[reflect.Type]struct{})
	}
}

func (s *MutStdTypeSet) Add(value reflect.Type) {
	s.s.m[value] = struct{}{}
}

func (s *MutStdTypeSet) TryAdd(value reflect.Type) bool {
	if s.Contains(value) {
		return false
	}
	s.s.m[value] = struct{}{}
	return true
}

func (s *MutStdTypeSet) Remove(value reflect.Type) {
	delete(s.s.m, value)
}

func (s *MutStdTypeSet) TryRemove(value reflect.Type) bool {
	if s.Contains(value) {
		return false
	}
	delete(s.s.m, value)
	return true
}

var _ Decay[Set[reflect.Type]] = &MutStdTypeSet{}

func (s *MutStdTypeSet) Decay() Set[reflect.Type] { return s.s }
