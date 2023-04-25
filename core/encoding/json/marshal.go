package json

import (
	msh "github.com/wrmsr/bane/core/marshal"
	rfl "github.com/wrmsr/bane/core/reflect"
	bt "github.com/wrmsr/bane/core/types"
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

func ToMarshal(n Node) (msh.Value, error) {
	switch n := n.(type) {

	case Object:
		kvs := make([]bt.Kv[msh.Value, msh.Value], len(n.Pairs))
		for i, p := range n.Pairs {
			k := msh.MakeString(p.K)
			v, err := ToMarshal(p.V)
			if err != nil {
				return nil, err
			}
			kvs[i] = bt.KvOf[msh.Value, msh.Value](k, v)
		}
		return msh.MakeObject(kvs...), nil

	case Array:
		mvs := make([]msh.Value, len(n.Values))
		for i, e := range n.Values {
			mv, err := ToMarshal(e)
			if err != nil {
				return nil, err
			}
			mvs[i] = mv
		}
		return msh.MakeArray(mvs...), nil

	case String:
		return msh.MakeString(n.S), nil

	case Number:
		return msh.JsonUnmarshal([]byte(n.S))

	case True:
		return msh.MakeBool(true), nil

	case False:
		return msh.MakeBool(false), nil

	case Null:
		return msh.MakeNull(), nil

	}
	panic(n)
}

func FromMarshal(mv msh.Value) (Node, error) {
	switch mv := mv.(type) {

	case msh.Null:
		return Null{}, nil

	case msh.Bool:
		if mv.V() {
			return True{}, nil
		}
		return False{}, nil

	case msh.Int:
		return Number{S: mv.String()}, nil

	case msh.Float:
		return Number{S: mv.String()}, nil

	case msh.Number:
		return Number{S: mv.String()}, nil

	case msh.String:
		return String{S: mv.V()}, nil

	case msh.Bytes:
		panic(mv)

	case msh.Array:
		mvs := mv.V()
		ns := make([]Node, len(mvs))
		for i, e := range mvs {
			n, err := FromMarshal(e)
			if err != nil {
				return nil, err
			}
			ns[i] = n
		}
		return Array{Values: ns}, nil

	case msh.Object:
		kvs := mv.V()
		ps := make([]Pair, len(kvs))
		for i, kv := range kvs {
			k := kv.K.(msh.String).V()
			v, err := FromMarshal(kv.V)
			if err != nil {
				return nil, err
			}
			ps[i] = bt.KvOf(k, v)
		}
		return Object{Pairs: ps}, nil

	case msh.Any:
		panic(mv)

	}
	panic(mv)
}
