package c

import "testing"

func TestParse(t *testing.T) {
	src := `void Sleep(int ms);`

	cp := CPState{
		i: []rune(src),
	}

	cp_init(&cp)
	cp_decl_multi(&cp)
}
