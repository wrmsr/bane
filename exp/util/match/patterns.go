package match

import "reflect"

//

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

//

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

