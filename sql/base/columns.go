package base

import (
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Column struct {
	Name   string
	DbType string
	Type   reflect.Type
}

//

type Columns struct {
	s  []Column
	m  map[string]*Column
	im map[string]int
}

func ColumnsOf(s ...Column) (*Columns, error) {
	m := make(map[string]*Column, len(s))
	im := make(map[string]int)
	for i := range s {
		c := &s[i]
		if _, ok := m[c.Name]; ok {
			return nil, DuplicateColumnError{c.Name}
		}
		m[c.Name] = c
		im[c.Name] = i
	}
	return &Columns{s: s, m: m, im: im}, nil
}

func (cs Columns) At(i int) Column {
	return cs.s[i]
}

func (cs Columns) Index(k string) (int, bool) {
	i, ok := cs.im[k]
	return i, ok
}

var _ ctr.Map[string, Column] = Columns{}

func (cs Columns) Len() int {
	return len(cs.s)
}

func (cs Columns) Contains(k string) bool {
	_, ok := cs.m[k]
	return ok
}

func (cs Columns) Get(k string) Column {
	if c, ok := cs.m[k]; ok {
		return *c
	}
	return Column{}
}

func (cs Columns) TryGet(k string) (Column, bool) {
	if c, ok := cs.m[k]; ok {
		return *c, true
	}
	return Column{}, false
}

func (cs Columns) Iterate() bt.Iterator[bt.Kv[string, Column]] {
	return its.MakeKvs(its.OfSlice(cs.s), func(c Column) string { return c.Name }).Iterate()
}

func (cs Columns) ForEach(fn func(kv bt.Kv[string, Column]) bool) bool {
	for _, c := range cs.s {
		if !fn(bt.KvOf(c.Name, c)) {
			return false
		}
	}
	return true
}
