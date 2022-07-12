package config

import (
	"fmt"
	"strconv"
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

func (f Flattening) Flatten(unflattened map[string]any) (map[string]any, error) {
	ret := make(map[string]any)

	var rec func([]string, any) error
	rec = func(prefix []string, value any) error {
		switch value := value.(type) {

		case map[string]any:
			for k, v := range value {
				if err := rec(slices.MakeAppend(prefix, k), v); err != nil {
					return err
				}
			}

		case []any:
			check.Condition(len(prefix) > 0)
			for i, v := range value {
				nk := fmt.Sprintf("%s%s%d%s", prefix[len(prefix)-1], f.cfg.IndexOpen, i, f.cfg.IndexClose)
				if err := rec(slices.MakeAppend(prefix[:len(prefix)-1], nk), v); err != nil {
					return err
				}
			}

		default:
			k := strings.Join(prefix, f.cfg.Delimiter)
			if _, ok := ret[k]; ok {
				return fmt.Errorf("duplicate key: %s", k)
			}
			ret[k] = value

		}
		return nil
	}
	if err := rec(nil, unflattened); err != nil {
		return nil, err
	}
	return ret, nil
}

//

type missing struct{}

type unflattenNode interface {
	get(key any) (any, error)
	put(key, value any) error
	build() any
}

func maybeUnflattenBuild(v any) any {
	if v, ok := v.(unflattenNode); ok {
		return v.build()
	}
	return v
}

//

type unflattenMap struct {
	m map[string]any
}

func newUnflattenMap() *unflattenMap {
	return &unflattenMap{m: make(map[string]any)}
}

var _ unflattenNode = &unflattenMap{}

func (n *unflattenMap) get(key any) (any, error) {
	if v, ok := n.m[key.(string)]; ok {
		return v, nil
	}
	return missing{}, nil
}

func (n *unflattenMap) put(key, value any) error {
	ks, ok := key.(string)
	if !ok {
		return fmt.Errorf("invalid key: %v", key)
	}

	if _, ok := n.m[ks]; ok {
		return fmt.Errorf("duplicate key: %s", ks)
	}
	n.m[ks] = value
	return nil
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

func (n *unflattenSlice) get(key any) (any, error) {
	ki, ok := key.(int)
	if !ok {
		return nil, fmt.Errorf("invalid key: %v", key)
	}

	if ki < 0 {
		return nil, fmt.Errorf("invalid slice index: %d", ki)
	}

	if ki < len(n.s) {
		return n.s[ki], nil
	}

	return missing{}, nil
}

func (n *unflattenSlice) put(key, value any) error {
	ki, ok := key.(int)
	if !ok {
		return fmt.Errorf("invalid key: %v", key)
	}

	if ki < 0 {
		return fmt.Errorf("invalid slice index: %d", ki)
	}

	for ki >= len(n.s) {
		n.s = append(n.s, missing{})
	}

	if _, ok := n.s[ki].(missing); !ok {
		return fmt.Errorf("duplicate slice index: %d", ki)
	}

	n.s[ki] = value
	return nil
}

func (n *unflattenSlice) build() any {
	s := make([]any, len(n.s))
	for i, e := range n.s {
		s[i] = maybeUnflattenBuild(e)
	}
	return s
}

//

func (f Flattening) Unflatten(flattened map[string]any) (map[string]any, error) {
	var root unflattenNode = newUnflattenMap()

	splitKeys := func(fkey string) ([]any, error) {
		var ret []any
		for _, part := range strings.Split(fkey, f.cfg.Delimiter) {
			if strings.Contains(part, f.cfg.IndexOpen) {
				if !strings.HasSuffix(part, f.cfg.IndexClose) {
					return nil, fmt.Errorf("invalid key: %s", fkey)
				}

				pos := strings.Index(part, f.cfg.IndexOpen)
				ret = append(ret, part[:pos])
				ps1 := part[pos+len(f.cfg.IndexOpen) : len(part)-len(f.cfg.IndexClose)]
				for _, p := range strings.Split(ps1, f.cfg.IndexClose+f.cfg.IndexOpen) {
					ret = append(ret, check.Must1(strconv.Atoi(p)))
				}

			} else {
				if strings.Contains(part, f.cfg.IndexClose) {
					return nil, fmt.Errorf("invalid key: %s", fkey)
				}
				ret = append(ret, part)

			}
		}
		return ret, nil
	}

	setDefault := func(n unflattenNode, key any, fn func() unflattenNode) (unflattenNode, error) {
		ret, err := n.get(key)
		if err != nil {
			return nil, err
		}

		if _, ok := ret.(missing); ok {
			ret = fn()
			if err := n.put(key, ret); err != nil {
				return nil, err
			}
		}

		retn, ok := ret.(unflattenNode)
		if !ok {
			return nil, fmt.Errorf("invalid node: %v", ret)
		}

		return retn, nil
	}

	for fk, v := range flattened {
		node := root
		fks, err := splitKeys(fk)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(fks)-1; i++ {
			key := fks[i]
			nkey := fks[i+1]
			switch nkey.(type) {

			case string:
				node, err = setDefault(node, key, func() unflattenNode { return newUnflattenMap() })

			case int:
				node, err = setDefault(node, key, func() unflattenNode { return newUnflattenSlice() })

			default:
				return nil, fmt.Errorf("invalid key part: %v", nkey)

			}
			if err != nil {
				return nil, err
			}
		}
		if err := node.put(fks[len(fks)-1], v); err != nil {
			return nil, err
		}
	}

	return root.build().(map[string]any), nil
}
