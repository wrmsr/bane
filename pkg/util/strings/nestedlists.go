package strings

type NestedList = any // string | []NestedList

func ParseNestedList(s string, l, r, d rune) (NestedList, error) {
	rs := []rune(s)
	i := 0
	var rec func() (NestedList, error)
	rec = func() (NestedList, error) {
		b := i
		for i < len(rs) {
			c := rs[i]
			i++
			switch c {
			case l:
				c, err := rec()
				if err != nil {
					return nil, err
				}
				_ = c
				panic("nyi")
			case r:
				break
			case d:
				panic("nyi")
			}
		}
		_ = b
		panic("nyi")
	}
	return rec()
}
