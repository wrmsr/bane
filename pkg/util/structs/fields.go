package structs

import (
	"fmt"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
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

func (sf StructFields) Root() ctr.List[*FieldInfo] { return sf.root }
func (sf StructFields) All() ctr.List[*FieldInfo]  { return sf.all }

func (sf StructFields) ByName() ctr.OrderedMap[string, ctr.List[*FieldInfo]] { return sf.byName }
func (sf StructFields) ByPath() ctr.OrderedMap[string, *FieldInfo]           { return sf.byPath }

func (sf StructFields) Flat() ctr.List[*FieldInfo]                 { return sf.flat }
func (sf StructFields) ByFlat() ctr.OrderedMap[string, *FieldInfo] { return sf.byFlat }

func (sf StructFields) Dupe() ctr.List[*FieldInfo]                           { return sf.dupe }
func (sf StructFields) ByDupe() ctr.OrderedMap[string, ctr.List[*FieldInfo]] { return sf.byDupe }

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

	byName := ctr.NewMutOrderedMap[string, ctr.MutList[*FieldInfo]](nil)
	byPath := ctr.NewMutOrderedMap[string, *FieldInfo](nil)
	for _, fi := range all {
		ctr.GetOrMake[string, ctr.MutList[*FieldInfo]](byName, fi.name.s, func() ctr.MutList[*FieldInfo] {
			return ctr.NewSliceMutList[*FieldInfo](nil)
		}).Append(fi)
		if byPath.Contains(fi.path) {
			panic(fmt.Errorf("duplicate field path: %s", fi.path))
		}
		byPath.Put(fi.path, fi)
	}

	var flat []*FieldInfo
	byFlat := ctr.NewMutOrderedMap[string, *FieldInfo](nil)
	var dupe []*FieldInfo
	byDupe := ctr.NewMutOrderedMap[string, ctr.List[*FieldInfo]](nil)
	byName.ForEach(func(kv bt.Kv[string, ctr.MutList[*FieldInfo]]) bool {
		s := kv.V
		if its.Any((*FieldInfo).Anonymous, s.(its.Iterable[*FieldInfo])) {
			if s.Len() != 1 {
				panic(fmt.Errorf("anonymous field name collision: %s", s.Get(0).Name()))
			}
			return true
		}

		var w []*FieldInfo
		var wd int
		s.ForEach(func(c *FieldInfo) bool {
			if w == nil {
				w = []*FieldInfo{c}
				wd = w[0].Depth()
			} else {
				cd := c.Depth()
				if cd < wd {
					w = []*FieldInfo{c}
				} else if cd == wd {
					w = append(w, c)
				}
			}
			return true
		})

		if len(w) > 1 {
			dupe = append(dupe, w...)
			byDupe.Put(w[0].name.s, ctr.NewSliceList(its.OfSlice(w)))
		} else {
			f := w[0]
			f.flat = true
			flat = append(flat, f)
			byFlat.Put(f.name.s, f)
		}

		return true
	})

	return StructFields{
		root: ctr.NewSliceList(its.OfSlice(root)),
		all:  ctr.NewSliceList(its.OfSlice(all)),

		byName: ctr.NewOrderedMap(its.MapValues[string, ctr.MutList[*FieldInfo], ctr.List[*FieldInfo]](byName, ctr.DecayList[*FieldInfo])),
		byPath: byPath,

		flat:   ctr.NewSliceList(its.OfSlice(flat)),
		byFlat: byFlat,

		dupe:   ctr.WrapSlice(dupe),
		byDupe: byDupe,
	}
}
