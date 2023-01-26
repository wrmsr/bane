package lifecycles

type Controller struct {
	fn  Handler
	st  State
	cbs []Callback
}

var _ Lifecycle = &Controller{}

func (c Controller) State() State { return c.st }

func (c Controller) AddCallback(cb Callback) {
	c.cbs = append(c.cbs, cb)
}

func (c *Controller) Construct() error { return c.advance(&construct) }
func (c *Controller) Start() error     { return c.advance(&start) }
func (c *Controller) Stop() error      { return c.advance(&stop) }
func (c *Controller) Destroy() error   { return c.advance(&destroy) }

func (c *Controller) advance(t *transition) error {
	if !t.old.Contains(c.st) {
		return StateError{c.st, t.old}
	}

	c.st = t.intermediate
	for _, cb := range c.cbs {
		cb(c.st)
	}

	err := c.fn(t.new)

	if err != nil {
		c.st = t.failed
	} else {
		c.st = t.new
	}
	for _, cb := range c.cbs {
		cb(c.st)
	}

	return err
}
