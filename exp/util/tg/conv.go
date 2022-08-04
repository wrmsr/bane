package tg

import (
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
)

type ConvArgs struct {
	h, w Dim

	groups int

	rcout int
	cin   int

	oy, ox Dim
	iy, ix Dim
	sy, sx Dim

	bs int

	cout int

	py, py_ Dim

	px, px_ Dim

	dy, dx Dim

	outShape Shape
}

type ConvOpts struct {
	Stride   []Dim
	Groups   int
	Padding  []Dim
	Dilation []Dim
	OutShape Shape
}

func BuildConvArgs(xsh, wsh Shape, opts ConvOpts) ConvArgs {
	if len(opts.Stride) < 1 {
		opts.Stride = []Dim{1, 1}
	}
	if len(opts.Stride) == 1 {
		opts.Stride = slices.Repeat(opts.Stride, 2)
	}
	if opts.Groups == 0 {
		opts.Groups = 1
	}
	if len(opts.Dilation) < 1 {
		opts.Dilation = []Dim{1}
	}

	check.Condition(len(wsh) == 4)
	cout, cin, h, w := wsh[0], wsh[1], wsh[2], wsh[3]

	check.Condition(len(opts.Stride) == 2)
	sx, sy := opts.Stride[0], opts.Stride[1]

	var px, px_, py, py_ Dim
	switch len(opts.Padding) {
	case 4:
		px, px_, py, py_ = opts.Padding[0], opts.Padding[1], opts.Padding[2], opts.Padding[3]
	case 2:
		px, px_, py, py_ = opts.Padding[0], opts.Padding[0], opts.Padding[1], opts.Padding[1]
	case 1:
		px, px_, py, py_ = opts.Padding[0], opts.Padding[0], opts.Padding[0], opts.Padding[0]
	default:
		panic("unhandled padding")
	}

	var dy, dx Dim
	switch len(opts.Dilation) {
	case 2:
		dy, dx = opts.Dilation[0], opts.Dilation[1]
	case 1:
		dy, dx = opts.Dilation[0], opts.Dilation[0]
	default:
		panic("unhandled dilation")
	}

	check.Condition(len(xsh) == 4)
	bs, cin_, iy, ix := xsh[0], xsh[1], xsh[2], xsh[3]

	/*
		      # this can change px_ and py_ to make the out_shape right
		      if len(opts.OutShape) > 0 {
				check.Condition(len(ou) == 4)
		          py_ = (out_shape[2] - 1) * sy + 1 + dy * (h - 1) - iy - py
		          px_ = (out_shape[3] - 1) * sx + 1 + dx * (w - 1) - ix - px
			}

		      # TODO: should be easy to support asymmetric padding by changing output size
		      # https://pytorch.org/docs/stable/generated/torch.nn.Conv2d.html describes these sizes well
		      oy = (iy + py + py_ - dy * (h - 1) - 1) // sy + 1
		      ox = (ix + px + px_ - dx * (w - 1) - 1) // sx + 1
		      if cin * groups != cin_: raise Exception(
		          f"Input Tensor shape {x_shape} does not match the shape of the weights {w_shape}. ({cin * groups} vs. {cin_})")
		      assert cout % groups == 0
		      assert out_shape is None or out_shape == (bs, cout, oy, ox)
		      return ConvArgs(
		   h,
		   w,
		   groups,
		   cout // groups,
		   cin,
		   oy,
		   ox,
		   iy,
		   ix,
		   sy,
		   sx,
		   bs,
		   cout,
		   py,
		   py_,
		   px,
		   px_,
		   dy,
		   dx,
		   (bs, cout, oy, ox))
	*/
	panic("nyi")
}
