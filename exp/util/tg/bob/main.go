package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime/pprof"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/exp/util/numpy"
	"github.com/wrmsr/bane/exp/util/tg"
	tgdev "github.com/wrmsr/bane/exp/util/tg/dev"
	"github.com/wrmsr/bane/pkg/util/check"
	ju "github.com/wrmsr/bane/pkg/util/json"
	nd "github.com/wrmsr/bane/pkg/util/ndarray"
	"github.com/wrmsr/bane/pkg/util/slices"
	log "github.com/wrmsr/bane/pkg/util/slog"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func readTgFile(name string) []byte {
	p := filepath.Join(
		check.Must1(os.Getwd()),
		fmt.Sprintf("../../geohot/tinygrad/%s", name),
	)
	return check.Must1(os.ReadFile(p))
}

func readTgNpy(name string) nd.NdArray[float32] {
	p := filepath.Join(
		check.Must1(os.Getwd()),
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

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		log.IfFatal(pprof.StartCPUProfile(f))
		defer pprof.StopCPUProfile()
	}

	l1t := tg.NewTensor(tg.MakeLoadBuffer(tg.BufferOfNd(
		//readTgMat2("l1.json"),
		readTgNpy("l1.npy").Reshape(nd.ShapeOf(784, 128)),
	)), true)
	l2t := tg.NewTensor(tg.MakeLoadBuffer(tg.BufferOfNd(
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
	samps = samps[:10]

	fmt.Println(tgdev.FetchDatasets())

	for e, samp := range samps {
		x := nd.New[float32](nd.ShapeOf(69, 784))
		y := nd.New[float32](nd.ShapeOf(69, nd.Dim(num_classes)))

		for i, s := range samp {
			x.Slice(i).Assign(x_train.Slice(s).Reshape(nd.ShapeOf(784)))
			c := y_train.Get(nd.AsDims(s))
			*y.At(nd.AsDims(i, c)) = -float32(num_classes)
		}

		xt := tg.NewTensor(tg.MakeLoadBuffer(tg.BufferOfNd(x)), false)
		yt := tg.NewTensor(tg.MakeLoadBuffer(tg.BufferOfNd(y)), false)

		out := xt.Dot(l1t).Relu().Dot(l2t).LogSoftmax()

		loss := out.Mul(yt).Mean(nil, false)

		fmt.Printf("loss %d: %s\n", e, loss.Data().Realize().Nd())

		loss.Backward()
		if e == 0 {
			//renderDot(loss)
		}

		lr := float32(0.001)
		for i, t := range []*tg.Tensor{
			l1t,
			l2t,
		} {
			_ = i

			//fmt.Printf("step realize: %d\n", i)

			g := t.Grad()

			lrt := tg.NewTensor(
				tg.MakeLoadBuffer(
					tg.BufferOfNd(
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
