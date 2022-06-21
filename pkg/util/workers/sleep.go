package workers

import (
	"context"
	"time"
)

type SleepWorkerConfig struct {
	Interval time.Duration
}

type SleepWorker struct {
	cfg SleepWorkerConfig
}

func NewSleepWorker(cfg SleepWorkerConfig) *SleepWorker {
	return &SleepWorker{
		cfg: cfg,
	}
}

func (w *SleepWorker) Start() {
	panic("nyi")
}

func (w *SleepWorker) Stop(ctx context.Context) {
	panic("nyi")
}
