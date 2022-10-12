package impl

import (
	"fmt"
	"go/ast"
)

func (fg *FileGen) inlineFuncs() {
	//i := 0
	//for _, ss := range fg.ps.Structs() {
	//	if i > 0 {
	//		fg.initStmts.Append(gg.Blank{})
	//	}
	//	i++
	//	fg.genStruct(ss)
	//}

	for wi := range fg.ps.WithInlines() {
		switch wi := wi.(type) {
		case *ast.Ident:
			obj := fg.pkg.Types.Scope().Lookup(wi.Name)
			fmt.Println(obj)
		default:
			panic(wi)
		}
	}
}
