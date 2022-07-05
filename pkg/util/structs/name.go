package structs

type Name struct {
	s string
	b []byte

	eqFold func(s, t []byte) bool
}

func (n Name) String() string       { return n.s }
func (n Name) Bytes() []byte        { return n.b }
func (n Name) EqFold(o []byte) bool { return n.eqFold(n.b, o) }

func buildName(s string) Name {
	b := []byte(s)
	return Name{
		s: s,
		b: b,

		eqFold: foldFunc(b),
	}
}
