package c

import (
	"fmt"
	"testing"

	antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

	"github.com/wrmsr/bane/exp/util/c/parser"
	antlru "github.com/wrmsr/bane/pkg/util/antlr"
	"github.com/wrmsr/bane/pkg/util/check"
	ju "github.com/wrmsr/bane/pkg/util/json"
	msh "github.com/wrmsr/bane/pkg/util/marshal"
)

/*
int main(int argc, const char * const *argv) {
	int c = argc + 2 + 3;
	printf("hi %d\n", c);
	return 0;
}

clang -Xclang -ast-dump -fsyntax-only -fno-color-diagnostics a.c

TranslationUnitDecl 0x14c83f608 <<invalid sloc>> <invalid sloc>
|-TypedefDecl 0x14c8404f0 <<invalid sloc>> <invalid sloc> implicit __int128_t '__int128'
| `-BuiltinType 0x14c83fbb0 '__int128'
|-TypedefDecl 0x14c840560 <<invalid sloc>> <invalid sloc> implicit __uint128_t 'unsigned __int128'
| `-BuiltinType 0x14c83fbd0 'unsigned __int128'
|-TypedefDecl 0x14c8da4b0 <<invalid sloc>> <invalid sloc> implicit __NSConstantString 'struct __NSConstantString_tag'
| `-RecordType 0x14c8da280 'struct __NSConstantString_tag'
|   `-Record 0x14c8da200 '__NSConstantString_tag'
|-TypedefDecl 0x14c8da518 <<invalid sloc>> <invalid sloc> implicit __SVInt8_t '__SVInt8_t'
| `-BuiltinType 0x14c83fe30 '__SVInt8_t'
|-TypedefDecl 0x14c8da580 <<invalid sloc>> <invalid sloc> implicit __SVInt16_t '__SVInt16_t'
| `-BuiltinType 0x14c83fe50 '__SVInt16_t'
|-TypedefDecl 0x14c8da5e8 <<invalid sloc>> <invalid sloc> implicit __SVInt32_t '__SVInt32_t'
| `-BuiltinType 0x14c83fe70 '__SVInt32_t'
|-TypedefDecl 0x14c8da650 <<invalid sloc>> <invalid sloc> implicit __SVInt64_t '__SVInt64_t'
| `-BuiltinType 0x14c83fe90 '__SVInt64_t'
|-TypedefDecl 0x14c8da6b8 <<invalid sloc>> <invalid sloc> implicit __SVUint8_t '__SVUint8_t'
| `-BuiltinType 0x14c83feb0 '__SVUint8_t'
|-TypedefDecl 0x14c8da720 <<invalid sloc>> <invalid sloc> implicit __SVUint16_t '__SVUint16_t'
| `-BuiltinType 0x14c83fed0 '__SVUint16_t'
|-TypedefDecl 0x14c8da788 <<invalid sloc>> <invalid sloc> implicit __SVUint32_t '__SVUint32_t'
| `-BuiltinType 0x14c83fef0 '__SVUint32_t'
|-TypedefDecl 0x14c8da7f0 <<invalid sloc>> <invalid sloc> implicit __SVUint64_t '__SVUint64_t'
| `-BuiltinType 0x14c83ff10 '__SVUint64_t'
|-TypedefDecl 0x14c8da858 <<invalid sloc>> <invalid sloc> implicit __SVFloat16_t '__SVFloat16_t'
| `-BuiltinType 0x14c83ff30 '__SVFloat16_t'
|-TypedefDecl 0x14c8da8c0 <<invalid sloc>> <invalid sloc> implicit __SVFloat32_t '__SVFloat32_t'
| `-BuiltinType 0x14c83ff50 '__SVFloat32_t'
|-TypedefDecl 0x14c8da928 <<invalid sloc>> <invalid sloc> implicit __SVFloat64_t '__SVFloat64_t'
| `-BuiltinType 0x14c83ff70 '__SVFloat64_t'
|-TypedefDecl 0x14c8da990 <<invalid sloc>> <invalid sloc> implicit __SVBFloat16_t '__SVBFloat16_t'
| `-BuiltinType 0x14c83ff90 '__SVBFloat16_t'
|-TypedefDecl 0x14c8da9f8 <<invalid sloc>> <invalid sloc> implicit __clang_svint8x2_t '__clang_svint8x2_t'
| `-BuiltinType 0x14c83ffb0 '__clang_svint8x2_t'
|-TypedefDecl 0x14c8daa60 <<invalid sloc>> <invalid sloc> implicit __clang_svint16x2_t '__clang_svint16x2_t'
| `-BuiltinType 0x14c83ffd0 '__clang_svint16x2_t'
|-TypedefDecl 0x14c8daac8 <<invalid sloc>> <invalid sloc> implicit __clang_svint32x2_t '__clang_svint32x2_t'
| `-BuiltinType 0x14c83fff0 '__clang_svint32x2_t'
|-TypedefDecl 0x14c8dab30 <<invalid sloc>> <invalid sloc> implicit __clang_svint64x2_t '__clang_svint64x2_t'
| `-BuiltinType 0x14c840010 '__clang_svint64x2_t'
|-TypedefDecl 0x14c8dab98 <<invalid sloc>> <invalid sloc> implicit __clang_svuint8x2_t '__clang_svuint8x2_t'
| `-BuiltinType 0x14c840030 '__clang_svuint8x2_t'
|-TypedefDecl 0x14c8dac00 <<invalid sloc>> <invalid sloc> implicit __clang_svuint16x2_t '__clang_svuint16x2_t'
| `-BuiltinType 0x14c840050 '__clang_svuint16x2_t'
|-TypedefDecl 0x14c8dac68 <<invalid sloc>> <invalid sloc> implicit __clang_svuint32x2_t '__clang_svuint32x2_t'
| `-BuiltinType 0x14c840070 '__clang_svuint32x2_t'
|-TypedefDecl 0x14c8dacd0 <<invalid sloc>> <invalid sloc> implicit __clang_svuint64x2_t '__clang_svuint64x2_t'
| `-BuiltinType 0x14c840090 '__clang_svuint64x2_t'
|-TypedefDecl 0x14c8dad38 <<invalid sloc>> <invalid sloc> implicit __clang_svfloat16x2_t '__clang_svfloat16x2_t'
| `-BuiltinType 0x14c8400b0 '__clang_svfloat16x2_t'
|-TypedefDecl 0x14c8dada0 <<invalid sloc>> <invalid sloc> implicit __clang_svfloat32x2_t '__clang_svfloat32x2_t'
| `-BuiltinType 0x14c8400d0 '__clang_svfloat32x2_t'
|-TypedefDecl 0x14c8dae08 <<invalid sloc>> <invalid sloc> implicit __clang_svfloat64x2_t '__clang_svfloat64x2_t'
| `-BuiltinType 0x14c8400f0 '__clang_svfloat64x2_t'
|-TypedefDecl 0x14c8dae70 <<invalid sloc>> <invalid sloc> implicit __clang_svbfloat16x2_t '__clang_svbfloat16x2_t'
| `-BuiltinType 0x14c840110 '__clang_svbfloat16x2_t'
|-TypedefDecl 0x14c8daed8 <<invalid sloc>> <invalid sloc> implicit __clang_svint8x3_t '__clang_svint8x3_t'
| `-BuiltinType 0x14c840130 '__clang_svint8x3_t'
|-TypedefDecl 0x14c8daf40 <<invalid sloc>> <invalid sloc> implicit __clang_svint16x3_t '__clang_svint16x3_t'
| `-BuiltinType 0x14c840150 '__clang_svint16x3_t'
|-TypedefDecl 0x14c8dafa8 <<invalid sloc>> <invalid sloc> implicit __clang_svint32x3_t '__clang_svint32x3_t'
| `-BuiltinType 0x14c840170 '__clang_svint32x3_t'
|-TypedefDecl 0x14c8db010 <<invalid sloc>> <invalid sloc> implicit __clang_svint64x3_t '__clang_svint64x3_t'
| `-BuiltinType 0x14c840190 '__clang_svint64x3_t'
|-TypedefDecl 0x14c8db078 <<invalid sloc>> <invalid sloc> implicit __clang_svuint8x3_t '__clang_svuint8x3_t'
| `-BuiltinType 0x14c8401b0 '__clang_svuint8x3_t'
|-TypedefDecl 0x14c8db0e0 <<invalid sloc>> <invalid sloc> implicit __clang_svuint16x3_t '__clang_svuint16x3_t'
| `-BuiltinType 0x14c8401d0 '__clang_svuint16x3_t'
|-TypedefDecl 0x14c8db148 <<invalid sloc>> <invalid sloc> implicit __clang_svuint32x3_t '__clang_svuint32x3_t'
| `-BuiltinType 0x14c8401f0 '__clang_svuint32x3_t'
|-TypedefDecl 0x14c8dba00 <<invalid sloc>> <invalid sloc> implicit __clang_svuint64x3_t '__clang_svuint64x3_t'
| `-BuiltinType 0x14c840210 '__clang_svuint64x3_t'
|-TypedefDecl 0x14c8dba68 <<invalid sloc>> <invalid sloc> implicit __clang_svfloat16x3_t '__clang_svfloat16x3_t'
| `-BuiltinType 0x14c840230 '__clang_svfloat16x3_t'
|-TypedefDecl 0x14c8dbad0 <<invalid sloc>> <invalid sloc> implicit __clang_svfloat32x3_t '__clang_svfloat32x3_t'
| `-BuiltinType 0x14c840250 '__clang_svfloat32x3_t'
|-TypedefDecl 0x14c8dbb38 <<invalid sloc>> <invalid sloc> implicit __clang_svfloat64x3_t '__clang_svfloat64x3_t'
| `-BuiltinType 0x14c840270 '__clang_svfloat64x3_t'
|-TypedefDecl 0x14c8dbba0 <<invalid sloc>> <invalid sloc> implicit __clang_svbfloat16x3_t '__clang_svbfloat16x3_t'
| `-BuiltinType 0x14c840290 '__clang_svbfloat16x3_t'
|-TypedefDecl 0x14c8dbc08 <<invalid sloc>> <invalid sloc> implicit __clang_svint8x4_t '__clang_svint8x4_t'
| `-BuiltinType 0x14c8402b0 '__clang_svint8x4_t'
|-TypedefDecl 0x14c8dbc70 <<invalid sloc>> <invalid sloc> implicit __clang_svint16x4_t '__clang_svint16x4_t'
| `-BuiltinType 0x14c8402d0 '__clang_svint16x4_t'
|-TypedefDecl 0x14c8dbcd8 <<invalid sloc>> <invalid sloc> implicit __clang_svint32x4_t '__clang_svint32x4_t'
| `-BuiltinType 0x14c8402f0 '__clang_svint32x4_t'
|-TypedefDecl 0x14c8dbd40 <<invalid sloc>> <invalid sloc> implicit __clang_svint64x4_t '__clang_svint64x4_t'
| `-BuiltinType 0x14c840310 '__clang_svint64x4_t'
|-TypedefDecl 0x14c8dbda8 <<invalid sloc>> <invalid sloc> implicit __clang_svuint8x4_t '__clang_svuint8x4_t'
| `-BuiltinType 0x14c840330 '__clang_svuint8x4_t'
|-TypedefDecl 0x14c8dbe10 <<invalid sloc>> <invalid sloc> implicit __clang_svuint16x4_t '__clang_svuint16x4_t'
| `-BuiltinType 0x14c840350 '__clang_svuint16x4_t'
|-TypedefDecl 0x14c8dbe78 <<invalid sloc>> <invalid sloc> implicit __clang_svuint32x4_t '__clang_svuint32x4_t'
| `-BuiltinType 0x14c840370 '__clang_svuint32x4_t'
|-TypedefDecl 0x14c8dbee0 <<invalid sloc>> <invalid sloc> implicit __clang_svuint64x4_t '__clang_svuint64x4_t'
| `-BuiltinType 0x14c840390 '__clang_svuint64x4_t'
|-TypedefDecl 0x14c8dbf48 <<invalid sloc>> <invalid sloc> implicit __clang_svfloat16x4_t '__clang_svfloat16x4_t'
| `-BuiltinType 0x14c8403b0 '__clang_svfloat16x4_t'
|-TypedefDecl 0x14c8dbfb0 <<invalid sloc>> <invalid sloc> implicit __clang_svfloat32x4_t '__clang_svfloat32x4_t'
| `-BuiltinType 0x14c8403d0 '__clang_svfloat32x4_t'
|-TypedefDecl 0x14c8dc018 <<invalid sloc>> <invalid sloc> implicit __clang_svfloat64x4_t '__clang_svfloat64x4_t'
| `-BuiltinType 0x14c8403f0 '__clang_svfloat64x4_t'
|-TypedefDecl 0x14c8dc080 <<invalid sloc>> <invalid sloc> implicit __clang_svbfloat16x4_t '__clang_svbfloat16x4_t'
| `-BuiltinType 0x14c840410 '__clang_svbfloat16x4_t'
|-TypedefDecl 0x14c8dc0e8 <<invalid sloc>> <invalid sloc> implicit __SVBool_t '__SVBool_t'
| `-BuiltinType 0x14c840430 '__SVBool_t'
|-TypedefDecl 0x14c8dc188 <<invalid sloc>> <invalid sloc> implicit __builtin_ms_va_list 'char *'
| `-PointerType 0x14c8dc140 'char *'
|   `-BuiltinType 0x14c83f6b0 'char'
|-TypedefDecl 0x14c8dc1f8 <<invalid sloc>> <invalid sloc> implicit __builtin_va_list 'char *'
| `-PointerType 0x14c8dc140 'char *'
|   `-BuiltinType 0x14c83f6b0 'char'
|-FunctionDecl 0x14c8dc3f0 <a.c:1:1, line:5:1> line:1:5 main 'int (int, const char *const *)'
| |-ParmVarDecl 0x14c8dc268 <col:10, col:14> col:14 used argc 'int'
| |-ParmVarDecl 0x14c8dc318 <col:20, col:40> col:40 argv 'const char *const *'
| `-CompoundStmt 0x14c9027b0 <col:46, line:5:1>
|   |-DeclStmt 0x14c8dc650 <line:2:5, col:25>
|   | `-VarDecl 0x14c8dc530 <col:5, col:24> col:9 used c 'int' cinit
|   |   `-BinaryOperator 0x14c8dc630 <col:13, col:24> 'int' '+'
|   |     |-BinaryOperator 0x14c8dc5f0 <col:13, col:20> 'int' '+'
|   |     | |-ImplicitCastExpr 0x14c8dc5d8 <col:13> 'int' <LValueToRValue>
|   |     | | `-DeclRefExpr 0x14c8dc598 <col:13> 'int' lvalue ParmVar 0x14c8dc268 'argc' 'int'
|   |     | `-IntegerLiteral 0x14c8dc5b8 <col:20> 'int' 2
|   |     `-IntegerLiteral 0x14c8dc610 <col:24> 'int' 3
|   |-CallExpr 0x14c902708 <line:3:5, col:24> 'int'
|   | |-ImplicitCastExpr 0x14c9026f0 <col:5> 'int (*)(const char *, ...)' <FunctionToPointerDecay>
|   | | `-DeclRefExpr 0x14c902610 <col:5> 'int (const char *, ...)' Function 0x14c902470 'printf' 'int (const char *, ...)'
|   | |-ImplicitCastExpr 0x14c902750 <col:12> 'const char *' <NoOp>
|   | | `-ImplicitCastExpr 0x14c902738 <col:12> 'char *' <ArrayToPointerDecay>
|   | |   `-StringLiteral 0x14c902668 <col:12> 'char [7]' lvalue "hi %d\n"
|   | `-ImplicitCastExpr 0x14c902768 <col:23> 'int' <LValueToRValue>
|   |   `-DeclRefExpr 0x14c902688 <col:23> 'int' lvalue Var 0x14c8dc530 'c' 'int'
|   `-ReturnStmt 0x14c9027a0 <line:4:5, col:12>
|     `-IntegerLiteral 0x14c902780 <col:12> 'int' 0
`-FunctionDecl 0x14c902470 <line:3:5> col:5 implicit used printf 'int (const char *, ...)' extern
  |-ParmVarDecl 0x14c902568 <<invalid sloc>> <invalid sloc> 'const char *'
  |-BuiltinAttr 0x14c902510 <<invalid sloc>> Implicit 806
  `-FormatAttr 0x14c9025d8 <col:5> Implicit printf 1 2
*/

