package tpch

import (
	"fmt"
	"strings"

	au "github.com/wrmsr/bane/pkg/util/atomic"
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
)

//

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

func (td *textDist) size() int {
	return len(td.values)
}

//

type randomString struct {
	genRandom

	dist *textDist
}

func newRandomString(
	seed int64,
	dist *textDist,
	expectedRowCount int, // = 1
) *randomString {
	return &randomString{
		genRandom: genRandom{
			gen: newIntGen(seed, expectedRowCount),
		},

		dist: dist,
	}
}

func (r *randomString) nextValue() string {
	return r.dist.randomValue(r.gen)
}

//

type randomStringSequence struct {
	genRandom

	count int
	dist  *textDist
}

func newRandomStringSequence(
	seed int64,
	count int,
	dist *textDist,
	expectedRowCount int, // = 1
) *randomStringSequence {
	return &randomStringSequence{
		genRandom: genRandom{
			gen: newIntGen(seed, expectedRowCount),
		},

		count: count,
		dist:  dist,
	}
}

func (r *randomStringSequence) nextValue() string {
	check.Condition(r.count < len(r.dist.values))
	values := slices.Copy(r.dist.values)

	// randomize first 'count' elements of the string
	for currentPos := 0; currentPos < r.count; currentPos++ {
		swapPos := r.gen.rand(int64(currentPos), int64(len(values)-1))
		values[currentPos], values[swapPos] = values[swapPos], values[currentPos]
	}

	return strings.Join(values[:r.count], " ")
}

//

type textPoolGen struct {
	size              int
	maxSentenceLength int

	dists textDists

	buf  []byte
	pos  int
	last byte
	gen  gen
}

func genTextPool(
	size int,
	maxSentenceLength int,
	dists textDists,
) []byte {
	g := &textPoolGen{
		size:              size,
		maxSentenceLength: maxSentenceLength,

		dists: dists,

		buf: make([]byte, size+maxSentenceLength),
		gen: newIntGen(933588178, 0x7FFFFFFF),
	}

	i := 0
	for g.pos < size {
		g.generateSentence()
		i++
	}
	g.buf = g.buf[:size]

	return g.buf
}

func (g *textPoolGen) write(b []byte) {
	if len(b) < 1 {
		return
	}
	bl := len(b)
	copy(g.buf[g.pos:g.pos+bl], b)
	g.pos += bl
	g.last = b[bl-1]
}

func (g *textPoolGen) writeString(s string) {
	g.write([]byte(s))
}

func (g *textPoolGen) erase(i int) {
	check.Condition(i <= g.pos)
	g.pos--
	if g.pos > 0 {
		g.last = g.buf[g.pos-1]
	} else {
		g.last = 0
	}
}

func (g *textPoolGen) rand(dist *textDist) []byte {
	return dist.randomBytesValue(g.gen)
}

func (g *textPoolGen) generateVerbPhrase() {
	syntax := g.rand(g.dists.VerbPhrase())
	for i := 0; i < len(syntax); i += 2 {
		var source *textDist
		if syntax[i] == 'D' {
			source = g.dists.Adverbs()
		} else if syntax[i] == 'V' {
			source = g.dists.Verbs()
		} else if syntax[i] == 'X' {
			source = g.dists.Auxiliaries()
		} else {
			panic(fmt.Errorf("unknown token: %v", syntax[i]))
		}
		word := g.rand(source)
		g.write(word)
		g.writeString(" ")
	}
}

func (g *textPoolGen) generateNounPhrase() {
	syntax := g.rand(g.dists.NounPhrase())
	for i := 0; i < len(syntax); i++ {
		var source *textDist
		if syntax[i] == 'A' {
			source = g.dists.Articles()
		} else if syntax[i] == 'J' {
			source = g.dists.Adjectives()
		} else if syntax[i] == 'D' {
			source = g.dists.Adverbs()
		} else if syntax[i] == 'N' {
			source = g.dists.Nouns()
		} else if syntax[i] == ',' {
			g.erase(1)
			g.writeString(", ")
			continue
		} else if syntax[i] == ' ' {
			continue
		} else {
			panic(fmt.Errorf("unknown token: %v", syntax[i]))
		}
		word := g.rand(source)
		g.write(word)
		g.writeString(" ")
	}
}

func (g *textPoolGen) generateSentence() {
	syntax := g.rand(g.dists.Grammars())
	for i := 0; i < len(syntax); i += 2 {
		if syntax[i] == 'V' {
			g.generateVerbPhrase()
		} else if syntax[i] == 'N' {
			g.generateNounPhrase()
		} else if syntax[i] == 'P' {
			preposition := g.rand(g.dists.Prepositions())
			g.write(preposition)
			g.writeString(" the ")
			g.generateNounPhrase()
		} else if syntax[i] == 'T' {
			g.erase(1)
			terminator := g.rand(g.dists.Terminators())
			g.write(terminator)
		} else {
			panic(fmt.Errorf("unknown token: %v", syntax[i]))
		}
		if g.last != ' ' {
			g.writeString(" ")
		}
	}
}

//

const (
	textPoolDefaultSize       = 300 * 1024 * 1024
	textPoolMaxSentenceLength = 256
)

type textPool struct {
	buf []byte
}

func newTextPool(
	size int,
	dists textDists,
) *textPool {
	return &textPool{
		buf: genTextPool(size, textPoolMaxSentenceLength, dists),
	}
}

func (p textPool) size() int {
	return len(p.buf)
}

func (p textPool) getText(begin, end int) string {
	return string(p.buf[begin:end])
}

var defaultTextPoolLazy = au.NewLazy(func() *textPool {
	return newTextPool(textPoolDefaultSize, getTextDists())
})

func getDefaultTextPool() *textPool {
	return defaultTextPoolLazy.Get()
}

//

type randomText struct {
	genRandom

	pool *textPool

	minLength int
	maxLength int
}

const (
	randomTextLowLengthMultiplier  = 0.4
	randomTextHighLengthMultiplier = 1.6
)

func newRandomText(
	seed int64,
	pool *textPool,
	averageTextLength float32,
	expectedRowCount int, // = 1
) *randomText {
	return &randomText{
		genRandom: genRandom{
			gen: newIntGen(seed, expectedRowCount*2),
		},

		pool: pool,

		minLength: int(averageTextLength * randomTextLowLengthMultiplier),
		maxLength: int(averageTextLength * randomTextHighLengthMultiplier),
	}
}

func (r *randomText) nextValue() string {
	offset := r.gen.rand(0, int64(r.pool.size()-r.maxLength))
	length := r.gen.rand(int64(r.minLength), int64(r.maxLength))
	return r.pool.getText(int(offset), int(offset+length))
}
