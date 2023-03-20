package match

import "testing"

///

type Capture struct{}

type Captures struct {
	Capture *Capture
	Value   any
	Next    *Captures
}

///

type Match interface {
	Present() bool
	Value() any
	Captures() *Captures

	Filter(fn func(v any) bool) Match
	Map(fn func(v any) any) Match
	FlatMap(fn func(v any) any) Match
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

func (m PresentMatch) Filter(fn func(v any) bool) Match {
	panic("implement me")
}

func (m PresentMatch) Map(fn func(v any) any) Match {
	panic("implement me")
}

func (m PresentMatch) FlatMap(fn func(v any) any) Match {
	panic("implement me")
}

//

type EmptyMatch struct{}

var _ Match = EmptyMatch{}

func (e EmptyMatch) Present() bool       { return false }
func (e EmptyMatch) Value() any          { panic(e) }
func (e EmptyMatch) Captures() *Captures { panic(e) }

func (e EmptyMatch) Filter(fn func(v any) bool) Match {
	panic("implement me")
}

func (e EmptyMatch) Map(fn func(v any) any) Match {
	panic("implement me")
}

func (e EmptyMatch) FlatMap(fn func(v any) any) Match {
	panic("implement me")
}

///

func TestMatch(t *testing.T) {

}
