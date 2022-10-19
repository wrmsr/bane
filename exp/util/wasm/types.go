package wasm

import (
	"fmt"
	"io"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

//

type Type interface {
	isType()

	Renderer
	fmt.Stringer
}

type type_ struct{}

func (t type_) isType() {}

//

type None struct {
	type_
}

var _ Type = None{}

func (t None) Render(w io.Writer) { iou.Dw(w).String("none") }

func (t None) String() string { return RenderString(t) }

//

type NumType interface {
	Type
	isNumType()
}

type numType struct {
	type_
}

func (t numType) isNumType() {}

//

type I32 struct {
	numType
}

var _ NumType = I32{}

type I64 struct {
	numType
}

var _ NumType = I64{}

type F32 struct {
	numType
}

var _ NumType = F32{}

type F64 struct {
	numType
}

var _ NumType = F64{}

func (t I32) Render(w io.Writer) { iou.Dw(w).String("i32") }
func (t I64) Render(w io.Writer) { iou.Dw(w).String("i64") }
func (t F32) Render(w io.Writer) { iou.Dw(w).String("f32") }
func (t F64) Render(w io.Writer) { iou.Dw(w).String("f64") }

func (t I32) String() string { return RenderString(t) }
func (t I64) String() string { return RenderString(t) }
func (t F32) String() string { return RenderString(t) }
func (t F64) String() string { return RenderString(t) }

//

func ParseType(s string) Type {
	switch s {
	case "i32":
		return I32{}
	case "i64":
		return I64{}
	case "f32":
		return F32{}
	case "f64":
		return F64{}
	}
	panic(s)
}
