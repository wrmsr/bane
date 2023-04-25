package tg

import (
	"image"
	"image/color"

	nd "github.com/wrmsr/bane/pkg/util/ndarray"
)

type Float32GrayscaleImage struct {
	a nd.NdArray[float32]
}

var _ image.Image = Float32GrayscaleImage{}

func (i Float32GrayscaleImage) ColorModel() color.Model {
	return color.GrayModel
}

func (i Float32GrayscaleImage) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: 0,
		},
		Max: image.Point{
			X: int(i.a.View().Shape().Get(1)),
			Y: int(i.a.View().Shape().Get(0)),
		},
	}
}

func (i Float32GrayscaleImage) At(x, y int) color.Color {
	v := i.a.Get(nd.DimsOf(nd.Dim(y), nd.Dim(x)))
	return color.Gray{uint8(v)}
}

//func TestMnistData(t *testing.T) {
//	i := Float32GrayscaleImage{f.Slice(0)}
//	wb := bytes.NewBuffer(nil)
//	check.Must(jpeg.Encode(wb, i, nil))
//
//	tf := check.Must1(ioutil.TempFile("", "foo"))
//	//check.Must1(os.CreateTemp("", "foo.jpg"))
//
//	_ = check.Must1(tf.Write(check.Must1(io.ReadAll(wb))))
//
//	nn := tf.Name() + ".jpg"
//	check.Must(os.Rename(tf.Name(), nn))
//
//	check.Must(exec.Command("open", nn).Run())
//}
