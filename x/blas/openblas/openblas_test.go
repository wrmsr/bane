//
/*
https://github.com/xianyi/OpenBLAS/wiki/Developer-manual#optimizing-gemm-for-a-given-hardware
https://www.cs.utexas.edu/~flame/web/FLAMEPublications.html
https://www.cs.utexas.edu/~flame/pubs/GotoTOMS_revision.pdf

000000010008f100 <_github.com/wrmsr/bane/x/blas/openblas._Cfunc_cblas_sgemm.abi0>:
10008f100: 90 0b 40 f9  ldr     x16, [x28, #16]
10008f104: ff 63 30 eb  cmp     sp, x16
10008f108: 09 0c 00 54  b.ls    0x10008f288 <_github.com/wrmsr/bane/x/blas/openblas._Cfunc_cblas_sgemm.abi0+0x188>
10008f10c: fe 0f 1d f8  str     x30, [sp, #-48]!
10008f110: fd 83 1f f8  stur    x29, [sp, #-8]
10008f114: fd 23 00 d1  sub     x29, sp, #8
10008f118: e2 e3 00 91  add     x2, sp, #56
10008f11c: e2 13 00 f9  str     x2, [sp, #32]
10008f120: 1b 20 00 b0  adrp    x27, 0x100490000 <__cgo_ce7779e6fc99_Cfunc_cblas_dnrm2+0x4>
10008f124: 60 97 44 f9  ldr     x0, [x27, #2344]
10008f128: e1 13 40 f9  ldr     x1, [sp, #32]
10008f12c: d1 e1 fd 97  bl      0x100007870 <_runtime.cgocall>
10008f130: 1b 22 00 b0  adrp    x27, 0x1004d0000 <__cgo_ce7779e6fc99_Cfunc_cblas_dsbmv+0x38>
10008f134: 62 13 78 39  ldrb    w2, [x27, #3588]
10008f138: 22 0a 00 36  tbz     w2, #0, 0x10008f27c <_github.com/wrmsr/bane/x/blas/openblas._Cfunc_cblas_sgemm.abi0+0x17c>
10008f13c: e0 3b 40 b9  ldr     w0, [sp, #56]
10008f140: a4 f9 fd 97  bl      0x10000d7d0 <_runtime.convT32>
10008f144: e1 03 00 aa  mov     x1, x0
10008f148: c0 1b 00 b0  adrp    x0, 0x100408000 <__cgo_ce7779e6fc99_Cfunc_cblas_dcopy+0x14>
10008f14c: 00 80 12 91  add     x0, x0, #1184
10008f150: b8 e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f154: e0 3f 40 b9  ldr     w0, [sp, #60]
10008f158: 9e f9 fd 97  bl      0x10000d7d0 <_runtime.convT32>
10008f15c: e1 03 00 aa  mov     x1, x0
10008f160: c0 1b 00 b0  adrp    x0, 0x100408000 <__cgo_ce7779e6fc99_Cfunc_cblas_ddot+0x10>
10008f164: 00 80 12 91  add     x0, x0, #1184
10008f168: b2 e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f16c: e0 43 40 b9  ldr     w0, [sp, #64]
10008f170: 98 f9 fd 97  bl      0x10000d7d0 <_runtime.convT32>
10008f174: e1 03 00 aa  mov     x1, x0
10008f178: c0 1b 00 b0  adrp    x0, 0x100408000 <__cgo_ce7779e6fc99_Cfunc_cblas_ddot+0x28>
10008f17c: 00 80 12 91  add     x0, x0, #1184
10008f180: ac e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f184: e0 47 80 b9  ldrsw   x0, [sp, #68]
10008f188: 92 f9 fd 97  bl      0x10000d7d0 <_runtime.convT32>
10008f18c: e1 03 00 aa  mov     x1, x0
10008f190: c0 1b 00 90  adrp    x0, 0x100407000 <__cgo_ce7779e6fc99_Cfunc_cblas_ddot+0x3c>
10008f194: 00 80 34 91  add     x0, x0, #3360
10008f198: a6 e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f19c: e0 4b 80 b9  ldrsw   x0, [sp, #72]
10008f1a0: 8c f9 fd 97  bl      0x10000d7d0 <_runtime.convT32>
10008f1a4: e1 03 00 aa  mov     x1, x0
10008f1a8: c0 1b 00 90  adrp    x0, 0x100407000 <__cgo_ce7779e6fc99_Cfunc_cblas_ddot+0x54>
10008f1ac: 00 80 34 91  add     x0, x0, #3360
10008f1b0: a0 e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f1b4: e0 4f 80 b9  ldrsw   x0, [sp, #76]
10008f1b8: 86 f9 fd 97  bl      0x10000d7d0 <_runtime.convT32>
10008f1bc: e1 03 00 aa  mov     x1, x0
10008f1c0: c0 1b 00 90  adrp    x0, 0x100407000 <__cgo_ce7779e6fc99_Cfunc_cblas_dgbmv+0x14>
10008f1c4: 00 80 34 91  add     x0, x0, #3360
10008f1c8: 9a e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f1cc: e0 53 40 bd  ldr     s0, [sp, #80]
10008f1d0: 00 00 26 1e  fmov    w0, s0
10008f1d4: 7f f9 fd 97  bl      0x10000d7d0 <_runtime.convT32>
10008f1d8: e1 03 00 aa  mov     x1, x0
10008f1dc: c0 1b 00 90  adrp    x0, 0x100407000 <__cgo_ce7779e6fc99_Cfunc_cblas_dgbmv+0x30>
10008f1e0: 00 80 33 91  add     x0, x0, #3296
10008f1e4: 93 e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f1e8: a0 1b 00 d0  adrp    x0, 0x100405000 <__cgo_ce7779e6fc99_Cfunc_cblas_dgbmv+0x34>
10008f1ec: 00 80 02 91  add     x0, x0, #160
10008f1f0: e1 2f 40 f9  ldr     x1, [sp, #88]
10008f1f4: 8f e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f1f8: e0 63 80 b9  ldrsw   x0, [sp, #96]
10008f1fc: 75 f9 fd 97  bl      0x10000d7d0 <_runtime.convT32>
10008f200: e1 03 00 aa  mov     x1, x0
10008f204: c0 1b 00 90  adrp    x0, 0x100407000 <__cgo_ce7779e6fc99_Cfunc_cblas_dgbmv+0x58>
10008f208: 00 80 34 91  add     x0, x0, #3360
10008f20c: 89 e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f210: a0 1b 00 d0  adrp    x0, 0x100405000 <__cgo_ce7779e6fc99_Cfunc_cblas_dgeadd>
10008f214: 00 80 02 91  add     x0, x0, #160
10008f218: e1 37 40 f9  ldr     x1, [sp, #104]
10008f21c: 85 e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f220: e0 73 80 b9  ldrsw   x0, [sp, #112]
10008f224: 6b f9 fd 97  bl      0x10000d7d0 <_runtime.convT32>
10008f228: e1 03 00 aa  mov     x1, x0
10008f22c: c0 1b 00 90  adrp    x0, 0x100407000 <__cgo_ce7779e6fc99_Cfunc_cblas_dgeadd+0x24>
10008f230: 00 80 34 91  add     x0, x0, #3360
10008f234: 7f e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f238: e0 77 40 bd  ldr     s0, [sp, #116]
10008f23c: 00 00 26 1e  fmov    w0, s0
10008f240: 64 f9 fd 97  bl      0x10000d7d0 <_runtime.convT32>
10008f244: e1 03 00 aa  mov     x1, x0
10008f248: c0 1b 00 90  adrp    x0, 0x100407000 <__cgo_ce7779e6fc99_Cfunc_cblas_dgemm+0x18>
10008f24c: 00 80 33 91  add     x0, x0, #3296
10008f250: 78 e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f254: a0 1b 00 d0  adrp    x0, 0x100405000 <__cgo_ce7779e6fc99_Cfunc_cblas_dgemm+0x1c>
10008f258: 00 80 02 91  add     x0, x0, #160
10008f25c: e1 3f 40 f9  ldr     x1, [sp, #120]
10008f260: 74 e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f264: e0 83 80 b9  ldrsw   x0, [sp, #128]
10008f268: 5a f9 fd 97  bl      0x10000d7d0 <_runtime.convT32>
10008f26c: e1 03 00 aa  mov     x1, x0
10008f270: c0 1b 00 90  adrp    x0, 0x100407000 <__cgo_ce7779e6fc99_Cfunc_cblas_dgemm+0x40>
10008f274: 00 80 34 91  add     x0, x0, #3360
10008f278: 6e e1 fd 97  bl      0x100007830 <_runtime.cgoUse>
10008f27c: fd fb 7f a9  ldp     x29, x30, [sp, #-8]
10008f280: ff c3 00 91  add     sp, sp, #48
10008f284: c0 03 5f d6  ret
10008f288: e3 03 1e aa  mov     x3, x30
10008f28c: 65 3a ff 97  bl      0x10005dc20 <_runtime.morestack_noctxt.abi0>
10008f290: 9c ff ff 17  b       0x10008f100 <_github.com/wrmsr/bane/x/blas/openblas._Cfunc_cblas_sgemm.abi0>
                ...
000000010008f2a0 <_github.com/wrmsr/bane/x/blas/openblas.Sgemm>:
10008f2a0: 90 0b 40 f9  ldr     x16, [x28, #16]
10008f2a4: ff 63 30 eb  cmp     sp, x16
10008f2a8: 49 05 00 54  b.ls    0x10008f350 <_github.com/wrmsr/bane/x/blas/openblas.Sgemm+0xb0>
10008f2ac: fe 0f 1a f8  str     x30, [sp, #-96]!
10008f2b0: fd 83 1f f8  stur    x29, [sp, #-8]
10008f2b4: fd 23 00 d1  sub     x29, sp, #8
10008f2b8: e6 5b 00 f9  str     x6, [sp, #176]
10008f2bc: ea 6b 00 f9  str     x10, [sp, #208]
10008f2c0: ff 00 00 f1  cmp     x7, #0
10008f2c4: e9 03 00 54  b.ls    0x10008f340 <_github.com/wrmsr/bane/x/blas/openblas.Sgemm+0xa0>
10008f2c8: 7f 01 00 f1  cmp     x11, #0
10008f2cc: 49 03 00 54  b.ls    0x10008f334 <_github.com/wrmsr/bane/x/blas/openblas.Sgemm+0x94>
10008f2d0: e7 3b 40 f9  ldr     x7, [sp, #112]
10008f2d4: ff 00 00 f1  cmp     x7, #0
10008f2d8: 89 02 00 54  b.ls    0x10008f328 <_github.com/wrmsr/bane/x/blas/openblas.Sgemm+0x88>
10008f2dc: e0 0b 00 b9  str     w0, [sp, #8]
10008f2e0: e1 0f 00 b9  str     w1, [sp, #12]
10008f2e4: e2 13 00 b9  str     w2, [sp, #16]
10008f2e8: e3 17 00 b9  str     w3, [sp, #20]
10008f2ec: e4 1b 00 b9  str     w4, [sp, #24]
10008f2f0: e5 1f 00 b9  str     w5, [sp, #28]
10008f2f4: e0 23 00 bd  str     s0, [sp, #32]
10008f2f8: e6 17 00 f9  str     x6, [sp, #40]
10008f2fc: e9 33 00 b9  str     w9, [sp, #48]
10008f300: ea 1f 00 f9  str     x10, [sp, #56]
10008f304: ed 43 00 b9  str     w13, [sp, #64]
10008f308: e1 47 00 bd  str     s1, [sp, #68]
10008f30c: e0 37 40 f9  ldr     x0, [sp, #104]
10008f310: e0 27 00 f9  str     x0, [sp, #72]
10008f314: ee 53 00 b9  str     w14, [sp, #80]
10008f318: 7a ff ff 97  bl      0x10008f100 <_github.com/wrmsr/bane/x/blas/openblas._Cfunc_cblas_sgemm.abi0>
10008f31c: fd fb 7f a9  ldp     x29, x30, [sp, #-8]
10008f320: ff 83 01 91  add     sp, sp, #96
10008f324: c0 03 5f d6  ret
10008f328: e0 03 1f aa  mov     x0, xzr
10008f32c: e1 03 00 aa  mov     x1, x0
10008f330: 34 44 ff 97  bl      0x100060400 <_runtime.panicIndex>
10008f334: e0 03 1f aa  mov     x0, xzr
10008f338: e1 03 00 aa  mov     x1, x0
10008f33c: 31 44 ff 97  bl      0x100060400 <_runtime.panicIndex>
10008f340: e0 03 1f aa  mov     x0, xzr
10008f344: e1 03 00 aa  mov     x1, x0
10008f348: 2e 44 ff 97  bl      0x100060400 <_runtime.panicIndex>
10008f34c: 1f 20 03 d5  nop
10008f350: e0 23 00 b9  str     w0, [sp, #32]
10008f354: e1 27 00 b9  str     w1, [sp, #36]
10008f358: e2 2b 00 b9  str     w2, [sp, #40]
10008f35c: e3 1b 00 f9  str     x3, [sp, #48]
10008f360: e4 1f 00 f9  str     x4, [sp, #56]
10008f364: e5 23 00 f9  str     x5, [sp, #64]
10008f368: e0 4b 00 bd  str     s0, [sp, #72]
10008f36c: e6 2b 00 f9  str     x6, [sp, #80]
10008f370: e7 2f 00 f9  str     x7, [sp, #88]
10008f374: e8 33 00 f9  str     x8, [sp, #96]
10008f378: e9 37 00 f9  str     x9, [sp, #104]
10008f37c: ea 3b 00 f9  str     x10, [sp, #112]
10008f380: eb 3f 00 f9  str     x11, [sp, #120]
10008f384: ec 43 00 f9  str     x12, [sp, #128]
10008f388: ed 47 00 f9  str     x13, [sp, #136]
10008f38c: e1 93 00 bd  str     s1, [sp, #144]
10008f390: ee 4f 00 f9  str     x14, [sp, #152]
10008f394: e3 03 1e aa  mov     x3, x30
10008f398: 22 3a ff 97  bl      0x10005dc20 <_runtime.morestack_noctxt.abi0>
10008f39c: e0 23 80 b9  ldrsw   x0, [sp, #32]
10008f3a0: e1 27 80 b9  ldrsw   x1, [sp, #36]
10008f3a4: e2 2b 80 b9  ldrsw   x2, [sp, #40]
10008f3a8: e3 1b 40 f9  ldr     x3, [sp, #48]
10008f3ac: e4 1f 40 f9  ldr     x4, [sp, #56]
10008f3b0: e5 23 40 f9  ldr     x5, [sp, #64]
10008f3b4: e0 4b 40 bd  ldr     s0, [sp, #72]
10008f3b8: e6 2b 40 f9  ldr     x6, [sp, #80]
10008f3bc: e7 2f 40 f9  ldr     x7, [sp, #88]
10008f3c0: e8 33 40 f9  ldr     x8, [sp, #96]
10008f3c4: e9 37 40 f9  ldr     x9, [sp, #104]
10008f3c8: ea 3b 40 f9  ldr     x10, [sp, #112]
10008f3cc: eb 3f 40 f9  ldr     x11, [sp, #120]
10008f3d0: ec 43 40 f9  ldr     x12, [sp, #128]
10008f3d4: ed 47 40 f9  ldr     x13, [sp, #136]
10008f3d8: e1 93 40 bd  ldr     s1, [sp, #144]
10008f3dc: ee 4f 40 f9  ldr     x14, [sp, #152]
10008f3e0: b0 ff ff 17  b       0x10008f2a0 <_github.com/wrmsr/bane/x/blas/openblas.Sgemm>

...

10008f524: 6b 6b ff 97  bl      0x10006a2d0 <_bufio.(*Scanner).Scan>
10008f528: e0 53 40 f9  ldr     x0, [sp, #160]
10008f52c: e0 07 00 f9  str     x0, [sp, #8]
10008f530: 0c 20 a0 d2  mov     x12, #16777216
10008f534: ec 0b 00 f9  str     x12, [sp, #16]
10008f538: ec 0f 00 f9  str     x12, [sp, #24]
10008f53c: e1 0d 80 d2  mov     x1, #111
10008f540: e2 03 01 aa  mov     x2, x1
10008f544: e3 03 74 b2  orr     x3, xzr, #0x1000
10008f548: e4 03 03 aa  mov     x4, x3
10008f54c: e5 03 03 aa  mov     x5, x3
10008f550: 00 10 2e 1e  fmov    s0, #1.00000000
10008f554: e6 5b 40 f9  ldr     x6, [sp, #176]
10008f558: e7 03 0c aa  mov     x7, x12
10008f55c: e8 03 07 aa  mov     x8, x7
10008f560: e9 03 03 aa  mov     x9, x3
10008f564: ea 57 40 f9  ldr     x10, [sp, #168]
10008f568: eb 03 07 aa  mov     x11, x7
10008f56c: ed 03 03 aa  mov     x13, x3
10008f570: 01 40 20 1e  fmov    s1, s0
10008f574: ee 03 03 aa  mov     x14, x3
10008f578: c0 0c 80 d2  mov     x0, #102
10008f57c: 49 ff ff 97  bl      0x10008f2a0 <_github.com/wrmsr/bane/x/blas/openblas.Sgemm>
10008f580: 1f 20 03 d5  nop

...

0000000100090ff0 <__cgo_ce7779e6fc99_Cfunc_cblas_sgemm>:
100090ff0: ff c3 00 d1  sub     sp, sp, #48
100090ff4: fd 7b 02 a9  stp     x29, x30, [sp, #32]
100090ff8: fd 83 00 91  add     x29, sp, #32
100090ffc: 08 04 40 29  ldp     w8, w1, [x0]
100091000: 02 0c 41 29  ldp     w2, w3, [x0, #8]
100091004: 04 14 42 29  ldp     w4, w5, [x0, #16]
100091008: 00 18 40 bd  ldr     s0, [x0, #24]
10009100c: 06 10 40 f9  ldr     x6, [x0, #32]
100091010: 07 28 40 b9  ldr     w7, [x0, #40]
100091014: 09 18 40 f9  ldr     x9, [x0, #48]
100091018: 0a 38 40 b9  ldr     w10, [x0, #56]
10009101c: 01 3c 40 bd  ldr     s1, [x0, #60]
100091020: 0b 20 40 f9  ldr     x11, [x0, #64]
100091024: 0c 48 40 b9  ldr     w12, [x0, #72]
100091028: ec 1b 00 b9  str     w12, [sp, #24]
10009102c: eb 0b 00 f9  str     x11, [sp, #16]
100091030: ea 0b 00 b9  str     w10, [sp, #8]
100091034: e9 03 00 f9  str     x9, [sp]
100091038: e0 03 08 aa  mov     x0, x8
10009103c: 21 66 00 94  bl      0x1000aa8c0 <_cblas_sgemm>
100091040: fd 7b 42 a9  ldp     x29, x30, [sp, #32]
100091044: ff c3 00 91  add     sp, sp, #48
100091048: c0 03 5f d6  ret

...

00000001000aa8c0 <_cblas_sgemm>:
1000aa8c0: ff 83 04 d1  sub     sp, sp, #288
1000aa8c4: fc 6f 10 a9  stp     x28, x27, [sp, #256]
1000aa8c8: fd 7b 11 a9  stp     x29, x30, [sp, #272]
1000aa8cc: fd 43 04 91  add     x29, sp, #272
1000aa8d0: ac 0b 40 f9  ldr     x12, [x29, #16]
1000aa8d4: a8 1b 40 b9  ldr     w8, [x29, #24]
1000aa8d8: ab 13 40 f9  ldr     x11, [x29, #32]
1000aa8dc: aa 2b 40 b9  ldr     w10, [x29, #40]
1000aa8e0: a0 c3 1e b8  stur    w0, [x29, #-20]
1000aa8e4: a1 83 1e b8  stur    w1, [x29, #-24]

*/
package openblas

