package pcg64

type Rng struct {
	state Uint128
	inc   Uint128
}

var defaultMultiplier128 = Uint128{uint64(2549297995355413924), uint64(4865540595714422341)}
var defaultIncrement128 = Uint128{uint64(6364136223846793005), uint64(1442695040888963407)}

var defaultInitializer = []Uint128{
	{0x979c9a98d8462005, 0x7d3e9cb6cfe0549b},
	{0x0000000000000001, 0xda3e39cb94b95bdb},
}

func (rng *Rng) Step() {
	rng.state = Add128(Mul128(rng.state, defaultMultiplier128), rng.inc)
}

func (rng *Rng) Seed(initState Uint128, initSeq Uint128) {
	rng.state = Uint128{0, 0}
	rng.inc.High = initSeq.High << 1
	rng.inc.High |= initSeq.Low >> 63
	rng.inc.Low = (initSeq.Low << 1) | 1
	rng.Step()
	rng.state = Add128(rng.state, initState)
	rng.Step()
}

const cheapMultiplier = uint64(0xda942042e4dd58b5)

func (rng *Rng) CheapStep() {
	rng.state = Add128(Mul128Low(rng.state, cheapMultiplier), rng.inc)
}

func (rng *Rng) CheapSeed(initState Uint128, initSeq Uint128) {
	rng.state = Uint128{0, 0}
	rng.inc.High = initSeq.High << 1
	rng.inc.High |= initSeq.Low >> 63
	rng.inc.Low = (initSeq.Low << 1) | 1
	rng.CheapStep()
	rng.state = Add128(rng.state, initState)
	rng.CheapStep()
}

func (rng *Rng) CheapRand() uint64 {
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
	rng.state = Add128(Mul128Low(rng.state, cheapMultiplier), rng.inc)
	return hi
}

func (rng *Rng) Advance(delta Uint128) {
	rng.state = AdvanceLcg128(rng.state, delta, defaultMultiplier128, rng.inc)
}

func (rng *Rng) CheapAdvance(delta Uint128) {
	rng.state = AdvanceLcg128(rng.state, delta, Uint128{0, cheapMultiplier}, rng.inc)
}

func (rng *Rng) Rand() uint64 {
	rng.Step()
	return Ror64(rng.state.High^rng.state.Low, uint(rng.state.High>>58))
}

func (rng *Rng) BoundRand(bound uint64) uint64 {
	threshold := -bound % bound
	for {
		r := rng.Rand()
		if r >= threshold {
			return r % bound
		}
	}
}
