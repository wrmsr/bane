package container

import (
	"fmt"
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
	_ = bl

	ts := slices.Map(func(i int) *ListThing { return &ListThing{i: i} }, slices.Range1(0, 10))
	fmt.Println(ts)

	al.PushFront(ts[1])
	//bl.PushFront(ts[2])

	al.PushFront(ts[2])
	//bl.PushFront(ts[3])

	//ll := list.New()
	//ll.PushFront(ts[1])
	//ll.PushFront(ts[2])
	//ll.PushFront(ts[4])

	//al.PushFront(ts[4])
	//bl.PushFront(ts[5])

	al.verify()
}
