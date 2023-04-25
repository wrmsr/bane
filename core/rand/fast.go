/*
The MIT License (MIT)

# Copyright (c) 2017 Aliaksandr Valialkin

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package rand

import (
	"sync"
	"time"
)

// Uint32 returns pseudorandom uint32.
//
// It is safe calling this function from concurrent goroutines.
func FastUint32() uint32 {
	v := fastUint32RngPool.Get()
	if v == nil {
		v = &FastUint32Rng{}
	}
	r := v.(*FastUint32Rng)
	x := r.Uint32()
	fastUint32RngPool.Put(r)
	return x
}

var fastUint32RngPool sync.Pool

// Uint32n returns pseudorandom uint32 in the range [0..maxN).
//
// It is safe calling this function from concurrent goroutines.
func FastUint32n(maxN uint32) uint32 {
	x := FastUint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32((uint64(x) * uint64(maxN)) >> 32)
}

// RNG is a pseudorandom number generator.
//
// It is unsafe to call RNG methods from concurrent goroutines.
type FastUint32Rng struct {
	x uint32
}

// Uint32 returns pseudorandom uint32.
//
// It is unsafe to call this method from concurrent goroutines.
func (r *FastUint32Rng) Uint32() uint32 {
	for r.x == 0 {
		r.x = getRandomUint32()
	}

	// See https://en.wikipedia.org/wiki/Xorshift
	x := r.x
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	r.x = x
	return x
}

// Uint32n returns pseudorandom uint32 in the range [0..maxN).
//
// It is unsafe to call this method from concurrent goroutines.
func (r *FastUint32Rng) Uint32n(maxN uint32) uint32 {
	x := r.Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32((uint64(x) * uint64(maxN)) >> 32)
}

// Seed sets the r state to n.
func (r *FastUint32Rng) Seed(n uint32) {
	r.x = n
}

func getRandomUint32() uint32 {
	x := time.Now().UnixNano()
	return uint32((x >> 32) ^ x)
}
