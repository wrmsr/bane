/*
Copyright (c) 2009 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the
following conditions are met:

* Redistributions of source code must retain the above copyright notice, this list of conditions and the following
  disclaimer.
* Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following
  disclaimer in the documentation and/or other materials provided with the distribution.
* Neither the name of Google Inc. nor the names of its contributors may be used to endorse or promote products derived
  from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

// Package json implements encoding and decoding of JSON as defined in RFC 7159. The mapping between JSON and Go values
// is described in the documentation for the Marshal and Unmarshal functions.
//
// See "JSON and Go" for an introduction to this package: https://golang.org/doc/articles/json_and_go.html
package structs

import (
	"reflect"
	"sort"
)

//

type walkedField struct {
	field reflect.StructField

	name      string
	nameBytes []byte                 // []byte(name)
	equalFold func(s, t []byte) bool // bytes.EqualFold or equivalent

	index []int
	ty    reflect.Type
}

type walkedFields struct {
	list      []walkedField
	nameIndex map[string]int
}

//

type walkedFieldsByIndex []walkedField

func (x walkedFieldsByIndex) Len() int { return len(x) }

func (x walkedFieldsByIndex) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func (x walkedFieldsByIndex) Less(i, j int) bool {
	for k, xik := range x[i].index {
		if k >= len(x[j].index) {
			return false
		}
		if xik != x[j].index[k] {
			return xik < x[j].index[k]
		}
	}
	return len(x[i].index) < len(x[j].index)
}

//

// walkFields returns a list of fields that JSON should recognize for the given type. The algorithm is breadth-first
// search over the set of structs to include - the top struct and then any reachable anonymous structs.
func walkFields(t reflect.Type) walkedFields {
	// Anonymous fields to explore at the current level and the next.
	var current []walkedField
	next := []walkedField{{ty: t}}

	// Count of queued names for current level and the next.
	var count, nextCount map[reflect.Type]int

	// Types already visited at an earlier level.
	visited := map[reflect.Type]bool{}

	// Fields found.
	var fields []walkedField

	for len(next) > 0 {
		current, next = next, current[:0]
		count, nextCount = nextCount, map[reflect.Type]int{}

		for _, f := range current {
			if visited[f.ty] {
				continue
			}
			visited[f.ty] = true

			// Scan f.typ for fields to include.
			for i := 0; i < f.ty.NumField(); i++ {
				sf := f.ty.Field(i)
				if sf.Anonymous {
					t := sf.Type
					if t.Kind() == reflect.Pointer {
						t = t.Elem()
					}
					if !sf.IsExported() && t.Kind() != reflect.Struct {
						// Ignore embedded fields of unexported non-struct types.
						continue
					}
					// Do not ignore embedded fields of unexported struct types since they may have exported fields.
				} else if !sf.IsExported() {
					// Ignore unexported non-embedded fields.
					continue
				}
				tag := sf.Tag.Get("json")
				if tag == "-" {
					continue
				}
				index := make([]int, len(f.index)+1)
				copy(index, f.index)
				index[len(f.index)] = i

				ft := sf.Type
				if ft.Name() == "" && ft.Kind() == reflect.Pointer {
					// Follow pointer.
					ft = ft.Elem()
				}

				// Record found field and index sequence.
				if !sf.Anonymous || ft.Kind() != reflect.Struct {
					field := walkedField{
						field: sf,
						name:  sf.Name,
						index: index,
						ty:    ft,
					}
					field.nameBytes = []byte(field.name)
					field.equalFold = foldFunc(field.nameBytes)

					fields = append(fields, field)
					if count[f.ty] > 1 {
						// If there were multiple instances, add a second, so that the annihilation code will see a
						// duplicate. It only cares about the distinction between 1 or 2, so don't bother generating any
						// more copies.
						fields = append(fields, fields[len(fields)-1])
					}
					continue
				}

				// Record new anonymous struct to explore in next round.
				nextCount[ft]++
				if nextCount[ft] == 1 {
					next = append(next, walkedField{name: ft.Name(), index: index, ty: ft})
				}
			}
		}
	}

	sort.Slice(fields, func(i, j int) bool {
		x := fields
		// sort field by name, breaking ties with depth, then breaking ties with "name came from json tag", then
		//  breaking ties with index sequence.
		if x[i].name != x[j].name {
			return x[i].name < x[j].name
		}
		if len(x[i].index) != len(x[j].index) {
			return len(x[i].index) < len(x[j].index)
		}
		return walkedFieldsByIndex(x).Less(i, j)
	})

	// Delete all fields that are hidden by the Go rules for embedded fields, except that fields with JSON tags are
	// promoted.

	// The fields are sorted in primary order of name, secondary order of field index length. Loop over names; for each
	// name, delete hidden fields by choosing the one dominant field that survives.
	out := fields[:0]
	for advance, i := 0, 0; i < len(fields); i += advance {
		// One iteration per name.
		// Find the sequence of fields with the name of this first field.
		fi := fields[i]
		name := fi.name
		for advance = 1; i+advance < len(fields); advance++ {
			fj := fields[i+advance]
			if fj.name != name {
				break
			}
		}
		if advance == 1 { // Only one field with this name
			out = append(out, fi)
			continue
		}
		dominant, ok := dominantField(fields[i : i+advance])
		if ok {
			out = append(out, dominant)
		}
	}

	fields = out
	sort.Sort(walkedFieldsByIndex(fields))

	nameIndex := make(map[string]int, len(fields))
	for i, field := range fields {
		nameIndex[field.name] = i
	}
	return walkedFields{fields, nameIndex}
}

// dominantField looks through the fields, all of which are known to have the same name, to find the single field that
// dominates the others using Go's embedding rules, modified by the presence of JSON tags. If there are multiple
// top-level fields, the boolean will be false: This condition is an error in Go and we skip all the fields.
func dominantField(fields []walkedField) (walkedField, bool) {
	// The fields are sorted in increasing index-length order, then by presence of tag. That means that the first field
	// is the dominant one. We need only check for error cases: two fields at top level, either both tagged or neither
	// tagged.
	if len(fields) > 1 && len(fields[0].index) == len(fields[1].index) {
		return walkedField{}, false
	}
	return fields[0], true
}
