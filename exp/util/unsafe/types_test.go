package unsafe

import (
	"fmt"
	"testing"
)

func TestTypeOf(t *testing.T) {
	typ := GetType("runtime.g")
	if typ == nil {
		t.Error("runtime.g not found")
	}

	typ = GetType("*runtime.g")
	if typ == nil {
		t.Error("*runtime.g not found")
	}

	typ = typ.Elem()
	n := typ.NumField()
	for i := 0; i < n; i++ {
		f := typ.Field(i)
		fmt.Printf("%d %s %s\n", i, f.Name, f.Type)
	}
}
