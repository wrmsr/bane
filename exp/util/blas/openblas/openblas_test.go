//
/*
https://github.com/xianyi/OpenBLAS/wiki/Developer-manual#optimizing-gemm-for-a-given-hardware
https://www.cs.utexas.edu/~flame/web/FLAMEPublications.html
https://www.cs.utexas.edu/~flame/pubs/GotoTOMS_revision.pdf
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
