package opencl

/*
#cgo CFLAGS: -I CL
#cgo !darwin LDFLAGS: -lOpenCL
#cgo darwin LDFLAGS: -framework OpenCL

#ifdef __APPLE__
#include <OpenCL/opencl.h>
#else
#include <CL/cl.h>
#endif
*/
import "C"

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
)

func GetPlatforms() {
	var n C.cl_uint
	check.Must(toErr(C.clGetPlatformIDs(0, nil, &n)))

	platformIds := make([]C.cl_platform_id, int(n))
	check.Must(toErr(C.clGetPlatformIDs(n, &platformIds[0], nil)))

	for _, platformId := range platformIds {
		const maxDeviceCount = 64
		var deviceIds [maxDeviceCount]C.cl_device_id
		var numDevices C.cl_uint
		check.Must(toErr(C.clGetDeviceIDs(
			platformId,
			C.cl_device_type(C.CL_DEVICE_TYPE_ALL),
			C.cl_uint(maxDeviceCount),
			&deviceIds[0],
			&numDevices,
		)))
		if numDevices > maxDeviceCount {
			numDevices = maxDeviceCount
		}

		var err C.cl_int
		clContext := C.clCreateContext(
			nil,
			C.cl_uint(len(deviceIds)),
			&deviceIds[0],
			nil,
			nil,
			&err,
		)
		check.Must(toErr(err))

		fmt.Println(clContext)
	}
}
