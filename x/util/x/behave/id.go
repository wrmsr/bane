package behave

import (
	"sync/atomic"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type idGen struct {
	next int64
}

func (g *idGen) Next() uintptr {
	r := atomic.AddInt64(&g.next, 1)
	if r < 0 {
		panic("id overflow")
	}
	return uintptr(r)
}

var ids idGen

func NextId() uintptr {
	return ids.Next()
}

//

type generatedIdentity struct {
	_id bt.Optional[uintptr]
}

func (hi *generatedIdentity) Identity() uintptr {
	return bt.SetIfAbsent(&hi._id, NextId)
}
