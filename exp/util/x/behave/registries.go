package behave

import (
	"reflect"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/maps"
)

//

type Registry[K, O any] interface {
	Get(key K) map[uintptr]O
	Add(id uintptr, obj O, key K)
	Remove(id uintptr, key K)
	RemoveAll(id uintptr)
	Clear(key K)
	ClearAll()
}

//

type SimpleRegistry[K comparable, O any] struct {
	setsByKey   map[K]map[uintptr]O
	addCallback func(O, K)
}

var _ Registry[int, string] = &SimpleRegistry[int, string]{}

func NewSimpleRegistry[K, O comparable](addCallback func(O, K)) *SimpleRegistry[K, O] {
	return &SimpleRegistry[K, O]{
		setsByKey:   make(map[K]map[uintptr]O),
		addCallback: addCallback,
	}
}

func (r *SimpleRegistry[K, O]) Get(key K) map[uintptr]O {
	check.NotNil(key)
	return r.setsByKey[key]
}

func (r *SimpleRegistry[K, O]) Add(id uintptr, obj O, key K) {
	maps.ComputeDefault(r.setsByKey, key, maps.Make[uintptr, O])[id] = obj
	if r.addCallback != nil {
		r.addCallback(obj, key)
	}
}

func (r *SimpleRegistry[K, O]) Remove(id uintptr, key K) {
	if s, ok := r.setsByKey[key]; ok {
		delete(s, id)
	}
}

func (r *SimpleRegistry[K, O]) RemoveAll(id uintptr) {
	for _, s := range r.setsByKey {
		delete(s, id)
	}
}

func (r *SimpleRegistry[K, O]) Clear(key K) {
	delete(r.setsByKey, key)
}

func (r *SimpleRegistry[K, O]) ClearAll() {
	r.setsByKey = make(map[K]map[uintptr]O)
}

//

type TypeRegistry[O any] struct {
	setsByType  map[reflect.Type]map[uintptr]O
	addCallback func(O, reflect.Type)
}

var _ Registry[reflect.Type, string] = &TypeRegistry[string]{}

func NewTypeRegistry[O any](addCallback func(O, reflect.Type)) *TypeRegistry[O] {
	return &TypeRegistry[O]{
		setsByType:  make(map[reflect.Type]map[uintptr]O),
		addCallback: addCallback,
	}
}

func (r *TypeRegistry[O]) Get(key reflect.Type) map[uintptr]O {
	check.NotNil(key)
	return r.setsByType[key]
}

func (r *TypeRegistry[O]) Add(id uintptr, obj O, key reflect.Type) {
	var s map[uintptr]O
	var ok bool
	if _, ok = r.setsByType[key]; ok {
		s[id] = obj
	} else {
		r.setsByType[key] = map[uintptr]O{
			id: obj,
		}
	}
	if r.addCallback != nil {
		r.addCallback(obj, key)
	}
}

func (r *TypeRegistry[O]) Remove(id uintptr, key reflect.Type) {
	if s, ok := r.setsByType[key]; ok {
		delete(s, id)
	}
}

func (r *TypeRegistry[O]) RemoveAll(id uintptr) {
	for _, s := range r.setsByType {
		delete(s, id)
	}
}

func (r *TypeRegistry[O]) Clear(key reflect.Type) {
	delete(r.setsByType, key)
}

func (r *TypeRegistry[O]) ClearAll() {
	r.setsByType = make(map[reflect.Type]map[uintptr]O)
}
