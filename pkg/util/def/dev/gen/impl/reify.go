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
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
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
		opts := make([]def.StructOpt, len(call.Args))
		for i, arg := range call.Args {
			opts[i] = ReifyDef(arg, ti).(def.StructOpt)
		}

		idx := call.Fun.(*ast.IndexExpr)
		id := idx.Index.(*ast.Ident)

		return def.StructDef{
			Ty:   id.Name,
			Opts: opts,
		}

	case "Field":
		opts := make([]def.FieldOpt, len(call.Args)-1)

		hasTyOpt := false
		var dflTy bt.Optional[any]
		for i, arg := range call.Args[1:] {
			o := ReifyDef(arg, ti).(def.FieldOpt)
			if _, ok := o.(def.TypeOpt); ok {
				hasTyOpt = true
			}
			if o, ok := o.(def.DefaultOpt); ok {
				dflTy = bt.Just(o.Ty)
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
			Ty: inst.TypeArgs.At(0),
		}

	case "Default":
		if len(call.Args) != 1 {
			panic(call)
		}

		val := call.Args[0]
		ty := ti.Types[val]
		return def.DefaultOpt{
			Val: val,
			Ty:  ty.Type,
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

	case "Enum":
		var opts []def.EnumOpt
		if len(call.Args) > 0 {
			opts = make([]def.EnumOpt, len(call.Args))
			for i, arg := range call.Args[1:] {
				opts[i] = ReifyDef(arg, ti).(def.EnumOpt)
			}
		}

		idx := call.Fun.(*ast.IndexExpr)
		sel := idx.X.(*ast.SelectorExpr)
		inst := ti.Instances[sel.Sel]
		if inst.TypeArgs.Len() != 1 {
			panic(call)
		}

		return def.EnumDef{
			Ty:   inst.TypeArgs.At(0),
			Opts: opts,
		}

	case "Inline":
		return def.InlineDef{
			Fns: slices.AsAny(call.Args),
		}

	case "WithInline":
		return def.WithInlineDef{
			Fns: slices.AsAny(call.Args),
		}

	case "ConstGeneric":
		cvs := make([]any, len(call.Args))
		for i, arg := range call.Args {
			cvs[i] = arg.(*ast.Ident).Name
		}

		idx := call.Fun.(*ast.IndexExpr)
		var ty any
		switch id := idx.Index.(type) {
		case *ast.Ident:
			ty = id.Name
		case *ast.IndexExpr:
			ty = id.X.(*ast.Ident).Name
		default:
			panic(ty)
		}

		return def.ConstGenericDef{
			Ty:  ty,
			Cvs: cvs,
		}

	}

	panic(fmt.Errorf("unhandled type: %T", node))
}

func ReifyFunc(node ast.Node, ti *types.Info) any {
	fn := node.(*ast.FuncDecl)

	if strings.HasPrefix(fn.Name.Name, "_def_init_") {
		rcv := check.Single(fn.Recv.List).Type

		var sn string
		switch rcv := rcv.(type) {
		case *ast.StarExpr:
			sn = rcv.X.(*ast.Ident).Name
		case *ast.Ident:
			sn = rcv.Name
		default:
			panic(rcv)
		}

		return def.StructDef{
			Ty: sn,
			Opts: []def.StructOpt{
				def.InitOpt{
					Fn: fn,
				},
			},
		}
	}

	if strings.HasPrefix(fn.Name.Name, "_def_lazy_") {
		panic("fixme")
	}

	panic(fmt.Errorf("unhandled def func: %s", fn.Name.Name))
}
