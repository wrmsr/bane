package tg

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
	ju "github.com/wrmsr/bane/pkg/util/json"
	nd "github.com/wrmsr/bane/pkg/util/ndarray"
)

func readTgFile(name string) []byte {
	p := filepath.Join(
		paths.FindProjectRoot(),
		fmt.Sprintf("../../geohot/tinygrad/%s", name),
	)
	return check.Must1(os.ReadFile(p))
}

func readTgMat2(name string) nd.NdArray[float32] {
	return nd.Unnest2(check.Must1(ju.UnmarshalAs[[][]float32](readTgFile(name))))
}

func readMnistImages(name string) nd.NdArray[float32] {
	gr := check.Must1(gzip.NewReader(bytes.NewReader(readTgFile(name))))
	b := check.Must1(io.ReadAll(gr))[0x10:]
	bf := make([]float32, len(b))
	for i, v := range b {
		bf[i] = float32(v)
	}
	return nd.Maker[float32]{Data: bf}.Make().Reshape(nd.ShapeOf(-1, 28, 28))
}

func readMnistLabels(name string) nd.NdArray[byte] {
	return nd.Maker[byte]{Data: readTgFile(name)[8:]}.Make()
}

func TestBobNet2(t *testing.T) {
	l1t := readTgMat2("l1.json")
	l2t := readTgMat2("l2.json")
	samp := check.Must1(ju.UnmarshalAs[[]int](readTgFile("samp.json")))

	x_train := readMnistImages("datasets/mnist/train-images-idx3-ubyte.gz")
	y_train := readMnistLabels("datasets/mnist/train-labels-idx1-ubyte.gz")

	fmt.Println(l1t.View())
	fmt.Println(l2t.View())
	fmt.Println(samp)

	fmt.Println(x_train.View())
	fmt.Println(y_train.View())

	//zt := xt.Dot(l1t).Relu().Dot(l2t).LogSoftmax()
	////zt := xt.Dot(l1t).Relu().Dot(l2t)
	//fmt.Println("====")
	//
	//dumpObj(zt)
	//
	////fmt.Println(zt.Mean(nil, false).Data().Realize())
	////fmt.Println(zt.Mean(nil, false).Data().Realize())
	//
	//// garbage
	//
	//scc := func(out *Tensor, y_ []int) *Tensor {
	//	check.Equal(Dim(len(y_)), out.Shape()[0])
	//	check.Equal(out.Shape()[0], Dim(len(y_)))
	//	num_classes := out.Shape()[len(out.Shape())-1]
	//	yb := NewBuffer(out.Shape())
	//	for i := Dim(0); i < Dim(len(y_)); i++ {
	//		yb.set(float32(-1*num_classes), i, Dim(y_[i]))
	//	}
	//	y := NewTensor(MakeLoadBuffer(yb), false)
	//	return out.Mul(y).Mean(nil, false)
	//}
	//
	//y := make([]int, 3)
	//for i := 0; i < 3; i++ {
	//	y[i] = i % 3
	//}
	//
	//z := scc(
	//	zt,
	//	y,
	//)
	//
	//fmt.Println(z.Data().Realize().Nd())
	//z.Backward()
}

//
