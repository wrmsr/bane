/*
MIT License

# Copyright (c) 2023 blast-go

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
package openblas

/*
#include "cblas.h"
#include <stdlib.h>
*/
import "C"

const (
	Sequential = 0
	Thread     = 1
	OpenMP     = 2
)

type Order int32

const (
	Rowmajor Order = C.CblasRowMajor
	Colmajor Order = C.CblasColMajor
)

type Transpose int32

const (
	Notrans     Transpose = C.CblasNoTrans
	Trans       Transpose = C.CblasTrans
	Conjtrans   Transpose = C.CblasConjTrans
	Conjnotrans Transpose = C.CblasConjNoTrans
)

type UpLo int32

const (
	Upper UpLo = C.CblasUpper
	Lower UpLo = C.CblasLower
)

type Diag int32

const (
	NonUnit Diag = C.CblasNonUnit
	Unit    Diag = C.CblasUnit
)

type Side int32

const (
	Left  Side = C.CblasLeft
	Right Side = C.CblasRight
)

type Layout int32

type Bfloat16 uint16
