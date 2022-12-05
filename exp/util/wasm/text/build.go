package text

import "fmt"

//

func L(os ...any) List {
	var l List
	for _, o := range os {
		l.Append(Of(o))
	}
	return l
}

func (e *List) Append(os ...any) *List {
	for _, o := range os {
		if c := Of(o); c != nil {
			e.Ps = append(e.Ps, Of(o))
		}
	}
	return e
}

func Q(o any) Quote {
	switch o := o.(type) {

	case Quote:
		return o

	case string:
		return Quote{S: o}

	default:
		return Quote{S: fmt.Sprintf("%v", o)}

	}
}

func Of(o any) Element {
	if o == nil {
		return nil
	}

	if n, ok := o.(Element); ok {
		return n
	}

	switch o := o.(type) {

	case string:
		return Atom{S: o}

	case []Element:
		return List{Ps: o}

	case []any:
		ps := make([]Element, len(o))
		for i, c := range o {
			ps[i] = Of(c)
		}
		return List{Ps: ps}

	}
	panic(o)
}
