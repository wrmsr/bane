/*
Copyright (c) 2013, Jason Moiron

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
// Package reflectx implements extensions to the standard reflect lib suitable for implementing marshalling and
// unmarshalling packages.  The main Mapper type allows for Go-compatible named attribute access, including accessing
// embedded struct attributes and the ability to use  functions and struct tags to customize field names.
package structs

import (
	"reflect"
)

type FieldMap struct {
	Index    []int
	Path     string
	Field    reflect.StructField
	Zero     reflect.Value
	Name     string
	Embedded bool
	Children []*FieldMap
	Parent   *FieldMap
}

type StructMap struct {
	Tree  *FieldMap
	Index []*FieldMap
	Paths map[string]*FieldMap
	Names map[string]*FieldMap
}

func (f StructMap) GetByPath(path string) *FieldMap {
	return f.Paths[path]
}

func (f StructMap) GetByTraversal(index []int) *FieldMap {
	if len(index) == 0 {
		return nil
	}

	tree := f.Tree
	for _, i := range index {
		if i >= len(tree.Children) || tree.Children[i] == nil {
			return nil
		}
		tree = tree.Children[i]
	}
	return tree
}

func Deref(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

// -- helpers & utilities --

type typeQueue struct {
	t  reflect.Type
	fi *FieldMap
	pp string // Parent path
}

// A copying append that creates a new slice each time.
func apnd(is []int, i int) []int {
	x := make([]int, len(is)+1)
	copy(x, is)
	x[len(x)-1] = i
	return x
}

// getMapping returns a mapping for the t type, using the tagName, mapFunc and tagMapFunc to determine the canonical
// names of fields.
func getMapping(t reflect.Type) *StructMap {
	m := []*FieldMap{}

	root := &FieldMap{}
	queue := []typeQueue{}
	queue = append(queue, typeQueue{Deref(t), root, ""})

QueueLoop:
	for len(queue) != 0 {
		// pop the first item off of the queue
		tq := queue[0]
		queue = queue[1:]

		// ignore recursive field
		for p := tq.fi.Parent; p != nil; p = p.Parent {
			if tq.fi.Field.Type == p.Field.Type {
				continue QueueLoop
			}
		}

		nChildren := 0
		if tq.t.Kind() == reflect.Struct {
			nChildren = tq.t.NumField()
		}
		tq.fi.Children = make([]*FieldMap, nChildren)

		// iterate through all of its fields
		for fieldPos := 0; fieldPos < nChildren; fieldPos++ {
			f := tq.t.Field(fieldPos)

			fi := FieldMap{
				Field: f,
				Name:  f.Name,
				Zero:  reflect.New(f.Type).Elem(),
			}

			// if the path is empty this path is just the name
			if tq.pp == "" {
				fi.Path = fi.Name
			} else {
				fi.Path = tq.pp + "." + fi.Name
			}

			// skip unexported fields
			if len(f.PkgPath) != 0 && !f.Anonymous {
				continue
			}

			// bfs search of anonymous embedded structs
			if f.Anonymous {
				fi.Embedded = true
				fi.Index = apnd(tq.fi.Index, fieldPos)
				nc := 0
				ft := Deref(f.Type)
				if ft.Kind() == reflect.Struct {
					nc = ft.NumField()
				}
				fi.Children = make([]*FieldMap, nc)
				queue = append(queue, typeQueue{Deref(f.Type), &fi, tq.pp})

			} else if fi.Zero.Kind() == reflect.Struct ||
				(fi.Zero.Kind() == reflect.Ptr && fi.Zero.Type().Elem().Kind() == reflect.Struct) {
				fi.Index = apnd(tq.fi.Index, fieldPos)
				fi.Children = make([]*FieldMap, Deref(f.Type).NumField())
				queue = append(queue, typeQueue{Deref(f.Type), &fi, fi.Path})

			}

			fi.Index = apnd(tq.fi.Index, fieldPos)
			fi.Parent = tq.fi
			tq.fi.Children[fieldPos] = &fi
			m = append(m, &fi)
		}
	}

	flds := &StructMap{
		Index: m,
		Tree:  root,
		Paths: map[string]*FieldMap{},
		Names: map[string]*FieldMap{},
	}
	for _, fi := range flds.Index {
		fld, ok := flds.Paths[fi.Path]
		if !ok || fld.Embedded {
			flds.Paths[fi.Path] = fi
			if fi.Name != "" && !fi.Embedded {
				flds.Names[fi.Path] = fi
			}
		}
	}

	return flds
}
