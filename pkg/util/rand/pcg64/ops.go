package pcg64

import "fmt"

//

type Uint128 struct {
	High uint64
	Low  uint64
}

func (v Uint128) String() string {
	return fmt.Sprintf("{%x, %x}", v.High, v.Low)
}

//

func Ror64(value uint64, rot uint) uint64 {
	return (value >> rot) | (value << ((-rot) & 63))
}

func Add128(a Uint128, b Uint128) Uint128 {
	var result Uint128

	result.Low = a.Low + b.Low
	result.High = a.High + b.High
	if result.Low < b.Low {
		result.High++
	}
	return result
}

func Mul128Parts(x uint64, y uint64, z1 *uint64, z0 *uint64) {
	var x0, x1, y0, y1 uint64
	var w0, w1, w2, t uint64

	// Lower 64 bits are straightforward clock-arithmetic.
	*z0 = x * y

	x0 = x & 0xFFFFFFFF
	x1 = x >> 32
	y0 = y & 0xFFFFFFFF
	y1 = y >> 32
	w0 = x0 * y0
	t = x1*y0 + (w0 >> 32)
	w1 = t & 0xFFFFFFFF
	w2 = t >> 32
	w1 += x0 * y1
	*z1 = x1*y1 + w2 + (w1 >> 32)
}

func Mul128(a Uint128, b Uint128) Uint128 {
	var h1 uint64
	var result Uint128

	h1 = a.High*b.Low + a.Low*b.High
	Mul128Parts(a.Low, b.Low, &(result.High), &(result.Low))
	result.High += h1
	return result
}

func Mul128Low(a Uint128, b uint64) Uint128 {
	var h1 uint64
	var result Uint128

	h1 = a.High * b
	Mul128Parts(a.Low, b, &(result.High), &(result.Low))
	result.High += h1
	return result
}
