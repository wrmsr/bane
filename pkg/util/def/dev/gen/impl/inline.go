//go:build !nodev

package impl

import (
	"go/ast"
	"go/printer"
	"go/token"
	"go/types"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/def"
	gg "github.com/wrmsr/bane/pkg/util/go/gen"
	"github.com/wrmsr/bane/pkg/util/maps"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func findFuncNames(decl *ast.FuncDecl) maps.Set[string] {
	names := maps.MakeSet[string]()

	if decl.Recv != nil && len(decl.Recv.List) > 0 {
		rn := check.Single(check.Single(decl.Recv.List).Names).Name
		names.Add(rn)
	}

	ast.Inspect(decl.Body, func(node ast.Node) bool {
		switch node := node.(type) {

		case *ast.ValueSpec:
			for _, n := range node.Names {
				names.Add(n.Name)
			}

		case *ast.AssignStmt:
			if node.Tok != token.DEFINE {
				break
			}
			for _, e := range node.Lhs {
				names.Add(e.(*ast.Ident).Name)
			}

		}
		return true
	})

	return names
}

type funcInliner struct {
	decl *FuncDecl
	ti   *types.Info
	im   map[string]*FuncDecl

	names  maps.Set[string]
	nameCt int
}

func (fil *funcInliner) nextName() string {
	namePfx := "__def_inl_"
	for {
		n := namePfx + strconv.Itoa(fil.nameCt)
		fil.nameCt++
		if !fil.names.Contains(n) {
			fil.names.Add(n)
			return n
		}
	}
}

func defVar(n string, t ast.Expr) ast.Stmt {
	return &ast.DeclStmt{
		Decl: &ast.GenDecl{
			Tok: token.VAR,
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{
						{Name: n},
					},
					Type: t,
				},
			},
		},
	}
}

func defAssign(n string, e ast.Expr) ast.Stmt {
	return &ast.AssignStmt{
		Tok: token.DEFINE,
		Lhs: []ast.Expr{
			&ast.Ident{Name: n},
		},
		Rhs: []ast.Expr{
			e,
		},
	}
}

func replaceIdents(n ast.Node, m map[string]string) ast.Node {
	return astutil.Apply(n, nil, func(cursor *astutil.Cursor) bool {
		switch n := cursor.Node().(type) {
		case *ast.Ident:
			if w, ok := m[n.Name]; ok {
				cursor.Replace(&ast.Ident{Name: w})
			}
		}
		return true
	})
}

func returnsToGotos(body *ast.BlockStmt) {

}

func (fil *funcInliner) doTransform(idecl *FuncDecl, outn string, paramns []string, argns []string) []ast.Stmt {
	el := fil.nextName()

	stmts := fil.doBlock(idecl.Decl.Body)

	for i, s := range stmts {
		stmts[i] = astutil.Apply(s, nil, func(cursor *astutil.Cursor) bool {
			switch n := cursor.Node().(type) {

			case *ast.ReturnStmt:
				cursor.Replace(
					&ast.BlockStmt{
						List: []ast.Stmt{
							&ast.AssignStmt{
								Tok: token.ASSIGN,
								Lhs: []ast.Expr{
									&ast.Ident{Name: outn},
								},
								Rhs: []ast.Expr{
									check.Single(n.Results),
								},
							},
							&ast.BranchStmt{
								Tok:   token.GOTO,
								Label: &ast.Ident{Name: el},
							},
						},
					})

			case *ast.Ident:
				for i, pn := range paramns {
					if n.Name == pn {
						cursor.Replace(&ast.Ident{Name: argns[i]})
						break
					}
				}

			}

			return true
		}).(ast.Stmt)
	}

	stmts = append(stmts, &ast.LabeledStmt{
		Label: &ast.Ident{Name: el},
		Stmt:  &ast.EmptyStmt{},
	})

	return stmts
}

