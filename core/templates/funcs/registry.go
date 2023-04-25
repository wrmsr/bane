package funcs

import (
	"fmt"
	"text/template"
)

type FuncRegistry template.FuncMap

func (fr FuncRegistry) Add(name string, fn any) FuncRegistry {
	if _, ok := fr[name]; ok {
		panic(fmt.Errorf("function already registered: %s", name))
	}
	fr[name] = fn
	return fr
}

func (fr FuncRegistry) AddAll(m template.FuncMap) FuncRegistry {
	for name, fn := range m {
		fr.Add(name, fn)
	}
	return fr
}

func (fr FuncRegistry) Update(m template.FuncMap) FuncRegistry {
	for name, fn := range m {
		fr[name] = fn
	}
	return fr
}
