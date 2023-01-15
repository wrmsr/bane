package c

import (
	"fmt"
	"testing"

	antlr "github.com/wrmsr/bane/pkg/util/antlr/runtime"

	"github.com/wrmsr/bane/exp/util/c/parser"
	antlru "github.com/wrmsr/bane/pkg/util/antlr"
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
		{`int main(int argc, const char * const *argv) {
			printf("hi\n");
			return 0;
		}`, cu},
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
		{`
/*
** 2001 September 15
**
** The author disclaims copyright to this source code.  In place of
** a legal notice, here is a blessing:
**
**    May you do good and not evil.
**    May you find forgiveness for yourself and forgive others.
**    May you share freely, never taking more than you give.
**
*************************************************************************
** This header file defines the interface that the sqlite B-Tree file
** subsystem.  See comments in the source code for a detailed description
** of what each interface routine does.
*/
#ifndef SQLITE_BTREE_H
#define SQLITE_BTREE_H

/* TODO: This definition is just included so other modules compile. It
** needs to be revisited.
*/
#define SQLITE_N_BTREE_META 16

/*
** If defined as non-zero, auto-vacuum is enabled by default. Otherwise
** it must be turned on for each database using "PRAGMA auto_vacuum = 1".
*/
#ifndef SQLITE_DEFAULT_AUTOVACUUM
  #define SQLITE_DEFAULT_AUTOVACUUM 0
#endif

#define BTREE_AUTOVACUUM_NONE 0        /* Do not do auto-vacuum */
#define BTREE_AUTOVACUUM_FULL 1        /* Do full auto-vacuum */
#define BTREE_AUTOVACUUM_INCR 2        /* Incremental vacuum */

/*
** Forward declarations of structure
*/
typedef struct Btree Btree;
typedef struct BtCursor BtCursor;
typedef struct BtShared BtShared;
typedef struct BtreePayload BtreePayload;


int sqlite3BtreeOpen(
  sqlite3_vfs *pVfs,       /* VFS to use with this b-tree */
  const char *zFilename,   /* Name of database file to open */
  sqlite3 *db,             /* Associated database connection */
  Btree **ppBtree,         /* Return open Btree* here */
  int flags,               /* Flags */
  int vfsFlags             /* Flags passed through to VFS open */
);

/* The flags parameter to sqlite3BtreeOpen can be the bitwise or of the
** following values.
**
** NOTE:  These values must match the corresponding PAGER_ values in
** pager.h.
*/
#define BTREE_OMIT_JOURNAL  1  /* Do not create or use a rollback journal */
#define BTREE_MEMORY        2  /* This is an in-memory DB */
#define BTREE_SINGLE        4  /* The file contains at most 1 b-tree */
#define BTREE_UNORDERED     8  /* Use of a hash implementation is OK */

int sqlite3BtreeClose(Btree*);
int sqlite3BtreeSetCacheSize(Btree*,int);
int sqlite3BtreeSetSpillSize(Btree*,int);
#if SQLITE_MAX_MMAP_SIZE>0
  int sqlite3BtreeSetMmapLimit(Btree*,sqlite3_int64);
#endif
int sqlite3BtreeSetPagerFlags(Btree*,unsigned);
int sqlite3BtreeSetPageSize(Btree *p, int nPagesize, int nReserve, int eFix);
int sqlite3BtreeGetPageSize(Btree*);
Pgno sqlite3BtreeMaxPageCount(Btree*,Pgno);
Pgno sqlite3BtreeLastPage(Btree*);
int sqlite3BtreeSecureDelete(Btree*,int);
int sqlite3BtreeGetRequestedReserve(Btree*);
int sqlite3BtreeGetReserveNoMutex(Btree *p);
int sqlite3BtreeSetAutoVacuum(Btree *, int);
int sqlite3BtreeGetAutoVacuum(Btree *);
int sqlite3BtreeBeginTrans(Btree*,int,int*);
int sqlite3BtreeCommitPhaseOne(Btree*, const char*);
int sqlite3BtreeCommitPhaseTwo(Btree*, int);
int sqlite3BtreeCommit(Btree*);
int sqlite3BtreeRollback(Btree*,int,int);
int sqlite3BtreeBeginStmt(Btree*,int);
int sqlite3BtreeCreateTable(Btree*, Pgno*, int flags);
int sqlite3BtreeTxnState(Btree*);
int sqlite3BtreeIsInBackup(Btree*);

void *sqlite3BtreeSchema(Btree *, int, void(*)(void *));
int sqlite3BtreeSchemaLocked(Btree *pBtree);
#ifndef SQLITE_OMIT_SHARED_CACHE
int sqlite3BtreeLockTable(Btree *pBtree, int iTab, u8 isWriteLock);
#endif

/* Savepoints are named, nestable SQL transactions mostly implemented */ 
/* in vdbe.c and pager.c See https://sqlite.org/lang_savepoint.html */
int sqlite3BtreeSavepoint(Btree *, int, int);

/* "Checkpoint" only refers to WAL. See https://sqlite.org/wal.html#ckpt */
#ifndef SQLITE_OMIT_WAL
  int sqlite3BtreeCheckpoint(Btree*, int, int *, int *);  
#endif

const char *sqlite3BtreeGetFilename(Btree *);
const char *sqlite3BtreeGetJournalname(Btree *);
int sqlite3BtreeCopyFile(Btree *, Btree *);

int sqlite3BtreeIncrVacuum(Btree *);

/* The flags parameter to sqlite3BtreeCreateTable can be the bitwise OR
** of the flags shown below.
**
** Every SQLite table must have either BTREE_INTKEY or BTREE_BLOBKEY set.
** With BTREE_INTKEY, the table key is a 64-bit integer and arbitrary data
** is stored in the leaves.  (BTREE_INTKEY is used for SQL tables.)  With
** BTREE_BLOBKEY, the key is an arbitrary BLOB and no content is stored
** anywhere - the key is the content.  (BTREE_BLOBKEY is used for SQL
** indices.)
*/
#define BTREE_INTKEY     1    /* Table has only 64-bit signed integer keys */
#define BTREE_BLOBKEY    2    /* Table has keys only - no data */

int sqlite3BtreeDropTable(Btree*, int, int*);
int sqlite3BtreeClearTable(Btree*, int, i64*);
int sqlite3BtreeClearTableOfCursor(BtCursor*);
int sqlite3BtreeTripAllCursors(Btree*, int, int);

void sqlite3BtreeGetMeta(Btree *pBtree, int idx, u32 *pValue);
int sqlite3BtreeUpdateMeta(Btree*, int idx, u32 value);

int sqlite3BtreeNewDb(Btree *p);

/*
** The second parameter to sqlite3BtreeGetMeta or sqlite3BtreeUpdateMeta
** should be one of the following values. The integer values are assigned 
** to constants so that the offset of the corresponding field in an
** SQLite database header may be found using the following formula:
**
**   offset = 36 + (idx * 4)
**
** For example, the free-page-count field is located at byte offset 36 of
** the database file header. The incr-vacuum-flag field is located at
** byte offset 64 (== 36+4*7).
**
** The BTREE_DATA_VERSION value is not really a value stored in the header.
** It is a read-only number computed by the pager.  But we merge it with
** the header value access routines since its access pattern is the same.
** Call it a "virtual meta value".
*/
#define BTREE_FREE_PAGE_COUNT     0
#define BTREE_SCHEMA_VERSION      1
#define BTREE_FILE_FORMAT         2
#define BTREE_DEFAULT_CACHE_SIZE  3
#define BTREE_LARGEST_ROOT_PAGE   4
#define BTREE_TEXT_ENCODING       5
#define BTREE_USER_VERSION        6
#define BTREE_INCR_VACUUM         7
#define BTREE_APPLICATION_ID      8
#define BTREE_DATA_VERSION        15  /* A virtual meta-value */

/*
** Kinds of hints that can be passed into the sqlite3BtreeCursorHint()
** interface.
**
** BTREE_HINT_RANGE  (arguments: Expr*, Mem*)
**
**     The first argument is an Expr* (which is guaranteed to be constant for
**     the lifetime of the cursor) that defines constraints on which rows
**     might be fetched with this cursor.  The Expr* tree may contain
**     TK_REGISTER nodes that refer to values stored in the array of registers
**     passed as the second parameter.  In other words, if Expr.op==TK_REGISTER
**     then the value of the node is the value in Mem[pExpr.iTable].  Any
**     TK_COLUMN node in the expression tree refers to the Expr.iColumn-th
**     column of the b-tree of the cursor.  The Expr tree will not contain
**     any function calls nor subqueries nor references to b-trees other than
**     the cursor being hinted.
**
**     The design of the _RANGE hint is aid b-tree implementations that try
**     to prefetch content from remote machines - to provide those
**     implementations with limits on what needs to be prefetched and thereby
**     reduce network bandwidth.
**
** Note that BTREE_HINT_FLAGS with BTREE_BULKLOAD is the only hint used by
** standard SQLite.  The other hints are provided for extentions that use
** the SQLite parser and code generator but substitute their own storage
** engine.
*/
#define BTREE_HINT_RANGE 0       /* Range constraints on queries */

/*
** Values that may be OR'd together to form the argument to the
** BTREE_HINT_FLAGS hint for sqlite3BtreeCursorHint():
**
** The BTREE_BULKLOAD flag is set on index cursors when the index is going
** to be filled with content that is already in sorted order.
**
** The BTREE_SEEK_EQ flag is set on cursors that will get OP_SeekGE or
** OP_SeekLE opcodes for a range search, but where the range of entries
** selected will all have the same key.  In other words, the cursor will
** be used only for equality key searches.
**
*/
#define BTREE_BULKLOAD 0x00000001  /* Used to full index in sorted order */
#define BTREE_SEEK_EQ  0x00000002  /* EQ seeks only - no range seeks */

/* 
** Flags passed as the third argument to sqlite3BtreeCursor().
**
** For read-only cursors the wrFlag argument is always zero. For read-write
** cursors it may be set to either (BTREE_WRCSR|BTREE_FORDELETE) or just
** (BTREE_WRCSR). If the BTREE_FORDELETE bit is set, then the cursor will
** only be used by SQLite for the following:
**
**   * to seek to and then delete specific entries, and/or
**
**   * to read values that will be used to create keys that other
**     BTREE_FORDELETE cursors will seek to and delete.
**
** The BTREE_FORDELETE flag is an optimization hint.  It is not used by
** by this, the native b-tree engine of SQLite, but it is available to
** alternative storage engines that might be substituted in place of this
** b-tree system.  For alternative storage engines in which a delete of
** the main table row automatically deletes corresponding index rows,
** the FORDELETE flag hint allows those alternative storage engines to
** skip a lot of work.  Namely:  FORDELETE cursors may treat all SEEK
** and DELETE operations as no-ops, and any READ operation against a
** FORDELETE cursor may return a null row: 0x01 0x00.
*/
#define BTREE_WRCSR     0x00000004     /* read-write cursor */
#define BTREE_FORDELETE 0x00000008     /* Cursor is for seek/delete only */

int sqlite3BtreeCursor(
  Btree*,                              /* BTree containing table to open */
  Pgno iTable,                         /* Index of root page */
  int wrFlag,                          /* 1 for writing.  0 for read-only */
  struct KeyInfo*,                     /* First argument to compare function */
  BtCursor *pCursor                    /* Space to write cursor structure */
);
BtCursor *sqlite3BtreeFakeValidCursor(void);
int sqlite3BtreeCursorSize(void);
void sqlite3BtreeCursorZero(BtCursor*);
void sqlite3BtreeCursorHintFlags(BtCursor*, unsigned);
#ifdef SQLITE_ENABLE_CURSOR_HINTS
void sqlite3BtreeCursorHint(BtCursor*, int, ...);
#endif

int sqlite3BtreeCloseCursor(BtCursor*);
int sqlite3BtreeTableMoveto(
  BtCursor*,
  i64 intKey,
  int bias,
  int *pRes
);
int sqlite3BtreeIndexMoveto(
  BtCursor*,
  UnpackedRecord *pUnKey,
  int *pRes
);
int sqlite3BtreeCursorHasMoved(BtCursor*);
int sqlite3BtreeCursorRestore(BtCursor*, int*);
int sqlite3BtreeDelete(BtCursor*, u8 flags);

/* Allowed flags for sqlite3BtreeDelete() and sqlite3BtreeInsert() */
#define BTREE_SAVEPOSITION 0x02  /* Leave cursor pointing at NEXT or PREV */
#define BTREE_AUXDELETE    0x04  /* not the primary delete operation */
#define BTREE_APPEND       0x08  /* Insert is likely an append */
#define BTREE_PREFORMAT    0x80  /* Inserted data is a preformated cell */

/* An instance of the BtreePayload object describes the content of a single
** entry in either an index or table btree.
**
** Index btrees (used for indexes and also WITHOUT ROWID tables) contain
** an arbitrary key and no data.  These btrees have pKey,nKey set to the
** key and the pData,nData,nZero fields are uninitialized.  The aMem,nMem
** fields give an array of Mem objects that are a decomposition of the key.
** The nMem field might be zero, indicating that no decomposition is available.
**
** Table btrees (used for rowid tables) contain an integer rowid used as
** the key and passed in the nKey field.  The pKey field is zero.  
** pData,nData hold the content of the new entry.  nZero extra zero bytes
** are appended to the end of the content when constructing the entry.
** The aMem,nMem fields are uninitialized for table btrees.
**
** Field usage summary:
**
**               Table BTrees                   Index Btrees
**
**   pKey        always NULL                    encoded key
**   nKey        the ROWID                      length of pKey
**   pData       data                           not used
**   aMem        not used                       decomposed key value
**   nMem        not used                       entries in aMem
**   nData       length of pData                not used
**   nZero       extra zeros after pData        not used
**
** This object is used to pass information into sqlite3BtreeInsert().  The
** same information used to be passed as five separate parameters.  But placing
** the information into this object helps to keep the interface more 
** organized and understandable, and it also helps the resulting code to
** run a little faster by using fewer registers for parameter passing.
*/
struct BtreePayload {
  const void *pKey;       /* Key content for indexes.  NULL for tables */
  sqlite3_int64 nKey;     /* Size of pKey for indexes.  PRIMARY KEY for tabs */
  const void *pData;      /* Data for tables. */
  sqlite3_value *aMem;    /* First of nMem value in the unpacked pKey */
  u16 nMem;               /* Number of aMem[] value.  Might be zero */
  int nData;              /* Size of pData.  0 if none. */
  int nZero;              /* Extra zero data appended after pData,nData */
};

int sqlite3BtreeInsert(BtCursor*, const BtreePayload *pPayload,
                       int flags, int seekResult);
int sqlite3BtreeFirst(BtCursor*, int *pRes);
int sqlite3BtreeLast(BtCursor*, int *pRes);
int sqlite3BtreeNext(BtCursor*, int flags);
int sqlite3BtreeEof(BtCursor*);
int sqlite3BtreePrevious(BtCursor*, int flags);
i64 sqlite3BtreeIntegerKey(BtCursor*);
void sqlite3BtreeCursorPin(BtCursor*);
void sqlite3BtreeCursorUnpin(BtCursor*);
#ifdef SQLITE_ENABLE_OFFSET_SQL_FUNC
i64 sqlite3BtreeOffset(BtCursor*);
#endif
int sqlite3BtreePayload(BtCursor*, u32 offset, u32 amt, void*);
const void *sqlite3BtreePayloadFetch(BtCursor*, u32 *pAmt);
u32 sqlite3BtreePayloadSize(BtCursor*);
sqlite3_int64 sqlite3BtreeMaxRecordSize(BtCursor*);

char *sqlite3BtreeIntegrityCheck(sqlite3*,Btree*,Pgno*aRoot,int nRoot,int,int*);
struct Pager *sqlite3BtreePager(Btree*);
i64 sqlite3BtreeRowCountEst(BtCursor*);

#ifndef SQLITE_OMIT_INCRBLOB
int sqlite3BtreePayloadChecked(BtCursor*, u32 offset, u32 amt, void*);
int sqlite3BtreePutData(BtCursor*, u32 offset, u32 amt, void*);
void sqlite3BtreeIncrblobCursor(BtCursor *);
#endif
void sqlite3BtreeClearCursor(BtCursor *);
int sqlite3BtreeSetVersion(Btree *pBt, int iVersion);
int sqlite3BtreeCursorHasHint(BtCursor*, unsigned int mask);
int sqlite3BtreeIsReadonly(Btree *pBt);
int sqlite3HeaderSizeBtree(void);

#ifdef SQLITE_DEBUG
sqlite3_uint64 sqlite3BtreeSeekCount(Btree*);
#else
# define sqlite3BtreeSeekCount(X) 0
#endif

#ifndef NDEBUG
int sqlite3BtreeCursorIsValid(BtCursor*);
#endif
int sqlite3BtreeCursorIsValidNN(BtCursor*);

int sqlite3BtreeCount(sqlite3*, BtCursor*, i64*);

#ifdef SQLITE_TEST
int sqlite3BtreeCursorInfo(BtCursor*, int*, int);
void sqlite3BtreeCursorList(Btree*);
#endif

#ifndef SQLITE_OMIT_WAL
  int sqlite3BtreeCheckpoint(Btree*, int, int *, int *);
#endif

int sqlite3BtreeTransferRow(BtCursor*, BtCursor*, i64);

void sqlite3BtreeClearCache(Btree*);

/*
** If we are not using shared cache, then there is no need to
** use mutexes to access the BtShared structures.  So make the
** Enter and Leave procedures no-ops.
*/
#ifndef SQLITE_OMIT_SHARED_CACHE
  void sqlite3BtreeEnter(Btree*);
  void sqlite3BtreeEnterAll(sqlite3*);
  int sqlite3BtreeSharable(Btree*);
  void sqlite3BtreeEnterCursor(BtCursor*);
  int sqlite3BtreeConnectionCount(Btree*);
#else
# define sqlite3BtreeEnter(X) 
# define sqlite3BtreeEnterAll(X)
# define sqlite3BtreeSharable(X) 0
# define sqlite3BtreeEnterCursor(X)
# define sqlite3BtreeConnectionCount(X) 1
#endif

#if !defined(SQLITE_OMIT_SHARED_CACHE) && SQLITE_THREADSAFE
  void sqlite3BtreeLeave(Btree*);
  void sqlite3BtreeLeaveCursor(BtCursor*);
  void sqlite3BtreeLeaveAll(sqlite3*);
#ifndef NDEBUG
  /* These routines are used inside assert() statements only. */
  int sqlite3BtreeHoldsMutex(Btree*);
  int sqlite3BtreeHoldsAllMutexes(sqlite3*);
  int sqlite3SchemaMutexHeld(sqlite3*,int,Schema*);
#endif
#else

# define sqlite3BtreeLeave(X)
# define sqlite3BtreeLeaveCursor(X)
# define sqlite3BtreeLeaveAll(X)

# define sqlite3BtreeHoldsMutex(X) 1
# define sqlite3BtreeHoldsAllMutexes(X) 1
# define sqlite3SchemaMutexHeld(X,Y,Z) 1
#endif


#endif /* SQLITE_BTREE_H */
`, cu},
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

		var nn Node
		nn = n.(Node)
		fmt.Println(nn)

		//m := check.Must1(msh.Marshal(&nn))
		//fmt.Println(check.Must1(ju.MarshalPretty(m)))
		//
		//var nn2 Node
		//check.Must(msh.Unmarshal(m, &nn2))
		//fmt.Println(check.Must1(ju.MarshalPretty(check.Must1(msh.Marshal(&nn2)))))

		fmt.Println()
	}
}
