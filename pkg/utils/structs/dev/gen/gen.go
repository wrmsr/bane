//go:build !nodev

package main

import (
	"bufio"
	"errors"
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/types/typeutil"

	"github.com/wrmsr/bane/pkg/utils/check"
	ctr "github.com/wrmsr/bane/pkg/utils/containers"
	eu "github.com/wrmsr/bane/pkg/utils/errors"
)

func findDirWithFile(cd, fn string) (string, error) {
	for {
		if _, err := os.Stat(filepath.Join(cd, fn)); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return "", err
			}
		} else {
			return cd, nil
		}

		nd := filepath.Dir(cd)
		if nd == cd {
			return "", fmt.Errorf("file %s not found from dir %s", fn, cd)
		}
		cd = nd
	}
}

func readFirstLine(fp string) (s string, err error) {
	f, err := os.Open(fp)
	if err != nil {
		return
	}
	defer eu.AppendInvoke(&err, eu.Close(f))

	r := bufio.NewReader(f)
	l, p, err := r.ReadLine()
	if err != nil {
		return
	}
	if p {
		err = errors.New("line too big")
		return
	}

	s = string(l)
	return
}

var dontFixRetract modfile.VersionFixer = func(_, vers string) (string, error) {
	return vers, nil
}

//

type Pass struct {
	Files     []*ast.File    // the abstract syntax tree of each file
	Pkg       *types.Package // type information about the package
	TypesInfo *types.Info    // type information about the syntax trees
}

type printfWrapper struct {
	obj     *types.Func
	fdecl   *ast.FuncDecl
	format  *types.Var
	args    *types.Var
	callers []printfCaller
	failed  bool // if true, not a printf wrapper
}

type printfCaller struct {
	w    *printfWrapper
	call *ast.CallExpr
}

// maybePrintfWrapper decides whether decl (a declared function) may be a wrapper
// around a fmt.Printf or fmt.Print function. If so it returns a printfWrapper
// function describing the declaration. Later processing will analyze the
// graph of potential printf wrappers to pick out the ones that are true wrappers.
// A function may be a Printf or Print wrapper if its last argument is ...any.
// If the next-to-last argument is a string, then this may be a Printf wrapper.
// Otherwise it may be a Print wrapper.
func maybePrintfWrapper(info *types.Info, decl ast.Decl) *printfWrapper {
	// Look for functions with final argument type ...any.
	fdecl, ok := decl.(*ast.FuncDecl)
	if !ok || fdecl.Body == nil {
		return nil
	}
	fn, ok := info.Defs[fdecl.Name].(*types.Func)
	// Type information may be incomplete.
	if !ok {
		return nil
	}

	sig := fn.Type().(*types.Signature)
	if !sig.Variadic() {
		return nil // not variadic
	}

	params := sig.Params()
	nparams := params.Len() // variadic => nonzero

	args := params.At(nparams - 1)
	iface, ok := args.Type().(*types.Slice).Elem().(*types.Interface)
	if !ok || !iface.Empty() {
		return nil // final (args) param is not ...any
	}

	// Is second last param 'format string'?
	var format *types.Var
	if nparams >= 2 {
		if p := params.At(nparams - 2); p.Type() == types.Typ[types.String] {
			format = p
		}
	}

	return &printfWrapper{
		obj:    fn,
		fdecl:  fdecl,
		format: format,
		args:   args,
	}
}

// findPrintfLike scans the entire package to find printf-like functions.
func findPrintfLike(pass *Pass) (any, error) {
	// Gather potential wrappers and call graph between them.
	byObj := make(map[*types.Func]*printfWrapper)
	var wrappers []*printfWrapper
	for _, file := range pass.Files {
		for _, decl := range file.Decls {
			w := maybePrintfWrapper(pass.TypesInfo, decl)
			if w == nil {
				continue
			}
			byObj[w.obj] = w
			wrappers = append(wrappers, w)
		}
	}

	// Walk the graph to figure out which are really printf wrappers.
	for _, w := range wrappers {
		// Scan function for calls that could be to other printf-like functions.
		ast.Inspect(w.fdecl.Body, func(n ast.Node) bool {
			if w.failed {
				return false
			}

			// TODO: Relax these checks; issue 26555.
			if assign, ok := n.(*ast.AssignStmt); ok {
				for _, lhs := range assign.Lhs {
					if match(pass.TypesInfo, lhs, w.format) ||
						match(pass.TypesInfo, lhs, w.args) {
						// Modifies the format
						// string or args in
						// some way, so not a
						// simple wrapper.
						w.failed = true
						return false
					}
				}
			}
			if un, ok := n.(*ast.UnaryExpr); ok && un.Op == token.AND {
				if match(pass.TypesInfo, un.X, w.format) ||
					match(pass.TypesInfo, un.X, w.args) {
					// Taking the address of the
					// format string or args,
					// so not a simple wrapper.
					w.failed = true
					return false
				}
			}

			call, ok := n.(*ast.CallExpr)
			if !ok || len(call.Args) == 0 || !match(pass.TypesInfo, call.Args[len(call.Args)-1], w.args) {
				return true
			}

			fn, kind := printfNameAndKind(pass, call)
			if kind != 0 {
				//checkPrintfFwd(pass, w, call, kind)
				fmt.Printf("%v %v %v %v\n", pass, w, call, kind)
				return true
			}

			// If the call is to another function in this package,
			// maybe we will find out it is printf-like later.
			// Remember this call for later checking.
			if fn != nil && fn.Pkg() == pass.Pkg && byObj[fn] != nil {
				callee := byObj[fn]
				callee.callers = append(callee.callers, printfCaller{w, call})
			}

			return true
		})
	}
	return nil, nil
}

