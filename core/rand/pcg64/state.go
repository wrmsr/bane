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

type State struct {
	rng   *Rng
	hasUi int
	ui    uint32
}

func New() *State {
	return &State{
		rng: &Rng{},
	}
}

func (st *State) Next64() uint64 {
	return st.rng.Rand()
}

func (st *State) Next32() uint32 {
	var next uint64
	if st.hasUi != 0 {
		st.hasUi = 0
		return st.ui
	}
	next = st.rng.Rand()
	st.hasUi = 1
	st.ui = uint32(next >> 32)
	return uint32(next & 0xffffffff)
}

func (st *State) CheapNext64() uint64 {
	return st.rng.CheapRand()
}

func (st *State) CheapNext32() uint32 {
	var next uint64
	if st.hasUi != 0 {
		st.hasUi = 0
		return st.ui
	}
	next = st.rng.CheapRand()
	st.hasUi = 1
	st.ui = uint32(next >> 32)
	return uint32(next & 0xffffffff)
}

func (st *State) Advance(step []uint64) {
	var delta Uint128
	delta.High = step[0]
	delta.Low = step[1]
	st.rng.Advance(delta)
}

func (st *State) CheapAdvance(step []uint64) {
	var delta Uint128
	delta.High = step[0]
	delta.Low = step[1]
	st.rng.CheapAdvance(delta)
}

func (st *State) Seed(seed []uint64, inc []uint64) {
	var s, i Uint128
	s.High = seed[0]
	s.Low = seed[1]
	i.High = inc[0]
	i.Low = inc[1]
	st.rng.Seed(s, i)
}

func (st *State) GetState(state []uint64, hasUi *int, ui *uint32) {
	// state_arr contains state.high, state.low, inc.high, inc.low which are interpreted as the upper 64 bits (high) or
	// lower 64 bits of a uint128_t variable
	state[0] = st.rng.state.High
	state[1] = st.rng.state.Low
	state[2] = st.rng.inc.High
	state[3] = st.rng.inc.Low
	*hasUi = st.hasUi
	*ui = st.ui
}

func (st *State) SetState(state []uint64, hasUi int, ui uint32) {
	// state_arr contains state.high, state.low, inc.high, inc.low which are interpreted as the upper 64 bits (high) or
	// lower 64 bits of a uint128_t variable
	st.rng.state.High = state[0]
	st.rng.state.Low = state[1]
	st.rng.inc.High = state[2]
	st.rng.inc.Low = state[3]
	st.hasUi = hasUi
	st.ui = ui
}
