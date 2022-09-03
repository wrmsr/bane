/*
The MIT License

PCG Random Number Generation for C.

Copyright 2014 Melissa O'Neill <oneill@pcg-random.org>

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
package pcg64

//

//

type StateSetseq128 struct {
	state Uint128
	inc   Uint128
}

const DEFAULT_MULTIPLIER_HIGH = uint64(2549297995355413924)
const DEFAULT_MULTIPLIER_LOW = uint64(4865540595714422341)

var DEFAULT_MULTIPLIER_128 = Uint128{DEFAULT_MULTIPLIER_HIGH, DEFAULT_MULTIPLIER_LOW}
var DEFAULT_INCREMENT_128 = Uint128{uint64(6364136223846793005), uint64(1442695040888963407)}

var STATE_SETSEQ_128_INITIALIZER = []Uint128{
	{0x979c9a98d8462005, 0x7d3e9cb6cfe0549b},
	{0x0000000000000001, 0xda3e39cb94b95bdb},
}

func (rng *StateSetseq128) PcgSetseq128Step() {
	rng.state = Add128(Mul128(rng.state, DEFAULT_MULTIPLIER_128), rng.inc)
}

func PcgOutputXslRr128_64(state Uint128) uint64 {
	return Ror64(state.High^state.Low, uint(state.High>>58))
}

func (rng *StateSetseq128) PcgSetseq128SrandomR(initstate Uint128, initseq Uint128) {
	rng.state = Uint128{0, 0}
	rng.inc.High = initseq.High << 1
	rng.inc.High |= initseq.Low >> 63
	rng.inc.Low = (initseq.Low << 1) | 1
	rng.PcgSetseq128Step()
	rng.state = Add128(rng.state, initstate)
	rng.PcgSetseq128Step()
}

func (rng *StateSetseq128) pcg_setseq_128_xsl_rr_64_random_r() uint64 {
	rng.PcgSetseq128Step()
	return PcgOutputXslRr128_64(rng.state)
}

const CHEAP_MULTIPLIER_128 = uint64(0xda942042e4dd58b5)

func (rng *StateSetseq128) pcg_cm_step_r() {
	rng.state = Add128(Mul128Low(rng.state, CHEAP_MULTIPLIER_128), rng.inc)
}

func (rng *StateSetseq128) pcg_cm_srandom_r(initstate Uint128, initseq Uint128) {
	rng.state = Uint128{0, 0}
	rng.inc.High = initseq.High << 1
	rng.inc.High |= initseq.Low >> 63
	rng.inc.Low = (initseq.Low << 1) | 1
	rng.pcg_cm_step_r()
	rng.state = Add128(rng.state, initstate)
	rng.pcg_cm_step_r()
}

func (rng *StateSetseq128) pcg_cm_random_r() uint64 {
	// Lots of manual inlining to help out certain compilers to generate performant code.
	hi := rng.state.High
	lo := rng.state.Low

	// Run the DXSM output function on the pre-iterated state.
	lo |= 1
	hi ^= hi >> 32
	hi *= 0xda942042e4dd58b5
	hi ^= hi >> 48
	hi *= lo

	// Run the CM step.
	rng.state = Add128(Mul128Low(rng.state, CHEAP_MULTIPLIER_128), rng.inc)
	return hi
}

func (rng *StateSetseq128) pcg_setseq_128_xsl_rr_64_boundedrand_r(bound uint64) uint64 {
	threshold := -bound % bound
	for {
		r := rng.pcg_setseq_128_xsl_rr_64_random_r()
		if r >= threshold {
			return r % bound
		}
	}
}

func (rng *StateSetseq128) pcg_setseq_128_advance_r(delta Uint128) {
	rng.state = pcg_advance_lcg_128(rng.state, delta, DEFAULT_MULTIPLIER_128, rng.inc)
}

func (rng *StateSetseq128) pcg_cm_advance_r(delta Uint128) {
	rng.state = pcg_advance_lcg_128(rng.state, delta, Uint128{0, CHEAP_MULTIPLIER_128}, rng.inc)
}

type pcg64_random_t = StateSetseq128

var pcg64_random_r = (*StateSetseq128).pcg_setseq_128_xsl_rr_64_random_r
var pcg64_boundedrand_r = (*StateSetseq128).pcg_setseq_128_xsl_rr_64_boundedrand_r
var pcg64_srandom_r = (*StateSetseq128).PcgSetseq128SrandomR
var pcg64_advance_r = (*StateSetseq128).pcg_setseq_128_advance_r
var PCG64_INITIALIZER = STATE_SETSEQ_128_INITIALIZER

type State struct {
	pcg_state  *pcg64_random_t
	has_uint32 int
	uinteger   uint32
}

func (state *State) pcg64_next64() uint64 {
	return pcg64_random_r(state.pcg_state)
}

func (state *State) pcg64_next32() uint32 {
	var next uint64
	if state.has_uint32 != 0 {
		state.has_uint32 = 0
		return state.uinteger
	}
	next = pcg64_random_r(state.pcg_state)
	state.has_uint32 = 1
	state.uinteger = uint32(next >> 32)
	return uint32(next & 0xffffffff)
}

func (state *State) pcg64_cm_next64() uint64 {
	return state.pcg_state.pcg_cm_random_r()
}

func (state *State) pcg64_cm_next32() uint32 {
	var next uint64
	if state.has_uint32 != 0 {
		state.has_uint32 = 0
		return state.uinteger
	}
	next = state.pcg_state.pcg_cm_random_r()
	state.has_uint32 = 1
	state.uinteger = uint32(next >> 32)
	return uint32(next & 0xffffffff)
}

// Muli-step advance functions (jump-ahead, jump-back)
//
// The method used here is based on Brown, "Random Number Generation with Arbitrary Stride,", Transactions of the
// American Nuclear Society (Nov. 1994).  The algorithm is very similar to fast exponentiation.
//
// Even though delta is an unsigned integer, we can pass a signed integer to go backwards, it just goes "the long way
// round".
func pcg_advance_lcg_128(state Uint128, delta Uint128, cur_mul Uint128, cur_plus Uint128) Uint128 {
	acc_mul := Uint128{0, 1}
	acc_plus := Uint128{0, 0}
	for (delta.High > 0) || (delta.Low > 0) {
		if (delta.Low & 1) != 0 {
			acc_mul = Mul128(acc_mul, cur_mul)
			acc_plus = Add128(Mul128(acc_plus, cur_mul), cur_plus)
		}
		cur_plus = Mul128(Add128(cur_mul, Uint128{0, 1}), cur_plus)
		cur_mul = Mul128(cur_mul, cur_mul)
		delta.Low = (delta.Low >> 1) | (delta.High << 63)
		delta.High >>= 1
	}
	return Add128(Mul128(acc_mul, state), acc_plus)
}

//

func pcg64_advance(state *State, step []uint64) {
	var delta Uint128
	delta.High = step[0]
	delta.Low = step[1]
	pcg64_advance_r(state.pcg_state, delta)
}

func pcg64_cm_advance(state *State, step []uint64) {
	var delta Uint128
	delta.High = step[0]
	delta.Low = step[1]
	state.pcg_state.pcg_cm_advance_r(delta)
}

func pcg64_set_seed(state *State, seed []uint64, inc []uint64) {
	var s, i Uint128
	s.High = seed[0]
	s.Low = seed[1]
	i.High = inc[0]
	i.Low = inc[1]
	pcg64_srandom_r(state.pcg_state, s, i)
}

func pcg64_get_state(state *State, state_arr []uint64, has_uint32 *int, uinteger []uint32) {
	// state_arr contains state.high, state.low, inc.high, inc.low which are interpreted as the upper 64 bits (high) or
	// lower 64 bits of a uint128_t variable
	state_arr[0] = state.pcg_state.state.High
	state_arr[1] = state.pcg_state.state.Low
	state_arr[2] = state.pcg_state.inc.High
	state_arr[3] = state.pcg_state.inc.Low
	*has_uint32 = state.has_uint32
	uinteger[0] = state.uinteger
}

func pcg64_set_state(state *State, state_arr []uint64, has_uint32 int, uinteger uint32) {
	// state_arr contains state.high, state.low, inc.high, inc.low which are interpreted as the upper 64 bits (high) or
	// lower 64 bits of a uint128_t variable
	state.pcg_state.state.High = state_arr[0]
	state.pcg_state.state.Low = state_arr[1]
	state.pcg_state.inc.High = state_arr[2]
	state.pcg_state.inc.Low = state_arr[3]
	state.has_uint32 = has_uint32
	state.uinteger = uinteger
}
