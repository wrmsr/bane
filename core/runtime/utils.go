package runtime

import (
	"reflect"
	"runtime"
)

func MaxProcs() int {
	return runtime.GOMAXPROCS(0)
}

func GetFuncName(v reflect.Value) string {
	return runtime.FuncForPC(v.Pointer()).Name()
}