func TestC(t *testing.T) {
	cu := func(p *parser.CParser) antlr.ParseTree { return p.CompilationUnit() }
	ex := func(p *parser.CParser) antlr.ParseTree { return p.PrimaryExpression() }
	ts := func(p *parser.CParser) antlr.ParseTree { return p.TypeSpecifier() }

	_ = cu
	_ = ex
	_ = ts

	for _, src := range []struct {
		s   string
		pfn func(p *parser.CParser) antlr.ParseTree
	}{
		{`1`, ex},
		{`int`, ts},
		{`int foo(int x) {
			return x;
		}`, cu},
		//{`int main(int argc, const char * const *argv) {
		//	printf("hi\n");
		//	return 0;
		//}`, cu},
		{`int main(int argc, const char * const *argv) {
			int c = argc + 2 + 3;
			printf("hi %d\n", c);
			return 0;
		}`, cu},
		{`struct tm {
			int tm_sec;
			int tm_min;
			int tm_hour;
			int tm_mday;
			int tm_mon;
			int tm_year;
			int tm_wday;
			int tm_yday;
			int tm_isdst;
			long tm_gmtoff;
			char *tm_zone;
		};`, cu},
	} {
		fmt.Println(src.s)

		is := antlr.NewInputStream(src.s)
		lexer := parser.NewCLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewCParser(stream)
		p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
		p.BuildParseTrees = true

		tree := src.pfn(p)

		fmt.Println(tree)

		antlru.Dump(tree, "")

		//tu.AssertNoErr(t, dot.Open(context.Background(), dot.RenderString(antlru.Dot(tree))))

		v := &parseVisitor{}
		//v.BaseParseTreeVisitor = v

		n := tree.Accept(v)
		fmt.Printf("%#v\n", n)

		m := check.Must1(msh.Marshal(n))
		fmt.Println(check.Must1(ju.MarshalPretty(m)))

		fmt.Println()
	}
}
