package grad

import (
	"fmt"
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func randMat64(r, c int) *mat.Dense {
	data := make([]float64, r*c)
	for i := range data {
		data[i] = rand.NormFloat64()
	}
	return mat.NewDense(r, c, data)
}

func TestSimple(t *testing.T) {
	a := randMat64(10, 10)
	b := randMat64(10, 10)

	c := mat.NewDense(10, 10, nil)
	c.Add(a, b)

	fmt.Println(c)
}

func TestBackward(t *testing.T) {
	xInit := randMat64(1, 3)
	//uInit := randMat64(3, 3)
	//vInit := randMat64(3, 3)
	//wInit := randMat64(3, 3)
	//mInit := randMat64(1, 3)

	var x mat.Dense
	x.CloneFrom(xInit)
	//var w mat.Dense
	//w.CloneFrom(wInit)

	fmt.Println(x)
}
