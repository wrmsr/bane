package tpch

import (
	"errors"
	"strings"
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
	_INT_GEN_MULTIPLIER = 16807
	_INT_GEN_MODULUS    = 2147483647
)

func (g *intGen) next() int64 {
	if g.usage >= g.expectedUsagePerRow {
		panic(errors.New("unexpected usages"))
	}
	g.seed = (g.seed * _INT_GEN_MULTIPLIER) % _INT_GEN_MODULUS
	g.usage++
	return g.seed
}

func (g *intGen) rand(lowValue, highValue int64) int64 {
	g.next()
	intRange := int32(highValue) - int32(lowValue) + 1
	doubleRange := float64(intRange)
	valueInRange := int32((1.0 * float64(g.seed) / float64(_INT_GEN_MODULUS)) * doubleRange)
	return int64(int32(lowValue) + valueInRange)
}

func (g *intGen) advanceSeed(count int) {
	multiplier := _INT_GEN_MULTIPLIER
	for count > 0 {
		if count%2 != 0 {
			g.seed = (int64(multiplier) * g.seed) % _INT_GEN_MODULUS
		}
		count = count / 2
		multiplier = (multiplier * multiplier) % _INT_GEN_MODULUS
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

//type longGen struct {
//	seed                int
//	expectedUsagePerRow int
//
//	usage int
//}
//
//func newLongGen(seed, expectedUsagePerRow int) *longGen {
//	return &longGen{
//		seed:                seed,
//		expectedUsagePerRow: expectedUsagePerRow,
//	}
//}
//
//var _ gen = &longGen{}
//
//const (
//	_LONG_GEN_MULTIPLIER = 6364136223846793005
//	_LONG_GEN_INCREMENT  = 1
//)
//
//func (g *longGen) next() int {
//	if g.usage >= g.expectedUsagePerRow {
//		panic(errors.New("unexpected usages"))
//	}
//	g.seed = (g.seed * _LONG_GEN_MULTIPLIER) + _LONG_GEN_INCREMENT
//	g.usage++
//	return g.seed
//}
//
//func (g *longGen) rand(lowValue, highValue int) int {
//	g.next()
//	valueInRange := mathu.AbsInt(g.seed) % (highValue - lowValue + 1)
//	return lowValue + valueInRange
//}
//
//const (
//	_LONG_GEN_MULTIPLIER_32 = 16807
//	_LONG_GEN_MODULUS_32    = 2147483647
//)
//
//func (g *longGen) advanceSeed(count int) {
//	multiplier := _LONG_GEN_MULTIPLIER_32
//	for count > 0 {
//		if count%2 != 0 {
//			g.seed = (multiplier * g.seed) % _LONG_GEN_MODULUS_32
//		}
//		count = count / 2
//		multiplier = (multiplier * multiplier) % _LONG_GEN_MODULUS_32
//	}
//}
//
//func (g *longGen) rowFinished() {
//	g.advanceSeed(g.expectedUsagePerRow - g.usage)
//	g.usage = 0
//}
//
//func (g *longGen) advanceRows(rowCount int) {
//	if g.usage > 0 {
//		g.rowFinished()
//	}
//	g.advanceSeed(g.expectedUsagePerRow * rowCount)
//}

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
	_ALPHA_NUMERIC_LOW_LENGTH_MULTIPLIER  = 0.4
	_ALPHA_NUMERIC_HIGH_LENGTH_MULTIPLIER = 1.6

	_ALPHA_NUMERIC_USAGE_PER_ROW = 9
)

func newRandomAlphaNumeric(
	seed int64,
	averageLength int,
	expectedRowCount int, // = 1,
) *randomAlphaNumeric {
	return &randomAlphaNumeric{
		genRandom: genRandom{
			gen: newIntGen(seed, _ALPHA_NUMERIC_USAGE_PER_ROW*expectedRowCount),
		},

		minLength: int(float64(averageLength) * _ALPHA_NUMERIC_LOW_LENGTH_MULTIPLIER),
		maxLength: int(float64(averageLength) * _ALPHA_NUMERIC_HIGH_LENGTH_MULTIPLIER),
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

/*
class RandomBoundedInt(GenRandom):

    def __init__(
            self,
            seed: int,
            lowValue: int,
            highValue: int,
            expectedRowCount: int = 1,
    ) -> None:
        super().__init__(IntGen(seed, expectedRowCount))

        self._lowValue = lowValue
        self._highValue = highValue

    def next_value(self) -> int:
        return self._gen.rand(self._lowValue, self._highValue)


class RandomBoundedLong(GenRandom):

    def __init__(
            self,
            seed: int,
            use_64_bits: bool,
            lowValue: int,
            highValue: int,
            expectedRowCount: int = 1,
    ) -> None:
        gen_cls = LongGen if use_64_bits else IntGen
        super().__init__(gen_cls(seed, expectedRowCount))

        self._lowValue = lowValue
        self._highValue = highValue

    def next_value(self) -> int:
        return self._gen.rand(self._lowValue, self._highValue)


class RandomPhoneNumber(GenRandom):

    _NATIONS_MAX = 90  # limited by country codes in phone numbers

    def __init__(self, seed: int, expectedRowCount: int = 1) -> None:
        super().__init__(IntGen(seed, 3 * expectedRowCount))

    def next_value(self, nation_key: int) -> str:
        return '%02d-%03d-%03d-%04d' % (
            (10 + (nation_key % self._NATIONS_MAX)),
            self._gen.rand(100, 999),
            self._gen.rand(100, 999),
            self._gen.rand(1000, 9999),
        )

*/
