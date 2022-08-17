package pagelets

import "github.com/wrmsr/bane/pkg/util/check"

//

type Pagelet struct {
	name string
}

func NewPagelet(name string) *Pagelet {
	return &Pagelet{
		name: check.NotEqual(name, ""),
	}
}

//

type PageletRegistry struct {
	m map[string]*Pagelet
}

func NewPageletRegistry() *PageletRegistry {
	return &PageletRegistry{
		m: make(map[string]*Pagelet),
	}
}
