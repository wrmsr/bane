package text

import (
	"fmt"

	wr "github.com/wrmsr/bane/x/wasm/rendering"
)

//

type Element interface {
	isElement()

	wr.Renderer
	fmt.Stringer
}

type element struct{}

func (p element) isElement() {}

//

type Atom struct {
	element
	S string
}

var _ Element = Atom{}

//

type Quote struct {
	element
	S string
}

var _ Element = Quote{}

//

type List struct {
	element
	Ps []Element
}

var _ Element = List{}
