//go:build !nodev

package impl

import (
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/types/typeutil"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/def"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
)

//

func reifyIdentStr(node ast.Node) string {
	if node, ok := node.(*ast.BasicLit); ok {
		if node.Kind == token.STRING {
			s := node.Value
			if s[0:1] != "\"" || s[len(s)-1:] != "\"" {
				panic(fmt.Errorf("illegal ident str: %s", s))
			}
			s = s[1 : len(s)-1]
			if strings.Index(s, "/") >= 0 {
				panic(fmt.Errorf("illegal ident str: %s", s))
			}
			return s
		}
	}
	panic(fmt.Errorf("illegal ident str: %v", node))
}

func ReifyDef(node ast.Node, ti *types.Info) any {
	call := node.(*ast.CallExpr)
	fn, _ := typeutil.Callee(ti, call).(*types.Func)
	check.NotNil(fn)

	pn := rtu.ParseName(fn.FullName())
	switch pn.Obj {

	case "Struct":
		opts := make([]def.StructOpt, len(call.Args)-1)
		for i, arg := range call.Args[1:] {
			opts[i] = ReifyDef(arg, ti).(def.StructOpt)
		}
		return def.StructDef{
			Name: reifyIdentStr(call.Args[0]),
			Opts: opts,
		}

	case "Field":
		opts := make([]def.FieldOpt, len(call.Args)-1)

		hasTyOpt := false
		var dflTy opt.Optional[TypeRef]
		for i, arg := range call.Args[1:] {
			o := ReifyDef(arg, ti).(def.FieldOpt)
			if _, ok := o.(def.TypeOpt); ok {
				hasTyOpt = true
			}
			if o, ok := o.(def.DefaultOpt); ok {
				dflTy = opt.Just(o.Ty.(TypeRef))
			}
			opts[i] = o
		}

		if !hasTyOpt {
			if !dflTy.Present() {
				panic(errors.New("no type or default for field"))
			}
			opts = append(opts, def.TypeOpt{Ty: dflTy.Value()})
		}

		return def.FieldDef{
			Name: reifyIdentStr(call.Args[0]),
			Opts: opts,
		}

	case "Type":
		if len(call.Args) != 0 {
			panic(call)
		}

		idx := call.Fun.(*ast.IndexExpr)
		sel := idx.X.(*ast.SelectorExpr)
		inst := ti.Instances[sel.Sel]
		if inst.TypeArgs.Len() != 1 {
			panic(call)
		}

		return def.TypeOpt{
			Ty: typeRef(inst.TypeArgs.At(0)),
		}

	case "Default":
		if len(call.Args) != 1 {
			panic(call)
		}

		val := call.Args[0]
		ty := ti.Types[val]
		return def.DefaultOpt{
			Val: val,
			Ty:  typeRef(ty.Type),
		}

	case "Init":
		if len(call.Args) != 1 {
			panic(call)
		}

		return def.InitOpt{
			Fn: call.Args[0],
		}

	case "Meta":
		if len(call.Args) != 1 {
			panic(call)
		}

		return def.MetaOpt{
			Value: call.Args[0],
		}
	}

	panic(fmt.Errorf("unhandled type: %T", node))
}
