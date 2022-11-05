package json

import (
	msh "github.com/wrmsr/bane/pkg/util/marshal"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

var _ = msh.Register(rfl.TypeOf[Node](),
	msh.SetImplOf[Object](),
	msh.SetImplOf[Array](),
	msh.SetImplOf[String](),
	msh.SetImplOf[Number](),
	msh.SetImplOf[True](),
	msh.SetImplOf[False](),
	msh.SetImplOf[Null](),
)

func ToMarshal(n Node) msh.Value {
	switch n := n.(type) {
	case Object:
		kvs := make([]bt.Kv[msh.Value, msh.Value], len(n.Pairs))
		for i, p := range n.Pairs {
			kvs[i] = bt.KvOf[msh.Value, msh.Value](msh.MakeString(p.K), ToMarshal(p.V))
		}
		return msh.MakeObject(kvs...)
	case Array:
		mvs := make([]msh.Value, len(n.Values))
		for i, e := range n.Values {
			mvs[i] = ToMarshal(e)
		}
		return msh.MakeArray(mvs...)
	case String:
		return msh.MakeString(n.S)
	case Number:
		panic(n)
	case True:
		return msh.MakeBool(true)
	case False:
		return msh.MakeBool(false)
	case Null:
		return msh.MakeNull()
	}
	panic(n)
}
