.PHONY: all
all: go tinygo

.PHONY: go
go:
	GOOS=js GOARCH=wasm go build -o main.wasm ./main.go
	~/src/WebAssembly/wabt/build/wasm2wat main.wasm > main.wat
	~/src/WebAssembly/wabt/build/wasm2c main.wasm > main.c

.PHONY: tinygo
tinygo:
	tinygo build -o main_tiny.wasm -target wasm ./main.go
	~/src/WebAssembly/wabt/build/wasm2wat main_tiny.wasm > main_tiny.wat