func match(info *types.Info, arg ast.Expr, param *types.Var) bool {
	id, ok := arg.(*ast.Ident)
	return ok && info.ObjectOf(id) == param
}

// Kind is a kind of fmt function behavior.
type Kind int

const (
	KindNone   Kind = iota // not a fmt wrapper function
	KindPrint              // function behaves like fmt.Print
	KindPrintf             // function behaves like fmt.Printf
	KindErrorf             // function behaves like fmt.Errorf
)

func (kind Kind) String() string {
	switch kind {
	case KindPrint:
		return "print"
	case KindPrintf:
		return "printf"
	case KindErrorf:
		return "errorf"
	}
	return ""
}

// isWrapper is a fact indicating that a function is a print or printf wrapper.
type isWrapper struct{ Kind Kind }

func (f *isWrapper) AFact() {}

func (f *isWrapper) String() string {
	switch f.Kind {
	case KindPrintf:
		return "printfWrapper"
	case KindPrint:
		return "printWrapper"
	case KindErrorf:
		return "errorfWrapper"
	default:
		return "unknownWrapper"
	}
}

func printfNameAndKind(pass *Pass, call *ast.CallExpr) (fn *types.Func, kind Kind) {
	fn, _ = typeutil.Callee(pass.TypesInfo, call).(*types.Func)
	if fn == nil {
		return nil, 0
	}

	if isPrint.Contains(fn.FullName()) || isPrint.Contains(strings.ToLower(fn.Name())) {
		if fn.FullName() == "fmt.Errorf" {
			kind = KindErrorf
		} else if strings.HasSuffix(fn.Name(), "f") {
			kind = KindPrintf
		} else {
			kind = KindPrint
		}
		return fn, kind
	}

	//var fact isWrapper
	//if pass.ImportObjectFact(fn, &fact) {
	//	return fn, fact.Kind
	//}

	return fn, KindNone
}

var isPrint = ctr.NewSetOf(
	"fmt.Errorf",
	"fmt.Fprint",
	"fmt.Fprintf",
	"fmt.Fprintln",
	"fmt.Print",
	"fmt.Printf",
	"fmt.Println",
	"fmt.Sprint",
	"fmt.Sprintf",
	"fmt.Sprintln",

	"runtime/trace.Logf",

	"log.Print",
	"log.Printf",
	"log.Println",
	"log.Fatal",
	"log.Fatalf",
	"log.Fatalln",
	"log.Panic",
	"log.Panicf",
	"log.Panicln",
	"(*log.Logger).Fatal",
	"(*log.Logger).Fatalf",
	"(*log.Logger).Fatalln",
	"(*log.Logger).Panic",
	"(*log.Logger).Panicf",
	"(*log.Logger).Panicln",
	"(*log.Logger).Print",
	"(*log.Logger).Printf",
	"(*log.Logger).Println",

	"(*testing.common).Error",
	"(*testing.common).Errorf",
	"(*testing.common).Fatal",
	"(*testing.common).Fatalf",
	"(*testing.common).Log",
	"(*testing.common).Logf",
	"(*testing.common).Skip",
	"(*testing.common).Skipf",
	// *testing.T and B are detected by induction, but testing.TB is
	// an interface and the inference can't follow dynamic calls.
	"(testing.TB).Error",
	"(testing.TB).Errorf",
	"(testing.TB).Fatal",
	"(testing.TB).Fatalf",
	"(testing.TB).Log",
	"(testing.TB).Logf",
	"(testing.TB).Skip",
	"(testing.TB).Skipf",
)

// stringSet is a set-of-nonempty-strings-valued flag.
// Note: elements without a '.' get lower-cased.
type stringSet map[string]bool

func (ss stringSet) String() string {
	var list []string
	for name := range ss {
		list = append(list, name)
	}
	sort.Strings(list)
	return strings.Join(list, ",")
}

func (ss stringSet) Set(flag string) error {
	for _, name := range strings.Split(flag, ",") {
		if len(name) == 0 {
			return fmt.Errorf("empty string")
		}
		if !strings.Contains(name, ".") {
			name = strings.ToLower(name)
		}
		ss[name] = true
	}
	return nil
}

//

func main() {
	cd := check.Must(os.Getwd())
	rd := check.Must(findDirWithFile(cd, "go.mod"))
	if !strings.HasPrefix(cd, rd) {
		panic(fmt.Errorf("can't find path %s from root %s", cd, rd))
	}
	do := cd[len(rd)+1:]
	fmt.Println(do)

	mf := filepath.Join(rd, "go.mod")
	mc := check.Must(ioutil.ReadFile(mf))
	mo := check.Must(modfile.Parse(mf, mc, dontFixRetract))
	mp := mo.Module.Mod.Path
	fmt.Println(mp)

	cfg := &packages.Config{
		Mode: packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
	}

	pn := fmt.Sprintf("%s/%s", mp, do)
	pkgs := check.Must(packages.Load(cfg, pn))

	for _, pkg := range pkgs {
		pass := Pass{
			Files:     pkg.Syntax,
			Pkg:       pkg.Types,
			TypesInfo: pkg.TypesInfo,
		}
		_ = check.Must(findPrintfLike(&pass))
	}
}
