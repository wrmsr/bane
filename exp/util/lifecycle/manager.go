//
/*
TODO:
 - threadsafety lol
 - DEFERRED DEPS
*/
package lifecycles

type managerEntry struct {
	obj any
	Controller

	dependencies,
	dependents map[*managerEntry]struct{}
}

type Manager struct {
	st  State
	cbs []Callback
	m   map[any]*managerEntry
}

func NewManager() *Manager {
	m := &Manager{
		m: make(map[any]*managerEntry),
	}
	return m
}

var _ Lifecycle = &Manager{}

func (m *Manager) State() State { return m.st }

func (m *Manager) AddCallback(cb Callback) {
	m.cbs = append(m.cbs, cb)
}

func (m *Manager) Construct() error { return m.advance(&construct) }
func (m *Manager) Start() error     { return m.advance(&start) }
func (m *Manager) Stop() error      { return m.advance(&stop) }
func (m *Manager) Destroy() error   { return m.advance(&destroy) }

var dependableStates = New.Mask() |
	Constructing.Mask() |
	Constructed.Mask() |
	Starting.Mask() |
	Started.Mask()

func (m *Manager) addDependencies(e *managerEntry, deps []any) error {
	if len(deps) < 1 {
		return nil
	}

	if e.dependencies == nil {
		e.dependencies = make(map[*managerEntry]struct{})
	}

	for _, dep := range deps {
		de := m.m[dep]
		if de == nil {
			return DependencyNotFoundError{e.obj, dep}
		}

		e.dependencies[de] = struct{}{}
		if de.dependents == nil {
			de.dependents = make(map[*managerEntry]struct{})
		}
		de.dependents[e] = struct{}{}
	}

	return nil
}

func (m *Manager) AddDependencies(obj any, deps []any) error {
	e := m.m[obj]
	if e == nil {
		return ObjectNotRegisteredError{obj}
	}

	return m.addDependencies(e, deps)
}

func (m *Manager) addInternal(obj any, fn func(State) error, deps []any) (*managerEntry, error) {
	if !dependableStates.Contains(m.st) {
		return nil, StateError{m.st, dependableStates}
	}

	e := m.m[obj]
	if e != nil {
		return nil, ObjectAlreadyRegisteredError{obj}
	}

	e = &managerEntry{
		obj: obj,
		Controller: Controller{
			fn: fn,
		},
	}
	m.m[obj] = e

	return e, m.addDependencies(e, deps)
}

func (m *Manager) advanceEntry(e *managerEntry, t *transition) error {
	if e.st >= t.new {
		return nil
	}

	if e.dependencies != nil {
		for dep := range e.dependencies {
			if err := m.advanceEntry(dep, t); err != nil {
				return err
			}
		}
	}

	return e.advance(t)
}

func (m *Manager) advance(t *transition) error {
	if !t.old.Contains(m.st) {
		return StateError{m.st, t.old}
	}

	m.st = t.intermediate
	for _, cb := range m.cbs {
		cb(m.st)
	}

	var err error
	for _, e := range m.m {
		if err = m.advanceEntry(e, t); err != nil {
			break
		}
	}

	if err != nil {
		m.st = t.failed
	} else {
		m.st = t.new
	}
	for _, cb := range m.cbs {
		cb(m.st)
	}

	return err
}

func (m *Manager) AddObject(obj any, fn func(State) error, deps []any) error {
	if !dependableStates.Contains(m.st) {
		return StateError{m.st, dependableStates}
	}

	e, err := m.addInternal(obj, fn, deps)
	if err != nil {
		return err
	}

	if m.st >= Constructing {
		if err := m.advanceEntry(e, &construct); err != nil {
			return err
		}
	}
	if m.st >= Starting {
		if err := m.advanceEntry(e, &start); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) AddObjectCallback(obj any, cb Callback) error {
	e := m.m[obj]
	if e == nil {
		return ObjectNotRegisteredError{obj}
	}

	e.cbs = append(e.cbs, cb)
	return nil
}
