package tpch

import (
	"strconv"
	"strings"

	"github.com/wrmsr/bane/pkg/util/check"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	stru "github.com/wrmsr/bane/pkg/util/strings"
)

type textDistLoading struct{}

func (tdl textDistLoading) loadDist(lines its.Iterator[string], name string) *textDist {
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
func (tdl textDistLoading) isEnd(name, line string) bool {
	parts := stru.TrimSpaceSplit(line, " ")
	if strings.ToUpper(parts[0]) == "END" {
		check.Condition(len(parts) == 2 && strings.ToLower(parts[1]) == strings.ToLower(name))
		return true
	}
	return false
}

/*
func (tdl textDistLoading) LoadDist(buf string) map[str]*TextDist {
	return tdl._loadDists((
		l
		for l in buf.splitlines()
		for l in [l.strip()]
		if not l.startswith('#')
	))
}

@classmethod
def loadDists(cls, lines: ta.Iterable[str]) -> ta.Mapping[str, TextDist]:
	dists: ta.Dict[str, TextDist] = {}

	lines = iter(lines)
	while True:
		line = next(lines, None)
		if line is None:
			break

		parts = line.strip().split()
		if len(parts) != 2:
			continue

		if parts[0].upper() == 'BEGIN':
			name = parts[1]
			dist = cls._load_dist(lines, name)
			dists[name.lower()] = dist

	return dists



class TextDists:

    def __init__(self, dists_by_name: ta.Mapping[str, TextDist]) -> None:
        super().__init__()

        self._dists_by_name = dists_by_name

    @classmethod
    def load_defaults(cls) -> ta.Mapping[str, TextDist]:
        return TextDistLoading.loadDist(
            pkg_resources.resource_string(__package__, 'dists.dss').decode('utf-8'))

    DEFAULT: 'TextDists' = properties.cached_class(lambda cls: cls(cls.load_defaults()))

    def __getitem__(self, name: str) -> TextDist:
        return self._dists_by_name[name]

    grammars: TextDist = property(lambda self: self['grammar'])
    noun_phrase: TextDist = property(lambda self: self['np'])
    verb_phrase: TextDist = property(lambda self: self['vp'])
    prepositions: TextDist = property(lambda self: self['prepositions'])
    nouns: TextDist = property(lambda self: self['nouns'])
    verbs: TextDist = property(lambda self: self['verbs'])
    articles: TextDist = property(lambda self: self['articles'])
    adjectives: TextDist = property(lambda self: self['adjectives'])
    adverbs: TextDist = property(lambda self: self['adverbs'])
    auxiliaries: TextDist = property(lambda self: self['auxillaries'])
    terminators: TextDist = property(lambda self: self['terminators'])
    order_priorities: TextDist = property(lambda self: self['o_oprio'])
    ship_instructions: TextDist = property(lambda self: self['instruct'])
    ship_modes: TextDist = property(lambda self: self['smode'])
    return_flags: TextDist = property(lambda self: self['rflag'])
    part_containers: TextDist = property(lambda self: self['p_cntr'])
    part_colors: TextDist = property(lambda self: self['colors'])
    part_types: TextDist = property(lambda self: self['p_types'])
    market_segments: TextDist = property(lambda self: self['msegmnt'])
    nations: TextDist = property(lambda self: self['nations'])
    regions: TextDist = property(lambda self: self['regions'])

*/
