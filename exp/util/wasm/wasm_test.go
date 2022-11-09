/*
https://webassembly.github.io/spec/core/_download/WebAssembly.pdf

##

	clang++ \
		--target=wasm32 \
		-nostdlib \
		-O3 \
		-Wl,--no-entry \
		-Wl,--export-all \
		-Wl,--lto-O3 \
		-Wl,--allow-undefined \
		-Wl,--import-memory \
		-o sqlite3.wasm \
		c/sqlite/sqlite3.c

==

https://github.com/WebAssembly/wasi-sdk/releases/download/wasi-sdk-16/wasi-sdk-16.0-linux.tar.gz
https://github.com/WebAssembly/wasi-sdk/releases/download/wasi-sdk-16/wasi-sysroot-16.0.tar.gz
https://github.com/WebAssembly/wabt/releases/download/1.0.30/wabt-1.0.30-ubuntu.tar.gz

wasi-sdk-16.0/bin/clang sqlite-amalgamation-3390400/sqlite3.c --sysroot wasi-sysroot -o sqlite3.wasm -DLONGDOUBLE_TYPE=double -DSQLITE_THREADSAFE=0 -D_WASI_EMULATED_MMAN

==

docker run --rm -v $(pwd):/src -u $(id -u):$(id -g) emscripten/emsdk emcc /src/c/sqlite/sqlite3.c -o sqlite3.wat

==

cd wasm-install

	bin/clang \
		-I sysroot/include \
		-emit-llvm \
		--target=wasm32 \
		-Oz \
		-c bin/emscripten/tests/sqlite/sqlite3.c \
		-o bin/emscripten/tests/sqlite/sqlite3.bc \
		-DSQLITE_DISABLE_LFS \
		-DLONGDOUBLE_TYPE=double \
		-DSQLITE_THREADSAFE=0

bin/llvm-dis bin/emscripten/tests/sqlite/sqlite3.bc
bin/llc -asm-verbose=false bin/emscripten/tests/sqlite/sqlite3.bc -o bin/emscripten/tests/sqlite/sqlite3.s
bin/s2wasm bin/emscripten/tests/sqlite/sqlite3.s > bin/emscripten/tests/sqlite/sqlite3.wast

==

wasm-install/bin/clang -I wasm-install/sysroot/include --target=wasm32 -emit-llvm -c a.c -o a.bc
wasm-install/bin/llvm-dis a.bc
wasm-install/bin/llc -asm-verbose=false a.bc -o a.s
wasm-install/bin/s2wasm a.s > a.wast

==

wasm-install/bin/clang -I wasm-install/sysroot/include --target=wasm32 -emit-llvm -O0 -fno-inline -c a.c -o a.bc && wasm-install/bin/llvm-dis a.bc && wasm-install/bin/llc -asm-verbose=false a.bc -o a.s && wasm-install/bin/s2wasm a.s > a.wast

==

git clone https://github.com/WebAssembly/musl/
cd musl
git checkout wasm-prototype-1
./libc.py --clang_dir /root/wasm-install/bin --binaryen_dir /root/wasm-install/bin --sexpr_wasm /root/wasm-install/bin/sexpr-wasm --musl /root/musl
x86_64-linux-gnu/asm/unistd_32.h

==

https://aransentin.github.io/cwasm/

-nostartfiles -nostdlib

	docker run --rm -v $(pwd):/src -u $(id -u):$(id -g) emscripten/emsdk emcc \
			--no-entry \
			-flto \
	        -O3 \
			-Wl,--import-memory \
	        -Wl,--no-entry \
	        -Wl,--export-all \
	        -Wl,--lto-O3 \
	        -Wl,--allow-undefined \
	        -o sqlite3.wasm \
	        /src/c/sqlite/sqlite3.c

TODO:
  - https://github.com/WebAssembly/testsuite
*/
package wasm

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	wp "github.com/wrmsr/bane/exp/util/wasm/parsing"
	wr "github.com/wrmsr/bane/exp/util/wasm/rendering"
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
)

func testWasm(src string) {
	r := wp.NewReader(bufio.NewReader(strings.NewReader(src)))
	root := check.NotNil(r.ReadElement()).(wp.List)
	check.Nil(r.ReadElement())

	fmt.Println(wr.RenderString(root))

	BuildModule(root)
}

func TestWasm(t *testing.T) {
	src := `
(module
  (memory 4096 4096
    (segment 1026 "\14\00")
  )
  (func $big_negative
    (local $temp f64)
    (block $block0
      (set_local $temp
        (f64.const -2147483648)
      )
      (set_local $temp
        (f64.const -2147483648)
      )
      (set_local $temp
        (f64.const -21474836480)
      )
      (set_local $temp
        (f64.const 0.039625)
      )
      (set_local $temp
        (f64.const -0.039625)
      )
    )
  )
)
`

	testWasm(src)
}

func TestWasmBig(t *testing.T) {
	src := string(check.Must1(os.ReadFile(filepath.Join(paths.FindProjectRoot(), "exp", "util", "wasm", "test", "boilerplate.wast"))))
	testWasm(src)
}
