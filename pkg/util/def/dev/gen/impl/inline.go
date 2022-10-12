//go:build !nodev

package impl

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
	"strconv"

	"golang.org/x/tools/go/ast/astutil"

	"github.com/wrmsr/bane/pkg/util/check"
	astu "github.com/wrmsr/bane/pkg/util/go/ast"
	"github.com/wrmsr/bane/pkg/util/maps"
)

func (fg *FileGen) inlineFunc(decl *FuncDecl, im map[string]*FuncDecl) {
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

	doExprStmt := func(stmt *ast.ExprStmt) []ast.Stmt {
		var stmts []ast.Stmt

		expr := astutil.Apply(stmt.X, nil, func(cursor *astutil.Cursor) bool {
			var idecl *FuncDecl
			switch node := cursor.Node().(type) {
			case *ast.CallExpr:
				fnam := nameFuncRef(node.Fun)
				idecl = im[fnam]
				if idecl == nil {
					return true
				}
			default:
				return true
			}

			fmt.Println(nextName())
			panic(idecl)

			return true
		})

		out := *stmt
		out.X = expr.(ast.Expr)
		stmts = append(stmts, &out)
		return stmts
	}

	var doBlock func(stmt *ast.BlockStmt) []ast.Stmt

	doStmt := func(stmt ast.Stmt) []ast.Stmt {
		switch stmt := stmt.(type) {
		case *ast.ExprStmt:
			return doExprStmt(stmt)
		case *ast.BlockStmt:
			return doBlock(stmt)
		default:
			return []ast.Stmt{stmt}
		}
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
	out.Body = &body

	astu.Dump(os.Stdout, &out)
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

		//nn := "_def_inl_" + n
		fg.inlineFunc(decl, im)
	}
}
