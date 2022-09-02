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
	gr := check.Must1(gzip.NewReader(bytes.NewReader(readTgFile(name))))
	b := check.Must1(io.ReadAll(gr))[8:]
	return nd.Maker[byte]{Data: b}.Make()
}

func TestBobNet2(t *testing.T) {
	l1t := NewTensor(MakeLoadBuffer(BufferOfNd(readTgMat2("l1.json"))), true)
	l2t := NewTensor(MakeLoadBuffer(BufferOfNd(readTgMat2("l2.json"))), true)
	samp := check.Must1(ju.UnmarshalAs[[]int](readTgFile("samp.json")))

	x_train := readMnistImages("datasets/mnist/train-images-idx3-ubyte.gz")
	y_train := readMnistLabels("datasets/mnist/train-labels-idx1-ubyte.gz")

	num_classes := 10
	x := nd.New[float32](nd.ShapeOf(69, 784))
	y := nd.New[float32](nd.ShapeOf(69, nd.Dim(num_classes)))
	for i, s := range samp {
		x.Slice(i).Assign(x_train.Slice(s).Reshape(nd.ShapeOf(784)))
		c := y_train.Get(nd.Dim(s))
		*y.At(nd.Dim(i), nd.Dim(c)) = -float32(num_classes)
	}

	xt := NewTensor(MakeLoadBuffer(BufferOfNd(x)), false)
	yt := NewTensor(MakeLoadBuffer(BufferOfNd(y)), false)

	out := xt.Dot(l1t).Relu().Dot(l2t).LogSoftmax()

	loss := out.Mul(yt).Mean(nil, false)
	loss.Backward()

	//lr := MakeConstBuffer(0.001, Shape{1})

	fmt.Println(l1t.grad.data.Realize().Nd().Slice([]any{100, 130}))
	fmt.Println(l2t.grad.data.Realize().Nd().Slice([]any{100, 130}))

	//l1t.grad.Mul(lr)
}
