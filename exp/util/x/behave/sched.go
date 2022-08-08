package behave

import (
	opt "github.com/wrmsr/bane/pkg/util/optional"
	"github.com/wrmsr/bane/pkg/util/slices"
)

//

type Schedulable interface {
	Run(ttl float32)
}

type Scheduler interface {
	Schedulable
	Add(schedulable Schedulable, frequency int, phase int)
}

//

type SchedulerEntry interface {
	Schedulable() Schedulable
	Frequency() int
	Phase() int
}

type BaseSchedulerEntry struct {
	schedulable Schedulable
	frequency   int
	phase       int
}

var _ SchedulerEntry = BaseSchedulerEntry{}

func (b BaseSchedulerEntry) Schedulable() Schedulable { return b.schedulable }
func (b BaseSchedulerEntry) Frequency() int           { return b.frequency }
func (b BaseSchedulerEntry) Phase() int               { return b.phase }

//

type BaseScheduler[T SchedulerEntry] struct {
	schedulable   []T
	phaseCounters []int
	dryRunFrames  int
	frame         int
}

func (bs *BaseScheduler[T]) calculatePhase(frequency int) int {
	bs.phaseCounters = slices.Repeat([]int{0}, frequency)

	for frame := 0; frame < bs.dryRunFrames; frame++ {
		slot := frame % frequency
		for _, entry := range bs.schedulable {
			if (frame-entry.Phase())%entry.Frequency() == 0 {
				bs.phaseCounters[slot] += 1
			}
		}
	}

	var minValue opt.Optional[int]
	minValueAt := -1
	for i := 0; i < frequency; i++ {
		if !minValue.Present() || bs.phaseCounters[i] < minValue.Value() {
			minValue = opt.Just(bs.phaseCounters[i])
			minValueAt = i
		}
	}
	return minValueAt
}

/*
class PriorityScheduler(BaseScheduler['PriorityScheduler.Entry']):

    class Entry(BaseScheduler.Entry, frozen=True):
        priority: float

    def __init__(self, dry_run_frames: int = 0) -> None:
        super().__init__(dry_run_frames)

    def run(self, ttl: float) -> None:
        ttl = float(ttl)
        self._frame += 1

        runnable = []
        total_priority = 0.
        for entry in self._schedulable:
            if (self._frame + entry.phase) % entry.frequency == 0:
                runnable.append(entry)
                total_priority += entry.priority

        last_time = time.time()
        for entry in runnable:
            now = time.time()
            ttl -= now - last_time
            available_time = (ttl * entry.priority / total_priority)

            entry.schedulable.run(available_time)
            last_time = now

    def add(
            self,
            schedulable: Schedulable,
            frequency: int,
            *,
            phase: int = None,
            priority: float = 1.,
    ) -> None:
        self._schedulable.append(
            self.Entry(
                schedulable,
                frequency,
                phase if phase is not None else self._calculate_phase(frequency),
                priority,
            ))


class LoadBalancingScheduler(BaseScheduler[BaseScheduler.Entry]):

    def run(self, ttl: float) -> None:
        ttl = float(ttl)
        self._frame += 1

        runnable = []
        for entry in self._schedulable:
            if (self._frame + entry.phase) % entry.frequency == 0:
                runnable.append(entry)

        last_time = time.time()
        for entry in runnable:
            now = time.time()
            ttl -= now - last_time
            available_time = ttl / (len(runnable) - 1)

            entry.schedulable.run(available_time)
            last_time = now

    def add(
            self,
            schedulable: Schedulable,
            frequency: int,
            *,
            phase: int = None,
    ) -> None:
        self._schedulable.append(
            self.Entry(
                schedulable,
                frequency,
                phase if phase is not None else self._calculate_phase(frequency),
            ))
*/
