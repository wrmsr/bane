package tpch

import (
	"embed"
	"strconv"
	"strings"

	au "github.com/wrmsr/bane/pkg/util/atomic"
	"github.com/wrmsr/bane/pkg/util/check"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	stru "github.com/wrmsr/bane/pkg/util/strings"
)

//

type textDists map[string]*textDist

func (td textDists) grammars() *textDist          { return td["grammar"] }
func (td textDists) noun_phrase() *textDist       { return td["np"] }
func (td textDists) verb_phrase() *textDist       { return td["vp"] }
func (td textDists) prepositions() *textDist      { return td["prepositions"] }
func (td textDists) nouns() *textDist             { return td["nouns"] }
func (td textDists) verbs() *textDist             { return td["verbs"] }
func (td textDists) articles() *textDist          { return td["articles"] }
func (td textDists) adjectives() *textDist        { return td["adjectives"] }
func (td textDists) adverbs() *textDist           { return td["adverbs"] }
func (td textDists) auxiliaries() *textDist       { return td["auxillaries"] }
func (td textDists) terminators() *textDist       { return td["terminators"] }
func (td textDists) order_priorities() *textDist  { return td["o_oprio"] }
func (td textDists) ship_instructions() *textDist { return td["instruct"] }
func (td textDists) ship_modes() *textDist        { return td["smode"] }
func (td textDists) return_flags() *textDist      { return td["rflag"] }
func (td textDists) part_containers() *textDist   { return td["p_cntr"] }
func (td textDists) part_colors() *textDist       { return td["colors"] }
func (td textDists) part_types() *textDist        { return td["p_types"] }
func (td textDists) market_segments() *textDist   { return td["msegmnt"] }
func (td textDists) nations() *textDist           { return td["nations"] }
func (td textDists) regions() *textDist           { return td["regions"] }

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
	members := make(map[string]int)
	for lines.HasNext() {
		line := lines.Next()

		if tdl.isEnd(name, line) {
			check.Condition(count == len(members))
			return newTextDist(name, members)
		}

		parts := stru.TrimSpaceSplit(line, "|")
		check.Condition(len(parts) == 2)

		value := parts[0]
		weight := check.Must(strconv.Atoi(parts[1]))

		if strings.ToLower(value) == "count" {
			count = weight
		} else {
			members[value] = weight
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
	buf := string(check.Must(textDistsEmbed.ReadFile("textdists.dss")))
	return textDistsLoader{}.LoadDists(buf)
})

func getTextDists() textDists {
	return textDistsLazy.Get()
}
