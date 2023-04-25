package behave

import (
	"time"

	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Schedulable interface {
	Run(ttl float32)
}

type Scheduler interface {
	Schedulable
	Add(schedulable Schedulable, frequency int, phase bt.Optional[int])
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

func (bsch *BaseScheduler[T]) calculatePhase(frequency int) int {
	bsch.phaseCounters = slices.Repeat([]int{0}, frequency)

	for frame := 0; frame < bsch.dryRunFrames; frame++ {
		slot := frame % frequency
		for _, entry := range bsch.schedulable {
			if (frame-entry.Phase())%entry.Frequency() == 0 {
				bsch.phaseCounters[slot]++
			}
		}
	}

	var minValue bt.Optional[int]
	minValueAt := -1
	for i := 0; i < frequency; i++ {
		if !minValue.Present() || bsch.phaseCounters[i] < minValue.Value() {
			minValue = bt.Just(bsch.phaseCounters[i])
			minValueAt = i
		}
	}
	return minValueAt
}

//

type PrioritySchedulerEntry struct {
	BaseSchedulerEntry
	priority float32
}

type PriorityScheduler struct {
	BaseScheduler[PrioritySchedulerEntry]
}

var _ Scheduler = &PriorityScheduler{}

func (sch *PriorityScheduler) Run(ttl float32) {
	sch.frame++

	var runnable []PrioritySchedulerEntry
	var totalPriority float32
	for _, entry := range sch.schedulable {
		if (sch.frame+entry.Phase())%entry.Frequency() == 0 {
			runnable = append(runnable, entry)
			totalPriority += entry.priority
		}
	}

	lastTime := time.Now()
	for _, entry := range runnable {
		now := time.Now()
		ttl -= float32(now.Sub(lastTime).Seconds())
		availableTime := (ttl * entry.priority) / totalPriority

		entry.schedulable.Run(availableTime)
		lastTime = now
	}
}

func (sch *PriorityScheduler) Add(schedulable Schedulable, frequency int, phase bt.Optional[int]) {
	sch.AddPriority(schedulable, 1, frequency, phase)
}

func (sch *PriorityScheduler) AddPriority(schedulable Schedulable, priority float32, frequency int, phase bt.Optional[int]) {
	var ph int
	if phase.Present() {
		ph = phase.Value()
	} else {
		ph = sch.calculatePhase(frequency)
	}

	sch.schedulable = append(sch.schedulable, PrioritySchedulerEntry{
		BaseSchedulerEntry: BaseSchedulerEntry{
			schedulable: schedulable,
			frequency:   frequency,
			phase:       ph,
		},
		priority: priority,
	})
}

//

type LoadBalancingScheduler struct {
	BaseScheduler[BaseSchedulerEntry]
}

var _ Scheduler = &LoadBalancingScheduler{}

func (sch *LoadBalancingScheduler) Run(ttl float32) {
	sch.frame++

	var runnable []BaseSchedulerEntry
	for _, entry := range sch.schedulable {
		if (sch.frame+entry.Phase())%entry.Frequency() == 0 {
			runnable = append(runnable, entry)
		}
	}

	lastTime := time.Now()
	for _, entry := range runnable {
		now := time.Now()
		ttl -= float32(now.Sub(lastTime).Seconds())
		availableTime := ttl / float32(len(runnable)-1)

		entry.schedulable.Run(availableTime)
		lastTime = now
	}
}

func (sch *LoadBalancingScheduler) Add(schedulable Schedulable, frequency int, phase bt.Optional[int]) {
	var ph int
	if phase.Present() {
		ph = phase.Value()
	} else {
		ph = sch.calculatePhase(frequency)
	}

	sch.schedulable = append(sch.schedulable,
		BaseSchedulerEntry{
			schedulable: schedulable,
			frequency:   frequency,
			phase:       ph,
		})
}
