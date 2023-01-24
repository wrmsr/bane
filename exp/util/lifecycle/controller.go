package lifecycles

import (
	"github.com/wrmsr/bane/pkg/util/check"
)

type controller struct {
	st  State
	cbs []Callback
}

func (c *controller) addCallback(cb Callback) {
	c.cbs = append(c.cbs, cb)
}

type transition struct {
	old StateMask

	intermediate,
	failed,
	new State
}

var (
	construct = transition{
		old: New.Mask(),

		intermediate: Constructing,
		failed:       FailedConstructing,
		new:          Constructed,
	}

	start = transition{
		old: Constructed.Mask(),

		intermediate: Starting,
		failed:       FailedStarting,
		new:          Started,
	}

	stop = transition{
		old: Started.Mask(),

		intermediate: Stopping,
		failed:       FailedStopping,
		new:          Stopped,
	}

	destroy = transition{
		old: New.Mask() |
			Constructing.Mask() |
			FailedConstructing.Mask() |
			Constructed.Mask() |
			Starting.Mask() |
			FailedStarting.Mask() |
			Started.Mask() |
			Stopping.Mask() |
			FailedStopping.Mask() |
			Stopped.Mask(),

		intermediate: Destroying,
		failed:       FailedDestroying,
		new:          Destroyed,
	}
)

func (c *controller) advance(lc *Lifecycle, t *transition) error {
	if !t.old.Contains(c.st) {
		return StateError{Current: c.st, Expected: t.old}
	}

	c.st = t.intermediate
	for _, cb := range c.cbs {
		cb(lc, t.intermediate)
	}
	check.Equal(c.st, t.intermediate)

	if err := lc.Fn(t.intermediate); err != nil {
		c.st = t.failed
		return err
	}

	c.st = t.new
	for _, cb := range c.cbs {
		cb(lc, t.new)
	}
	check.Equal(c.st, t.new)

	return nil
}
