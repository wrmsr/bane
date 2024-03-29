//go:build darwin || linux

/*
The MIT License (MIT)

# Copyright (c) 2016 Achille

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package dl

/*
#include <dlfcn.h>
#include <stdlib.h>
*/
import "C"

import (
	"runtime"
	"sync"
	"syscall"
	"unsafe"
)

type dylib struct {
	mutex  sync.RWMutex
	handle unsafe.Pointer
}

func open(path string, mode Mode) (lib *dylib, err error) {
	var handle unsafe.Pointer
	var flags int

	if (mode&Lazy) == 0 && (mode&Now) == 0 {
		mode |= Now
	}

	if (mode & Lazy) != 0 {
		flags |= C.RTLD_LAZY
	}

	if (mode & Now) != 0 {
		flags |= C.RTLD_NOW
	}

	if (mode & Global) != 0 {
		flags |= C.RTLD_GLOBAL
	}

	if (mode & Local) != 0 {
		flags |= C.RTLD_LOCAL
	}

	if handle, err = dlopen(path, flags); err != nil {
		return
	}

	lib = &dylib{
		handle: handle,
	}

	runtime.SetFinalizer(lib, (*dylib).close)
	return
}

func (lib *dylib) Close() (err error) {
	lib.mutex.Lock()
	defer lib.mutex.Unlock()
	return lib.close()
}

func (lib *dylib) Symbol(name string) (addr uintptr, err error) {
	var handle unsafe.Pointer
	var sym unsafe.Pointer

	lib.mutex.RLock()
	defer lib.mutex.RUnlock()

	if handle = lib.handle; handle == nil {
		err = syscall.EINVAL
		return
	}

	if sym, err = dlsym(handle, name); err != nil {
		return
	}

	addr = uintptr(sym)
	return
}

func (lib *dylib) close() (err error) {
	var handle unsafe.Pointer

	if handle, lib.handle = lib.handle, nil; handle != nil {
		err = dlclose(handle)
	} else {
		err = syscall.EINVAL
	}

	return
}

var dlmtx sync.Mutex

func dlopen(path string, flags int) (lib unsafe.Pointer, err error) {
	var f = C.int(flags)
	var s *C.char

	if len(path) != 0 {
		s = C.CString(path)
		defer C.free(unsafe.Pointer(s))
	}

	dlmtx.Lock()
	defer dlmtx.Unlock()

	if lib = C.dlopen(s, f); lib == nil {
		err = dlerror()
	}

	return
}

func dlclose(lib unsafe.Pointer) (err error) {
	dlmtx.Lock()
	defer dlmtx.Unlock()

	if C.dlclose(lib) != 0 {
		err = dlerror()
	}

	return
}

func dlsym(lib unsafe.Pointer, name string) (addr unsafe.Pointer, err error) {
	var s = C.CString(name)
	defer C.free(unsafe.Pointer(s))

	dlmtx.Lock()
	defer dlmtx.Unlock()

	if addr = C.dlsym(lib, s); addr == nil {
		err = dlerror()
	}

	return
}

func dlerror() error {
	return &Error{
		Message: C.GoString(C.dlerror()),
	}
}
