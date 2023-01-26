//
/*
TODO:
 - DEFERRED DEPS
 - split Lifecycle into Lifecycle = any -> (lc Lifecylce, fn func(State) error) ?
*/
package lifecycles

import (
	"fmt"
	"sync"
)

type managerEntry struct {
	obj any
	fn  func(State) error
	st  State

	dependencies,
	dependents map[*managerEntry]struct{}
}

type Manager struct {
	mtx sync.Mutex
	st  State
	m   map[any]*managerEntry
}

func NewManager() *Manager {
	m := &Manager{
		m: make(map[any]*managerEntry),
	}
	return m
}

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
			return fmt.Errorf("lifecylce dependency not found: %v -> %v", e.obj, dep)
		}

		e.dependencies[de] = struct{}{}
		if de.dependents == nil {
			de.dependents = make(map[*managerEntry]struct{})
		}
		de.dependents[e] = struct{}{}
	}

	return nil
}

func (m *Manager) addInternal(obj any, fn func(State) error, deps []any) (*managerEntry, error) {
	if !dependableStates.Contains(m.st) {
		return nil, StateError{Current: m.st, Expected: dependableStates}
	}

	e := m.m[obj]
	if e != nil {
		return nil, fmt.Errorf("lifecycle already registered: %v", obj)
	}

	e = &managerEntry{
		obj: obj,
		fn:  fn,
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

	e.st = t.intermediate
	if err := e.fn(t.new); err != nil {
		e.st = t.failed
		return err
	}
	e.st = t.new

	return nil
}

func (m *Manager) advance(t *transition) error {
	if !t.old.Contains(m.st) {
		return StateError{Current: m.st, Expected: t.old}
	}

	m.st = t.intermediate
	for _, e := range m.m {
		if err := m.advanceEntry(e, t); err != nil {
			m.st = t.failed
			return err
		}
	}
	m.st = t.new

	return nil
}

func (m *Manager) Add(obj any, fn func(State) error, deps []any) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	if !dependableStates.Contains(m.st) {
		return StateError{Current: m.st, Expected: dependableStates}
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

/*
   def _do_lifecycle_construct(self) -> None:
       def rec(entry: LifecycleManager.Entry) -> None:
           for dep in entry.dependencies:
               rec(dep)
           if entry.controller.state.is_failed:
               raise LifecycleStateException(entry.controller)
           if entry.controller.state.phase < LifecycleStates.CONSTRUCTED.phase:
               entry.controller.lifecycle_construct()
       for entry in self._entries_by_lifecycle.values():
           rec(entry)

   def start(self) -> None:
       self.lifecycle_start()

   def _do_lifecycle_start(self) -> None:
       def rec(entry: LifecycleManager.Entry) -> None:
           for dep in entry.dependencies:
               rec(dep)
           if entry.controller.state.is_failed:
               raise LifecycleStateException(entry.controller)
           if entry.controller.state.phase < LifecycleStates.CONSTRUCTED.phase:
               entry.controller.lifecycle_construct()
           if entry.controller.state.phase < LifecycleStates.STARTED.phase:
               entry.controller.lifecycle_start()
       for entry in self._entries_by_lifecycle.values():
           rec(entry)

   def stop(self) -> None:
       self.lifecycle_stop()

   def _do_lifecycle_stop(self) -> None:
       def rec(entry: LifecycleManager.Entry) -> None:
           for dep in entry.dependents:
               rec(dep)
           if entry.controller.state.is_failed:
               raise LifecycleStateException(entry.controller)
           if entry.controller.state is LifecycleStates.STARTED:
               entry.controller.lifecycle_stop()
       for entry in self._entries_by_lifecycle.values():
           rec(entry)

   def destroy(self) -> None:
       self.lifecycle_destroy()

   def _do_lifecycle_destroy(self) -> None:
       def rec(entry: LifecycleManager.Entry) -> None:
           for dep in entry.dependents:
               rec(dep)
           if entry.controller.state.phase < LifecycleStates.DESTROYED.phase:
               entry.controller.lifecycle_destroy()
       for entry in self._entries_by_lifecycle.values():
           rec(entry)
*/
