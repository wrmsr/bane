package behave

import (
	"reflect"

	"github.com/wrmsr/bane/pkg/util/check"
	ctr "github.com/wrmsr/bane/pkg/util/container"
	"github.com/wrmsr/bane/pkg/util/maps"
)

//

type Registry[K, O any] interface {
	Get(key K) ctr.Set[O]
	Add(obj O, key K)
	Remove(obj O, key K)
	Clear(key K)
}

//

type SimpleRegistry[K, O comparable] struct {
	setsByKey   map[K]maps.Set[O]
	addCallback func(O, K)
}

var _ Registry[int, string] = &SimpleRegistry[int, string]{}

func NewSimpleRegistry[K, O comparable](addCallback func(O, K)) *SimpleRegistry[K, O] {
	return &SimpleRegistry[K, O]{
		setsByKey:   make(map[K]maps.Set[O]),
		addCallback: addCallback,
	}
}

func (r *SimpleRegistry[K, O]) Get(key K) ctr.Set[O] {
	check.NotNil(key)
	return ctr.WrapSet(r.setsByKey[key])
}

func (r *SimpleRegistry[K, O]) Add(obj O, key K) {
	maps.ComputeDefault(r.setsByKey, key, maps.MakeSet[O]).Add(obj)
	if r.addCallback != nil {
		r.addCallback(obj, key)
	}
}

func (r *SimpleRegistry[K, O]) Remove(obj O, key K) {
	if key != nil {
		if s, ok := r.setsByKey[key]; ok {
			s.Remove(obj)
		}
	} else {
		for _, s := range r.setsByKey {
			s.Remove(obj)
		}
	}
}

func (r *SimpleRegistry[K, O]) Clear(key K) {
	if key != nil {
		delete(r.setsByKey, key)
	} else {
		r.setsByKey = make(map[K]maps.Set[O])
	}
}

//

type TypeRegistry[O comparable] struct {
	setsByType  map[reflect.Type]maps.Set[O]
	addCallback func(O, reflect.Type)
}

var _ Registry[reflect.Type, string] = &TypeRegistry[string]{}

func NewTypeRegistry[O comparable](addCallback func(O, reflect.Type)) *TypeRegistry[O] {
	return &TypeRegistry[O]{
		setsByType:  make(map[reflect.Type]maps.Set[O]),
		addCallback: addCallback,
	}
}

func (r *TypeRegistry[O]) Get(key reflect.Type) maps.Set[O] {
	check.NotNil(key)
	return r.setsByType[key]
}

func (r *TypeRegistry[O]) Add(obj O, key reflect.Type) {
	maps.ComputeDefault(r.setsByType, key, maps.MakeSet[O]).Add(obj)
	if r.addCallback != nil {
		r.addCallback(obj, key)
	}
}

func (r *TypeRegistry[O]) Remove(obj O, key reflect.Type) {
	if key != nil {
		if s, ok := r.setsByType[key]; ok {
			s.Remove(obj)
		}
	} else {
		for _, s := range r.setsByType {
			s.Remove(obj)
		}
	}
}

func (r *TypeRegistry[O]) Clear(key reflect.Type) {
	if key != nil {
		delete(r.setsByType, key)
	} else {
		r.setsByType = make(map[reflect.Type]maps.Set[O])
	}
}
