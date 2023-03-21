package match

import (
	"fmt"
	"reflect"
	"sync/atomic"
	"testing"
)

///

type Capture struct {
	seq int64
}

var captureSeq atomic.Int64

func NewCapture() *Capture {
	return &Capture{
		seq: captureSeq.Add(1),
	}
}

//

type Captures struct {
	capture *Capture
	value   any
	next    *Captures
}

func MakeCaptures(c *Capture, v any) *Captures {
	if v == nil {
		return nil
	}
	return &Captures{
		capture: c,
		value:   v,
	}
}

func (cs *Captures) AddAll(o *Captures) *Captures {
	if cs == nil {
		return o
	}
	return &Captures{
		capture: cs.capture,
		value:   cs.value,
		next:    cs.next.AddAll(o),
	}
}

///

type Match interface {
	Present() bool
	Value() any
	Captures() *Captures

	Filter(fn func(any) bool) Match
	Map(fn func(any) any) Match
}

//

type PresentMatch struct {
	value    any
	captures *Captures
}

var _ Match = PresentMatch{}

func (m PresentMatch) Present() bool       { return true }
func (m PresentMatch) Value() any          { return m.value }
func (m PresentMatch) Captures() *Captures { return m.captures }

func (m PresentMatch) Filter(fn func(any) bool) Match {
	if fn(m.value) {
		return m
	}
	return EmptyMatch{}
}

func (m PresentMatch) Map(fn func(any) any) Match {
	return PresentMatch{
		value:    fn(m.value),
		captures: m.captures,
	}
}

//

type EmptyMatch struct{}

var _ Match = EmptyMatch{}

func (m EmptyMatch) Present() bool       { return false }
func (m EmptyMatch) Value() any          { panic(m) }
func (m EmptyMatch) Captures() *Captures { panic(m) }

func (m EmptyMatch) Filter(fn func(any) bool) Match {
	return m
}

func (m EmptyMatch) Map(fn func(any) any) Match {
	return m
}

///

type Pattern interface {
	Previous() Pattern

	isPattern()
}

type pattern struct {
	previous Pattern
}

func (p pattern) isPattern() {}

func (p pattern) Previous() Pattern { return p.previous }

//

type CapturePattern struct {
	pattern
	capture *Capture
}

var _ Pattern = CapturePattern{}

//

type EqualsPattern struct {
	pattern
	value any
}

var _ Pattern = EqualsPattern{}

//

type FilterPattern struct {
	pattern
	fn func(any) bool
}

var _ Pattern = FilterPattern{}

//

type TypePattern struct {
	pattern
	rtype reflect.Type
}

var _ Pattern = TypePattern{}

///

type Property struct {
	name string
	fn   func(any) any
}

type PropertyPattern struct {
	prop Property
	pat  Pattern
}

type WithPattern struct {
	pattern
	propPat PropertyPattern
}

var _ Pattern = WithPattern{}

///

func DefaultMatch(p Pattern, v any, cs *Captures) Match {
	if p.Previous() != nil {
		m := DefaultMatch(p.Previous(), v, cs)
		if !m.Present() {
			return m
		}
		cs = m.Captures()
	}

	switch p := p.(type) {

	case CapturePattern:
		return PresentMatch{
			value:    v,
			captures: cs.AddAll(MakeCaptures(p.capture, v)),
		}

	case EqualsPattern:
		return PresentMatch{
			value:    v,
			captures: cs,
		}.Filter(func(a any) bool {
			return reflect.DeepEqual(a, p.value)
		})

	case FilterPattern:
		return PresentMatch{
			value:    v,
			captures: cs,
		}.Filter(p.fn)

	case TypePattern:
		return PresentMatch{
			value:    v,
			captures: cs,
		}.Filter(func(a any) bool {
			return reflect.TypeOf(a) == p.rtype
		})

	case WithPattern:
		pv := p.propPat.prop.fn(v)
		return DefaultMatch(p.propPat.pat, pv, cs)

	}
	panic(p)
}

///

func TestMatch(t *testing.T) {
	c := NewCapture()
	var p Pattern
	p = TypePattern{rtype: reflect.TypeOf(4)}
	p = CapturePattern{pattern: pattern{previous: p}, capture: c}
	m := DefaultMatch(p, 5, nil)
	fmt.Printf("%#v\n", m)
}

func TestPropMatch(t *testing.T) {
	strPat := TypePattern{rtype: reflect.TypeOf("")}
	lenProp := Property{name: "len", fn: func(v any) any { return len(v.(string)) }}
	len1 := PropertyPattern{prop: lenProp, pat: EqualsPattern{value: 1}}
	withPat := WithPattern{pattern: pattern{previous: strPat}, propPat: len1}
	m := DefaultMatch(withPat, "a", nil)
	fmt.Printf("%#v\n", m)
	m = DefaultMatch(withPat, "aa", nil)
	fmt.Printf("%#v\n", m)
}
