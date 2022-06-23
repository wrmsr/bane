package tpch

import (
	"errors"
	"fmt"
	"strings"

	mathu "github.com/wrmsr/bane/pkg/util/math"
)

//

type byRow interface {
	rowFinished()
	advanceRows(rowCount int)
}

type gen interface {
	byRow

	next() int64
	rand(lowValue, highValue int64) int64
}

//

type intGen struct {
	seed                int64
	expectedUsagePerRow int

	usage int
}

func newIntGen(seed int64, expectedUsagePerRow int) *intGen {
	return &intGen{
		seed:                seed,
		expectedUsagePerRow: expectedUsagePerRow,
	}
}

var _ gen = &intGen{}

const (
	intGenMultiplier = int64(16807)
	intGenModulus    = int64(2147483647)
)

func (g *intGen) next() int64 {
	if g.usage >= g.expectedUsagePerRow {
		panic(errors.New("unexpected usages"))
	}
	g.seed = (g.seed * intGenMultiplier) % intGenModulus
	g.usage++
	return g.seed
}

func (g *intGen) rand(lowValue, highValue int64) int64 {
	g.next()
	intRange := highValue - lowValue + 1
	doubleRange := float64(intRange)
	valueInRange := int32((1.0 * float64(g.seed) / float64(intGenModulus)) * doubleRange)
	return int64(int32(lowValue) + valueInRange)
}

func (g *intGen) advanceSeed(count int) {
	multiplier := intGenMultiplier
	for count > 0 {
		if count%2 != 0 {
			g.seed = (multiplier * g.seed) % intGenModulus
		}
		count = count / 2
		multiplier = (multiplier * multiplier) % intGenModulus
	}
}

func (g *intGen) rowFinished() {
	g.advanceSeed(g.expectedUsagePerRow - g.usage)
	g.usage = 0
}

func (g *intGen) advanceRows(rowCount int) {
	if g.usage > 0 {
		g.rowFinished()
	}
	g.advanceSeed(g.expectedUsagePerRow * rowCount)
}

//

type longGen struct {
	seed                int64
	expectedUsagePerRow int

	usage int
}

func newLongGen(seed int64, expectedUsagePerRow int) *longGen {
	return &longGen{
		seed:                seed,
		expectedUsagePerRow: expectedUsagePerRow,
	}
}

var _ gen = &longGen{}

const (
	longGenMultiplier = 6364136223846793005
	longGenIncrement  = 1
)

func (g *longGen) next() int64 {
	if g.usage >= g.expectedUsagePerRow {
		panic(errors.New("unexpected usages"))
	}
	g.seed = (g.seed * longGenMultiplier) + longGenIncrement
	g.usage++
	return g.seed
}

func (g *longGen) rand(lowValue, highValue int64) int64 {
	g.next()
	valueInRange := mathu.AbsInt64(g.seed) % (highValue - lowValue + 1)
	return lowValue + valueInRange
}

const (
	_LONG_GEN_MULTIPLIER_32 = int64(16807)
	_LONG_GEN_MODULUS_32    = int64(2147483647)
)

func (g *longGen) advanceSeed(count int) {
	multiplier := _LONG_GEN_MULTIPLIER_32
	for count > 0 {
		if count%2 != 0 {
			g.seed = (multiplier * g.seed) % _LONG_GEN_MODULUS_32
		}
		count = count / 2
		multiplier = (multiplier * multiplier) % _LONG_GEN_MODULUS_32
	}
}

func (g *longGen) rowFinished() {
	g.advanceSeed(g.expectedUsagePerRow - g.usage)
	g.usage = 0
}

func (g *longGen) advanceRows(rowCount int) {
	if g.usage > 0 {
		g.rowFinished()
	}
	g.advanceSeed(g.expectedUsagePerRow * rowCount)
}

//

type genRandom struct {
	gen gen
}

var _ byRow = genRandom{}

func (g genRandom) rowFinished() {
	g.gen.rowFinished()
}

func (g genRandom) advanceRows(rowCount int) {
	g.gen.advanceRows(rowCount)
}

//

type randomAlphaNumeric struct {
	genRandom

	minLength int
	maxLength int
}

var _ALPHA_NUMERIC = strings.Split("0123456789abcdefghijklmnopqrstuvwxyz ABCDEFGHIJKLMNOPQRSTUVWXYZ,", "")

const (
	alphaNumericLowLengthMultiplier  = 0.4
	alphaNumericHighLengthMultiplier = 1.6

	alphaNumericUsagePerRow = 9
)

func newRandomAlphaNumeric(
	seed int64,
	averageLength int,
	expectedRowCount int, // = 1,
) *randomAlphaNumeric {
	return &randomAlphaNumeric{
		genRandom: genRandom{
			gen: newIntGen(seed, alphaNumericUsagePerRow*expectedRowCount),
		},

		minLength: int(float64(averageLength) * alphaNumericLowLengthMultiplier),
		maxLength: int(float64(averageLength) * alphaNumericHighLengthMultiplier),
	}
}

func (r *randomAlphaNumeric) nextValue() string {
	sz := int(r.gen.rand(int64(r.minLength), int64(r.maxLength)))
	var sb strings.Builder
	charIndex := 0
	for i := 0; i < sz; i++ {
		if i%5 == 0 {
			charIndex = int(r.gen.rand(0, 0x7FFFFFFF))
		}
		sb.WriteString(_ALPHA_NUMERIC[charIndex&0x3F])
		charIndex = (charIndex >> 6) & 0x7FFFFFFF
	}

	return sb.String()
}

//

type randomBoundedInt struct {
	genRandom

	lowValue  int64
	highValue int64
}

func newRandomBoundedInt(
	seed int64,
	lowValue int,
	highValue int,
	expectedRowCount int, // = 1,
) *randomBoundedInt {
	return &randomBoundedInt{
		genRandom: genRandom{
			gen: newIntGen(seed, expectedRowCount),
		},

		lowValue:  int64(lowValue),
		highValue: int64(highValue),
	}
}

func (r *randomBoundedInt) nextValue() int32 {
	return int32(r.gen.rand(r.lowValue, r.highValue))
}

//

type randomBoundedLong struct {
	genRandom

	lowValue  int64
	highValue int64
}

func newRandomBoundedLong(
	seed int64,
	use64bits bool,
	lowValue int,
	highValue int,
	expectedRowCount int, // = 1,
) *randomBoundedLong {
	var gen gen
	if use64bits {
		gen = newLongGen(seed, expectedRowCount)
	} else {
		gen = newIntGen(seed, expectedRowCount)
	}

	return &randomBoundedLong{
		genRandom: genRandom{
			gen: gen,
		},

		lowValue:  int64(lowValue),
		highValue: int64(highValue),
	}
}

func (r *randomBoundedLong) nextValue() int64 {
	return r.gen.rand(r.lowValue, r.highValue)
}

//

type randomPhoneNumber struct {
	genRandom
}

func newRandomPhoneNumber(
	seed int64,
	expectedRowCount int, // = 1,
) *randomPhoneNumber {
	return &randomPhoneNumber{
		genRandom: genRandom{
			gen: newIntGen(seed, 3*expectedRowCount),
		},
	}
}

const phoneNumberNationsMax = 90 // limited by country codes in phone numbers

func (r *randomPhoneNumber) nextValue(nationKey int) string {
	return fmt.Sprintf(
		"%02d-%03d-%03d-%04d",
		10+(nationKey%phoneNumberNationsMax),
		r.gen.rand(100, 999),
		r.gen.rand(100, 999),
		r.gen.rand(1000, 9999),
	)
}
