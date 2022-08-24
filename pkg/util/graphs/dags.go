package graphs

import (
	"github.com/wrmsr/bane/pkg/util/check"
	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	"github.com/wrmsr/bane/pkg/util/maps"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Dag[T comparable] struct {
	inputSetsByOutput ctr.Map[T, ctr.Set[T]]

	all bt.Optional[ctr.Set[T]]

	outputSetsByOutput bt.Optional[ctr.Map[T, ctr.Set[T]]]
}

func SliceEdges[T comparable](m map[T][]T) ctr.Map[T, ctr.Set[T]] {
	return ctr.NewStdMap(its.OfMap(maps.MapValues(func(s []T) ctr.Set[T] { return ctr.NewStdSet(its.OfSlice(s)) }, m)))
}

func ItEdges[T comparable](it bt.Iterable[bt.Kv[T, bt.Iterable[T]]]) ctr.Map[T, ctr.Set[T]] {
	return ctr.NewStdMap[T, ctr.Set[T]](its.MapValues(it, func(i bt.Iterable[T]) ctr.Set[T] { return ctr.NewStdSet(i) }))
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
	return bt.SetIfAbsent(&d.all, func() ctr.Set[T] {
		return ctr.NewStdSet(ctr.Keys(d.inputSetsByOutput))
	})
}

func (d *Dag[T]) OutputSetsByOutput() ctr.Map[T, ctr.Set[T]] {
	return bt.SetIfAbsent(&d.outputSetsByOutput, func() ctr.Map[T, ctr.Set[T]] {
		m := ctr.NewMutStdMap[T, ctr.MutSet[T]](nil)
		d.All().ForEach(func(t T) bool {
			m.Put(t, ctr.NewMutStdSet[T](nil))
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

func (d *Dag[T]) Subdag(targets bt.Iterable[T], ignored bt.Iterable[T]) *Subdag[T] {
	return &Subdag[T]{
		dag: d,

		targets: ctr.NewStdSet[T](targets),
		ignored: ctr.NewStdSet[T](ignored),
	}
}

func TraverseLinks[T comparable](data ctr.Map[T, ctr.Set[T]], keys bt.Iterable[T]) ctr.Set[T] {
	keySet := ctr.NewStdSet(keys)
	todoSet := ctr.NewMutStdSet[T](keySet)
	seenSet := ctr.NewMutStdSet[T](nil)
	for todoSet.Len() > 0 {
		key := check.Ok1(ctr.SetPop[T](todoSet))
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

	inputs  bt.Optional[ctr.Set[T]]
	outputs bt.Optional[ctr.Set[T]]

	outputInputs bt.Optional[ctr.Set[T]]

	all bt.Optional[ctr.Set[T]]
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
	return bt.SetIfAbsent(&sd.inputs, func() ctr.Set[T] {
		return ctr.Difference(TraverseLinks[T](sd.dag.InputSetsByOutput(), sd.targets), sd.ignored)
	})
}

func (sd *Subdag[T]) Outputs() ctr.Set[T] {
	return bt.SetIfAbsent(&sd.outputs, func() ctr.Set[T] {
		return ctr.Difference(TraverseLinks[T](sd.dag.OutputSetsByOutput(), sd.targets), sd.ignored)

	})
}

func (sd *Subdag[T]) OutputInputs() ctr.Set[T] {
	return bt.SetIfAbsent(&sd.outputInputs, func() ctr.Set[T] {
		return ctr.Difference(TraverseLinks[T](sd.dag.InputSetsByOutput(), sd.Outputs()), sd.ignored)
	})
}

func (sd *Subdag[T]) All() ctr.Set[T] {
	return bt.SetIfAbsent(&sd.all, func() ctr.Set[T] {
		return ctr.Union(sd.targets, sd.Inputs(), sd.Outputs(), sd.OutputInputs())
	})
}
