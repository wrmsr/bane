// Copyright 2009 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package binary

import (
	"reflect"
	"sync"
)

// Size returns how many bytes Write would generate to encode the value v, which must be a fixed-size value or a slice
// of fixed-size values, or a pointer to such data. If v is neither of these, Size returns -1.
func Size(v any) int {
	return ValueSize(reflect.Indirect(reflect.ValueOf(v)))
}

var structSize sync.Map // map[reflect.Type]int

// ValueSize returns the number of bytes the actual data represented by v occupies in memory. For compound structures,
// it sums the sizes of the elements. Thus, for instance, for a slice it returns the length of the slice times the
// element size and does not count the memory occupied by the header. If the type of v is not acceptable, ValueSize
// returns -1.
func ValueSize(v reflect.Value) int {
	switch v.Kind() {
	case reflect.Slice:
		if s := SizeOf(v.Type().Elem()); s >= 0 {
			return s * v.Len()
		}
		return -1

	case reflect.Struct:
		t := v.Type()
		if size, ok := structSize.Load(t); ok {
			return size.(int)
		}
		size := SizeOf(t)
		structSize.Store(t, size)
		return size

	default:
		return SizeOf(v.Type())
	}
}

// SizeOf returns the size >= 0 of variables for the given type or -1 if the type is not acceptable.
func SizeOf(t reflect.Type) int {
	switch t.Kind() {
	case reflect.Array:
		if s := SizeOf(t.Elem()); s >= 0 {
			return s * t.Len()
		}

	case reflect.Struct:
		sum := 0
		for i, n := 0, t.NumField(); i < n; i++ {
			s := SizeOf(t.Field(i).Type)
			if s < 0 {
				return -1
			}
			sum += s
		}
		return sum

	case reflect.Bool,
		reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return int(t.Size())
	}

	return -1
}

// IntDataSize returns the size of the data required to represent the data when encoded. It returns zero if the type
// cannot be implemented by the fast path in Read or Write.
func IntDataSize(data any) int {
	switch data := data.(type) {
	case bool, int8, uint8, *bool, *int8, *uint8:
		return 1
	case []bool:
		return len(data)
	case []int8:
		return len(data)
	case []uint8:
		return len(data)
	case int16, uint16, *int16, *uint16:
		return 2
	case []int16:
		return 2 * len(data)
	case []uint16:
		return 2 * len(data)
	case int32, uint32, *int32, *uint32:
		return 4
	case []int32:
		return 4 * len(data)
	case []uint32:
		return 4 * len(data)
	case int64, uint64, *int64, *uint64:
		return 8
	case []int64:
		return 8 * len(data)
	case []uint64:
		return 8 * len(data)
	case float32, *float32:
		return 4
	case float64, *float64:
		return 8
	case []float32:
		return 4 * len(data)
	case []float64:
		return 8 * len(data)
	}
	return 0
}
