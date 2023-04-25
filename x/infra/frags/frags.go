package frags

import "github.com/wrmsr/bane/core/slices"

//

type Frag interface {
	isFrag()
}

type frag struct{}

func (f frag) isFrag() {}

//

func of(o any) Frag {
	if o, ok := o.(Frag); ok {
		return o
	}
	if o, ok := o.(string); ok {
		return ShFrag{S: o}
	}
	panic(o)
}

func Of(s ...any) Frag {
	if len(s) < 1 {
		return nil
	}
	if len(s) < 2 {
		return of(s[0])
	}
	return SeqFrag{C: slices.Map(of, s)}
}

//

type ShFrag struct {
	S string
	frag
}

func Sh(o any) ShFrag {
	panic("nyi")
}

//

type ArgFrag struct {
	S string
	frag
}

func Arg(o any) ArgFrag {
	panic("nyi")
}

//

type SeqFrag struct {
	C []Frag
	frag
}

func Seq(s ...any) SeqFrag {
	panic("nyi")
}
