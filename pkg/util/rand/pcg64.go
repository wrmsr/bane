package rand

/*
 * PCG64 Random Number Generation for C.
 *
 * Copyright 2014 Melissa O'Neill <oneill@pcg-random.org>
 * Copyright 2015 Robert Kern <robert.kern@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * For additional information about the PCG random number generation scheme,
 * including its license and other licensing options, visit
 *
 *     http://www.pcg-random.org
 *
 * Relicensed MIT in May 2019
 *
 * The MIT License
 *
 * PCG Random Number Generation for C.
 *
 * Copyright 2014 Melissa O'Neill <oneill@pcg-random.org>
 *
 * Permission is hereby granted, free of charge, to any person obtaining
 * a copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

type pcg128_t struct {
	high uint64
	low  uint64
}

func PCG_128BIT_CONSTANT(high uint64, low uint64) pcg128_t {
	return pcg128_t{
		high: high,
		low:  low,
	}
}

type pcg_state_128 struct{ state pcg128_t }

type pcg_state_setseq_128 struct {
	state pcg128_t
	inc   pcg128_t
}

const PCG_DEFAULT_MULTIPLIER_HIGH = uint64(2549297995355413924)
const PCG_DEFAULT_MULTIPLIER_LOW = uint64(4865540595714422341)

var PCG_DEFAULT_MULTIPLIER_128 = PCG_128BIT_CONSTANT(PCG_DEFAULT_MULTIPLIER_HIGH, PCG_DEFAULT_MULTIPLIER_LOW)
var PCG_DEFAULT_INCREMENT_128 = PCG_128BIT_CONSTANT(uint64(6364136223846793005), uint64(1442695040888963407))

var PCG_STATE_SETSEQ_128_INITIALIZER = []pcg128_t{
	PCG_128BIT_CONSTANT(0x979c9a98d8462005, 0x7d3e9cb6cfe0549b),
	PCG_128BIT_CONSTANT(0x0000000000000001, 0xda3e39cb94b95bdb),
}

const PCG_CHEAP_MULTIPLIER_128 = uint64(0xda942042e4dd58b5)

func pcg_rotr_64(value uint64, rot uint) uint64 {
	return (value >> rot) | (value << ((-rot) & 63))
}

func pcg128_add(a pcg128_t, b pcg128_t) pcg128_t {
	var result pcg128_t

	result.low = a.low + b.low
	result.high = a.high + b.high
	if result.low < b.low {
		result.high++
	}
	return result
}

func _pcg_mult64(x uint64, y uint64, z1 *uint64, z0 *uint64) {
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

func pcg128_mult(a pcg128_t, b pcg128_t) pcg128_t {
	var h1 uint64
	var result pcg128_t

	h1 = a.high*b.low + a.low*b.high
	_pcg_mult64(a.low, b.low, &(result.high), &(result.low))
	result.high += h1
	return result
}

func pcg_setseq_128_step_r(rng *pcg_state_setseq_128) {
	rng.state = pcg128_add(pcg128_mult(rng.state, PCG_DEFAULT_MULTIPLIER_128), rng.inc)
}

func pcg_output_xsl_rr_128_64(state pcg128_t) uint64 {
	return pcg_rotr_64(state.high^state.low, uint(state.high>>58))
}

func pcg_setseq_128_srandom_r(rng *pcg_state_setseq_128, initstate pcg128_t, initseq pcg128_t) {
	rng.state = PCG_128BIT_CONSTANT(0, 0)
	rng.inc.high = initseq.high << 1
	rng.inc.high |= initseq.low >> 63
	rng.inc.low = (initseq.low << 1) | 1
	pcg_setseq_128_step_r(rng)
	rng.state = pcg128_add(rng.state, initstate)
	pcg_setseq_128_step_r(rng)
}

func pcg_setseq_128_xsl_rr_64_random_r(rng *pcg_state_setseq_128) uint64 {
	pcg_setseq_128_step_r(rng)
	return pcg_output_xsl_rr_128_64(rng.state)
}

func pcg128_mult_64(a pcg128_t, b uint64) pcg128_t {
	var h1 uint64
	var result pcg128_t

	h1 = a.high * b
	_pcg_mult64(a.low, b, &(result.high), &(result.low))
	result.high += h1
	return result
}

func pcg_cm_step_r(rng *pcg_state_setseq_128) {
	rng.state = pcg128_add(pcg128_mult_64(rng.state, PCG_CHEAP_MULTIPLIER_128), rng.inc)
}

func pcg_cm_srandom_r(rng *pcg_state_setseq_128, initstate pcg128_t, initseq pcg128_t) {
	rng.state = PCG_128BIT_CONSTANT(0, 0)
	rng.inc.high = initseq.high << 1
	rng.inc.high |= initseq.low >> 63
	rng.inc.low = (initseq.low << 1) | 1
	pcg_cm_step_r(rng)
	rng.state = pcg128_add(rng.state, initstate)
	pcg_cm_step_r(rng)
}

func pcg_cm_random_r(rng *pcg_state_setseq_128) uint64 {
	// Lots of manual inlining to help out certain compilers to generate performant code.
	hi := rng.state.high
	lo := rng.state.low

	// Run the DXSM output function on the pre-iterated state.
	lo |= 1
	hi ^= hi >> 32
	hi *= 0xda942042e4dd58b5
	hi ^= hi >> 48
	hi *= lo

	// Run the CM step.
	rng.state = pcg128_add(pcg128_mult_64(rng.state, PCG_CHEAP_MULTIPLIER_128),
		rng.inc)
	return hi
}

func pcg_setseq_128_xsl_rr_64_boundedrand_r(rng *pcg_state_setseq_128, bound uint64) uint64 {
	threshold := -bound % bound
	for {
		r := pcg_setseq_128_xsl_rr_64_random_r(rng)
		if r >= threshold {
			return r % bound
		}
	}
}

func pcg_setseq_128_advance_r(rng *pcg_state_setseq_128, delta pcg128_t) {
	rng.state = pcg_advance_lcg_128(rng.state, delta, PCG_DEFAULT_MULTIPLIER_128, rng.inc)
}

func pcg_cm_advance_r(rng *pcg_state_setseq_128, delta pcg128_t) {
	rng.state = pcg_advance_lcg_128(rng.state, delta, PCG_128BIT_CONSTANT(0, PCG_CHEAP_MULTIPLIER_128), rng.inc)
}

type pcg64_random_t = pcg_state_setseq_128
var pcg64_random_r = pcg_setseq_128_xsl_rr_64_random_r
var pcg64_boundedrand_r = pcg_setseq_128_xsl_rr_64_boundedrand_r
var pcg64_srandom_r = pcg_setseq_128_srandom_r
var pcg64_advance_r = pcg_setseq_128_advance_r
var PCG64_INITIALIZER = PCG_STATE_SETSEQ_128_INITIALIZER

type pcg64_state struct {
	pcg_state  *pcg64_random_t
	has_uint32 int
	uinteger   uint32
}

func pcg64_next64(state *pcg64_state) uint64 {
	return pcg64_random_r(state.pcg_state)
}

func pcg64_next32(state *pcg64_state) uint32 {
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

func pcg64_cm_next64(state *pcg64_state) uint64 {
	return pcg_cm_random_r(state.pcg_state)
}

func pcg64_cm_next32(state *pcg64_state) uint32 {
	var next uint64
	if state.has_uint32 != 0 {
		state.has_uint32 = 0
		return state.uinteger
	}
	next = pcg_cm_random_r(state.pcg_state)
	state.has_uint32 = 1
	state.uinteger = uint32(next >> 32)
	return uint32(next & 0xffffffff)
}

void pcg64_advance(state *pcg64_state, uint64_t *step);
void pcg64_cm_advance(state *pcg64_state, uint64_t *step);

void pcg64_set_seed(state *pcg64_state, uint64_t *seed, uint64_t *inc);

void pcg64_get_state(state *pcg64_state, uint64_t *state_arr, int *has_uint32,
uint32_t *uinteger);

void pcg64_set_state(state *pcg64_state, uint64_t *state_arr, int has_uint32,
uint32_t uinteger);

#endif /* PCG64_H_INCLUDED */

