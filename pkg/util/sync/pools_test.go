package sync

import (
	"fmt"
	"testing"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

func TestPool(t *testing.T) {
	p := NewTrackingPool[*int](
		NewSyncPool(func() *int { return new(int) }),
		func(p *int) uintptr { return rfl.PointerOf(p) },
	)

	xs := make([]*int, 10)
	for i := range xs {
		xs[i] = p.Get()
		*xs[i] = 420 + i
	}
	for i := range xs {
		p.Put(xs[i])
	}

	for i := 0; i < len(xs)+3; i++ {
		x := p.Get()
		fmt.Printf("%d %x %d\n", i, x, *x)
	}
}
