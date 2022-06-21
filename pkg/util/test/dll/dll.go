/*
go build -buildmode c-shared -tags nodev -o bin/libfoo.so "github.com/wrmsr/bane/pkg/util/test/dll"

import ctypes as ct
l = ct.CDLL('bin/libfoo.so')
l.cgoRandom.restype = ct.c_int
l.cgoRandom.argtypes = [ct.c_int]
l.cgoRandom(10)
*/
package main

import "C"
import (
	"math/rand"
	"time"
)

//export cgoCurrentMillis
func cgoCurrentMillis() C.long {
	return C.long(time.Now().Unix())
}

//export cgoSeed
func cgoSeed(m C.long) {
	rand.Seed(int64(m))
}

//export cgoRandom
func cgoRandom(m C.int) C.int {
	return C.int(rand.Intn(int(m)))
}

func main() {}
