package testing

import "testing"

func AssertEqual[T comparable](t *testing.T, ex T, ac T) {
	if ex != ac {
		t.Fatalf("%v != %v", ex, ac)
	}
}

func AssertNoErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("%v", err)
	}
}
