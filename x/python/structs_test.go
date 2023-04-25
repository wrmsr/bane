package python

import (
	"reflect"

	rfl "github.com/wrmsr/bane/core/reflect"
)

//

type StructSize int8

const (
	StandardSize StructSize = iota
	NativeSize
)

//

type StructAlignment int8

const (
	NoneAlignment StructAlignment = 0
	NativeAlignment
)

//

type StructOrdering struct {
	Glyph rune
	Name  string

	Size      StructSize
	Alignment StructAlignment
}

var StructOrderings = []StructOrdering{
	{'@', "native", NativeSize, NativeAlignment},
	{'=', "native", StandardSize, NoneAlignment},
	{'<', "little-endian", StandardSize, NoneAlignment},
	{'>', "big-endian", StandardSize, NoneAlignment},
	{'!', "network", StandardSize, NoneAlignment},
}

//

type StructSpecialSize int8

const (
	PointerSize StructSpecialSize = iota + 1
	VariableSize
)

//

type StructFormat struct {
	Glyph rune
	Name  string

	Type reflect.Type

	Size        int
	SpecialSize StructSpecialSize

	Unsigned bool
}

var StructFormats = []StructFormat{
	{Glyph: 'x', Name: "pad"},
	{Glyph: 'c', Name: "char", Type: rfl.TypeOf[byte](), Size: 1},
	{Glyph: 'b', Name: "signed char", Type: rfl.TypeOf[int8](), Size: 1},
	{Glyph: 'B', Name: "unsigned char", Type: rfl.TypeOf[uint8](), Size: 1, Unsigned: true},
	{Glyph: '?', Name: "bool", Type: rfl.TypeOf[bool](), Size: 1},
	{Glyph: 'h', Name: "short", Type: rfl.TypeOf[int16](), Size: 2},
	{Glyph: 'H', Name: "unsigned short", Type: rfl.TypeOf[uint16](), Size: 2, Unsigned: true},
	{Glyph: 'i', Name: "int", Type: rfl.TypeOf[int32](), Size: 4},
	{Glyph: 'I', Name: "unsigned int", Type: rfl.TypeOf[uint32](), Size: 4, Unsigned: true},
	{Glyph: 'l', Name: "long", Type: rfl.TypeOf[int32](), Size: 4},
	{Glyph: 'L', Name: "unsigned long", Type: rfl.TypeOf[uint32](), Size: 4, Unsigned: true},
	{Glyph: 'q', Name: "long long", Type: rfl.TypeOf[int64](), Size: 8},
	{Glyph: 'Q', Name: "unsigned long long", Type: rfl.TypeOf[uint64](), Size: 8, Unsigned: true},
	{Glyph: 'n', Name: "ssize_t", Type: rfl.TypeOf[int](), SpecialSize: PointerSize},
	{Glyph: 'N', Name: "size_t", Type: rfl.TypeOf[int](), SpecialSize: PointerSize},
	{Glyph: 'e', Name: "float16", Type: rfl.TypeOf[float32](), Size: 2},
	{Glyph: 'f', Name: "float", Type: rfl.TypeOf[float32](), Size: 4},
	{Glyph: 'd', Name: "double", Type: rfl.TypeOf[float64](), Size: 8},
	{Glyph: 's', Name: "char[]", Type: rfl.TypeOf[[]byte](), SpecialSize: VariableSize},
	{Glyph: 'p', Name: "char[]", Type: rfl.TypeOf[[]byte](), SpecialSize: VariableSize},
	{Glyph: 'P', Name: "void *", Type: rfl.TypeOf[uintptr](), SpecialSize: PointerSize},
}

//

type BitStructFormat struct {
	StructFormat
}

var BitStructFormats = []BitStructFormat{
	{StructFormat{Glyph: 'j', Name: "packed bytes", Type: rfl.TypeOf[[]byte](), SpecialSize: VariableSize}},
	{StructFormat{Glyph: 't', Name: "packed int", Type: rfl.TypeOf[int](), SpecialSize: VariableSize}},
	{StructFormat{Glyph: 'T', Name: "unsigned packed int", Type: rfl.TypeOf[int](), SpecialSize: VariableSize, Unsigned: true}},
	{StructFormat{Glyph: 'X', Name: "pad bits", Type: rfl.TypeOf[int](), SpecialSize: VariableSize}},
}
