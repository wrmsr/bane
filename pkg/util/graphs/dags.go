package graphs

import (
	"github.com/wrmsr/bane/pkg/util/check"
	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	"github.com/wrmsr/bane/pkg/util/maps"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Dag[T comparable] struct {
	inputSetsByOutput ctr.Map[T, ctr.Set[T]]

	all opt.Optional[ctr.Set[T]]

	outputSetsByOutput opt.Optional[ctr.Map[T, ctr.Set[T]]]
}

func SliceEdges[T comparable](m map[T][]T) ctr.Map[T, ctr.Set[T]] {
	return ctr.NewMap(its.OfMap(maps.MapValues(func(s []T) ctr.Set[T] { return ctr.NewSet(its.OfSlice(s)) }, m)))
}

func ItEdges[T comparable](it its.Iterable[bt.Kv[T, its.Iterable[T]]]) ctr.Map[T, ctr.Set[T]] {
	return ctr.NewMap[T, ctr.Set[T]](its.MapValues(it, func(i its.Iterable[T]) ctr.Set[T] { return ctr.NewSet(i) }))
}

func NewDag[T comparable](inputSetsByOutput ctr.Map[T, ctr.Set[T]]) *Dag[T] {
	return &Dag[T]{
		inputSetsByOutput: inputSetsByOutput,
	}
}

func (d *Dag[T]) InputSetsByOutput() ctr.Map[T, ctr.Set[T]] {
	return d.inputSetsByOutput
}

func (d *Dag[T]) All() ctr.Set[T] {
	return opt.SetIfAbsent(&d.all, func() ctr.Set[T] {
		return ctr.NewSet(ctr.Keys(d.inputSetsByOutput))
	})
}

func (d *Dag[T]) OutputSetsByOutput() ctr.Map[T, ctr.Set[T]] {
	return opt.SetIfAbsent(&d.outputSetsByOutput, func() ctr.Map[T, ctr.Set[T]] {
		m := ctr.NewMutMap[T, ctr.MutSet[T]](nil)
		d.All().ForEach(func(t T) bool {
			m.Put(t, ctr.NewMutSet[T](nil))
			return true
		})
		d.inputSetsByOutput.ForEach(func(kv bt.Kv[T, ctr.Set[T]]) bool {
			kv.V.ForEach(func(v T) bool {
				m.Get(v).Add(kv.K)
				return true
			})
			return true
		})
		return ctr.MapValues[T, ctr.MutSet[T], ctr.Set[T]](func(s ctr.MutSet[T]) ctr.Set[T] { return s }, m)
	})
}

func (d *Dag[T]) Subdag(targets its.Iterable[T], ignored its.Iterable[T]) *Subdag[T] {
	return &Subdag[T]{
		dag: d,

		targets: ctr.NewSet[T](targets),
		ignored: ctr.NewSet[T](ignored),
	}
}

func TraverseLinks[T comparable](data ctr.Map[T, ctr.Set[T]], keys its.Iterable[T]) ctr.Set[T] {
	keySet := ctr.NewSet(keys)
	todoSet := ctr.NewMutSet[T](keySet)
	seenSet := ctr.NewMutSet[T](nil)
	for todoSet.Len() > 0 {
		key := check.Ok1(ctr.SetPop(todoSet))
		seenSet.Add(key)
		if cur, ok := data.TryGet(key); ok {
			cur.ForEach(func(v T) bool {
				if !seenSet.Contains(v) {
					todoSet.Add(v)
				}
				return true
			})
		}
	}
	return ctr.Difference[T](seenSet, keySet)
}

//

type Subdag[T comparable] struct {
	dag *Dag[T]

	targets ctr.Set[T]
	ignored ctr.Set[T]

	inputs  opt.Optional[ctr.Set[T]]
	outputs opt.Optional[ctr.Set[T]]

	outputInputs opt.Optional[ctr.Set[T]]

	all opt.Optional[ctr.Set[T]]
}

func (sd *Subdag[T]) Dag() *Dag[T] {
	return sd.dag
}

func (sd *Subdag[T]) Targets() ctr.Set[T] {
	return sd.targets
}

func (sd *Subdag[T]) Ignored() ctr.Set[T] {
	return sd.ignored
}

func (sd *Subdag[T]) Inputs() ctr.Set[T] {
	return opt.SetIfAbsent(&sd.inputs, func() ctr.Set[T] {
		return ctr.Difference(TraverseLinks[T](sd.dag.InputSetsByOutput(), sd.targets), sd.ignored)
	})
}

func (sd *Subdag[T]) Outputs() ctr.Set[T] {
	return opt.SetIfAbsent(&sd.outputs, func() ctr.Set[T] {
		return ctr.Difference(TraverseLinks[T](sd.dag.OutputSetsByOutput(), sd.targets), sd.ignored)

	})
}

func (sd *Subdag[T]) OutputInputs() ctr.Set[T] {
	return opt.SetIfAbsent(&sd.outputInputs, func() ctr.Set[T] {
		return ctr.Difference(TraverseLinks[T](sd.dag.InputSetsByOutput(), sd.Outputs()), sd.ignored)
	})
}

func (sd *Subdag[T]) All() ctr.Set[T] {
	return opt.SetIfAbsent(&sd.all, func() ctr.Set[T] {
		return ctr.Union(sd.targets, sd.Inputs(), sd.Outputs(), sd.OutputInputs())
	})
}
