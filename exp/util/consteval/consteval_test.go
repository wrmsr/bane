//
/*
https://github.com/golang/tools/tree/master/go/ssa/interp
*/
package consteval

//

type Value interface {
	isValue()
}

type value struct{}

func (v value) isValue() {}

//

type Primitive struct {
	value
	S string
}

//

type Type struct {
	value
	N string
}

//

type Struct struct {
	value
	N string
	M map[string]Value
}
