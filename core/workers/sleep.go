package workers

import (
	"errors"
	"time"
)

type SleepWorkerConfig struct {
	Interval time.Duration
	//Timeout  time.Duration
}

type SleepWorker struct {
	fn  func() bool
	cfg SleepWorkerConfig

	stop    chan struct{}
	stopped chan struct{}
}

func NewSleepWorker(fn func() bool, cfg SleepWorkerConfig) *SleepWorker {
	return &SleepWorker{
		fn:  fn,
		cfg: cfg,
	}
}

func (w *SleepWorker) Start() {
	if w.stop != nil {
		panic(errors.New("already started"))
	}
	w.stop = make(chan struct{})
	w.stopped = make(chan struct{})
	go w.run()
}

func (w *SleepWorker) run() {
	defer func() {
		close(w.stopped)
	}()
loop:
	for {
		if !w.fn() {
			break
		}
		select {
		case <-w.stop:
			break loop
		case <-time.After(w.cfg.Interval):
		}
	}
}

func (w *SleepWorker) Stop() {
	if w.stop == nil {
		panic(errors.New("not started"))
	}
	close(w.stop)
	select {
	case <-w.stopped:
	}
}
