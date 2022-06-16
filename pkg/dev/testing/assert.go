//go:build !nodev

package testing

import (
	"reflect"
	"testing"
)

func AssertEqual[T comparable](t *testing.T, ex T, ac T) {
	if ex != ac {
		t.Fatalf("%v != %v", ex, ac)
	}
}

func AssertDeepEqual(t *testing.T, ex any, ac any) {
	if !reflect.DeepEqual(ex, ac) {
		t.Fatalf("%v != %v", ex, ac)
	}
}

func AssertNoErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("%v", err)
	}
}
