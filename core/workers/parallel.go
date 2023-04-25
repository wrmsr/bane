package workers

import (
	"context"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/wrmsr/bane/core/slices"
)

type WorkerId int

func DoParallel[T any](n int, fns []func() T) []T {
	if len(fns) < 1 {
		return nil
	}

	if n < 2 {
		return slices.Map(func(fn func() T) T { return fn() }, fns)
	}

	ret := make([]T, len(fns))

	ids := make(chan WorkerId, n)
	for i := 0; i < n; i++ {
		ids <- WorkerId(i)
	}

	var wg sync.WaitGroup
	wg.Add(len(fns))
	for i, fn := range fns {
		i, fn := i, fn
		go func() {
			worker := <-ids
			defer func() { ids <- worker }()

			ret[i] = fn()
			wg.Done()
		}()
	}

	wg.Wait()
	return ret
}

func DoParallelErrGroup[T any](ctx context.Context, n int, fns []func(ctx context.Context) (T, error)) ([]T, error) {
	if len(fns) < 1 {
		return nil, nil
	}

	ret := make([]T, len(fns))

	if n < 2 {
		for i, fn := range fns {
			r, err := fn(ctx)
			if err != nil {
				return nil, err
			}
			ret[i] = r
		}
		return ret, nil
	}

	ids := make(chan WorkerId, n)
	for i := 0; i < n; i++ {
		ids <- WorkerId(i)
	}

	eg, ctx := errgroup.WithContext(ctx)
	for i, fn := range fns {
		i, fn := i, fn
		eg.Go(func() error {
			// FIXME: select
			worker := <-ids
			defer func() { ids <- worker }()

			r, err := fn(ctx)
			if err != nil {
				return err
			}
			ret[i] = r
			return nil
		})
	}

	return ret, eg.Wait()
}
