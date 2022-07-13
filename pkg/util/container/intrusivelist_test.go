package container

import (
	"testing"

	"github.com/wrmsr/bane/pkg/util/slices"
)

type ListThing struct {
	i int

	al, bl IntrusiveListNode[ListThing]
}

var (
	listThingAOps = NewIntrusiveListOps(func(p *ListThing) *IntrusiveListNode[ListThing] { return &p.al })
	listThingBOps = NewIntrusiveListOps(func(p *ListThing) *IntrusiveListNode[ListThing] { return &p.bl })
)

func TestIntrusiveList(t *testing.T) {
	al := NewIntrusiveList[ListThing](listThingAOps)
	bl := NewIntrusiveList[ListThing](listThingAOps)

	tS := slices.Range[int](0, 10, 1)

}
