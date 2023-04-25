/*
The MIT License (MIT)

# Copyright (c) 2019 The inspect Project

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
package unsafe

import (
	"reflect"
	"sync"
	"unsafe"
)

//

//go:linkname typelinks reflect.typelinks
func typelinks() (sections []unsafe.Pointer, offset [][]int32)

//go:linkname resolveTypeOff reflect.resolveTypeOff
func resolveTypeOff(rtype unsafe.Pointer, off int32) unsafe.Pointer

type iface struct {
	tab, data unsafe.Pointer
}

func readAllTypes() []reflect.Type {
	s := make([]reflect.Type, 0, 1024)
	ty := reflect.TypeOf(0)
	f := (*iface)(unsafe.Pointer(&ty))
	sections, offset := typelinks()
	for i, offs := range offset {
		rodata := sections[i]
		for _, off := range offs {
			f.data = resolveTypeOff(rodata, off)
			s = append(s, ty)
		}
	}
	return s
}

func makeTypeNameMap(tys []reflect.Type) map[string]reflect.Type {
	m := make(map[string]reflect.Type, len(tys))
	for _, ty := range tys {
		if ty.Kind() == reflect.Ptr && len(ty.Elem().Name()) > 0 {
			m[ty.String()] = ty
			m[ty.Elem().String()] = ty.Elem()
		}
	}
	return m
}

//

var (
	allTypesOnce   sync.Once
	allTypes       []reflect.Type
	allTypesByName map[string]reflect.Type
)

func AllTypes() []reflect.Type {
	allTypesOnce.Do(func() {
		allTypes = readAllTypes()
		allTypesByName = makeTypeNameMap(allTypes)
	})
	return allTypes
}

func AllTypesByName() map[string]reflect.Type {
	AllTypes()
	return allTypesByName
}

func GetType(name string) reflect.Type {
	if typ, ok := AllTypesByName()[name]; ok {
		return typ
	}
	return nil
}
