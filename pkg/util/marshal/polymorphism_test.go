package marshal

import (
	"testing"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

type PolyI interface{ isPolyI() }

type PolyA struct{ I int }
type PolyB struct{ S string }
type PolyC struct{ L, R PolyI }

func (p PolyA) isPolyI() {}
func (p PolyB) isPolyI() {}
func (p PolyC) isPolyI() {}

func TestPolymorphism(t *testing.T) {
	i := PolyC{
		L: PolyA{
			I: 420,
		},
		R: PolyC{
			L: PolyA{
				I: 421,
			},
			R: PolyB{
				S: "four twenty three",
			},
		},
	}

	p := NewPolymorphism(rfl.TypeOf[PolyI](),
		PolymorphismEntry{Tag: "a", Impl: rfl.TypeOf[PolyA]()},
		PolymorphismEntry{Tag: "b", Impl: rfl.TypeOf[PolyB]()},
		PolymorphismEntry{Tag: "c", Impl: rfl.TypeOf[PolyC]()},
	)
}