func (fil *funcInliner) doInline(call *ast.CallExpr, idecl *FuncDecl) (ast.Expr, []ast.Stmt) {
	var stmts []ast.Stmt

	params := check.Single(idecl.Decl.Type.Params.List).Names
	check.Equal(len(params), len(call.Args))
	m := make(map[string]string)
	paramns := make([]string, 0, len(params)+1)
	var tn string
	if idecl.Decl.Recv != nil && len(idecl.Decl.Recv.List) > 0 {
		rn := check.Single(check.Single(idecl.Decl.Recv.List).Names).Name
		sx := call.Fun.(*ast.SelectorExpr).X
		var si *ast.Ident
		for {
			if si2, ok := sx.(*ast.Ident); ok {
				si = si2
				break
			}
			sx = sx.(*ast.SelectorExpr).X
		}
		tn = si.Name
		paramns = append(paramns, rn)
	}
	for _, p := range params {
		paramns = append(paramns, p.Name)
	}

	outn := fil.nextName()
	stmts = append(stmts, defVar(outn, check.Single(idecl.Decl.Type.Results.List).Type))

	argns := make([]string, 0, len(call.Args)+1)
	if tn != "" {
		tan := fil.nextName()
		argns = append(argns, tan)
		stmts = append(stmts, defAssign(tan, &ast.Ident{Name: tn}))
	}
	for i := 0; i < len(call.Args); i++ {
		argns = append(argns, fil.nextName())
	}

	for i, a := range call.Args {
		an := argns[i+bt.Choose(tn != "", 1, 0)]
		ae, as := fil.doExpr(a)

		stmts = append(stmts, as...)
		stmts = append(stmts, defAssign(an, ae))
	}

	istmts := fil.doTransform(idecl, outn, paramns, argns)

	stmts = append(stmts, &ast.BlockStmt{List: istmts})

	return &ast.Ident{Name: outn}, stmts
}

func (fil *funcInliner) doExpr(expr ast.Expr) (ast.Expr, []ast.Stmt) {
	var stmts []ast.Stmt

	out := astutil.Apply(expr, nil, func(cursor *astutil.Cursor) bool {
		switch node := cursor.Node().(type) {
		case *ast.CallExpr:
			fnam := nameFuncRef(node.Fun, fil.ti)
			idecl := fil.im[fnam]
			if idecl == nil {
				return true
			}

			inl, inlstmts := fil.doInline(node, idecl)
			stmts = append(stmts, inlstmts...)
			cursor.Replace(inl)
		}
		return true
	})

	return out.(ast.Expr), stmts
}

func (fil *funcInliner) doStmt(stmt ast.Stmt) []ast.Stmt {
	var stmts []ast.Stmt

	out := astutil.Apply(stmt, nil, func(cursor *astutil.Cursor) bool {
		n := cursor.Node()
		if e, ok := n.(ast.Expr); ok {
			o, s := fil.doExpr(e)
			stmts = append(stmts, s...)
			cursor.Replace(o)
		}
		return true
	})

	stmts = append(stmts, out.(ast.Stmt))
	return stmts
}

func (fil *funcInliner) doBlock(block *ast.BlockStmt) []ast.Stmt {
	var stmts []ast.Stmt

	for _, stmt := range block.List {
		stmts = append(stmts, fil.doStmt(stmt)...)
	}

	return stmts
}

func (fil *funcInliner) inlineFunc() *ast.FuncDecl {
	fil.names = findFuncNames(fil.decl.Decl)

	stmts := fil.doBlock(fil.decl.Decl.Body)
	body := *fil.decl.Decl.Body
	body.List = stmts

	out := *fil.decl.Decl
	out.Name = &ast.Ident{
		Name: "_def_inl_" + out.Name.Name,
	}
	out.Body = &body

	return &out
}

type inliner struct {
	pkg *packages.Package
	ps  *def.PackageSpec
	rsv *Resolver
}

func (il *inliner) inlineFuncs() gg.Decls {
	var decls gg.Decls

	im := make(map[string]*FuncDecl)
	for id := range il.ps.Inlines() {
		n := nameFuncRef(id, il.pkg.TypesInfo)
		decl := il.rsv.FuncDecl(n)
		if decl == nil {
			panic(n)
		}
		im[n] = decl
	}

	for wid := range il.ps.WithInlines() {
		n := nameFuncRef(wid, il.pkg.TypesInfo)
		decl := il.rsv.FuncDecl(n)
		if decl == nil {
			panic(n)
		}

		id := (&funcInliner{
			decl: decl,
			ti:   il.pkg.TypesInfo,
			im:   im,
		}).inlineFunc()

		var sb strings.Builder
		_ = printer.Fprint(&sb, il.pkg.Fset, id)

		decls.Append(gg.Raw{Raw: sb.String()})
	}

	return decls
}