extern inline void pcg_setseq_128_step_r(rng *pcg_state_setseq_128);
extern inline uint64_t pcg_output_xsl_rr_128_64(pcg128_t state);
extern inline void pcg_setseq_128_srandom_r(rng *pcg_state_setseq_128,
pcg128_t initstate,
pcg128_t initseq);
extern inline uint64_t
pcg_setseq_128_xsl_rr_64_random_r(rng *pcg_state_setseq_128);
extern inline uint64_t
pcg_setseq_128_xsl_rr_64_boundedrand_r(rng *pcg_state_setseq_128,
uint64_t bound);
extern inline void pcg_setseq_128_advance_r(rng *pcg_state_setseq_128,
pcg128_t delta);
extern inline uint64_t pcg_cm_random_r(rng *pcg_state_setseq_128);
extern inline void pcg_cm_step_r(rng *pcg_state_setseq_128);
extern inline uint64_t pcg_output_cm_128_64(pcg128_t state);
extern inline void pcg_cm_srandom_r(rng *pcg_state_setseq_128, pcg128_t initstate, pcg128_t initseq);

/* Multi-step advance functions (jump-ahead, jump-back)
 *
 * The method used here is based on Brown, "Random Number Generation
 * with Arbitrary Stride,", Transactions of the American Nuclear
 * Society (Nov. 1994).  The algorithm is very similar to fast
 * exponentiation.
 *
 * Even though delta is an unsigned integer, we can pass a
 * signed integer to go backwards, it just goes "the long way round".
 */

#ifndef PCG_EMULATED_128BIT_MATH

