/*
TODO:
  - github.com/mitchellh/reflectwalk
*/
package reflect

import (
	"reflect"
	"testing"
)

//

type Cursor struct {
	v reflect.Value
	k any
	p *Cursor
}

func Walk(v reflect.Value) {

}

//

type foo struct {
	i int
}

type bar struct {
	f  foo
	ss []string
}

type baz struct {
	bar
	fs []foo
	bm map[string]*bar
	b  bool
}

//

func TestWalk(t *testing.T) {

}
