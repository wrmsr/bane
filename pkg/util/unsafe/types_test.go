package unsafe

import (
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
}
