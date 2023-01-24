package lifecycles

import (
	"fmt"
	"sync"
)

type Controller struct {
	lc *Lifecycle

	mtx sync.Mutex
	st  State
	cbs []Callback
}

func (c *Controller) State() State {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	return c.st
}

func (c *Controller) AddCallback(cb Callback) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.cbs = append(c.cbs, cb)
}

type transition struct {
	old StateMask

	intermediate State
	failed       State
	new          State

	fn func(*Lifecycle) LifecycleFn
}

var (
	construct = transition{
		old: New.Mask(),

		intermediate: Constructing,
		failed:       FailedConstructing,
		new:          Constructed,

		fn: func(lc *Lifecycle) LifecycleFn {
			return lc.Construct
		},
	}

	start = transition{
		old: Constructed.Mask(),

		intermediate: Starting,
		failed:       FailedStarting,
		new:          Started,

		fn: func(lc *Lifecycle) LifecycleFn {
			return lc.Start
		},
	}

	stop = transition{
		old: Started.Mask(),

		intermediate: Stopping,
		failed:       FailedStopping,
		new:          Stopped,

		fn: func(lc *Lifecycle) LifecycleFn {
			return lc.Stop
		},
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

		fn: func(lc *Lifecycle) LifecycleFn {
			return lc.Destroy
		},
	}
)

func (c *Controller) advance(t *transition) error {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if !t.old.Contains(c.st) {
		return fmt.Errorf("illegal state transition: %s -> %s", c.st, t.new)
	}

	for _, cb := range c.cbs {
		cb(c.lc, Before, t.new)
	}

	c.st = t.intermediate

	if fn := t.fn(c.lc); fn != nil {
		if err := fn(); err != nil {
			c.st = t.failed
			return err
		}
	}

	c.st = t.new

	for _, cb := range c.cbs {
		cb(c.lc, After, t.new)
	}

	return nil
}

func (c *Controller) Construct() error { return c.advance(&construct) }
func (c *Controller) Start() error     { return c.advance(&start) }
func (c *Controller) Stop() error      { return c.advance(&stop) }
func (c *Controller) Destroy() error   { return c.advance(&destroy) }
