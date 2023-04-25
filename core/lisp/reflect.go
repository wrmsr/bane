package lisp

import (
	"fmt"
	"reflect"
)

type Reflect struct {
	v reflect.Value
}

func (r Reflect) isValue() {}

var _ Value = Reflect{}

func (r Reflect) IsIdentity() bool {
	return false
}

func (r Reflect) String() string {
	return fmt.Sprintf("<<<%v>>>", r.v)
}
