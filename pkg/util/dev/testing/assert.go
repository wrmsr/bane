//go:build !nodev

package testing

import (
	"reflect"
	"runtime"
	"testing"
)

func OnAssertFailure() {
	runtime.Gosched()
}

func AssertEqual[T comparable](t *testing.T, ex T, ac T) {
	if ex != ac {
		OnAssertFailure()
		t.Fatalf("%v != %v", ex, ac)
	}
}

func AssertDeepEqual(t *testing.T, ex any, ac any) {
	if !reflect.DeepEqual(ex, ac) {
		OnAssertFailure()
		t.Fatalf("%v != %v", ex, ac)
	}
}

func AssertNoErr(t *testing.T, err error) {
	if err != nil {
		OnAssertFailure()
		t.Fatalf("%v", err)
	}
}
