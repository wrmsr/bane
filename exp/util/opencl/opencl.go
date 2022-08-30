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

type OpenClError struct {
	Code int
}

func (e OpenClError) Error() string {
	return fmt.Sprintf("opencl error: %d", e.Code)
}

func toErr(code C.cl_int) error {
	if code == 0 {
		return nil
	}
	return OpenClError{Code: int(code)}
}

func GetPlatforms() {
	var n C.cl_uint
	check.Must(toErr(C.clGetPlatformIDs(0, nil, &n)))

	platformIds := make([]C.cl_platform_id, int(n))
	check.Must(toErr(C.clGetPlatformIDs(n, &platformIds[0], nil)))

	//fmt.Println(platformIds)
	//var devices []*Device
	//for _, p := range platformIds {
	//	var n C.cl_uint
	//	err = toErr(C.clGetDeviceIDs(p, C.cl_device_type(deviceType), 0, nil, &n))
	//	if err != nil {
	//		return nil, err
	//	}
	//	deviceIds := make([]C.cl_device_id, int(n))
	//	err = toErr(C.clGetDeviceIDs(p, C.cl_device_type(deviceType), n, &deviceIds[0], nil))
	//	if err != nil {
	//		return nil, err
	//	}
	//	for _, d := range deviceIds {
	//		device, err := newDevice(d)
	//		if err != nil {
	//			return nil, err
	//		}
	//		devices = append(devices, device)
	//	}
	//}
}
