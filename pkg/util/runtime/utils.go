package runtime

import "runtime"

func MaxProcs() int {
	return runtime.GOMAXPROCS(0)
}
