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
package ffi

// #cgo LDFLAGS: -lffi
//
// #include <ffi.h>
//
// extern void GoClosureCallback(ffi_cif *, void *, void **, void *);
//
// typedef void (*closure)(ffi_cif *, void *, void **, void *);
import "C"
import "unsafe"

func constructClosure(fn *function) (err error) {
	var closure C.ffi_closure
	var fptr unsafe.Pointer
	var mptr unsafe.Pointer

	if mptr, err = C.ffi_closure_alloc(C.size_t(unsafe.Sizeof(closure)), &fptr); mptr == nil {
		return
	}

	if status := Status(C.ffi_prep_closure_loc((*C.ffi_closure)(mptr), &fn.Interface.ffi_cif, C.closure(C.GoClosureCallback), unsafe.Pointer(fn), fptr)); status != OK {
		C.ffi_closure_free(mptr)
		err = status
		return
	}

	fn.fptr = fptr
	fn.mptr = mptr
	return nil
}

func destroyClosure(fn *function) {
	C.ffi_closure_free(fn.mptr)
}
