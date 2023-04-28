package annotate

import "reflect"

//

type typeIndexedCollection struct {
	s []any
	m map[reflect.Type][]any
}

func (c *typeIndexedCollection) add(vs ...any) {
	if len(vs) < 1 {
		return
	}
	if c.m == nil {
		c.m = make(map[reflect.Type][]any)
	}
	c.s = append(c.s, vs...)
	for _, a := range vs {
		ty := reflect.TypeOf(a)
		c.m[ty] = append(c.m[ty], a)
	}
}

//

type keyedCollection[K comparable, V any] struct {
	s []V
	m map[K]V
}

func (c *keyedCollection[K, V]) getOrMake(k K, f func() V) V {
	if c.m == nil {
		c.m = make(map[K]V)
	}
	v, ok := c.m[k]
	if !ok {
		v = f()
		c.m[k] = v
		c.s = append(c.s, v)
	}
	return v
}