import (
	"fmt"
	"testing"
)

func TestOpenblas1(t *testing.T) {
	x := make([]float32, 1024*1024)
	for i := range x {
		x[i] = float32(i)
	}
	r := Ssum(len(x), x, 1)
	fmt.Println(r)
}

func TestOpenblas2(t *testing.T) {
	/*
			C := alpha*op(A)*op(B) + beta*C

			op(A) is an m-by-k matrix,
			op(B) is a k-by-n matrix,
			C is an m-by-n matrix.

		void cblas_sgemm (
			const CBLAS_LAYOUT Layout,
			const CBLAS_TRANSPOSE transa,
			const CBLAS_TRANSPOSE transb,
			const MKL_INT m,
			const MKL_INT n,
			const MKL_INT k,
			const float alpha,
			const float *a,
			const MKL_INT lda,
			const float *b,
			const MKL_INT ldb,
			const float beta,
			float *c,
			const MKL_INT ldc,
		);
	*/
	m := 16
	n := 16
	k := 16
	a := make([]float32, m*k)
	for i := range a {
		a[i] = float32(i)
	}
	b := make([]float32, k*n)
	for i := range b {
		b[i] = float32(i)
	}
	c := make([]float32, m*n)
	Sgemm(
		Rowmajor,
		Notrans,
		Notrans,
		m,
		n,
		k,
		1,
		a,
		1,
		b,
		1,
		1,
		c,
		1,
	)
	fmt.Println(c)
}

func TestOpenblas3(t *testing.T) {
	x := make([]float32, 16)
	for i := range x {
		x[i] = float32(i)
	}
	y := make([]float32, 16)
	fmt.Println(y)
	Scopy(len(x), x, 1, y, 1)
	fmt.Println(y)
}
