package lisp

import "strconv"

//

type Number interface {
	Value
}

//

type Int int64

var _ Value = Int(0)

func (v Int) isValue() {}

func (v Int) String() string { return strconv.Itoa(int(v)) }
