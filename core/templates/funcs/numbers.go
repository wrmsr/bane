package funcs

import "text/template"

var _ = Std.AddAll(template.FuncMap{
	"intRange": IntRange,
})

func IntRange(start, end int) []int {
	n := end - start
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = start + i
	}
	return r
}
