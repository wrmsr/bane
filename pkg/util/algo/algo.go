package algo

import (
	"fmt"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	"github.com/wrmsr/bane/pkg/util/maps"
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
		ret = append(ret, ctr.NewSetOf[T](ord.Slice()...))
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

	return ctr.NewListOf(ret...), nil
}

func Histogram[T comparable](it its.Iterable[T]) ctr.Map[T, int] {
	m := ctr.NewMutMap[T, int](nil)
	for it := it.Iterate(); it.HasNext(); {
		k := it.Next()
		m.Put(k, m.Get(k)+1)
	}
	return m
}

/*
func Unify(sets it.Iterable[ctr.Set[T]]) ctr.List[ctr.Set[T]] {
    rem: ta.List[ta.Set[T]] = [set(s) for s in sets]
    ret: ta.List[ta.Set[T]] = []

    for rem {
        cur = rem.pop()
        while True {
            moved = False
            i = len(rem) - 1
            while i >= 0 {
                if any(e in cur for e in rem[i]) {
                    cur |= rem[i]
                    del rem[i]
                    moved = True
				}
                i -= 1
			}
            if not moved {
                break
			}
		}
        ret.append(cur)
	}

    if ret {
        ret_set = {e for s in ret for e in s}
        ret_len = sum(map(len, ret))
        if ret_len != len(ret_set) {
            raise ValueError(ret)
		}
	}

    return ret
}
*/