pcg128_t pcg_advance_lcg_128(pcg128_t state, pcg128_t delta, pcg128_t cur_mult,
pcg128_t cur_plus) {
pcg128_t acc_mult = 1u;
pcg128_t acc_plus = 0u;
while (delta > 0) {
if (delta & 1) {
acc_mult *= cur_mult;
acc_plus = acc_plus * cur_mult + cur_plus;
}
cur_plus = (cur_mult + 1) * cur_plus;
cur_mult *= cur_mult;
delta /= 2;
}
return acc_mult * state + acc_plus;
}

#else

pcg128_t pcg_advance_lcg_128(pcg128_t state, pcg128_t delta, pcg128_t cur_mult,
pcg128_t cur_plus) {
pcg128_t acc_mult = PCG_128BIT_CONSTANT(0u, 1u);
pcg128_t acc_plus = PCG_128BIT_CONSTANT(0u, 0u);
while ((delta.high > 0) || (delta.low > 0)) {
if (delta.low & 1) {
acc_mult = pcg128_mult(acc_mult, cur_mult);
acc_plus = pcg128_add(pcg128_mult(acc_plus, cur_mult), cur_plus);
}
cur_plus = pcg128_mult(pcg128_add(cur_mult, PCG_128BIT_CONSTANT(0u, 1u)),
cur_plus);
cur_mult = pcg128_mult(cur_mult, cur_mult);
delta.low = (delta.low >> 1) | (delta.high << 63);
delta.high >>= 1;
}
return pcg128_add(pcg128_mult(acc_mult, state), acc_plus);
}

#endif

extern inline uint64_t pcg64_next64(state *pcg64_state);
extern inline uint32_t pcg64_next32(state *pcg64_state);

extern inline uint64_t pcg64_cm_next64(state *pcg64_state);
extern inline uint32_t pcg64_cm_next32(state *pcg64_state);

extern void pcg64_advance(state *pcg64_state, uint64_t *step) {
pcg128_t delta;
#ifndef PCG_EMULATED_128BIT_MATH
delta = (((pcg128_t)step[0]) << 64) | step[1];
#else
delta.high = step[0];
delta.low = step[1];
#endif
pcg64_advance_r(state.pcg_state, delta);
}

extern void pcg64_cm_advance(state *pcg64_state, uint64_t *step) {
pcg128_t delta;
#ifndef PCG_EMULATED_128BIT_MATH
delta = (((pcg128_t)step[0]) << 64) | step[1];
#else
delta.high = step[0];
delta.low = step[1];
#endif
pcg_cm_advance_r(state.pcg_state, delta);
}

extern void pcg64_set_seed(state *pcg64_state, uint64_t *seed, uint64_t *inc) {
pcg128_t s, i;
#ifndef PCG_EMULATED_128BIT_MATH
s = (((pcg128_t)seed[0]) << 64) | seed[1];
i = (((pcg128_t)inc[0]) << 64) | inc[1];
#else
s.high = seed[0];
s.low = seed[1];
i.high = inc[0];
i.low = inc[1];
#endif
pcg64_srandom_r(state.pcg_state, s, i);
}

extern void pcg64_get_state(state *pcg64_state, uint64_t *state_arr,
int *has_uint32, uint32_t *uinteger) {
/*
 * state_arr contains state.high, state.low, inc.high, inc.low
 *    which are interpreted as the upper 64 bits (high) or lower
 *    64 bits of a uint128_t variable
 *
 */
#ifndef PCG_EMULATED_128BIT_MATH
state_arr[0] = (uint64_t)(state.pcg_state.state >> 64);
state_arr[1] = (uint64_t)(state.pcg_state.state & 0xFFFFFFFFFFFFFFFFULL);
state_arr[2] = (uint64_t)(state.pcg_state.inc >> 64);
state_arr[3] = (uint64_t)(state.pcg_state.inc & 0xFFFFFFFFFFFFFFFFULL);
#else
state_arr[0] = (uint64_t)state.pcg_state.state.high;
state_arr[1] = (uint64_t)state.pcg_state.state.low;
state_arr[2] = (uint64_t)state.pcg_state.inc.high;
state_arr[3] = (uint64_t)state.pcg_state.inc.low;
#endif
has_uint32[0] = state.has_uint32;
uinteger[0] = state.uinteger;
}

extern void pcg64_set_state(state *pcg64_state, uint64_t *state_arr,
int has_uint32, uint32_t uinteger) {
/*
 * state_arr contains state.high, state.low, inc.high, inc.low
 *    which are interpreted as the upper 64 bits (high) or lower
 *    64 bits of a uint128_t variable
 *
 */
#ifndef PCG_EMULATED_128BIT_MATH
state.pcg_state.state = (((pcg128_t)state_arr[0]) << 64) | state_arr[1];
state.pcg_state.inc = (((pcg128_t)state_arr[2]) << 64) | state_arr[3];
#else
state.pcg_state.state.high = state_arr[0];
state.pcg_state.state.low = state_arr[1];
state.pcg_state.inc.high = state_arr[2];
state.pcg_state.inc.low = state_arr[3];
#endif
state.has_uint32 = has_uint32;
state.uinteger = uinteger;
}

