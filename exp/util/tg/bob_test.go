package tg

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/wrmsr/bane/exp/util/numpy"
	tgdev "github.com/wrmsr/bane/exp/util/tg/dev"
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/dev/paths"
	"github.com/wrmsr/bane/pkg/util/graphs/dot"
	iou "github.com/wrmsr/bane/pkg/util/io"
	ju "github.com/wrmsr/bane/pkg/util/json"
	nd "github.com/wrmsr/bane/pkg/util/ndarray"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func readTgFile(name string) []byte {
	p := filepath.Join(
		paths.FindProjectRoot(),
		fmt.Sprintf("../../geohot/tinygrad/%s", name),
	)
	return check.Must1(os.ReadFile(p))
}

func readTgNpy(name string) nd.NdArray[float32] {
	p := filepath.Join(
		paths.FindProjectRoot(),
		fmt.Sprintf("../../geohot/tinygrad/%s", name),
	)
	return nd.Maker[float32]{Data: numpy.ShittyReadFloat32s(p)}.Make()
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

/*
Fma:
[ 2.5077548 ]
[ 2.299427 ]
[ 2.8177679 ]
[ 1.8951459 ]
[ 2.7053564 ]
[ 2.2945564 ]
[ 1.388959 ]
[ 4.608702 ]
[ 1.5387887 ]
[ 2.633925 ]

Sep:
[ 2.5077548 ]
[ 2.299427 ]
[ 2.8177683 ]
[ 1.8951463 ]
[ 2.7053568 ]
[ 2.294557 ]
[ 1.3889599 ]
[ 4.608704 ]
[ 1.5387893 ]
[ 2.6339264 ]
*/

func renderDot(t *Tensor) {
	g := dot.Graph{}

	name := func(l Lazy) string {
		switch l := l.(type) {
		case *LazyBuffer:
			return fmt.Sprintf("b%x", rfl.AddressOf(l))
		case *LazyOp:
			return fmt.Sprintf("o%x", rfl.AddressOf(l))
		default:
			panic(l)
		}
	}

	seen := make(map[uintptr]struct{})
	var rec func(Lazy)
	rec = func(l Lazy) {
		a := rfl.AddressOf(l)
		if _, ok := seen[a]; ok {
			return
		}
		seen[a] = struct{}{}

		switch l := l.(type) {

		case *LazyBuffer:
			n := name(l)
			g.Stmts = append(g.Stmts,
				dot.NewNode(n).SetAttr("label", fmt.Sprintf("%s %v", n, l.Shape())),
				dot.NewEdge(n, name(l.op)).SetAttr("dir", "back"),
			)
			rec(l.op)

		case *LazyOp:
			n := name(l)
			g.Stmts = append(g.Stmts,
				dot.NewNode(n).SetAttr("label", fmt.Sprintf("%s %s", n, l.op.String())),
			)
			for _, s := range l.srcs {
				g.Stmts = append(g.Stmts,
					dot.NewEdge(n, name(s)).SetAttr("dir", "back"),
				)
			}
			for _, s := range l.srcs {
				rec(s)
			}

		default:
			panic(l)
		}
	}
	rec(t.data)

	dot.NewEdge("a", "b")

	var sb strings.Builder
	dot.Render(iou.NewDiscardStringWriter(&sb), g)

	check.Must(dot.Open(context.Background(), sb.String()))
}

func TestBobNet2(t *testing.T) {
	l1t := NewTensor(MakeLoadBuffer(BufferOfNd(
		//readTgMat2("l1.json"),
		readTgNpy("l1.npy").Reshape(nd.ShapeOf(784, 128)),
	)), true)
	l2t := NewTensor(MakeLoadBuffer(BufferOfNd(
		//readTgMat2("l2.json"),
		readTgNpy("l2.npy").Reshape(nd.ShapeOf(128, 10)),
	)), true)

	x_train := readMnistImages("datasets/mnist/train-images-idx3-ubyte.gz")
	y_train := readMnistLabels("datasets/mnist/train-labels-idx1-ubyte.gz")

	num_classes := 10

	var samps [][]int
	for _, l := range strings.Split(strings.TrimSpace(string(readTgFile("samp_log.txt"))), "\n") {
		samps = append(samps, slices.Map(func(s string) int { return check.Must1(strconv.Atoi(s)) }, strings.Split(l, ",")))
	}

	//num_epochs := 10

	fmt.Println(tgdev.FetchDatasets())

	for e, samp := range samps {
		x := nd.New[float32](nd.ShapeOf(69, 784))
		y := nd.New[float32](nd.ShapeOf(69, nd.Dim(num_classes)))

		//var samp []int
		//if e < 10 {
		//	samp = check.Must1(ju.UnmarshalAs[[]int](readTgFile(fmt.Sprintf("samp_%d.json", e))))
		//} else {
		//	samp = make([]int, 69)
		//	for i := 0; i < len(samp); i++ {
		//		for {
		//			s := rand.Int() % 60_000
		//			if !slices.Contains(samp, s) {
		//				samp[i] = s
		//				break
		//			}
		//		}
		//	}
		//}

		for i, s := range samp {
			x.Slice(i).Assign(x_train.Slice(s).Reshape(nd.ShapeOf(784)))
			c := y_train.Get(nd.Dim(s))
			*y.At(nd.Dim(i), nd.Dim(c)) = -float32(num_classes)
		}

		xt := NewTensor(MakeLoadBuffer(BufferOfNd(x)), false)
		yt := NewTensor(MakeLoadBuffer(BufferOfNd(y)), false)

		out := xt.Dot(l1t).Relu().Dot(l2t).LogSoftmax()

		loss := out.Mul(yt).Mean(nil, false)

		fmt.Printf("loss %d: %s\n", e, loss.Data().Realize().Nd())

		loss.Backward()
		if e == 0 {
			//renderDot(loss)
		}

		lr := float32(0.001)
		for i, t := range []*Tensor{
			l1t,
			l2t,
		} {
			_ = i

			//fmt.Printf("step realize: %d\n", i)

			g := t.grad

			lrt := NewTensor(
				MakeLoadBuffer(
					BufferOfNd(
						nd.Maker[float32]{
							Shape:   nd.ShapeOf(g.Shape()...),
							Default: bt.Just[float32](lr),
						}.Make(),
					),
				),
				false,
			)

			t.Assign(t.Sub(g.Mul(lrt)))
			t.Data().Realize()
			t.ClearGrad()
		}
	}
}
