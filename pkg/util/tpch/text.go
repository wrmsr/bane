package tpch

import "github.com/wrmsr/bane/pkg/util/check"

type textDist struct {
	name string

	values  []string
	weights []int

	maxWeight int

	seq      []string
	bytesSeq [][]byte
}

func newTextDist(
	name string,
	weightsByValue map[string]int,
) *textDist {
	td := &textDist{
		name: name,

		values:  make([]string, len(weightsByValue)),
		weights: make([]int, len(weightsByValue)),
	}

	runningWeight := 0
	isValid := true
	i := 0
	for k, v := range weightsByValue {
		td.values[i] = k
		runningWeight += v
		td.weights[i] = runningWeight
		isValid = isValid && v > 0
		i++
	}

	//  "nations" is hack and not a valid dist so we need to skip it
	if isValid {
		td.maxWeight = td.weights[len(td.weights)-1]
		td.seq = make([]string, td.maxWeight)
		td.bytesSeq = make([][]byte, td.maxWeight)

		i := 0
		for _, value := range td.values {
			bytesValue := []byte(value)
			check.Equal(len(bytesValue), len(value))
			for j := weightsByValue[value]; j > 0; j-- {
				td.seq[i] = value
				td.bytesSeq[i] = bytesValue
				i++
			}
		}

	} else {
		td.maxWeight = -1
	}

	return td
}

func (td *textDist) randomValue(gen gen) string {
	randomValue := gen.rand(0, int64(td.maxWeight-1))
	return td.seq[randomValue]
}

func (td *textDist) randomBytesValue(gen gen) []byte {
	randomValue := gen.rand(0, int64(td.maxWeight-1))
	return td.bytesSeq[randomValue]
}
