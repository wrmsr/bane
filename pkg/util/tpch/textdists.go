package tpch

import (
	"embed"
	"strconv"
	"strings"

	au "github.com/wrmsr/bane/pkg/util/atomic"
	"github.com/wrmsr/bane/pkg/util/check"
	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	stru "github.com/wrmsr/bane/pkg/util/strings"
)

//

type textDists map[string]*textDist

func (td textDists) Grammars() *textDist         { return td["grammar"] }
func (td textDists) NounPhrase() *textDist       { return td["np"] }
func (td textDists) VerbPhrase() *textDist       { return td["vp"] }
func (td textDists) Prepositions() *textDist     { return td["prepositions"] }
func (td textDists) Nouns() *textDist            { return td["nouns"] }
func (td textDists) Verbs() *textDist            { return td["verbs"] }
func (td textDists) Articles() *textDist         { return td["articles"] }
func (td textDists) Adjectives() *textDist       { return td["adjectives"] }
func (td textDists) Adverbs() *textDist          { return td["adverbs"] }
func (td textDists) Auxiliaries() *textDist      { return td["auxillaries"] }
func (td textDists) Terminators() *textDist      { return td["terminators"] }
func (td textDists) OrderPriorities() *textDist  { return td["o_oprio"] }
func (td textDists) ShipInstructions() *textDist { return td["instruct"] }
func (td textDists) ShipModes() *textDist        { return td["smode"] }
func (td textDists) ReturnFlags() *textDist      { return td["rflag"] }
func (td textDists) PartContainers() *textDist   { return td["p_cntr"] }
func (td textDists) PartColors() *textDist       { return td["colors"] }
func (td textDists) PartTypes() *textDist        { return td["p_types"] }
func (td textDists) MarketSegments() *textDist   { return td["msegmnt"] }
func (td textDists) Nations() *textDist          { return td["nations"] }
func (td textDists) Regions() *textDist          { return td["regions"] }

//

type textDistsLoader struct{}

func (tdl textDistsLoader) LoadDists(buf string) textDists {
	lines := its.OfSlice(stru.TrimSpaceSplit(buf, "\n"))
	lines = its.Filter(lines, func(s string) bool { return !strings.HasPrefix(s, "#") })
	return tdl.loadDists(lines.Iterate())
}

func (tdl textDistsLoader) loadDists(lines its.Iterator[string]) textDists {
	dists := make(textDists)

	for lines.HasNext() {
		line := lines.Next()

		parts := stru.TrimSpaceSplit(line, " ")
		if len(parts) != 2 {
			continue
		}

		if strings.ToUpper(parts[0]) == "BEGIN" {
			name := parts[1]
			dist := tdl.loadDist(lines, name)
			dists[strings.ToLower(name)] = dist
		}
	}

	return dists
}

func (tdl textDistsLoader) loadDist(lines its.Iterator[string], name string) *textDist {
	count := -1
	members := ctr.NewMutSliceMap[string, int](nil)
	for lines.HasNext() {
		line := lines.Next()

		if tdl.isEnd(name, line) {
			check.Condition(count == members.Len())
			return newTextDist(name, members)
		}

		parts := stru.TrimSpaceSplit(line, "|")
		check.Condition(len(parts) == 2)

		value := parts[0]
		weight := check.Must1(strconv.Atoi(parts[1]))

		if strings.ToLower(value) == "count" {
			count = weight
		} else {
			members.Put(value, weight)
		}
	}

	panic(name)
}

func (tdl textDistsLoader) isEnd(name, line string) bool {
	parts := stru.TrimSpaceSplit(line, " ")
	if strings.ToUpper(parts[0]) == "END" {
		check.Condition(len(parts) == 2 && strings.ToLower(parts[1]) == strings.ToLower(name))
		return true
	}
	return false
}

//

//go:embed textdists.dss
var textDistsEmbed embed.FS

var textDistsLazy = au.NewLazy(func() textDists {
	buf := string(check.Must1(textDistsEmbed.ReadFile("textdists.dss")))
	return textDistsLoader{}.LoadDists(buf)
})

func getTextDists() textDists {
	return textDistsLazy.Get()
}
