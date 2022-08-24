package math

import (
	"math"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

func Quantile(sortedData []float64, q float64) float64 {
	check.Condition(0.0 <= q && q <= 1.0)
	data, n := sortedData, len(sortedData)
	idx := q / 1.0 * (float64(n) - 1)
	idxf, idxc := math.Floor(idx), math.Ceil(idx)
	if idxf == idxc {
		return data[int(idxf)]
	}
	return (data[int(idxf)] * (idxc - idx)) + (data[int(idxc)] * (idx - idxf))
}

//

type Stats struct {
	data []float64

	dfl float64
	_eq func(l, r float64) bool

	min  bt.Optional[float64]
	max  bt.Optional[float64]
	sum  bt.Optional[float64]
	mean bt.Optional[float64]

	sorted bt.Optional[[]float64]
}

func NewStats(data []float64) Stats {
	return Stats{data: data}
}

func (s Stats) Data() []float64  { return s.data }
func (s Stats) Default() float64 { return s.dfl }

func (s *Stats) SetDefault(dfl float64) *Stats {
	s.dfl = dfl
	return s
}

func (s *Stats) SetEq(eq func(l, r float64) bool) *Stats {
	s._eq = eq
	return s
}

func (s Stats) Eq(l, r float64) bool {
	if s._eq != nil {
		return s._eq(l, r)
	}
	return l == r
}

func (s Stats) Len() int { return len(s.data) }

func (s *Stats) Min() float64 {
	return bt.SetIfAbsent(&s.min, func() float64 {
		return bt.Min(s.data...)
	})
}

func (s *Stats) Max() float64 {
	return bt.SetIfAbsent(&s.min, func() float64 {
		return bt.Max(s.data...)
	})
}

func (s *Stats) Sum() float64 {
	return bt.SetIfAbsent(&s.sum, func() float64 {
		return bt.Sum(s.data...)
	})
}

func (s *Stats) Mean() float64 {
	return bt.SetIfAbsent(&s.mean, func() float64 {
		return s.Sum() / float64(len(s.data))
	})
}

func (s *Stats) Sorted() []float64 {
	return bt.SetIfAbsent(&s.sorted, func() []float64 {
		return slices.Sort(slices.Clone(s.data))
	})
}

func (s *Stats) Quantile(q float64) float64 {
	return Quantile(s.Sorted(), q)
}
