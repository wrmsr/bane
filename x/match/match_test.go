package match

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMatch(t *testing.T) {
	c := NewCapture()
	var p Pattern
	p = TypePattern{rtype: reflect.TypeOf(4)}
	p = CapturePattern{pattern: pattern{previous: p}, capture: c}
	m := Run(p, 5, nil)
	fmt.Printf("%#v\n", m)
}

func TestPropMatch(t *testing.T) {
	strPat := TypePattern{rtype: reflect.TypeOf("")}
	lenProp := Property{name: "len", fn: func(v any) any { return len(v.(string)) }}
	len1 := PropertyPattern{prop: lenProp, pat: EqualsPattern{value: 1}}
	withPat := WithPattern{pattern: pattern{previous: strPat}, propPat: len1}
	m := Run(withPat, "a", nil)
	fmt.Printf("%#v\n", m)
	m = Run(withPat, "aa", nil)
	fmt.Printf("%#v\n", m)
}
