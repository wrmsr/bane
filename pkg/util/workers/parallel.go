package workers

import (
	"sync"
)

func DoParallel(nWorkers int, fns []func()) {
	if len(fns) < 1 {
		return
	}

	queue := func(work func(int)) {
		work(0)
	}

	if nWorkers > 1 {
		workerIDs := make(chan int, nWorkers)
		for i := 0; i < nWorkers; i++ {
			workerIDs <- i
		}

		queue = func(work func(int)) {
			go func() {
				worker := <-workerIDs
				work(worker)
				workerIDs <- worker
			}()
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(fns))
	for _, fn := range fns {
		fn := fn
		queue(func(worker int) {
			fn()
			wg.Done()
		})
	}

	wg.Wait()
}
