package structs

import "fmt"

type StructFields struct {
	root []*FieldInfo
	all  []*FieldInfo

	byName map[string][]*FieldInfo
	byPath map[string]*FieldInfo

	flat   []*FieldInfo
	byFlat map[string]*FieldInfo

	dupe   []*FieldInfo
	byDupe map[string][]*FieldInfo
}

func (s StructFields) Root() []*FieldInfo { return s.root }
func (s StructFields) All() []*FieldInfo  { return s.all }

func (s StructFields) ByName() map[string][]*FieldInfo { return s.byName }
func (s StructFields) ByPath() map[string]*FieldInfo   { return s.byPath }

func (s StructFields) Flat() []*FieldInfo            { return s.flat }
func (s StructFields) ByFlat() map[string]*FieldInfo { return s.byFlat }

func (s StructFields) Dupe() []*FieldInfo              { return s.dupe }
func (s StructFields) ByDupe() map[string][]*FieldInfo { return s.byDupe }

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

	byName := make(map[string][]*FieldInfo)
	byPath := make(map[string]*FieldInfo)
	for _, fi := range all {
		byName[fi.name.s] = append(byName[fi.name.s], fi)
		if _, ok := byPath[fi.path]; ok {
			panic(fmt.Errorf("duplicate field path: %s", fi.path))
		}
		byPath[fi.path] = fi
	}

	var flat []*FieldInfo
	byFlat := make(map[string]*FieldInfo)
	var dupe []*FieldInfo
	byDupe := make(map[string][]*FieldInfo)
	for _, s := range byName {
		FIXME fi.anonymous!!
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
			byFlat[f.name.s] = f
		}
	}

	return StructFields{
		root: root,
		all:  all,

		byName: byName,
		byPath: byPath,

		flat:   flat,
		byFlat: byFlat,

		dupe:   dupe,
		byDupe: byDupe,
	}
}
