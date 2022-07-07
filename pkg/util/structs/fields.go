package structs

import (
	"fmt"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type StructFields struct {
	root ctr.List[*FieldInfo]
	all  ctr.List[*FieldInfo]

	byName ctr.OrderedMap[string, ctr.List[*FieldInfo]]
	byPath ctr.OrderedMap[string, *FieldInfo]

	flat   ctr.List[*FieldInfo]
	byFlat ctr.OrderedMap[string, *FieldInfo]

	dupe   ctr.List[*FieldInfo]
	byDupe ctr.OrderedMap[string, ctr.List[*FieldInfo]]
}

func (s StructFields) Root() ctr.List[*FieldInfo] { return s.root }
func (s StructFields) All() ctr.List[*FieldInfo]  { return s.all }

func (s StructFields) ByName() ctr.OrderedMap[string, ctr.List[*FieldInfo]] { return s.byName }
func (s StructFields) ByPath() ctr.OrderedMap[string, *FieldInfo]           { return s.byPath }

func (s StructFields) Flat() ctr.List[*FieldInfo]                 { return s.flat }
func (s StructFields) ByFlat() ctr.OrderedMap[string, *FieldInfo] { return s.byFlat }

func (s StructFields) Dupe() ctr.List[*FieldInfo]                           { return s.dupe }
func (s StructFields) ByDupe() ctr.OrderedMap[string, ctr.List[*FieldInfo]] { return s.byDupe }

func buildStructFields(root []*FieldInfo) StructFields {
	var all []*FieldInfo
	var rec func(*FieldInfo)
	rec = func(fi *FieldInfo) {
		all = append(all, fi)
		for _, cfi := range fi.children {
			rec(cfi)
		}
	}
	for _, fi := range root {
		rec(fi)
	}

	byName := ctr.NewMutMap[string, []*FieldInfo](nil)
	byPath := ctr.NewMutMap[string, *FieldInfo](nil)
	for _, fi := range all {
		byName.Put(fi.name.s, append(byName.Get(fi.name.s), fi))
		if byPath.Contains(fi.path) {
			panic(fmt.Errorf("duplicate field path: %s", fi.path))
		}
		byPath.Put(fi.path, fi)
	}

	var flat []*FieldInfo
	byFlat := ctr.NewMutMap[string, *FieldInfo](nil)
	var dupe []*FieldInfo
	byDupe := make(map[string][]*FieldInfo)
	byName.ForEach(func(kv bt.Kv[string, []*FieldInfo]) bool {
		s := kv.V
		if slices.Any(s, (*FieldInfo).Anonymous) {
			if len(s) != 1 {
				panic(fmt.Errorf("anonymous field name collision: %s", s[0].Name()))
			}
			return true
		}

		w := []*FieldInfo{s[0]}
		wd := w[0].Depth()
		for _, c := range s[1:] {
			cd := c.Depth()
			if cd < wd {
				w = []*FieldInfo{c}
			} else if cd == wd {
				w = append(w, c)
			}
		}

		if len(w) > 1 {
			dupe = append(dupe, w...)
			byDupe[w[0].name.s] = w
		} else {
			f := w[0]
			flat = append(flat, f)
			byFlat.Put(f.name.s, f)
		}

		return true
	})

	return StructFields{
		root: ctr.NewList(its.OfSlice(root)),
		all:  ctr.NewList(its.OfSlice(all)),

		byName: byName,
		byPath: byPath,

		flat:   ctr.NewList(its.OfSlice(flat)),
		byFlat: byFlat,

		dupe:   dupe,
		byDupe: byDupe,
	}
}
