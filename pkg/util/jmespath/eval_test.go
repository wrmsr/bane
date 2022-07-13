package jmespath

import "testing"

func TestEval(t *testing.T) {
	type o = map[string]any
	type a = []any

	d := o{
		"people": a{
			o{
				"age":   20,
				"other": "foo",
				"name":  "Bob",
			},
			o{
				"age":   25,
				"other": "bar",
				"name":  "Fred",
			},
			o{
				"age":   30,
				"other": "baz",
				"name":  "George",
			},
		},
	}

	e := a{
		a{
			"Fred",
			25,
		},
		a{
			"George",
			30,
		},
	}

	_ = d
	_ = e
}
