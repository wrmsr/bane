package tg

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
)

type ConvArgs struct {
	h, w Dim

	groups Dim

	rcout Dim
	cin   Dim

	oy, ox Dim
	iy, ix Dim
	sy, sx Dim

	bs Dim

	cout Dim

	py, py_ Dim

	px, px_ Dim

	dy, dx Dim

	outShape Shape
}

type ConvOpts struct {
	Stride   []Dim
	Groups   Dim
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

	cout, cin, h, w := slices.Unpack4(wsh)
	sx, sy := slices.Unpack2(opts.Stride)

	var px, px_, py, py_ Dim
	switch len(opts.Padding) {
	case 4:
		px, px_, py, py_ = slices.Unpack4(opts.Padding)
	case 2:
		px, py = slices.Unpack2(opts.Padding)
		px_, py_ = px, py
	case 1:
		px, px_, py, py_ = slices.Dup4(opts.Padding[0])
	default:
		panic("unhandled padding")
	}

	var dy, dx Dim
	switch len(opts.Dilation) {
	case 2:
		dy, dx = slices.Unpack2(opts.Dilation)
	case 1:
		dy, dx = slices.Dup2(opts.Dilation[0])
	default:
		panic("unhandled dilation")
	}

	bs, cin_, iy, ix := slices.Unpack4(xsh)

	if len(opts.OutShape) > 0 {
		check.Condition(len(opts.OutShape) == 4)
		py_ = (opts.OutShape[2]-1)*sy + 1 + dy*(h-1) - iy - py
		px_ = (opts.OutShape[3]-1)*sx + 1 + dx*(w-1) - ix - px
	}

	oy := (iy+py+py_-dy*(h-1)-1)/sy + 1
	ox := (ix+px+px_-dx*(w-1)-1)/sx + 1

	if cin*opts.Groups != cin_ {
		panic(fmt.Errorf(
			"input Tensor shape %v does not match the shape of the weights %v. (%v vs. %v)",
			xsh,
			wsh,
			cin*opts.Groups,
			cin_,
		))
	}

	check.Condition(cout%opts.Groups == 0)
	check.Condition(opts.OutShape == nil || opts.OutShape.Equals(Shape{bs, cout, oy, ox}))

	return ConvArgs{
		h:        h,
		w:        w,
		groups:   opts.Groups,
		rcout:    cout / opts.Groups,
		cin:      cin,
		oy:       oy,
		ox:       ox,
		iy:       iy,
		ix:       ix,
		sy:       sy,
		sx:       sx,
		bs:       bs,
		cout:     cout,
		py:       py,
		py_:      py_,
		px:       px,
		px_:      px_,
		dy:       dy,
		dx:       dx,
		outShape: Shape{bs, cout, oy, ox},
	}
}
