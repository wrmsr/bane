//go:build !nodev

package impl

import (
	"fmt"
	"reflect"

	"golang.org/x/tools/go/ast/astutil"
)

func (fg *FileGen) inlineFuncs() {
	im := make(map[string]*FuncDecl)
	for id := range fg.ps.Inlines() {
		n := nameFuncRef(id)
		decl := fg.rsv.FuncDecl(n)
		if decl == nil {
			panic(n)
		}
		im[n] = decl
	}

	for wid := range fg.ps.WithInlines() {
		n := nameFuncRef(wid)
		decl := fg.rsv.FuncDecl(n)
		if decl == nil {
			panic(n)
		}

		//nn := "_def_inl_" + n

		astutil.Apply(decl.Decl.Body, nil, func(cursor *astutil.Cursor) bool {
			fmt.Println(reflect.TypeOf(cursor.Node()))
			return true
		})
	}
}
