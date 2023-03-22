package match

//

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
