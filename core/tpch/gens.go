package tpch

import (
	"fmt"

	its "github.com/wrmsr/bane/core/iterators"
	bt "github.com/wrmsr/bane/core/types"
)

func calculateRowCount(scaleBase int, scaleFactor float32, part int, partCount int) int {
	totalRowCount := int(float32(scaleBase) * scaleFactor)
	rowCount := totalRowCount / partCount
	if part == partCount {
		// for the last part, add the remainder rows
		rowCount += totalRowCount % partCount
	}
	return rowCount
}

func calculateStartIndex(scaleBase int, scaleFactor float32, part int, partCount int) int {
	totalRowCount := int(float32(scaleBase) * scaleFactor)
	rowsPerPart := totalRowCount // partCount
	return rowsPerPart * (part - 1)
}

const (
	generatedDateEpochOffset = 83966 // The value of 1970-01-01 in the date generator system
	minGenerateDate          = 92001
	currentDate              = 95168
	totalDateRange           = 2557
)

var monthYearDayStart = []int{
	0,
	31,
	59,
	90,
	120,
	151,
	181,
	212,
	243,
	273,
	304,
	334,
	365,
}

func isLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0
}

func leapYearAdjustment(year, month int) int {
	if isLeapYear(year) && (month) >= 2 {
		return 1
	}
	return 0
}

func julian(date int) int {
	offset := date - minGenerateDate
	result := minGenerateDate
	for {
		year := result / 1000
		yearEnd := year*1000 + 365
		if isLeapYear(year) {
			yearEnd++
		}
		if result+offset <= yearEnd {
			break
		}
		offset -= yearEnd - result + 1
		result += 1000
	}
	return result + offset
}

func makeDate(index int) string {
	y := julian(index+minGenerateDate-1) / 1000
	d := julian(index+minGenerateDate-1) % 1000
	m := 0
	for d > monthYearDayStart[m]+leapYearAdjustment(y, m) {
		m++
	}
	dy := d - monthYearDayStart[m-1]
	if isLeapYear(y) && m > 2 {
		dy--
	}
	return fmt.Sprintf("19%02d-%02d-%02d", y, m, dy)
}

var dateStringIndex = its.Seq(its.Map(its.Range(0, totalDateRange, 1), func(i int) string {
	return makeDate(i + 1)
}))

func toEpochDate(generated_date int) int {
	return generated_date - generatedDateEpochOffset
}

func formatDate(epoch_date int) string {
	return dateStringIndex[epoch_date-(minGenerateDate-generatedDateEpochOffset)]
}

func isInPast(date int) bool {
	return julian(date) <= currentDate
}

const regionCommentAverageLength = 72

func generateRegions(
	textDists textDists,
	textPool *textPool,
) bt.Iterable[map[string]any] {
	if textDists == nil {
		textDists = getTextDists()
	}
	if textPool == nil {
		textPool = getDefaultTextPool()
	}

	commentRandom := newRandomText(1500869201, textPool, regionCommentAverageLength, 1)

	return its.Map(its.Range(0, textDists.Regions().size(), 1), func(i int) map[string]any {
		r := map[string]any{
			"r_regionkey": i,
			"r_name":      textDists.Regions().values[i],
			"r_comment":   commentRandom.nextValue(),
		}
		commentRandom.rowFinished()
		return r
	})
}
