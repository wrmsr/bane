//go:build !nodev

package impl

import (
	"go/ast"
	"go/printer"
	"go/token"
	"strconv"
	"strings"

	"golang.org/x/tools/go/ast/astutil"

	"github.com/wrmsr/bane/pkg/util/check"
	gg "github.com/wrmsr/bane/pkg/util/go/gen"
	"github.com/wrmsr/bane/pkg/util/maps"
)

func (fg *FileGen) inlineFunc(decl *FuncDecl, im map[string]*FuncDecl) *ast.FuncDecl {
	names := maps.MakeSet[string]()

	if decl.Decl.Recv != nil && len(decl.Decl.Recv.List) > 0 {
		rn := check.Single(check.Single(decl.Decl.Recv.List).Names).Name
		names.Add(rn)
	}

	ast.Inspect(decl.Decl.Body, func(node ast.Node) bool {
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

	namePfx := "__def_inl_"
	nameCt := 0
	nextName := func() string {
		for {
			n := namePfx + strconv.Itoa(nameCt)
			nameCt++
			if !names.Contains(n) {
				names.Add(n)
				return n
			}
		}
	}

	defVar := func(n string, t ast.Expr) ast.Stmt {
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

	defAssign := func(n string, e ast.Expr) ast.Stmt {
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

	var doExpr func(expr ast.Expr) (ast.Expr, []ast.Stmt)

	var doBlock func(stmt *ast.BlockStmt) []ast.Stmt

	doTransform := func(idecl *FuncDecl, outn string, paramns []string, argns []string) []ast.Stmt {
		el := nextName()
		stmts := doBlock(idecl.Decl.Body)
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

	doInline := func(call *ast.CallExpr, idecl *FuncDecl) (ast.Expr, []ast.Stmt) {
		var stmts []ast.Stmt

		params := check.Single(idecl.Decl.Type.Params.List).Names
		check.Equal(len(params), len(call.Args))
		paramns := make([]string, len(params))
		for i, p := range params {
			paramns[i] = p.Name
		}

		outn := nextName()
		stmts = append(stmts, defVar(outn, check.Single(idecl.Decl.Type.Results.List).Type))

		argns := make([]string, len(call.Args))
		for i := range call.Args {
			argns[i] = nextName()
		}

		for i, a := range call.Args {
			an := argns[i]
			ae, as := doExpr(a)

			stmts = append(stmts, as...)
			stmts = append(stmts, defAssign(an, ae))
		}

		istmts := doTransform(idecl, outn, paramns, argns)

		stmts = append(stmts, &ast.BlockStmt{List: istmts})

		return &ast.Ident{Name: outn}, stmts
	}

	doExpr = func(expr ast.Expr) (ast.Expr, []ast.Stmt) {
		var stmts []ast.Stmt
		out := astutil.Apply(expr, nil, func(cursor *astutil.Cursor) bool {
			switch node := cursor.Node().(type) {
			case *ast.CallExpr:
				fnam := nameFuncRef(node.Fun)
				idecl := im[fnam]
				if idecl == nil {
					return true
				}

				inl, inlstmts := doInline(node, idecl)
				stmts = append(stmts, inlstmts...)
				cursor.Replace(inl)
			}
			return true
		})
		return out.(ast.Expr), stmts
	}

	doStmt := func(stmt ast.Stmt) []ast.Stmt {
		var stmts []ast.Stmt

		out := astutil.Apply(stmt, nil, func(cursor *astutil.Cursor) bool {
			n := cursor.Node()
			if e, ok := n.(ast.Expr); ok {
				o, s := doExpr(e)
				stmts = append(stmts, s...)
				cursor.Replace(o)
			}
			return true
		})

		stmts = append(stmts, out.(ast.Stmt))
		return stmts
	}

	doBlock = func(block *ast.BlockStmt) []ast.Stmt {
		var stmts []ast.Stmt
		for _, stmt := range block.List {
			stmts = append(stmts, doStmt(stmt)...)
		}
		return stmts
	}

	stmts := doBlock(decl.Decl.Body)
	body := *decl.Decl.Body
	body.List = stmts

	out := *decl.Decl
	out.Name = &ast.Ident{
		Name: "_def_inl_" + out.Name.Name,
	}
	out.Body = &body

	return &out
}

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

		id := fg.inlineFunc(decl, im)

		var sb strings.Builder
		_ = printer.Fprint(&sb, fg.pkg.Fset, id)

		fg.decls.Append(gg.Raw{Raw: sb.String()})
	}
}
