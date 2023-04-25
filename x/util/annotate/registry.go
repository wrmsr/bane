package annotate

import (
	"reflect"
	"sync"
)

//

type Registry interface {
	Get(obj any) []any
	GetOf(obj any, ty reflect.Type) []any
}

//

type objRegistry struct {
	obj any
	s   []any
	m   map[reflect.Type][]any
}

type registry struct {
	mtx sync.Mutex
	m   map[any]*objRegistry
	//ps  []*registry
}

var _ Registry = &registry{}

func newRegistry() *registry {
	return &registry{
		m: make(map[any]*objRegistry),
		//ps: parents,
	}
}

func (r *registry) register(obj any, items ...any) *registry {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	or := r.m[obj]
	if or == nil {
		or = &objRegistry{
			obj: obj,
			m:   make(map[reflect.Type][]any),
		}
		r.m[obj] = or
	}

	so := len(or.s)
	s := make([]any, so+len(items))
	copy(s, or.s)
	copy(s[so:], items)
	or.s = s

	cm := make(map[reflect.Type]int)
	for _, e := range items {
		ty := reflect.TypeOf(e)
		cm[ty] = cm[ty] + 1
	}
	for ty, c := range cm {
		s0 := or.m[ty]
		s1 := make([]any, len(s0), len(s0)+c)
		copy(s1, s0)
		or.m[ty] = s1
	}
	for _, e := range items {
		ty := reflect.TypeOf(e)
		or.m[ty] = append(or.m[ty], e)
	}

	return r
}

func (r *registry) Get(obj any) (ret []any) {
	r.mtx.Lock()
	e := r.m[obj]
	if e != nil {
		ret = e.s
	}
	r.mtx.Unlock()
	return
}

func (r *registry) GetOf(obj any, ty reflect.Type) (ret []any) {
	r.mtx.Lock()
	e := r.m[obj]
	if e != nil {
		ret = e.m[ty]
	}
	r.mtx.Unlock()
	return
}
