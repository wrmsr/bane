// hi
/*
hi there
*/
package algo

import (
	"fmt"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	"github.com/wrmsr/bane/pkg/util/maps"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TopoSort[T comparable](data ctr.Map[T, ctr.Set[T]]) (ctr.List[ctr.Set[T]], error) {
	m := make(map[T]maps.Set[T])
	a := maps.MakeSet[T]()
	data.ForEach(func(kv bt.Kv[T, ctr.Set[T]]) bool {
		m[kv.K] = maps.MakeSetOf(its.Seq[T](kv.V))
		m[kv.K].Remove(kv.K)
		a.Add(kv.K)
		a.AddAll(its.Seq[T](kv.V))
		return true
	})

	a.RemoveAll(maps.Keys(m))
	for k := range a {
		m[k] = maps.MakeSet[T]()
	}

	var ret []ctr.Set[T]
	for {
		ord := maps.MakeSet[T]()
		for k, v := range m {
			if len(v) < 1 {
				ord.Add(k)
			}
		}

		if len(ord) < 1 {
			break
		}

		rem := ord.Slice()
		ret = append(ret, ctr.NewStdSetOf[T](ord.Slice()...))
		for k, v := range m {
			if !ord.Contains(k) {
				v.RemoveAll(rem)
			} else {
				delete(m, k)
			}
		}
	}

	if len(m) > 0 {
		return nil, fmt.Errorf("cyclic dependencies exist among these items: %s", data)
	}

	return ctr.NewSliceListOf(ret...), nil
}

func Histogram[T comparable](it bt.Iterable[T]) ctr.Map[T, int] {
	m := ctr.NewMutStdMap[T, int](nil)
	for it := it.Iterate(); it.HasNext(); {
		k := it.Next()
		m.Put(k, m.Get(k)+1)
	}
	return m
}

func Unify[T comparable](sets []maps.Set[T]) []maps.Set[T] {
	rem := make([]maps.Set[T], len(sets))
	for i, s := range sets {
		rem[i] = maps.Clone(s)
	}
	var ret []maps.Set[T]

	for len(rem) > 0 {
		cur := slices.Pop(&rem)
		for {
			moved := false
			i := len(rem) - 1
			for !moved && i >= 0 {
				for e := range rem[i] {
					if cur.Contains(e) {
						cur.Update(rem[i])
						rem = slices.DeleteAt(rem, i)
						moved = true
						break
					}
				}
				i--
			}
			if !moved {
				break
			}
		}
		ret = append(ret, cur)
	}

	if len(ret) > 0 {
		retSet := maps.MakeSet[T]()
		retLen := 0
		for _, s := range ret {
			for e := range s {
				retSet.Add(e)
			}
			retLen += len(s)
		}
		if retLen != len(retSet) {
			panic(ret)
		}
	}

	return ret
}
