package c

import "testing"

func TestParse(t *testing.T) {
	src := `void Sleep(int ms);`

	cp := CPState{
		i:    []rune(src),
		cts:  lj_ctype_init(),
		mode: CPARSE_MODE_MULTI | CPARSE_MODE_DIRECT,
	}

	cp_init(&cp)
	cp_decl_multi(&cp)
}
