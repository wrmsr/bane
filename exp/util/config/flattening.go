package config

import (
	"fmt"
	"strings"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
)

type FlatteningConfig struct {
	Delimiter  string
	IndexOpen  string
	IndexClose string
}

func DefaultFlatteningConfig() FlatteningConfig {
	return FlatteningConfig{
		Delimiter:  ".",
		IndexOpen:  "(",
		IndexClose: ")",
	}
}

type Flattening struct {
	cfg FlatteningConfig
}

func NewFlattening(cfg FlatteningConfig) Flattening {
	return Flattening{cfg: cfg}
}

func (f Flattening) Flatten(unflattened map[string]any) map[string]any {
	ret := make(map[string]any)
	var rec func([]string, any)
	rec = func(prefix []string, value any) {
		switch value := value.(type) {
		case map[string]any:
			for k, v := range value {
				rec(slices.MakeAppend(prefix, k), v)
			}
		case []any:
			check.Condition(len(prefix) > 0)
			for i, v := range value {
				rec(slices.MakeAppend(prefix[:len(prefix)-1], fmt.Sprintf("%s%s%d%s", prefix[len(prefix)-1], f.cfg.IndexOpen, i, f.cfg.IndexClose)), v)
			}
		default:
			k := strings.Join(prefix, f.cfg.Delimiter)
			if _, ok := ret[k]; ok {
				panic(fmt.Errorf("duplicate key: %s", k))
			}
			ret[k] = value
		}
	}
	rec(nil, unflattened)
	return ret
}

//

type missing struct{}

type unflattenNode interface {
	get(key any) any
	put(key, value any)
	build() any
}

func maybeUnflattenBuild(v any) any {
	if v, ok := v.(unflattenNode); ok {
		return v.build()
	}
	return v
}

//def setdefault(self, key: K, supplier: ta.Callable[[], V]) -> V:
//	ret = self.get(key)
//	if ret is _MISSING:
//		ret = supplier()
//		self.put(key, ret)
//	return ret

//

type unflattenMap struct {
	m map[string]any
}

func newUnflattenMap() *unflattenMap {
	return &unflattenMap{m: make(map[string]any)}
}

var _ unflattenNode = &unflattenMap{}

func (n *unflattenMap) get(key any) any {
	if v, ok := n.m[key.(string)]; ok {
		return v
	}
	return missing{}
}

func (n *unflattenMap) put(key, value any) {
	ks := key.(string)
	if _, ok := n.m[ks]; ok {
		panic(fmt.Errorf("duplicate key: %s", ks))
	}
	n.m[ks] = value
}

func (n *unflattenMap) build() any {
	m := make(map[string]any)
	for k, v := range n.m {
		m[k] = maybeUnflattenBuild(v)
	}
	return m
}

//

type unflattenSlice struct {
	s []any
}

func newUnflattenSlice() *unflattenSlice {
	return &unflattenSlice{}
}

var _ unflattenNode = &unflattenSlice{}

func (n *unflattenSlice) get(key any) any {
	ki := key.(int)
	if ki < 0 {
		panic(fmt.Errorf("invalid slice index: %d", ki))
	}
	if ki < len(n.s) {
		return n.s[ki]
	}
	return missing{}
}

func (n *unflattenSlice) put(key, value any) {
	ki := key.(int)
	if ki < 0 {
		panic(fmt.Errorf("invalid slice index: %d", ki))
	}
	for ki >= len(n.s) {
		n.s = append(n.s, missing{})
	}
	if _, ok := n.s[ki].(missing); !ok {
		panic(fmt.Errorf("duplicate slice index: %d", ki))
	}
	n.s[ki] = value
}

func (n *unflattenSlice) build() any {
	s := make([]any, len(n.s))
	for i, e := range n.s {
		s[i] = maybeUnflattenBuild(e)
	}
	return s
}

/*
//

def unflatten(self, flattened: map[string]any) -> map[string]any:
	root = Flattening.UnflattenDict()

	def split_keys(fkey: str) -> ta.Iterable[ta.Union[str, int]]:
		for part in fkey.split(self._delimiter):
			if self._index_open in part:
				check.state(part.endswith(self._index_close))
				pos = part.index(self._index_open)
				yield part[:pos]
				for p in part[pos + len(self._index_open):-len(self._index_close)] \
						.split(self._index_close + self._index_open):
					yield int(p)
			else:
				check.state(')' not in part)
				yield part

	for fk, v in flattened.items():
		node = root
		fks = list(split_keys(fk))
		for key, nkey in zip(fks, fks[1:]):
			if isinstance(nkey, str):
				node = node.setdefault(key, Flattening.UnflattenDict)
			elif isinstance(nkey, int):
				node = node.setdefault(key, Flattening.UnflattenSlice)
			else:
				raise TypeError(key)
		node.put(fks[-1], v)

	return root.build()
*/
