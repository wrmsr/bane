package lisp

import (
	"fmt"
	"strings"
)

type Value interface {
	String() string

	isValue()
}

type Cons struct {
	Car, Cdr Value
}

func (c Cons) String() string {
	var sb strings.Builder
	c.BuildString(&sb)
	return sb.String()
}

func (c Cons) BuildString(sb *strings.Builder) {
	sb.WriteString("(")
	sb.WriteString(fmt.Sprintf("%v . %v", c.Car, c.Cdr))
	sb.WriteString(")")
}
